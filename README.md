# Micah's homelab setup

## Images
Images are x86_64 only.

### Alpine Linux

`public.ecr.aws/s3t5k9h7/alpine:3.13.0`

### Samba

`public.ecr.aws/s3t5k9h7/samba:4.13.3`

### Avahi

`public.ecr.aws/s3t5k9h7/avahi:0.8`

## Project Plan

Started
* Samba Time Machine backup server
* Avahi autodiscovery

TODO:
* K8s configs for samba/avahi
* K8s setup scripts
* Minecraft server setup on K8s
* Tailscale install
    * Or set up wireguard
* In-cluster CoreDNS as LAN DNS
* Personal [Prow](https://prow.k8s.io/) setup

## Hardware Inventory

* [Intel NUC
    i3](https://www.newegg.com/intel-bxnuc10i3fnh1/p/N82E16856102231?Item=N82E16856102231) 32GB DDR4 RAM
    * [500GB WD PCIe
        SSD](https://www.newegg.com/western-digital-black-sn750-nvme-500gb/p/N82E16820250109?Item=N82E16820250109)
    * [1TB WD SATA
        SSD](https://www.newegg.com/samsung-860-evo-series-1tb/p/N82E16820147673?Item=N82E16820147673)

## License

Apache 2.0 [LICENSE](./LICENSE)
