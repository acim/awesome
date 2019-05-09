# Kubernetes

* [Kubernetes by example](http://kubernetesbyexample.com/)
* [Istio](https://www.youtube.com/watch?v=s4qasWn_mFc)
* [Using NodeSelector to Schedule Deployments with large volumes of Stateful Data on Kubernetes](https://medium.com/@jmarhee/using-nodeselector-to-schedule-deployments-with-large-volumes-of-stateful-data-on-kubernetes-46bd3ac6059d)
* [Scalable microservices with Kubernetes](https://eu.udacity.com/course/scalable-microservices-with-kubernetes--ud615)
* [Kubernetes cheat sheet](https://codefresh.io/kubernetes-tutorial/kubernetes-cheat-sheet/)
* [SkyDNS](https://github.com/skynetservices/skydns)
* [Single master cluster with kubeadm](https://kubernetes.io/docs/setup/independent/create-cluster-kubeadm/)
* [HA Kubernetes cluster on bare metal](https://github.com/salmanb/Kubernetes-HA-on-baremetal)
* [Encryption of Kubernetes persistent local volumes](https://medium.com/@dfrnascimento/encryption-of-kubernetes-persistent-local-volumes-70da62e0ed68)
* [Static provisioner of local volumes](https://github.com/kubernetes-sigs/sig-storage-local-static-provisioner)
* [OpenEBS - Creating and attaching disk on GKE node](https://docs.openebs.io/docs/next/setupstoragepools.html#creating-and-attaching-a-disk-on-gke-node)
* [Kubespray - Deploy a production ready Kubernetes cluster](https://github.com/kubernetes-sigs/kubespray)
* [Community Operators](https://commons.openshift.org/sig/operators.html)
* [Ingress basic auth](https://kubernetes.github.io/ingress-nginx/examples/auth/basic/)
* [Generate RBAC policies based on Kubernetes audit logs](https://github.com/liggitt/audit2rbac)

## Blogs

* [Alex Ellis](https://blog.alexellis.io/)

## Tools

* [Clair - Vulnerability static analysis for containers](https://github.com/coreos/clair)
* [JenkinsX - Continuous delivery for Kubernetes](https://jenkins-x.io/demos/devoxx-uk-2018/)
* [ChartMuseum - Host your own Helm chart repository](https://github.com/helm/chartmuseum)
* [Kompose - translate docker-compose files to Kubernetes resources](https://github.com/kubernetes/kompose)
* [kuberhealthy - synthetic testing](https://github.com/Comcast/kuberhealthy)
* [Keel - Kubernetes Operator to automate Helm, DaemonSet, StatefulSet & Deployment updates](https://github.com/keel-hq/keel)
* [kontinuous - Kubernetes Continuous Integration & Delivery Platform](https://github.com/AcalephStorage/kontinuous)
* [M3 - Uber’s large-scale metrics platform for Prometheus](https://eng.uber.com/m3/)
* [kubectx & kubens - Switch faster between clusters and namespaces](https://github.com/ahmetb/kubectx)

## Courses

* [Stephen Grider - Udemy - Docker casts](https://github.com/StephenGrider/DockerCasts)
* [Ward Viaene - On premise or cloud-agnostic Kubernetes](https://github.com/wardviaene/on-prem-or-cloud-agnostic-kubernetes)
* [Ward Viaene - Kubernetes course](https://github.com/wardviaene/kubernetes-course)

## Helm

* [Helm hub](https://hub.helm.sh/)
* [Charts](https://github.com/helm/charts)
* [Solr](https://github.com/guigo2k/helm-solr)
* [Portus](https://github.com/kubic-project/caasp-services/tree/master/contrib/helm-charts/portus)
* [Plugins](https://docs.helm.sh/related/)
* [Template developer’s guide](https://helm.sh/docs/chart_template_guide/)
* [helmfile - Deploy Kubernetes Helm Charts](https://github.com/roboll/helmfile)

## Custom Resource Defintions

* [Accessing Kubernetes CRDs from client-go package](https://www.martin-helmich.de/en/blog/kubernetes-crd-client.html)
* [Extending Kubernetes APIs with Custom Resource Definitions](https://medium.com/velotio-perspectives/extending-kubernetes-apis-with-custom-resource-definitions-crds-139c99ed3477)

## Pin Docker version

apt-get install docker-ce="18.06.0~ce~3-0~ubuntu"

## kubectl Cheat Sheet

[Official Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)

### Autocomplete

* Turn on autocomplete in the current shell

```sh
source <(kubectl completion bash)
```

* Turn on autocomplete permanently

```sh
echo "source <(kubectl completion bash)" >> ~/.bashrc
```

* Make alias k and autocomplete for it

```sh
alias k=kubectl
complete -F __start_kubectl k
```

### Set context in kube-system namespace

```sh
kubectl set-context context_name --namespace kube-system
```

### Example commands

* Get pods with labels

```sh
kubectl get pods --show-labels
```

* Get all resources

```sh
kubectl get all
```

* Get multiple resources

```sh
kubectl get po,svc,pv,pvc,deploy,rs
```

* Output resource without creationTimestamp, selfLink, uid, ResourceVersion, etc.

```sh
kubectl get deployment deployment_name -o yaml --export
```

* Create service by exposing port

```sh
kubectl expose rc nginx --port=80 --target-port=8000
```

* Create service for pod

```sh
kubectl expose pod pod_name --port=444 --name=frontend
```

* Create ClusterIP service

```sh
kubectl create service clusterip service_name --tcp=5678:8080
```

* Set deployment's nginx container image to specific version and it's busybox container to busybox

```sh
kubectl set image deployment/nginx busybox=busybox nginx=nginx:1.9.1
```

* Update all deployments and replicasets using nginx to version 1.9.1

```sh
kubectl set image deployments,rc nginx=nginx:1.9.1 --all
```

* Update image of all containers of deamonset

```sh
kubectl set image deamonset ds_name *=nginx:1.9.1
```

* Create nginx deployment

```sh
kubectl run nginx --image=nginx
```

* Create nginx pod

```sh
kubectl run nginx --image=nginx --restart=Never
```

* Create nginx job

```sh
kubectl run nginx --image=nginx --restart=OnFailure
```

* Create nginx cronjob

```sh
kubectl run nginx --image=nginx --restart=OnFailure --schedule="* * * * *"
```

* Easily create pod resource yaml file

```sh
kubectl run nginx --image=nginx --restart=Never --port=80 --namespace=my_ns --command \
  --serviceaccount=ma_sa --env=HOSTNAME=local --labels=unit=finance,env=dev \
  --requests='cpu=100m,memory=256Mi' --limits='cpu=200m,memory=512Mi' \
  --dry-run -o yaml -- /bin/sh -c 'echo hello'
```

* Easily create deployment resource

```sh
kubectl run deploy_name --replicas=2 --labels=run=load-balancer-example --image=busybox --port=8080
```

* Easily create service

```sh
kubectl expose deyployment frontend --type=NodePort --name=frontend-service --port=6262 --target-port=8080
```

* Easily set service account for deployment

```sh
kubectl set serviceaccount deployment frontend user_name
```

* Easily create service resource

```sh
kubectl create service clusterip service_name --tcp=5678:8080 --dry-run -o yaml
```

* Rollout

```sh
kubectl rollout status
kubectl rollout history
kubectl rollout undo
```