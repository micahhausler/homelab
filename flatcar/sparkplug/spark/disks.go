package spark

import (
	"fmt"
	"github.com/kinvolk/ignition/config/v2_2/types"
	"path/filepath"
)

func partitionDisk(device string, partitionCount, partitionSize int) types.Disk {
	disk := types.Disk{
		Device:    device,
		WipeTable: true,
	}
	parts := []types.Partition{}

	for i := 0; i < partitionCount; i++ {
		part := types.Partition{
			Start:  i * partitionSize,
			Number: i + 1,
			Size:   partitionSize,
		}
		parts = append(parts, part)
	}
	disk.Partitions = parts
	return disk
}

func fileSystems(mountRoot, fstype, device string, partitionCount int) []types.Filesystem {
	fss := []types.Filesystem{}
	for i := 0; i < partitionCount; i++ {

		fss = append(fss, types.Filesystem{
			Path: strPtr(filepath.Join(mountRoot, fmt.Sprintf("p%d", i+1))),
			Name: fmt.Sprintf("local-storage-%d", i+1),
			Mount: &types.Mount{
				Device:         fmt.Sprintf("%s%d", device, i+1),
				Format:         fstype,
				WipeFilesystem: true,
			},
		})
	}
	return fss
}

func NewDiskOpt(mountRoot, fstype, device string, partitionCount, partitionSize int) ConfigOpt {
	return ConfigOpt(func(config *types.Config) error {
		disk := partitionDisk(device, partitionCount, partitionSize)
		disks := config.Storage.Disks
		config.Storage.Disks = append(disks, disk)

		config.Storage.Filesystems = append(config.Storage.Filesystems, fileSystems(mountRoot, fstype, device, partitionCount)...)

		return nil
	})
}
