package spark

import (
	"fmt"
	"path/filepath"

	"github.com/kinvolk/ignition/config/v2_2/types"
)

func partitionDisk(device string, partitionSizes []int) types.Disk {
	disk := types.Disk{
		Device:    device,
		WipeTable: true,
	}
	parts := []types.Partition{}

	var start = 0
	for i, pSize := range partitionSizes {
		pSizeGB := 1 << 30 * pSize
		part := types.Partition{
			Start:  start,
			Number: i + 1,
			Size:   pSizeGB,
		}
		parts = append(parts, part)
		start += pSizeGB
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

func NewDiskOpt(mountRoot, fstype, device string, partitionSizes []int) ConfigOpt {
	return ConfigOpt(func(config *types.Config) error {
		disk := partitionDisk(device, partitionSizes)
		disks := config.Storage.Disks
		config.Storage.Disks = append(disks, disk)

		config.Storage.Filesystems = append(config.Storage.Filesystems, fileSystems(mountRoot, fstype, device, len(partitionSizes))...)

		return nil
	})
}
