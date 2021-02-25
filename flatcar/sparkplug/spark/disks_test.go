package spark

import (
	"reflect"
	"testing"

	"github.com/kinvolk/ignition/config/v2_2/types"
)

func TestPartitionDisk(t *testing.T) {
	cases := []struct {
		Name           string
		Device         string
		PartitionCount int
		PartitionSize  int
		Want           types.Disk
	}{
		{
			"2 partition disk",
			"/dev/sdb",
			2,
			1 << 30 * 500,
			types.Disk{
				Device:    "/dev/sdb",
				WipeTable: true,
				Partitions: []types.Partition{
					types.Partition{
						Start:  0,
						Number: 1,
						Size:   1 << 30 * 500,
					},
					types.Partition{
						Start:  1 << 30 * 500,
						Number: 2,
						Size:   1 << 30 * 500,
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			got := partitionDisk(tc.Device, tc.PartitionCount, tc.PartitionSize)
			if !reflect.DeepEqual(got, tc.Want) {
				t.Errorf("Disk partitions did not match: Expected %#v, got %#v", tc.Want, got)
			}
		})
	}
}
