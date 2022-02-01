package cligen

type API struct {
	Name     string    `json:"name" yaml:"name"`
	Commands []Command `json:"commands" yaml:"commands"`
}

type Command struct {
	Name        string `json:"name" yaml:"name"`
	Alias       string `json:"alias" yaml:"alias"`
	Description string `json:"description" yaml:"description"`
}
