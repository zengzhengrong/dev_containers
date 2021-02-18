### Tips

add you localhost ip to dev container

```
code /etc/hosts 
```

```
127.0.0.1	localhost
192.168.2.105	localhost # add here
::1	localhost ip6-localhost ip6-loopback
fe00::0	ip6-localnet
ff00::0	ip6-mcastprefix
ff02::1	ip6-allnodes
ff02::2	ip6-allrouters
172.17.0.3	84bb13b90abd
```

### Create you cluster by k3d


```
k3d cluster create -p "8081:80@loadbalancer" --agents 2
```

update kubeconfig 0.0.0.0 to localhost

```
apiVersion: v1
clusters:
- cluster:
    server: https://localhost:38381 # 0.0.0.0 to localhost
  name: k3d-k3s-default
```


### mirrors 


```
mirrors:
  "docker.io":
    endpoint:
      - "https://fogjl973.mirror.aliyuncs.com"
      - "https://registry-1.docker.io"
```

### RUN 

```
k3d cluster create demo --volume "/etc/rancher/k3d/registries.yaml:/etc/rancher/k3s/registries.yaml@server[0]" -p "8081:80@loadbalancer" -a 2
```
if failed ,just create a cluster no use -v commmand, and docker cp registries.yaml to server and stop , start 
```
k3d cluster create demo -p "8080:80@loadbalancer" --agents 2
```

```
docker cp /etc/rancher/k3d/registries.yaml <container id>:/etc/rancher/k3s/registries.yaml
```

```
k3d cluster stop demo
```

```
k3d cluster start demo
```

check 
```
docker exec -it <container id> cat /etc/rancher/k3s/registries.yaml
```


### Use nginx ingress instead traefik

bug in k3d .......
```
k3d cluster create dev --k3s-server-arg '--disable traefik' -p "8080:80@loadbalancer" -a 2
```


### Test

```
kubectl create deployment nginx --image=nginx
```

```
kubectl create service clusterip nginx --tcp=80:80
```

```
cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: nginx
  annotations:
    ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: nginx
            port:
              number: 80
EOF
```

#### install rancher

1.
```
helm install rancher rancher-latest/rancher \
 --namespace cattle-system \
 --set hostname=localhost
```
update 
```
helm upgrade rancher rancher-latest/rancher \
 --namespace cattle-system \
 --set hostname=localhost
```

2.
```
docker run -d --privileged --restart=unless-stopped \
    -p 8080:80 -p 8443:443 \
    rancher/rancher:latest
```

### install krew plugin

```
PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"
```

kl krew install tree images ns rbac-view rbac-view pod-inspect

