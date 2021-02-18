### Install cert-manager

```
kubectl create namespace cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --version v1.1.0 \
  --set installCRDs=true \
  --set cainjector.image.repository=zengzhengrong889/cert-manager \
  --set cainjector.image.tag=v1.1.0
```

### Install ingress

```
kl apply -f .devcontainer/deploy-ingress.yaml
```


### Install rancher with Helm

```
kubectl create namespace cattle-system
```

```
helm install rancher rancher-latest/rancher \
 --namespace cattle-system \
 --set hostname=rancher.localhost
```

### Modify your hosts file 
