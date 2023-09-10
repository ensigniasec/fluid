package yaml

import (
	"bytes"

	"sigs.k8s.io/yaml"
)

func MarshalYAML(crds []any) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	for i, crd := range crds {
		o, err := yaml.Marshal(crd)
		if err != nil {
			return nil, err
		}

		_, err = buf.Write(o)
		if err != nil {
			return nil, err
		}

		if i != len(crds)-1 {
			_, err = buf.Write([]byte("---\n"))
			if err != nil {
				return nil, err
			}
		}
	}

	return buf.Bytes(), nil
}
