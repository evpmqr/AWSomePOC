package resources

// namespace: my-awsome-project
// params: some_param: some_param_string
// // Define function only, role gets created by default
// functions:
// 	- name: hello-world
// // resource name ends up being my-awesome-project-hello-word
// 	handler: some_name
// 	description: some descr
// 	codeuri: some/path/zip.zip
// 	runtime: go1.x
// 	timeout 5 // Default is 5
// 	events:
// 		eventname: name
// 		type: type
// 		properties: // depends on type    policies:  - PolicyDoc

type Lambda struct {
	Name         string   `yaml:"name"`
	ResourceName string   `yaml:"resourcename"`
	Handler      string   `yaml:"handler"`
	Description  string   `yaml:"description"`
	CodeURI      string   `yaml:"codeuri"`
	Runtime      string   `yaml:"runtime"`
	Timeout      int      `yaml:"timeout"`
	Event        Event    `yaml:"event"`
	Policies     []Policy `yaml:"policies"`
}

type Event struct {
	EventName  string     `yaml:"eventname"`
	Type       string     `yaml:"type"`
	Properties Properties `yaml:"properties"`
}

type Properties struct{}
