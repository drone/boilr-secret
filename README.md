# boilr-secret

This is a [boilr template](https://github.com/tmrts/boilr) for creating a secret extension. Use a secret extension to expose secrets from external sources. For example, expose secrets from a Vault server. Get started by cloning the project and installing the template:

```console
$ cd boilr-secret
$ boilr template save . drone-secret -f
```

create a project in directory my-secrets:

```console
$ boilr template use drone-secret my-secrets
```

enter the docker registry name for this project:

```text
[?] Please choose a value for "DockerRepository" [default: "foo/bar"]:
```

enter the go module name:

```text
[?] Please choose a value for "GoModule" [default: "github.com/foo/bar":
```
