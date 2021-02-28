# sparkplug

> The ignition config writer

sparkplug writes an ignition config JSON file with the things I think are cool.

```
Usage of sparkplug:
      --authorized-keys-file string   Authorized key file
      --disabled-units strings        Systemd untis to disable (default [tcsd.service])
      --fs-type string                The filesystem type (default "ext4")
  -g, --groups strings                Groups to add to the user (default [sudo,docker])
      --hostname string               Hostname to set for the config
      --indent                        Indent output with 2 spaces (default true)
      --metadata-script string        Path to metadata shell script source
      --mount-device string           The device to add to mounts and filesystems. An empty value omits partitions and filesystems
      --mount-partition-size ints     A list of partition sizes in GB. Partitions will be created in the order they are supplied
      --mount-root string             The directory to mount secondary partitions into (default "/mnt/csi-local-storage")
      --user string                   username to create (default "core")
```

### Example 

```bash
sparkplug  \
    --authorized-keys-file ./keys.txt  \
    --user micahhausler \
    --hostname nuc1 \
    --mount-device /dev/sda \
    --mount-partition-size 10,10,15,15,25,25,50,50,50,100,150,500 
```
