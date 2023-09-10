package fluid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildingDeployment(t *testing.T) {
	a := assert.New(t)

	result, err := NewDeployment().
		WithName("nginx").
		WithPodTemplate(
			NewPod().
				WithName("nginx").
				WithContainers(
					NewContainer().
						WithName("nginx").
						WithImage("nginx:latest").
						WithSecurityContext(
							NewSecurityContext().WithoutPrivilege(),
						),
				),
		).
		YAML()
	a.NoError(err)
	t.Log("\n", string(result))
}
