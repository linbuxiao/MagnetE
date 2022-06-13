package action

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/linbuxiao/magnete/model"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var YamlCommand = &cli.Command{
	Name: "yaml",
	Action: func(ctx *cli.Context) error {
		rulesByte, err := ioutil.ReadFile("./rules.json")
		if err != nil {
			return err
		}
		var rules []*model.RuleRepo
		if err = json.Unmarshal(rulesByte, &rules); err != nil {
			return err
		}
		out, err := yaml.Marshal(rules)
		if err != nil {
			return err
		}
		file, err := os.Create("rules.yaml")
		defer file.Close()
		if err != nil {
			return err
		}
		_, err = file.Write(out)
		if err != nil {
			return err
		}
		return nil
	},
}
