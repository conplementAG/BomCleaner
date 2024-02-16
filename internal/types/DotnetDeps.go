package types

type DotnetDeps struct {
	RuntimeTarget      DotnetDepsRuntime                     `json:"runtimeTarget,omitempty"`
	CompilationOptions map[string]interface{}                `json:"compilationOptions"`
	Targets            map[string]map[string]DotnetSubTarget `json:"targets,omitempty"`
	Libraries          map[string]DotnetDepsLibrary          `json:"libraries,omitempty"`
}

type DotnetSubTarget struct {
	Dependencies interface{}            `json:"dependencies,omitempty"`
	Runtime      map[string]interface{} `json:"runtime,omitempty"`
	Resources    interface{}            `json:"resources,omitempty"`
}

type DotnetDepsRuntime struct {
	Name      string `json:"name"`
	Signature string `json:"signature"`
}

type DotnetDepsLibrary struct {
	Type     string `json:"type"`
	Path     string `json:"path"`
	Sha512   string `json:"sha512"`
	HashPath string `json:"hashPath"`
}
