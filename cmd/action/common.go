package action

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

const DefaultFormat = "yaml"

func formatOutput(format string, content interface{}) ([]byte, error) {
	if format == "" {
		format = DefaultFormat
	}
	switch format {
	case "json":
		return json.MarshalIndent(content, "", "  ")
	case "yaml":
		return yaml.Marshal(content)
	}
	return nil, nil
}
