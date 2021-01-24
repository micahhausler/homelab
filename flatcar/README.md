# Flatcar setup

## To document

* Disable UEFI boot on Intel NUC 10th gen ([Intel
    Docs](https://www.intel.com/content/www/us/en/support/articles/000032529/intel-nuc.html))
* Install steps on baremetal (Basically use [flatcar
    documented method](https://kinvolk.io/docs/flatcar-container-linux/latest/installing/bare-metal/installing-to-disk/)
    from an ubuntu install USB stick shell)

## TODOs

1. Write metadata gather script that writes to a file on disk
      1. Get ipv4 addr
      1. Get ipv6 addr
      1. Get ipv6 delegated range
      1. Get hostname
1. Invoke metadata script in oneshot systemd unit
1. Get `kube{let,ctl,adm}`, write to `/opt/bin/`
1. Get CNI binaries, write to `/opt/cni/bin`
1. Write Kubelet systemd unit
      1. Refer to EKS AMI
         [kubelet.service](https://github.com/awslabs/amazon-eks-ami/blob/master/files/kubelet.service)
         and
         [bootstrap.sh](https://github.com/awslabs/amazon-eks-ami/blob/master/files/bootstrap.sh)
         for args
1. Compile all units into ignition (and a script/tool that gathers them? A CDK
    for ignition would be sweet)
1. Figure out continerd-only setup for flatcar (documented one didn't seem to
   work). Maybe source above metadata + one or two args?
1. Write kubeadm templated config
