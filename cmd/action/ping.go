package action

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/linbuxiao/magnete/gen/collector"
	"github.com/linbuxiao/magnete/model"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var PingCommand = &cli.Command{
	Name:   "ping",
	Action: ping,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "target",
			Aliases: []string{"t"},
		},
		&cli.StringFlag{
			Name:    "format",
			Aliases: []string{"f"},
		},
	},
}

type pingResultRepo struct {
	Name   string `json:"name" yaml:"name"`
	Result string `json:"result" yaml:"result"`
}

func (p *pingResultRepo) Format(f string) {
	out, err := formatOutput(f, p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func ping(ctx *cli.Context) error {
	allFlag := false
	t := ctx.String("target")
	f := ctx.String("format")
	if t == "" {
		allFlag = true
	}
	rulesByte, err := ioutil.ReadFile("./rules.yaml")
	if err != nil {
		return err
	}
	var rules []*model.RuleRepo
	err = yaml.Unmarshal(rulesByte, &rules)
	if err != nil {
		return nil
	}
	var wg sync.WaitGroup
	if allFlag {
		for _, v := range rules {
			wg.Add(1)
			go handlePingStatus(&wg, v.Name, f)
		}
	} else {
		handlePingStatus(&wg, t, f)
	}
	wg.Wait()
	return nil
}

func pingTarget(name string) error {
	query := gofakeit.Animal()
	_, err := collector.Get(name, query)
	return err
}

func handlePingStatus(wg *sync.WaitGroup, name string, format string) {
	err := pingTarget(name)
	res := &pingResultRepo{Name: name}
	if err != nil {
		res.Result = err.Error()
	} else {
		res.Result = "Pass"
	}
	res.Format(format)
	wg.Done()
}
