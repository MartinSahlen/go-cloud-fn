package main

import (
	"io/ioutil"
	"os"

	docopt "github.com/docopt/docopt-go"
)

func main() {
	usage := `go-cloud-fn.

Usage:
  go-cloud-fn init <function-name>

Options:
  -h --help     	Show this screen.
  --version     	Show version.`

	arguments, err := docopt.Parse(usage, nil, true, "go-cloud-fn 0.0 Pre-Alpha", false)

	if err != nil {
		panic(err)
	}

	functionName := arguments["<function-name>"].(string)

	conf := config{
		FunctionName:   functionName,
		TargetDir:      "target",
		TriggerHTTP:    true,
		StageBucket:    "your-bucket",
		ExecutableName: functionName,
		Production:     true,
	}

	buildOut, err := parseTemplate(buildSh, conf)
	if err != nil {
		panic(err)
	}

	httpOut, err := parseTemplate(httpJs, conf)
	if err != nil {
		panic(err)
	}
	deployOut, err := parseTemplate(deploySh, conf)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("deploy.sh", []byte(deployOut), os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("build.sh", []byte(buildOut), os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("index.js", []byte(httpOut), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
