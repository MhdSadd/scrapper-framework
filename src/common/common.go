package common

type Step struct {
	Number int `yaml:"number"`
	Website        string            `yaml:"website"` //Can't contain protocol (http or https)
	AttributePaths map[string]string `yaml:"attributePaths"`
}

type Extractor struct {
	Name string `yaml:"name"`
	Steps []Step `yaml:"steps"`
	ActiveSteps []int `yaml:"activeSteps"`
}
type Config struct {
	Extractors  []Extractor `yaml:"extractors"`
	Transfomers interface{} `yaml:"transformers"`
}
