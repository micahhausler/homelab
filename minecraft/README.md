# Minecraft server

First be sure to set up a CSI that can provision PVs. I used the [local volume
provisioner][local-provisioner]. Then you can create the Minecraft server:
```bash
kubectl apply -f minecraft-stateful-set.yaml
```

[local-provisioner]: ../csi
