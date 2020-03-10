package resources

// Lambda is struct representation of a lambda function declaration
type Lambda struct {
	Name         string            `yaml:"name"`
	ResourceName string            `yaml:"resourcename"`
	Handler      string            `yaml:"handler"`
	Description  string            `yaml:"description"`
	CodeURI      string            `yaml:"codeuri"`
	Runtime      string            `yaml:"runtime"`
	Timeout      int               `yaml:"timeout"`
	Event        Event             `yaml:"event"`
	Environment  map[string]string `yaml:"environment"`
	Policies     []Policy          `yaml:"policies"`
}

// Event is struct representation of event trigger type
type Event struct {
	EventName  string     `yaml:"eventname"`
	Type       string     `yaml:"type"`
	Properties Properties `yaml:"properties"`
}

// Properties is struct representation of an Event property
// Needs to be validated - ex: Can't have API Properties and SQS Properties
type Properties struct {
	// API Properties
	Path              string   `yaml:"path"`
	Method            string   `yaml:"method"`
	RestAPIID         string   `yaml:"restapiid"`
	RequestParamaters []string `yaml:"requestparameters"`

	// SQS Properties
	Queue     string `yaml:"queue"`
	BatchSize int    `yaml:"batchsize"`
	Enabled   bool   `yaml:"enabled"`

	//TODO Add more Event types
}

func ValidateLambda(lambda Lambda) {
	// TODO Do some validation, maybe look into JSON Validation Schemas
}
