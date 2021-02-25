# sparkplug

> The ignition config writer

sparkplug writes an ignition config JSON file with the things I think are cool.

```
Usage of sparkplug:
      --authorized-keys-file string   Authorized key file
      --fs-type string                The filesystem type (default "ext4")
  -g, --groups strings                Groups to add to the user (default [sudo,docker])
      --hostname string               Hostname to set for the config
      --indent                        Indent output with 2 spaces (default true)
      --metadata-script string        Path to metadata shell script source
      --mount-device string           The device to add to mounts and filesystems. An empty value omits partitions and filesystems
      --mount-partiition-count int    The number of partitions to create (default 2)
      --mount-partition-size int      The size (in GB) of partitions to create (default 500)
      --mount-root string             The directory to mount secondary partitions into (default "/mnt/csi-local-storage")
      --user string                   username to create (default "core")
```
