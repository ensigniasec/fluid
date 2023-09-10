package fluid

import (
	corev1 "k8s.io/api/core/v1"
)

type (
	PodBuilder struct {
		pod *corev1.PodTemplateSpec
	}
)

func NewPod() *PodBuilder {
	return &PodBuilder{
		pod: &corev1.PodTemplateSpec{},
	}
}

func (b *PodBuilder) WithName(name string) *PodBuilder {
	b.pod.Name = name
	return b
}

func (b *PodBuilder) WithNamespace(namespace string) *PodBuilder {
	b.pod.Namespace = namespace
	return b
}

func (b *PodBuilder) WithLabels(labels map[string]string) *PodBuilder {
	b.pod.Labels = labels
	return b
}

func (b *PodBuilder) WithAnnotations(annotations map[string]string) *PodBuilder {
	b.pod.Annotations = annotations
	return b
}

func (b *PodBuilder) WithContainer(container *ContainerBuilder) *PodBuilder {
	b.pod.Spec.Containers = append(b.pod.Spec.Containers, *container.container)
	return b
}

func (b *PodBuilder) WithContainers(containers ...*ContainerBuilder) *PodBuilder {
	for _, c := range containers {
		b.pod.Spec.Containers = append(b.pod.Spec.Containers, *c.container)
	}
	return b
}

// func (b *PodBuilder) WithVolume(volume *VolumeBuilder) *PodBuilder {
// 	b.pod.Spec.Volumes = append(b.pod.Spec.Volumes, *volume.volume)
// 	return b
// }
