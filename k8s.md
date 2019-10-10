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
* [Awesome Kubernetes](https://github.com/ramitsurana/awesome-kubernetes)
* [stern - tail multiple pods on Kubernetes and multiple containers within the pod](https://github.com/wercker/stern)
* [Kubernetes SIG](https://github.com/kubernetes-sigs)
* [12 Kubernetes configuration best practices](https://www.stackrox.com/post/2019/09/12-kubernetes-configuration-best-practices/)

## Blogs

* [Alex Ellis](https://blog.alexellis.io/)

## Tools

* [Clair - Vulnerability static analysis for containers](https://github.com/coreos/clair)
* [JenkinsX - Continuous delivery for Kubernetes](https://jenkins-x.io/demos/devoxx-uk-2018/)
* [Kompose - translate docker-compose files to Kubernetes resources](https://github.com/kubernetes/kompose)
* [kuberhealthy - synthetic testing](https://github.com/Comcast/kuberhealthy)
* [Keel - Kubernetes Operator to automate Helm, DaemonSet, StatefulSet & Deployment updates](https://github.com/keel-hq/keel)
* [Keel webhooks](https://keel.sh/v1/guide/documentation.html#Triggers)
* [M3 - Uber’s large-scale metrics platform for Prometheus](https://eng.uber.com/m3/)
* [kubectx & kubens - Switch faster between clusters and namespaces](https://github.com/ahmetb/kubectx)
* [kubeone - Lifecycle management tool for highly available Kubernetes clusters](https://github.com/kubermatic/kubeone)
* [TK8 - Multi-cloud, multi-cluster Kubernetes platform installation and integration tool](https://github.com/kubernauts/tk8)
* [kaniko - Build Images In Kubernetes](https://github.com/GoogleContainerTools/kaniko)
* [Kontemplate - Simple Kubernetes templater](https://github.com/tazjin/kontemplate)

## CI & CD

* [Tekton Pipelines](https://github.com/tektoncd/pipeline)
* [Skaffold](https://github.com/GoogleContainerTools/skaffold)
* [How to use Knative Pipelining Component to automate an Application Build and Deployment on Kubernetes](https://itnext.io/how-to-use-knative-pipelining-component-to-automate-an-application-build-and-deployment-on-442b0b1bebf)
* [Build cloud native CI/CD build pipeline from GIT webhook](https://medium.com/@pongsatt/build-cloud-native-ci-cd-build-pipeline-from-git-webhook-9cd9a57a32e1)
* [kontinuous - Kubernetes Continuous Integration & Delivery Platform](https://github.com/AcalephStorage/kontinuous)

## Backup

* [Velero (formerly Ark) - Backup and migrate Kubernetes applications and their persistent volumes](https://github.com/heptio/velero)
[Backup Kubernetes – how and why](https://elastisys.com/2018/12/10/backup-kubernetes-how-and-why/)
[How To Back Up and Restore Kubernetes Cluster using Ark](https://www.digitalocean.com/community/tutorials/how-to-back-up-and-restore-a-kubernetes-cluster-on-digitalocean-using-heptio-ark)
[Backup etcd cluster](https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/#backing-up-an-etcd-cluster)

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
* [Helmfile - deploy Kubernetes Helm Charts](https://github.com/roboll/helmfile)
* [Helmfile - chart deployment tool](https://medium.com/@devopseylife/helmfile-aka-helm-chart-deployment-tool-4e3378fad242)
* [helmfile  -  it’s like a Helm for Helm](https://medium.com/@naseem_60378/helmfile-its-like-a-helm-for-your-helm-74a908581599)
* [15+ useful Helm Charts tools](https://caylent.com/15-useful-helm-charts-tools/)
* [ChartMuseum - host your own Helm chart repository](https://github.com/helm/chartmuseum)
* [Using incubator/raw chart](https://github.com/roboll/helmfile/issues/494#issuecomment-474697430)
* [incubator/raw - chart for kubectl declarations](https://github.com/helm/charts/tree/master/incubator/raw)
* [Awesome Helm](https://github.com/cdwv/awesome-helm)
* [Helm Chart starter - better Helm chart boilerplate](https://github.com/sitewards/helm-chart)
* [Creating a better chart](https://medium.com/sitewards/deploying-on-kubernetes-2-scaffolding-6a54e5d181fb)

## Custom Resource Defintions and Development

* [Accessing Kubernetes CRDs from client-go package](https://www.martin-helmich.de/en/blog/kubernetes-crd-client.html)
* [Extending Kubernetes APIs with Custom Resource Definitions](https://medium.com/velotio-perspectives/extending-kubernetes-apis-with-custom-resource-definitions-crds-139c99ed3477)
* [Unit testing](https://medium.com/@e_frogers/unit-testing-with-kubernetes-client-go-283b11aaa7db)
* [Code Generation for CustomResources](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/)
* [How to generate client codes for Kubernetes Custom Resource Definitions](https://itnext.io/how-to-generate-client-codes-for-kubernetes-custom-resource-definitions-crd-b4b9907769ba)
* [Writing Kubernetes Custom Controllers](https://medium.com/@cloudark/kubernetes-custom-controllers-b6c7d0668fdf)
* [How did that sidecar get there?](https://medium.com/dowjones/how-did-that-sidecar-get-there-4dcd73f1a0a4)
* [k8s-sidecar-injector](https://github.com/tumblr/k8s-sidecar-injector)

## Security

* [Kubernetes Security Practices You Should Follow](https://blog.sonatype.com/kubesecops-kubernetes-security-practices-you-should-follow)
* [Sealed Secrets](https://github.com/bitnami-labs/sealed-secrets)
* [Vault](https://www.vaultproject.io/)

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

* Browse documentation

```sh
kubectl explain --api-version="batch/v1beta1" cronjobs.spec
```

* Reboot cluster

** stop kupe-apiserver on master
** stop kube-scheduler on master
** stop kube-controller on master
** stop kubelet on master (if applicable)
** kube-proxy on master (if applicable)
** stop federation-apiserver on master (if applicable)
** backup etcd
** stop etcd
** for each node stop kubelet and kube-proxy
