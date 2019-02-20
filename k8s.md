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

## Blogs

* [Alex Ellis](https://blog.alexellis.io/)

## Tools

* [Clair - Vulnerability static analysis for containers](https://github.com/coreos/clair)
* [JenkinsX - Continuous delivery for Kubernetes](https://jenkins-x.io/demos/devoxx-uk-2018/)
* [ChartMuseum - Host your own Helm chart repository](https://github.com/helm/chartmuseum)
* [Kompose - translate docker-compose files to Kubernetes resources](https://github.com/kubernetes/kompose)

## Courses

* [Stephen Grider - Udemy - Docker casts](https://github.com/StephenGrider/DockerCasts)
* [Ward Viaene - On premise or cloud-agnostic Kubernetes](https://github.com/wardviaene/on-prem-or-cloud-agnostic-kubernetes)
*[Ward Viaene - Kubernetes course](https://github.com/wardviaene/kubernetes-course)

## Helm

* [Helm hub](https://hub.helm.sh/)
* [Charts](https://github.com/helm/charts)
* [Solr](https://github.com/guigo2k/helm-solr)
* [Portus](https://github.com/kubic-project/caasp-services/tree/master/contrib/helm-charts/portus)
* [Plugins](https://docs.helm.sh/related/)

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

* Rollout

```sh
kubectl rollout status
kubectl rollout history
kubectl rollout undo
```