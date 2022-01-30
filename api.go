package cligen

type API struct {
	Name     string
	Commands []Command
}

type Command struct {
	Name        string
	Description string
}
