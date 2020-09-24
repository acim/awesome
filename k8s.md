# Kubernetes

## [ConfigMaps and Secrets as volumes](k8s-volumes.md)

## Documentation

- [Kubernetes by example](http://kubernetesbyexample.com/)
- [Istio](https://www.youtube.com/watch?v=s4qasWn_mFc)
- [Using NodeSelector to Schedule Deployments with large volumes of Stateful Data on Kubernetes](https://medium.com/@jmarhee/using-nodeselector-to-schedule-deployments-with-large-volumes-of-stateful-data-on-kubernetes-46bd3ac6059d)
- [Scalable microservices with Kubernetes](https://eu.udacity.com/course/scalable-microservices-with-kubernetes--ud615)
- [Kubernetes cheat sheet](https://codefresh.io/kubernetes-tutorial/kubernetes-cheat-sheet/)
- [SkyDNS](https://github.com/skynetservices/skydns)
- [Single master cluster with kubeadm](https://kubernetes.io/docs/setup/independent/create-cluster-kubeadm/)
- [HA Kubernetes cluster on bare metal](https://github.com/salmanb/Kubernetes-HA-on-baremetal)
- [Encryption of Kubernetes persistent local volumes](https://medium.com/@dfrnascimento/encryption-of-kubernetes-persistent-local-volumes-70da62e0ed68)
- [Static provisioner of local volumes](https://github.com/kubernetes-sigs/sig-storage-local-static-provisioner)
- [OpenEBS - Creating and attaching disk on GKE node](https://docs.openebs.io/docs/next/setupstoragepools.html#creating-and-attaching-a-disk-on-gke-node)
- [Kubespray - Deploy a production ready Kubernetes cluster](https://github.com/kubernetes-sigs/kubespray)
- [Ingress basic auth](https://kubernetes.github.io/ingress-nginx/examples/auth/basic/)
- [Generate RBAC policies based on Kubernetes audit logs](https://github.com/liggitt/audit2rbac)
- [Awesome Kubernetes](https://github.com/ramitsurana/awesome-kubernetes)
- [stern - tail multiple pods on Kubernetes and multiple containers within the pod](https://github.com/wercker/stern)
- [Kubernetes SIG](https://github.com/kubernetes-sigs)
- [12 Kubernetes configuration best practices](https://www.stackrox.com/post/2019/09/12-kubernetes-configuration-best-practices/)
- [10 most important differences between OpenShift and Kubernetes](https://cloudowski.com/articles/10-differences-between-openshift-and-kubernetes)
- [Copy files from Kubernetes to S3 and back](https://medium.com/nuvo-group-tech/copy-files-and-directories-between-kubernetes-and-s3-d290ded9a5e0)
- [Kubernetes shared storage with S3 backend](https://icicimov.github.io/blog/virtualization/Kubernetes-shared-storage-with-S3-backend/)
- [Kubernetes Examples - Minimal self-contained examples of standard Kubernetes features and patterns in YAML](https://github.com/ContainerSolutions/kubernetes-examples)
- [AWS Controllers for Kubernetes (ACK) - project enabling you to manage AWS services from Kubernetes](https://github.com/aws/aws-controllers-k8s)

## Blogs

- [Alex Ellis](https://blog.alexellis.io/)
- [IMTI - Architecting, Developing, nixCraft, DevOps, AI/ML, Blockchain](https://imti.co/)
- [Cloud Native Architect Blog](https://cloudowski.com/)
- [Igor Cicimov](https://icicimov.github.io/blog/)

## Administration

- [Adding a name or IP to the Kubernetes API Server certificate](https://blog.scottlowe.org/2019/07/30/adding-a-name-to-kubernetes-api-server-certificate/)
- [CORS on Kubernetes Ingress Nginx](https://imti.co/kubernetes-ingress-nginx-cors/)

## Operators

- [Community Operators](https://commons.openshift.org/sig/operators.html)
- [OperatorHub](https://operatorhub.io/)

## Load balancers

[MetalLB - load-balancer implementation for bare metal Kubernetes](https://github.com/metallb/metallb)
[Fabio is a fast, modern, zero-conf load balancing HTTPS and TCP router](https://github.com/fabiolb/fabio)

## Tools

- [JenkinsX - Continuous delivery for Kubernetes](https://jenkins-x.io/demos/devoxx-uk-2018/)
- [Kompose - translate docker-compose files to Kubernetes resources](https://github.com/kubernetes/kompose)
- [kuberhealthy - synthetic testing](https://github.com/Comcast/kuberhealthy)
- [Keel - Kubernetes Operator to automate Helm, DaemonSet, StatefulSet & Deployment updates](https://github.com/keel-hq/keel)
- [Keel webhooks](https://keel.sh/v1/guide/documentation.html#Triggers)
- [M3 - Uber’s large-scale metrics platform for Prometheus](https://eng.uber.com/m3/)
- [kubectx & kubens - Switch faster between clusters and namespaces](https://github.com/ahmetb/kubectx)
- [kubeone - Lifecycle management tool for highly available Kubernetes clusters](https://github.com/kubermatic/kubeone)
- [TK8 - Multi-cloud, multi-cluster Kubernetes platform installation and integration tool](https://github.com/kubernauts/tk8)
- [Kontemplate - Simple Kubernetes templater](https://github.com/tazjin/kontemplate)
- [webkubectl - kubectl in web browser](https://github.com/webkubectl/webkubectl)
- [Skbn - tool for copying files and directories between Kubernetes and cloud storage providers](https://github.com/maorfr/skbn)
- [kubenav - desktop and mobile navigator for your Kubernetes clusters](https://github.com/kubenav/kubenav)
- [Kube-Scan - Kubernetes risk assessment tool](https://github.com/octarinesec/kube-scan)
- [Permission Manager is a project that brings sanity to Kubernetes RBAC and Users management](https://github.com/sighupio/permission-manager)
- [KubeCarrier - open source system for managing applications and services across multiple Kubernetes Clusters](https://github.com/kubermatic/kubecarrier)
- [Kubeapps - web-based UI for deploying and managing applications in Kubernetes clusters](https://github.com/kubeapps/kubeapps)
- [Delete stale feature branches in your Kubernetes cluster](https://github.com/dmytrostriletskyi/stale-feature-branch-operator)
- [Kyverno - Kubernetes Native Policy Management](https://github.com/nirmata/kyverno)
- [KubiScan - tool to scan Kubernetes cluster for risky permissions](https://github.com/cyberark/KubiScan)

## CI & CD

- [Tekton Pipelines](https://github.com/tektoncd/pipeline)
- [Skaffold](https://github.com/GoogleContainerTools/skaffold)
- [How to use Knative Pipelining Component to automate an Application Build and Deployment on Kubernetes](https://itnext.io/how-to-use-knative-pipelining-component-to-automate-an-application-build-and-deployment-on-442b0b1bebf)
- [Build cloud native CI/CD build pipeline from GIT webhook](https://medium.com/@pongsatt/build-cloud-native-ci-cd-build-pipeline-from-git-webhook-9cd9a57a32e1)
- [kontinuous - Kubernetes Continuous Integration & Delivery Platform](https://github.com/AcalephStorage/kontinuous)
- [GitOps with Tekton and ArgoCD](https://github.com/RolandOrg/node-web-app)

## Backup

- [Velero (formerly Ark) - Backup and migrate Kubernetes applications and their persistent volumes](https://github.com/heptio/velero)
- [Backup Kubernetes – how and why](https://elastisys.com/2018/12/10/backup-kubernetes-how-and-why/)
- [How To Back Up and Restore Kubernetes Cluster using Ark](https://www.digitalocean.com/community/tutorials/how-to-back-up-and-restore-a-kubernetes-cluster-on-digitalocean-using-heptio-ark)
- [Backup etcd cluster](https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/#backing-up-an-etcd-cluster)

## Courses

- [Ward Viaene - On premise or cloud-agnostic Kubernetes](https://github.com/wardviaene/on-prem-or-cloud-agnostic-kubernetes)
- [Ward Viaene - Kubernetes course](https://github.com/wardviaene/kubernetes-course)

## Helm

- [Helm hub](https://hub.helm.sh/)
- [Charts](https://github.com/helm/charts)
- [Solr](https://github.com/guigo2k/helm-solr)
- [Portus](https://github.com/kubic-project/caasp-services/tree/master/contrib/helm-charts/portus)
- [Plugins](https://docs.helm.sh/related/)
- [Template developer’s guide](https://helm.sh/docs/chart_template_guide/)
- [Helmfile - deploy Kubernetes Helm Charts](https://github.com/roboll/helmfile)
- [Helmfile - chart deployment tool](https://medium.com/@devopseylife/helmfile-aka-helm-chart-deployment-tool-4e3378fad242)
- [helmfile  -  it’s like a Helm for Helm](https://medium.com/@naseem_60378/helmfile-its-like-a-helm-for-your-helm-74a908581599)
- [15+ useful Helm Charts tools](https://caylent.com/15-useful-helm-charts-tools/)
- [ChartMuseum - host your own Helm chart repository](https://github.com/helm/chartmuseum)
- [Using incubator/raw chart](https://github.com/roboll/helmfile/issues/494#issuecomment-474697430)
- [incubator/raw - chart for kubectl declarations](https://github.com/helm/charts/tree/master/incubator/raw)
- [Awesome Helm](https://github.com/cdwv/awesome-helm)
- [Helm Chart starter - better Helm chart boilerplate](https://github.com/sitewards/helm-chart)
- [Creating a better chart](https://medium.com/sitewards/deploying-on-kubernetes-2-scaffolding-6a54e5d181fb)
- [helm-docs - generates automatic documentation from helm charts into a markdown file](https://github.com/norwoodj/helm-docs)
- [Frigate - documentation generation tool for Kubernetes Helm Charts](https://medium.com/rapids-ai/introducing-frigate-a-documentation-generation-tool-for-kubernetes-1791854031a1)

## Custom Resource Definitions and Development of Kubernetes Applications

- [Accessing Kubernetes CRDs from client-go package](https://www.martin-helmich.de/en/blog/kubernetes-crd-client.html)
- [Extending Kubernetes APIs with Custom Resource Definitions](https://medium.com/velotio-perspectives/extending-kubernetes-apis-with-custom-resource-definitions-crds-139c99ed3477)
- [Unit testing](https://medium.com/@e_frogers/unit-testing-with-kubernetes-client-go-283b11aaa7db)
- [Code Generation for CustomResources](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/)
- [How to generate client codes for Kubernetes Custom Resource Definitions](https://itnext.io/how-to-generate-client-codes-for-kubernetes-custom-resource-definitions-crd-b4b9907769ba)
- [Writing Kubernetes Custom Controllers](https://medium.com/@cloudark/kubernetes-custom-controllers-b6c7d0668fdf)
- [How did that sidecar get there?](https://medium.com/dowjones/how-did-that-sidecar-get-there-4dcd73f1a0a4)
- [k8s-sidecar-injector](https://github.com/tumblr/k8s-sidecar-injector)
- [Controller to manage databases](https://github.com/kubehippie/database-controller)
- [KUDO](https://github.com/kudobuilder/kudo)
- [How to Build a Custom Kubernetes Ingress Controller in Go](http://www.doxsey.net/blog/how-to-build-a-custom-kubernetes-ingress-controller-in-go)
- [Programmatically Kubernetes port forward in Go](https://github.com/gianarb/kube-port-forward)

## Security

- [Kubernetes Security Practices You Should Follow](https://blog.sonatype.com/kubesecops-kubernetes-security-practices-you-should-follow)
- [Sealed Secrets](https://github.com/bitnami-labs/sealed-secrets)
- [Vault](https://www.vaultproject.io/)
- [Advanced Persistence Threats: The Future of Kubernetes Attacks](https://www.youtube.com/watch?v=CH7S5rE3j8w)
- [Rego policies collection](https://github.com/redhat-cop/rego-policies)
- [Kubernetes secrets store CSI driver](https://github.com/kubernetes-sigs/secrets-store-csi-driver/)

## Hetzner Cloud

- [Sysctl configuration for high performance](https://gist.github.com/techgaun/958e117b730634fa8128)
- [Installing kubernetes cluster with wireguard](https://propellered.com/posts/kubernetes/)
- [Install a Kubernetes cluster on cloud servers](https://community.hetzner.com/tutorials/install-kubernetes-cluster)
- [Load balancer Helm chart](https://github.com/exocode/helm-charts/tree/master/hetzner-failover-ip)
- [Install hcloud-cloud-controller-manager with network support](https://github.com/hetznercloud/hcloud-cloud-controller-manager/blob/master/docs/deploy_with_networks.md)
- [Creating a single control-plane cluster with kubeadm](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/)
- [Installing a pod network add-on](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/#pod-network)
- [Install Cilium](https://cilium.readthedocs.io/en/stable/gettingstarted/k8s-install-default/)
- [Kubernetes without kube-proxy](https://docs.cilium.io/en/stable/gettingstarted/kubeproxy-free/)
- [hcloud-k8s - Ansible playbook to install Kubernetes on Hetzner Cloud](https://github.com/gammpamm/hcloud-k8s)

## kubectl commands

- [Cheat sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [Example commands](kubectl.md)

### Autocomplete

- Turn on autocomplete in the current shell

```sh
source <(kubectl completion bash)
```

- Turn on autocomplete permanently

```sh
echo "source <(kubectl completion bash)" >> ~/.bashrc
```

- Make alias k and autocomplete for it

```sh
alias k=kubectl
complete -F __start_kubectl k
```
