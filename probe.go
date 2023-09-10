package fluid

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type ProbeBuilder struct {
	probe *v1.Probe
}

type Probe interface {
	isProbe()
}

type HTTPProbeBuilder struct {
	action *v1.HTTPGetAction
	Probe
}

func NewProbe() *ProbeBuilder {
	return &ProbeBuilder{
		probe: &v1.Probe{},
	}
}

func NewHTTPProbe(path string) *HTTPProbeBuilder {
	builder := &ProbeBuilder{
		probe: &v1.Probe{
			ProbeHandler: v1.ProbeHandler{
				HTTPGet: &v1.HTTPGetAction{
					Path: path,
				},
			},
		},
	}

	return &HTTPProbeBuilder{
		action: builder.probe.ProbeHandler.HTTPGet,
	}
}

func (h *HTTPProbeBuilder) Path(p string) *HTTPProbeBuilder {
	h.action.Path = p
	return h
}

func (h *HTTPProbeBuilder) WithPort(p int) *HTTPProbeBuilder {
	h.action.Port = intstr.IntOrString{Type: intstr.Int, IntVal: int32(p)}
	return h
}

func (h *HTTPProbeBuilder) isProbe() {}
