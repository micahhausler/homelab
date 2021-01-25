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
