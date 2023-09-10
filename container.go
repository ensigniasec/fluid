package fluid

import (
	"github.com/ensigniasec/fluid/pkg/yaml"
	corev1 "k8s.io/api/core/v1"
)

type (
	ContainerBuilder struct {
		container *corev1.Container
	}

	SecurityContextBuilder struct {
		sc *corev1.SecurityContext
	}
)

func NewContainer() *ContainerBuilder {
	return &ContainerBuilder{
		container: &corev1.Container{},
	}
}

// WithName
func (b *ContainerBuilder) WithName(name string) *ContainerBuilder {
	b.container.Name = name
	return b
}

// WithImage
func (b *ContainerBuilder) WithImage(image string) *ContainerBuilder {
	b.container.Image = image
	return b
}

// WithCommand
func (b *ContainerBuilder) WithCommand(command []string) *ContainerBuilder {
	b.container.Command = command
	return b
}

// WithArgs
func (b *ContainerBuilder) WithArgs(args []string) *ContainerBuilder {
	b.container.Args = args
	return b
}

// WithWorkingDir
func (b *ContainerBuilder) WithWorkingDir(workingDir string) *ContainerBuilder {
	b.container.WorkingDir = workingDir
	return b
}

// Ports
func (b *ContainerBuilder) AddPort(containerPort int32) *ContainerBuilder {
	b.container.Ports = append(b.container.Ports, corev1.ContainerPort{ContainerPort: containerPort})
	return b
}

func (b *ContainerBuilder) WithPorts(port *ContainerPortBuilder, more ...*ContainerPortBuilder) *ContainerBuilder {
	keyMap := make(map[int32]*corev1.ContainerPort)
	keyMap[port.ContainerPort.ContainerPort] = port.ContainerPort

	b.container.Ports = append(b.container.Ports, *port.ContainerPort)
	for _, p := range more {
		if _, ok := keyMap[p.ContainerPort.ContainerPort]; !ok {
			b.container.Ports = append(b.container.Ports, *p.ContainerPort)
			keyMap[port.ContainerPort.ContainerPort] = port.ContainerPort
		}
	}

	return b
}

func (b *ContainerBuilder) WithEnv(env *EnvBuilder, more ...*EnvBuilder) *ContainerBuilder {
	keyMap := make(map[string]*corev1.EnvVar)
	keyMap[env.EnvVar.Name] = env.EnvVar

	b.container.Env = append(b.container.Env, *env.EnvVar)
	for _, e := range more {
		if _, ok := keyMap[e.EnvVar.Name]; !ok {
			b.container.Env = append(b.container.Env, *e.EnvVar)
			keyMap[env.EnvVar.Name] = env.EnvVar
		}
	}
	return b
}

type EnvVarBuilder struct {
	value *corev1.EnvVar
}

func EnvVar(name string) *EnvVarBuilder {
	return &EnvVarBuilder{value: &corev1.EnvVar{Name: name}}
}

func (b *EnvVarBuilder) WithName(name string) *EnvVarBuilder {
	b.value.Name = name
	return b
}

func (b *EnvVarBuilder) WithValue(v string) *EnvVarBuilder {
	b.value.Value = v
	b.value.ValueFrom = nil
	return b
}

func (b *EnvVarBuilder) ValueFrom(src corev1.EnvVarSource) *EnvVarBuilder {
	b.value.Value = ""
	b.value.ValueFrom = &src
	return b
}

// WithValueFromSecret sets the value of the environment variable to the value of the secret key
// specified by secret key ref.
func (b *EnvVarBuilder) WithValueFromSecret(ref string) *EnvVarBuilder {
	b.value.Value = ""
	b.value.ValueFrom = &corev1.EnvVarSource{
		SecretKeyRef: &corev1.SecretKeySelector{
			Key: ref,
		},
	}
	return b
}

// WithResources
func (b *ContainerBuilder) WithResources(requests, limits corev1.ResourceList) *ContainerBuilder {
	b.container.Resources = corev1.ResourceRequirements{
		Requests: requests,
		Limits:   limits,
	}
	return b
}

// VolumeMounts
func (b *ContainerBuilder) AddVolumeMount(name, mountPath string) *ContainerBuilder {
	b.container.VolumeMounts = append(b.container.VolumeMounts, corev1.VolumeMount{Name: name, MountPath: mountPath})
	return b
}

// WithLivenessProbe
func (b *ContainerBuilder) WithLivenessProbe(handler Probe) *ContainerBuilder {
	switch v := handler.(type) {
	case *HTTPProbeBuilder:
		b.container.LivenessProbe = &corev1.Probe{ProbeHandler: corev1.ProbeHandler{
			HTTPGet: v.action,
		}}
	}
	return b
}

// WithReadinessProbe
func (b *ContainerBuilder) WithReadinessProbe(handler corev1.ProbeHandler) *ContainerBuilder {
	b.container.ReadinessProbe = &corev1.Probe{ProbeHandler: handler}
	return b
}

// WithStartupProbe
func (b *ContainerBuilder) WithStartupProbe(handler corev1.ProbeHandler) *ContainerBuilder {
	b.container.StartupProbe = &corev1.Probe{ProbeHandler: handler}
	return b
}

// WithLifecycle
func (b *ContainerBuilder) WithLifecycle(lifecycle *corev1.Lifecycle) *ContainerBuilder {
	b.container.Lifecycle = lifecycle
	return b
}

// WithTerminationMessagePath
func (b *ContainerBuilder) WithTerminationMessagePath(path string) *ContainerBuilder {
	b.container.TerminationMessagePath = path
	return b
}

// WithTerminationMessagePolicy
func (b *ContainerBuilder) WithTerminationMessagePolicy(policy corev1.TerminationMessagePolicy) *ContainerBuilder {
	b.container.TerminationMessagePolicy = policy
	return b
}

// WithImagePullPolicy
func (b *ContainerBuilder) WithImagePullPolicy(policy corev1.PullPolicy) *ContainerBuilder {
	b.container.ImagePullPolicy = policy
	return b
}

// WithSecurityContext
func (b *ContainerBuilder) WithSecurityContext(sc *SecurityContextBuilder) *ContainerBuilder {
	b.container.SecurityContext = sc.sc
	return b
}

func NewSecurityContext() *SecurityContextBuilder {
	return &SecurityContextBuilder{
		sc: &corev1.SecurityContext{},
	}
}

func (b *SecurityContextBuilder) WithoutPrivilege() *SecurityContextBuilder {
	no := false
	yes := true
	b.sc.Privileged = &no
	b.sc.AllowPrivilegeEscalation = &no
	b.sc.Capabilities = &corev1.Capabilities{
		Drop: []corev1.Capability{"ALL"},
	}
	b.sc.RunAsNonRoot = &yes
	return b
}

func (b *SecurityContextBuilder) Privileged(priv bool) *SecurityContextBuilder {
	b.sc.Privileged = &priv
	return b
}

// WithStdin
func (b *ContainerBuilder) WithStdin() *ContainerBuilder {
	b.container.Stdin = true
	return b
}

// WithoutStdin
func (b *ContainerBuilder) WithoutStdin() *ContainerBuilder {
	b.container.Stdin = false
	return b
}

// WithStdinOnce
func (b *ContainerBuilder) WithStdinOnce() *ContainerBuilder {
	b.container.StdinOnce = true
	return b
}

// WithTTY
func (b *ContainerBuilder) WithTTY() *ContainerBuilder {
	b.container.TTY = true
	return b
}

func (b *ContainerBuilder) Build() *corev1.Container {
	return b.container
}

func (b *ContainerBuilder) YAML() ([]byte, error) {
	return yaml.MarshalYAML([]any{b.container})
}
