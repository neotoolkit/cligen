package cligen

import "strings"

type CmdNameError struct {
	Name string
}

func (e *CmdNameError) Error() string {
	return "command name `" + e.Name + "` have spaces"
}

func Validate(cmds []Command) error {
	for _, c := range cmds {
		if strings.ContainsRune(c.Name, ' ') {
			return &CmdNameError{
				Name: c.Name,
			}
		}
	}

	return nil
}
