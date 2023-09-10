package fluid

import (
	"github.com/ensigniasec/fluid/pkg/yaml"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	DeploymentBuilder struct {
		deploy *appsv1.Deployment
	}
)

func NewDeployment() *DeploymentBuilder {
	return &DeploymentBuilder{
		deploy: &appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Deployment",
				APIVersion: "apps/v1",
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: ptr(int32(1)),
				Strategy: appsv1.DeploymentStrategy{
					Type: appsv1.RollingUpdateDeploymentStrategyType,
				},
				Template: corev1.PodTemplateSpec{},
			},
		},
	}
}

func (b *DeploymentBuilder) WithName(n string) *DeploymentBuilder {
	b.deploy.Name = n
	return b
}

func (b *DeploymentBuilder) WithPodTemplate(pod *PodBuilder) *DeploymentBuilder {
	b.deploy.Spec.Template = *pod.pod
	return b
}

func (b *DeploymentBuilder) Build() *appsv1.Deployment {
	return b.deploy
}

func (b *DeploymentBuilder) YAML() ([]byte, error) {
	return yaml.MarshalYAML([]any{b.deploy})
}
