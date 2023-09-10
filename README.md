# fluid

Fluid is a Kubernetes resource builder using Go instead of YAML. It is designed to be consumed by CLI tooling that needs to emit valid Kubeernetes YAML or talk directory to the Kube APIs to provision resources, without writing or managing YAML or Helm.

## Usage

```go
NewContainer().
  WithName("nginx").
  WithImage("nginx:latest").
  WithLivenessProbe(
    NewHTTPProbe("/healthz").WithPort(8081),
  ).
  WithEnv(
    NewEnvVar("KEY_1").WithValue("val1"),
    NewEnvVar("KEY_2").WithValue("val2"),
    NewEnvVar("KEY_3").WithValueFromSecret("secret-1"),
  ).
  WithPorts(
    NewTCPPort(8080).WithHostPort(80).WithName("http"),
  ).
  WithSecurityContext(
    NewSecurityContext().WithoutPrivilege(),
  ).
  YAML()
```

Results in:

```yaml
env:
  - name: KEY_1
    value: val1
  - name: KEY_2
    value: val2
  - name: KEY_3
    valueFrom:
      secretKeyRef:
        key: secret-1
image: nginx:latest
livenessProbe:
  httpGet:
    path: /healthz
    port: 8081
name: nginx
ports:
  - containerPort: 8080
    hostPort: 80
    name: http
    protocol: TCP
resources: {}
securityContext:
  allowPrivilegeEscalation: false
  capabilities:
    drop:
      - ALL
  privileged: false
  runAsNonRoot: true
```
