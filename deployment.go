package fluid

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	DeploymentBuilder struct {
		final *appsv1.Deployment
	}
)

func NewDeploymentBuilder() *DeploymentBuilder {
	return &DeploymentBuilder{
		final: &appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Deployment",
				APIVersion: "apps/v1",
			},
		},
	}
}

func (b *DeploymentBuilder) Name(n string) *DeploymentBuilder {
	b.final.Name = n
	return b
}
