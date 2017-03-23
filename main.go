package main

import (
	"log"
	"os/exec"

	docopt "github.com/docopt/docopt-go"
)

func main() {
	usage := `go-cloud-fn.

Usage:
	go-cloud-fn build <function-name>
	go-cloud-fn build <function-name> --production
  go-cloud-fn build <function-name> --development
	go-cloud-fn deploy <function-name> <stage-bucket>

Options:
  -h --help     	  Show this screen.
	-p --production   Deploy for production
	-d --development  Deploy for local development
  --version     	  Show version.`

	arguments, err := docopt.Parse(usage, nil, true, "go-cloud-fn 0.0 Pre-Alpha", false)

	if err != nil {
		panic(err)
	}

	functionName := arguments["<function-name>"].(string)

	//build := arguments["build"].(bool)
	//build := arguments["deploy"].(bool)

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
	} /*
		httpOut, err := parseTemplate(httpJs, conf)
		if err != nil {
			panic(err)
		}

		deployOut, err := parseTemplate(deploySh, conf)
		if err != nil {
			panic(err)
		}*/
	log.Println(buildOut)
	out, err := exec.Command(buildOut).Output()
	if err != nil {
		panic(err)
	}
	log.Println(out)
	/*
		err = ioutil.WriteFile("build.sh", []byte(buildOut), os.ModePerm)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(conf.TargetDir+"/index.js", []byte(httpOut), os.ModePerm)
		if err != nil {
			panic(err)
		}*/
}
