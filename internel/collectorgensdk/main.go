package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/linbuxiao/magnete/model"
	"gopkg.in/yaml.v3"
)

func main() {
	specFile := os.Args[1]
	tmpFile := os.Args[2]
	destFileRoot := os.Args[3]
	tmp := template.New("collector")
	cotmplByte, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		log.Fatal(err)
	}
	cotmpl, err := tmp.Parse(string(cotmplByte))
	if err != nil {
		log.Fatal("parse template err: ", err)
	}
	rulesByte, err := ioutil.ReadFile(specFile)
	if err != nil {
		log.Fatal(err)
	}
	var rules []*model.RuleRepo
	if err = yaml.Unmarshal(rulesByte, &rules); err != nil {
		log.Fatal(err)
	}
	destFileDir := fmt.Sprintf("%s/collector", destFileRoot)
	if !dirExist(destFileDir) {
		if err := os.Mkdir(destFileDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "mkdir '%s' for writing failed: %+v\n", destFileDir, err)
			os.Exit(1)
		}
	}
	destFileName := fmt.Sprintf("%s/rule.go", destFileDir)
	file, err := os.Create(destFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open '%s' for writing failed: %+v\n", destFileName, err)
		os.Exit(1)
	}
	w := bufio.NewWriter(file)
	if err = cotmpl.Execute(w, rules); err != nil {
		log.Fatal(err)
	}
	w.Flush()
	file.Close()
}

func dirExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
