package cligen_test

import (
	"github.com/go-dummy/cligen"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {
	got, err := cligen.Parse("./testdata/.cligen.yml")
	if err != nil {
		t.Fatal(err)
	}
	want := cligen.API{
		Name: "cli",
		Commands: []cligen.Command{
			{
				Name:        "first command",
				Alias:       "fc",
				Description: "first command description",
			},
			{
				Name:        "second command",
				Alias:       "sc",
				Description: "second command description",
			},
		},
	}
	require.Equal(t, want, got)
}
