# [ConfigMaps](https://kubernetes.io/docs/concepts/storage/volumes/#configmap) and [Secrets](https://kubernetes.io/docs/concepts/storage/volumes/#secret) as volumes

```yaml
spec:
    volumes:
    - name: foo-volume
    configMap:
        name: foo-config-map
        defaultMode: 0644
        items:
        - key: foo-config-map-key
          path: destination-filename.conf
        optional: false
```

```txt
RESOURCE: configMap <Object>

DESCRIPTION:
     ConfigMap represents a configMap that should populate this volume

     Adapts a ConfigMap into a volume. The contents of the target ConfigMap's
     Data field will be presented in a volume as files using the keys in the
     Data field as the file names, unless the items element is populated with
     specific mappings of keys to paths. ConfigMap volumes support ownership
     management and SELinux relabeling.

FIELDS:
   defaultMode  <integer>
     Optional: mode bits to use on created files by default. Must be a value
     between 0 and 0777. Defaults to 0644. Directories within the path are not
     affected by this setting. This might be in conflict with other options that
     affect the file mode, like fsGroup, and the result can be other mode bits
     set.

   items        <[]Object>
     If unspecified, each key-value pair in the Data field of the referenced
     ConfigMap will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the ConfigMap, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

   name         <string>
     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

   optional     <boolean>
     Specify whether the ConfigMap or its keys must be defined
```

## Difference for Secret (secretName instead of name)

```txt
   secretName   <string>
     Name of the secret in the pod's namespace to use. More info:
     https://kubernetes.io/docs/concepts/storage/volumes#secret
```

```txt
RESOURCE: items <[]Object>

DESCRIPTION:
     If unspecified, each key-value pair in the Data field of the referenced
     ConfigMap will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the ConfigMap, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

     Maps a string key to a path within a volume.

FIELDS:
   key  <string> -required-
     The key to project.

   mode <integer>
     Optional: mode bits to use on this file, must be a value between 0 and
     0777. If not specified, the volume defaultMode will be used. This might be
     in conflict with other options that affect the file mode, like fsGroup, and
     the result can be other mode bits set.

   path <string> -required-
     The relative path of the file to map the key to. May not be an absolute
     path. May not contain the path element '..'. May not start with the string
     '..'.
```

## Mount just to one file

```yaml
spec:
    containers:
    - volumeMounts:
      - name: foo-volume
        mountPath: /etc/destination-filename.conf
        subPath: source-filename.conf
        readOnly: true
```

```txt
RESOURCE: volumeMounts <[]Object>

DESCRIPTION:
     Pod volumes to mount into the container's filesystem. Cannot be updated.

     VolumeMount describes a mounting of a Volume within a container.

FIELDS:
   mountPath        <string> -required-
     Path within the container at which the volume should be mounted. Must not
     contain ':'.

   mountPropagation <string>
     mountPropagation determines how mounts are propagated from the host to
     container and the other way around. When not set, MountPropagationNone is
     used. This field is beta in 1.10.

   name             <string> -required-
     This must match the Name of a Volume.

   readOnly         <boolean>
     Mounted read-only if true, read-write otherwise (false or unspecified).
     Defaults to false.

   subPath          <string>
     Path within the volume from which the container's volume should be mounted.
     Defaults to "" (volume's root).

   subPathExpr      <string>
     Expanded path within the volume from which the container's volume should be
     mounted. Behaves similarly to SubPath but environment variable references
     $(VAR_NAME) are expanded using the container's environment. Defaults to ""
     (volume's root). SubPathExpr and SubPath are mutually exclusive. This field
     is beta in 1.15.
```

## [Projected volumes](https://kubernetes.io/docs/concepts/storage/volumes/#projected)

Maps several existing volume sources into the same directory

```yaml
spec:
    volumes:
    - name: all-in-one
    projected:
        defaultMode: 0644
        sources:
        - secret:
            name: mysecret
            items:
            - key: username
                path: my-group/my-username
        - downwardAPI:
            items:
            - path: "labels"
                fieldRef:
                fieldPath: metadata.labels
            - path: "cpu_limit"
                resourceFieldRef:
                containerName: container-test
                resource: limits.cpu
        - configMap:
            name: myconfigmap
            items:
            - key: config
                path: my-group/my-config
        - serviceAccountToken:
            path: token
```

```txt
RESOURCE: projected <Object>

DESCRIPTION:
     Items for all in one resources secrets, configmaps, and downward API

     Represents a projected volume source

FIELDS:
   defaultMode  <integer>
     Mode bits to use on created files by default. Must be a value between 0 and
     0777. Directories within the path are not affected by this setting. This
     might be in conflict with other options that affect the file mode, like
     fsGroup, and the result can be other mode bits set.

   sources      <[]Object> -required-
     list of volume projections
```

```txt
RESOURCE: sources <[]Object>

DESCRIPTION:
     list of volume projections

     Projection that may be projected along with other supported volume types

FIELDS:
   configMap            <Object>
     information about the configMap data to project

   downwardAPI          <Object>
     information about the downwardAPI data to project

   secret               <Object>
     information about the secret data to project

   serviceAccountToken  <Object>
     information about the serviceAccountToken data to project
```

```txt
RESOURCE: configMap <Object>

DESCRIPTION:
     information about the configMap data to project

     Adapts a ConfigMap into a projected volume. The contents of the target
     ConfigMap's Data field will be presented in a projected volume as files
     using the keys in the Data field as the file names, unless the items
     element is populated with specific mappings of keys to paths. Note that
     this is identical to a configmap volume source without the default mode.

FIELDS:
   items    <[]Object>
     If unspecified, each key-value pair in the Data field of the referenced
     ConfigMap will be projected into the volume as a file whose name is the key
     and content is the value. If specified, the listed keys will be projected
     into the specified paths, and unlisted keys will not be present. If a key
     is specified which is not present in the ConfigMap, the volume setup will
     error unless it is marked optional. Paths must be relative and may not
     contain the '..' path or start with '..'.

   name     <string>
     Name of the referent. More info:
     https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

   optional <boolean>
     Specify whether the ConfigMap or its keys must be defined
```
