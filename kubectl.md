# kubectl example commands

- Set context in kube-system namespace

```sh
kubectl set-context context_name --namespace kube-system
```

- Get pods with labels

```sh
kubectl get pods --show-labels
```

- Get all resources

```sh
kubectl get all
```

- Get multiple resources

```sh
kubectl get po,svc,pv,pvc,deploy,rs
```

- Output resource without creationTimestamp, selfLink, uid, ResourceVersion, etc.

```sh
kubectl get deployment deployment_name -o yaml --export
```

- Create service by exposing port

```sh
kubectl expose rc nginx --port=80 --target-port=8000
```

- Create service for pod

```sh
kubectl expose pod pod_name --port=444 --name=frontend
```

- Create ClusterIP service

```sh
kubectl create service clusterip service_name --tcp=5678:8080
```

- Set deployment's nginx container image to specific version and it's busybox container to busybox

```sh
kubectl set image deployment/nginx busybox=busybox nginx=nginx:1.9.1
```

- Update all deployments and replicasets using nginx to version 1.9.1

```sh
kubectl set image deployments,rc nginx=nginx:1.9.1 --all
```

- Update image of all containers of deamonset

```sh
kubectl set image deamonset ds_name *=nginx:1.9.1
```

- Create nginx deployment

```sh
kubectl run nginx --image=nginx
```

- Create nginx pod

```sh
kubectl run nginx --image=nginx --restart=Never
```

- Create nginx job

```sh
kubectl run nginx --image=nginx --restart=OnFailure
```

- Create nginx cronjob

```sh
kubectl run nginx --image=nginx --restart=OnFailure --schedule="* * * * *"
```

- Easily create pod resource yaml file

```sh
kubectl run nginx --image=nginx --restart=Never --port=80 --namespace=my_ns --command \
  --serviceaccount=ma_sa --env=HOSTNAME=local --labels=unit=finance,env=dev \
  --requests='cpu=100m,memory=256Mi' --limits='cpu=200m,memory=512Mi' \
  --dry-run -o yaml -- /bin/sh -c 'echo hello'
```

- Easily create deployment resource

```sh
kubectl run deploy_name --replicas=2 --labels=run=load-balancer-example --image=busybox --port=8080
```

- Easily create service

```sh
kubectl expose deyployment frontend --type=NodePort --name=frontend-service --port=6262 --target-port=8080
```

- Easily set service account for deployment

```sh
kubectl set serviceaccount deployment frontend user_name
```

- Easily create service resource

```sh
kubectl create service clusterip service_name --tcp=5678:8080 --dry-run -o yaml
```

- Rollout

```sh
kubectl rollout status
kubectl rollout history
kubectl rollout undo
```

- Browse documentation

```sh
kubectl explain --api-version="batch/v1beta1" cronjobs.spec
```

## Install from apt

```sh
wget -qO - https://packages.cloud.google.com/apt/doc/apt-key.gpg | gpg --dearmor | sudo dd of=/usr/share/keyrings/kubernetes-archive-keyring.gpg

echo 'deb [ arch=amd64 signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg ] http://apt.kubernetes.io/ kubernetes-xenial main' | sudo tee /etc/apt/sources.list.d/kubernetes.list

sudo apt-get update
sudo apt-get install kubectl
```
