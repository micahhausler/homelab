package spark

import (
	"fmt"
	"github.com/kinvolk/ignition/config/v2_2/types"
)

func metadata(source string) (types.File, error) {
	source, err := scriptToSource(source)
	if err != nil {
		return types.File{}, err
	}
	return types.File{
		types.Node{
			Filesystem: "root",
			Path:       "/opt/bin/metadata.sh",
		},
		types.FileEmbedded1{
			Mode: intPtr(755),
			Contents: types.FileContents{
				Source: source,
			},
		},
	}, nil
}

var metadataUnit = `[Unit]
Description=Metadata File provisioning

[Service]
Type=oneshot
ExecStart=/usr/bin/bash /opt/bin/metadata.sh
`

func MetadataConfigOpt(source string) ConfigOpt {
	fileFunc := metadata
	return ConfigOpt(func(config *types.Config) error {
		// Create metadata script
		file, err := fileFunc(source)
		if err != nil {
			return err
		}
		files := config.Storage.Files
		config.Storage.Files = append(files, file)

		// Create metadata systemd unit
		units := config.Systemd.Units
		config.Systemd.Units = append(units, types.Unit{
			Name:     "metadata.service",
			Enabled:  boolPtr(true),
			Contents: metadataUnit,
		})
		return nil
	})
}

func hostnameFile(hostname string) (types.File, error) {
	return types.File{
		types.Node{
			Filesystem: "root",
			Path:       "/etc/hostname",
		},
		types.FileEmbedded1{
			Mode: intPtr(420),
			Contents: types.FileContents{
				Source: fmt.Sprintf("data:,%s", hostname),
			},
		},
	}, nil
}

func HostnameConfigOpt(hostname string) ConfigOpt {
	return fileFromSource(hostname, hostnameFile)
}
