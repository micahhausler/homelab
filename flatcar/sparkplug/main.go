package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/micahhausler/homelab/flatcar/sparkplug/spark"
	flag "github.com/spf13/pflag"
)

func main() {
	user := flag.String("user", "core", "username to create")
	groups := flag.StringSliceP("groups", "g", []string{"sudo", "docker"}, "Groups to add to the user")
	authorizedKeys := flag.String("authorized-keys-file", "", "Authorized key file")
	hostname := flag.String("hostname", "", "Hostname to set for the config")
	metadataScript := flag.String("metadata-script", "", "Path to metadata shell script source")
	indent := flag.Bool("indent", true, "Indent output with 2 spaces")
	mountDevice := flag.String("mount-device", "", "The device to add to mounts and filesystems. An empty value omits partitions and filesystems")
	mountPartitionCount := flag.Int("mount-partiition-count", 2, "The number of partitions to create")
	mountPartitionSize := flag.Int("mount-partition-size", 500, "The size (in GB) of partitions to create")
	mountRoot := flag.String("mount-root", "/mnt/csi-local-storage", "The directory to mount secondary partitions into")
	fsType := flag.String("fs-type", "ext4", "The filesystem type")

	flag.Parse()

	opts := []spark.ConfigOpt{}
	if *metadataScript != "" {
		opts = append(opts, spark.MetadataConfigOpt(*metadataScript))
	}
	if *hostname != "" {
		opts = append(opts, spark.HostnameConfigOpt(*hostname))
	}
	if *user != "" {
		opts = append(opts, spark.NewUserOpt(*user, *groups, *authorizedKeys))
	}

	if *mountDevice != "" {
		opts = append(opts, spark.NewDiskOpt(
			*mountRoot,
			*fsType,
			*mountDevice,
			*mountPartitionCount,
			*mountPartitionSize,
		))
	}

	config, err := spark.NewConfig(opts...)
	if err != nil {
		fmt.Printf("Fatal error creating config: %v\n", err)
		os.Exit(1)
	}

	report := config.Validate()
	if len(report.Entries) > 0 {
		for _, entry := range report.Entries {
			fmt.Println(entry)
		}
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	if *indent {
		enc.SetIndent("", "  ")
	}

	err = enc.Encode(config)
	if err != nil {
		fmt.Printf("Error serializing output: %v\n", err)
		os.Exit(1)
	}

}
