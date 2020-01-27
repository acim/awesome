# CI/CD

## Drone

* [Build and deploy applications using Drone CI, Docker and Ansible](https://blog.maqpie.com/2017/03/21/build-and-deploy-applications-using-drone-ci-docker-and-ansible/)
* [Continuous Delivery with Drone CI](https://medium.com/@sergey.kolodyazhnyy/continuous-delivery-with-drone-ci-3a3fea5aa83)
* [The ultimate Drone CI caching guide](https://laszlo.cloud/the-ultimate-droneci-caching-guide)
* [Drone Environment Variables](https://laszlo.cloud/drone-environment-variables-three-tips)
* [Using Drone CI/CD to build, test your images and deploy to Kubernetes with Helm](https://vitobotta.com/2019/10/09/ci-cd-with-drone-for-deployment-to-kubernetes-with-helm/)
* [Painless Github releases with Drone and GoReleaser](https://dev.to/mstrsobserver/painless-github-releases-with-drone-and-goreleaser-45b7)

### [Parameters substitution](https://github.com/drone/docs/blob/master/content/usage/config/substitution.md)

### Conditional build if directory modified

* [Report files changed to limit testing scope](https://github.com/drone/drone/issues/1021)
* [Drone Config Plugin - Changeset Conditional](https://github.com/microadam/drone-config-changeset-conditional)
* [Drone Config Plugin - Pipeline](https://github.com/microadam/drone-config-plugin-pipeline)
* [boilr-config - Drone Official](https://github.com/drone/boilr-config)
* [drone-convert-starlark - Drone Official](https://github.com/drone/drone-convert-starlark)
* [Conversion plugin - Drone Official](https://github.com/drone/drone-go/tree/master/plugin/converter)

## Tekton Pipelines

### [Conditional build if directory modified](https://github.com/tektoncd/pipeline/issues/1922)

``` yaml
---
apiVersion: tekton.dev/v1alpha1
kind: Condition
metadata:
  name: is-nodejs-runtime
spec:
  params:
    - name: OW_APP_PATH
      type: string
  resources:
    - name: app-git
      type: git
  check:
    image: ubuntu
    command:
      - bash
    args:
      - -c
      - |
        echo "INFO: Identifying the language of the application source based on the file extension at:"
        echo $(params.OW_APP_PATH)
        echo "INFO: ls $(params.OW_APP_PATH):"
        ls $(params.OW_APP_PATH)
        echo "INFO: Find files with extension .js, sort displays js if one or more files found: "
        find $(params.OW_APP_PATH) -type f | sed -e 's/.*\.//' | sed -e 's/.*\///' | sort -u | grep ^js$
---
apiVersion: tekton.dev/v1alpha1
kind: Condition
metadata:
  name: is-python-runtime
spec:
  params:
    - name: OW_APP_PATH
      type: string
  resources:
    - name: app-git
      type: git
  check:
    image: ubuntu
    command:
      - bash
    args:
      - -c
      - |
        echo "INFO: Identifying the language of the application source based on the file extension at:"
        echo $(params.OW_APP_PATH)
        echo "INFO: ls $(params.OW_APP_PATH):"
        ls $(params.OW_APP_PATH)
        echo "INFO: Find files with extension .py, sort displays py if one or more files found: "
        find $(params.OW_APP_PATH) -type f | sed -e 's/.*\.//' | sed -e 's/.*\///' | sort -u | grep ^py$
---
tasks:
- name: build-backend
  conditions:
  - conditionRef: "is-python-runtime"
- name: build-frontend
  conditions:
  - conditionRef: "is-nodejs-runtime"
```
