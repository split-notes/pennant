package submodule_config

type Submodule struct {
	Language string `json:"LANGUAGE"`
	Transport string `json:"TRANSPORT"`
	ProjectPath string `json:"-"`
	ProjectName string `json:"-"`
}

