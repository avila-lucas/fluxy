package models

type Project struct {
	Annotations string `yaml:"annotations"`
	FluxPatch   string `yaml:"fluxPatch"`
}

type Config struct {
	Projects map[string]Project `yaml:"projects"`
}

type FluxPatch struct {
	Spec Spec `yaml:"spec"`
}

type Spec struct {
	Template Template `yaml:"template"`
}

type Template struct {
	Spec TemplateSpec `yaml:"spec"`
}

type TemplateSpec struct {
	Containers []Containers `yaml:"containers"`
}

type Containers struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
}
