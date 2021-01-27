### Kubesphere install 

doc https://kubesphere.com.cn/docs/quick-start/minimal-kubesphere-on-k8s/

sc is ready,before Install Kubesphere 

```
kubectl apply -f https://github.com/kubesphere/ks-installer/releases/download/v3.0.0/kubesphere-installer.yaml
   
kubectl apply -f https://github.com/kubesphere/ks-installer/releases/download/v3.0.0/cluster-configuration.yaml

```

check 

```
kubectl logs -n kubesphere-system $(kubectl get pod -n kubesphere-system -l app=ks-install -o jsonpath='{.items[0].metadata.name}') -f

```

If you open container port mapping 30880 to host , brower localhost:30880



### Through ingress

yaml

```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    ingress.kubernetes.io/ssl-redirect: "false"
  name: kubesphere
  namespace: kubesphere-system
spec:
  rules:
  - host: "ks.localhost"
    http:
      paths:
      - pathType: Prefix
        path: /
        backend:
            service:
              name: ks-console
              port:
                number: 80
```

### k8s v1.19 issues

kl logs ks-controller-manager-648657b54d-7svzh -n kubesphere-system

```
x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0, requeuing
```

solve

```
kubectl apply -f kubesphere/
kubectl -n kubesphere-system rollout restart deploy ks-controller-manager
```


### Enable Plugin

config cluster-configuration and set true for plugin





