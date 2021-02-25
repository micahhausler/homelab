# CSI Local Driver setup

I've spec'd out my NUC with a [1TB Western Digital SATA SSD][1tb-m2], which gets
assigned the device `/dev/sda`. I use [Ignition][ignition-storage] to partition,
create filesystems, and mount the device.

[1tb-m2]: https://www.newegg.com/samsung-860-evo-series-1tb/p/N82E16820147673?Item=N82E16820147673
[ignition-storage]: https://kinvolk.io/docs/flatcar-container-linux/latest/setup/storage/mounting-storage/

## Manual Disk partitioning

To manually create the disk partitions and mount them, you can run:
```bash
# Partition disk manually
sudo parted -s -a opt --script /dev/sda \
    mklabel gpt \
    mkpart primary 0 500G \
    mkpart primary 500GB 1000GB \
    align-check min 1 
# Make file systems
sudo mkfs -t ext4 /dev/sda1
sudo mkfs -t ext4 /dev/sda2

# Create mount points
sudo mkdir -p /mnt/csi-local-storage/{p1,p2}
# TODO: add to mtab/fstab
sudo mount /dev/sda1 /mnt/csi-local-storage/p1
sudo mount /dev/sda2 /mnt/csi-local-storage/p2
```

## Install the Kubernetes CSI local static provisioner

You can read the [K8s blog post][csi-blog] for a full explanation and
walk-through, but basically to get local persistent Kubernetes volumes, you have
to:
* Create a StorageClass for local disks
* Mount your disks into a directory on the host filesystem
* Run the provisioner as a DaemonSet and specify the mount directory

Do all that, and the provisioner will create Kubernetes Persistent Volumes (PVs)
that can be claimed by Persistent Volume Claims (PVCs) for pods.


[csi-blog]: https://kubernetes.io/blog/2019/04/04/kubernetes-1.14-local-persistent-volumes-ga/

To setup the provisioner, run:

```bash
git clone --depth=1 https://github.com/kubernetes-sigs/sig-storage-local-static-provisioner.git
# I use my own values here, you can use yours
helm template csi-local-provisioner \
    --namespace kube-system \
    ./sig-storage-local-static-provisioner/helm/provisioner \
    --values local-values.yaml > local-volume-provisioner.generated.yaml
kubectl apply -f local-volume-provisioner.generated.yaml
```
