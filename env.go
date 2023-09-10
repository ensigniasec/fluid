package fluid

import corev1 "k8s.io/api/core/v1"

type EnvBuilder struct {
	*corev1.EnvVar
}

func NewEnvVar(name string) *EnvBuilder {
	return &EnvBuilder{EnvVar: &corev1.EnvVar{Name: name}}
}

func (ev *EnvBuilder) WithValue(value string) *EnvBuilder {
	ev.Value = value
	ev.ValueFrom = nil
	return ev
}

func (ev *EnvBuilder) WithValueFrom(src corev1.EnvVarSource) *EnvBuilder {
	ev.Value = ""
	ev.ValueFrom = &src
	return ev
}

func (ev *EnvBuilder) WithValueFromSecret(key string) *EnvBuilder {
	ev.Value = ""
	ev.ValueFrom = &corev1.EnvVarSource{
		SecretKeyRef: &corev1.SecretKeySelector{
			Key: key,
		},
	}
	return ev
}
