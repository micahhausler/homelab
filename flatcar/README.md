# Flatcar setup

## To document

* Disable UEFI boot on Intel NUC 10th gen ([Intel Docs][intel docs])
* Install steps on baremetal (Basically use [flatcar documented
  method][flatcar-baremetal] from an Ubuntu install USB stick shell)

[intel docs]: https://www.intel.com/content/www/us/en/support/articles/000032529/intel-nuc.html
[flatcar-baremetal]: https://kinvolk.io/docs/flatcar-container-linux/latest/installing/bare-metal/installing-to-disk/

## TODOs

- [ ]. Research using [AWS ACM Private CA][acm-pca]
   for enabling SecureBoot ([Debian docs][debian-secure-boot]
- [x] Write metadata gather script that writes to files on disk
      - [x] Get ipv4 addr
      - [x] Get ipv6 addr
      - [x] Get ipv6 delegated range
      - [x] Get hostname
- [ ] Invoke metadata script in oneshot systemd unit
- [ ] Get `kube{let,ctl,adm}`, write to `/opt/bin/`
- [ ] Get CNI binaries, write to `/opt/cni/bin`
- [ ] Write Kubelet systemd unit
    - [ ] Refer to EKS AMI [kubelet.service][kubelet.service] and
       [bootstrap.sh][bootstrap.sh] for args
- [x] Compile all units into ignition (and a script/tool that gathers them? A CDK
    for ignition would be sweet)
    - [x] See [sparkplug](./sparkplug/)
- [ ] Figure out continerd-only setup for flatcar (documented one didn't seem to
   work). Maybe source above metadata + one or two args?
- [ ] Write kubeadm templated config

[acm-pca]: https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-private.html
[debian-secure-boot]: https://wiki.debian.org/SecureBoot#What_is_UEFI.3F
[kubelet.service]: https://github.com/awslabs/amazon-eks-ami/blob/master/files/kubelet.service
[bootstrap.sh]: https://github.com/awslabs/amazon-eks-ami/blob/master/files/bootstrap.sh

## Ignition

To re-run ignition with a new configuration on an already-running machine, 

## Kubernetes

### Kubeadm notes

```bash
#Generate random token
TOKEN=$(kubeadm token generate)

# Enable IPv6 Forwarding
# TODO figure out how to persist this across reboot
echo "1" | sudo tee  /proc/sys/net/ipv6/conf/default/forwarding

sudo mkdir -p /etc/systemd/system/kubelet.service.d/

docker pull public.ecr.aws/eks-distro/kubernetes/pause:v1.18.9-eks-1-18-1
docker tag public.ecr.aws/eks-distro/kubernetes/pause:v1.18.9-eks-1-18-1 public.ecr.aws/eks-distro/kubernetes/pause:3.2

# /usr is read-only
sed -i 's,path: /usr/libexec,path: /opt/libexec,g'  /etc/kubernetes/manifests/kube-controller-manager.yaml

```

* Service subnet is an ipv6 `/116`
* Pod subnet is an ipv6 `/68`

### K8s TODOs

kube-controller manager is exiting with log:
```
Controller: Invalid --cluster-cidr, mask size of cluster CIDR must be less than or equal to --node-cidr-mask-size configured for CIDR family
```


Maybe just use ipv4 to get it working first?


Blow everything away and start over
```
sudo rm -rf /etc/kubernetes/ /var/lib/kubelet /var/lib/etcd
```
