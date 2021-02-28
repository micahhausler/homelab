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
		PartitionSizes []int
		Want           types.Disk
	}{
		{
			"2 partition disk",
			"/dev/sda",
			[]int{10, 10, 15, 15, 25, 25, 50, 50, 50, 100, 150, 500},
			types.Disk{
				Device: "/dev/sda",
				Partitions: []types.Partition{
					types.Partition{Number: 1, Size: 1 << 30 * 10, Start: 0},
					types.Partition{Number: 2, Size: 1 << 30 * 10, Start: 1 << 30 * 10},
					types.Partition{Number: 3, Size: 1 << 30 * 15, Start: 1 << 30 * 20},
					types.Partition{Number: 4, Size: 1 << 30 * 15, Start: 1 << 30 * 35},
					types.Partition{Number: 5, Size: 1 << 30 * 25, Start: 1 << 30 * 50},
					types.Partition{Number: 6, Size: 1 << 30 * 25, Start: 1 << 30 * 75},
					types.Partition{Number: 7, Size: 1 << 30 * 50, Start: 1 << 30 * 100},
					types.Partition{Number: 8, Size: 1 << 30 * 50, Start: 1 << 30 * 150},
					types.Partition{Number: 9, Size: 1 << 30 * 50, Start: 1 << 30 * 200},
					types.Partition{Number: 10, Size: 1 << 30 * 100, Start: 1 << 30 * 250},
					types.Partition{Number: 11, Size: 1 << 30 * 150, Start: 1 << 30 * 350},
					types.Partition{Number: 12, Size: 1 << 30 * 500, Start: 1 << 30 * 500}},
				WipeTable: true},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			got := partitionDisk(tc.Device, tc.PartitionSizes)
			if !reflect.DeepEqual(got, tc.Want) {
				t.Errorf("Disk partitions did not match: Expected \n%#v \ngot \n%#v", tc.Want, got)
			}
		})
	}
}
