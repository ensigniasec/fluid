package fluid

import corev1 "k8s.io/api/core/v1"

type ContainerPortBuilder struct {
	*corev1.ContainerPort
}

func NewPort(port int32) *ContainerPortBuilder {
	return &ContainerPortBuilder{ContainerPort: &corev1.ContainerPort{ContainerPort: port}}
}

func NewTCPPort(port int32) *ContainerPortBuilder {
	return &ContainerPortBuilder{ContainerPort: &corev1.ContainerPort{ContainerPort: port, Protocol: corev1.ProtocolTCP}}
}

func (cp *ContainerPortBuilder) WithName(name string) *ContainerPortBuilder {
	cp.Name = name
	return cp
}

func (cp *ContainerPortBuilder) WithProtocol(protocol corev1.Protocol) *ContainerPortBuilder {
	cp.Protocol = protocol
	return cp
}

func (cp *ContainerPortBuilder) UseTCP() *ContainerPortBuilder {
	cp.Protocol = corev1.ProtocolTCP
	return cp
}

func (cp *ContainerPortBuilder) UseUDP() *ContainerPortBuilder {
	cp.Protocol = corev1.ProtocolUDP
	return cp
}

func (cp *ContainerPortBuilder) UseSCTP() *ContainerPortBuilder {
	cp.Protocol = corev1.ProtocolSCTP
	return cp
}

func (cp *ContainerPortBuilder) WithHostPort(port int32) *ContainerPortBuilder {
	cp.HostPort = port
	return cp
}

func (cp *ContainerPortBuilder) WithHostIP(ip string) *ContainerPortBuilder {
	cp.HostIP = ip
	return cp
}
