package fluid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildingContainer(t *testing.T) {
	a := assert.New(t)

	b := NewContainer().
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
			NewSecurityContext().WithoutPrivilege().Privileged(false),
		)
	c := b.Build()
	a.Equal("nginx", c.Name)
	a.Equal("nginx:latest", c.Image)
	a.NotNil(c.LivenessProbe)

	result, err := b.YAML()
	a.NoError(err)
	t.Log("\n", string(result))
}
