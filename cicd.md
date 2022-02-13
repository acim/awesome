# CI/CD

- [Codebeat - get instant feedback on your code](https://codebeat.co/)
- [Coveralls - code test coverage](https://coveralls.io/)
- [Codecov - code test coverage](https://about.codecov.io/)

## Drone

- [Build and deploy applications using Drone CI, Docker and Ansible](https://blog.maqpie.com/2017/03/21/build-and-deploy-applications-using-drone-ci-docker-and-ansible/)
- [Continuous Delivery with Drone CI](https://medium.com/@sergey.kolodyazhnyy/continuous-delivery-with-drone-ci-3a3fea5aa83)
- [The ultimate Drone CI caching guide](https://laszlo.cloud/the-ultimate-droneci-caching-guide)
- [Drone Environment Variables](https://laszlo.cloud/drone-environment-variables-three-tips)
- [Using Drone CI/CD to build, test your images and deploy to Kubernetes with Helm](https://vitobotta.com/2019/10/09/ci-cd-with-drone-for-deployment-to-kubernetes-with-helm/)

### Conditional build if directory modified

- [Drone Config Plugin - Changeset Conditional](https://github.com/microadam/drone-config-changeset-conditional)
- [Drone Config Plugin - Pipeline](https://github.com/microadam/drone-config-plugin-pipeline)
- [boilr-config - Drone Official](https://github.com/drone/boilr-config)
- [Conversion plugin - Drone Official](https://github.com/drone/drone-go/tree/master/plugin/converter)

### Use BuildKit

```yaml
steps:
  - name: docker-buildkit
    pull: always
    image: plugins/docker:linux-amd64
    environment:
      DOCKER_BUILDKIT: 1
    settings:
      daemon_off: false
```

## Tekton Pipelines

### [Conditional build if directory modified](https://github.com/tektoncd/pipeline/issues/1922)

```yaml
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
      - conditionRef: 'is-python-runtime'
  - name: build-frontend
    conditions:
      - conditionRef: 'is-nodejs-runtime'
```

## GitHub Actions

- [metadata-action - metadata from Git reference and events useful for Build Push action to tag and label Docker images](https://github.com/docker/metadata-action)
- [Creating PostgreSQL service containers](https://docs.github.com/en/actions/using-containerized-services/creating-postgresql-service-containers)
- [GitHub Actions by Example](https://www.actionsbyexample.com/)
- [How We Write GitHub Actions in Go](https://full-stack.blend.com/how-we-write-github-actions-in-go.html)

## Infrastructure as code

- [Crossplane - enables platform teams to assemble infrastructure from multiple vendors](https://github.com/crossplane/crossplane)
