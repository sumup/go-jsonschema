package generator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIdentifierFromFileName(t *testing.T) {
	g := Generator{
		config: Config{
			YAMLExtensions:    []string{"yaml", "yml"},
			ResolveExtensions: []string{"yaml", "yml"},
		},
	}

	tc := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "Remove extension",
			filename: "objectSchemaDefinition.yaml",
			want:     "ObjectSchemaDefinition",
		},
	}

	for _, tt := range tc {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, g.identifierFromFileName(tt.filename))
		})
	}
}
