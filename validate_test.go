package cligen_test

import (
	"github.com/neotoolkit/cligen"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCmdNameError(t *testing.T) {
	got := &cligen.CmdNameError{
		Name: "test command",
	}

	require.Equal(t, got.Error(), "command name `test command` have spaces")
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name string
		cmds []cligen.Command
		err  error
	}{
		{
			name: "",
			cmds: nil,
			err:  nil,
		},
		{
			name: "",
			cmds: []cligen.Command{
				{
					Name:        "test",
					Alias:       "test",
					Description: "test",
				},
			},
			err: nil,
		},
		{
			name: "",
			cmds: []cligen.Command{
				{
					Name:        "test command",
					Alias:       "test",
					Description: "test",
				},
			},
			err: &cligen.CmdNameError{
				Name: "test command",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := cligen.Validate(tc.cmds)
			if err != nil {
				require.EqualError(t, err, tc.err.Error())
			}
		})
	}
}
