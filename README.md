# Intro
Kubernetes Shell Operator.

Big shout out and thank you to the [Flant team](https://flant.github.io/shell-operator/index.html) for creating the Shell Operator.

# Some commands

```bash
make go-build-linux
make go-build-windows
```

```bash
docker build --no-cache -f ./Dockerfile -t kso:latest .
kind load docker-image kso:latest --name demo
```

```bash
# prep
kubectl apply -f ./manifests/ns_demo.yaml
kubectl apply -f ./manifests/crd_tenant.yaml

# operator
kubectl apply -f ./manifests/dpt_kso.yaml
kubectl logs -n demo --all-containers=true deployment/kso

# test
kubectl apply -f ./tests/hc_tenant.yaml
```


```bash
kubectl delete -f ./tests/hc_tenant.yaml
kubectl delete -f ./manifests/dpt_kso.yaml
kubectl delete -f ./manifests/crd_tenant.yaml
kubectl delete -f ./manifests/ns_demo.yaml
```
