package helm

import (
	"github.com/SUSE/fissile/model"
)

// MakeValues returns a ConfigObject with all default values for the Helm chart
func MakeValues(rolesManifest *model.RoleManifest, defaults map[string]string) (ConfigObject, error) {
	env := ConfigObject{}
	for name, cv := range model.MakeMapOfVariables(rolesManifest) {
		if !cv.Secret || cv.Generator == nil || cv.Generator.Type != model.GeneratorTypePassword {
			ok, value := cv.Value(defaults)
			if !ok {
				value = "~"
			}
			env.Add(name, NewConfigScalarWithComment(value, cv.Description))
		}
	}

	sc := ConfigObject{}
	sc.Add("persistent", NewConfigScalar("persistent"))
	sc.Add("shared", NewConfigScalar("shared"))

	kube := ConfigObject{}
	kube.Add("external_ip", NewConfigScalar("192.168.77.77"))
	kube.Add("storage_class", &sc)

	values := ConfigObject{}
	values.Add("env", &env)
	values.Add("kube", &kube)

	return values, nil
}
