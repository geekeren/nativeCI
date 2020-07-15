package model

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	str "strings"
)

type Pipeline struct {
	Name   string  `yaml:"name"`
	Stages []Stage `yaml:"stages"`
}

func (pipeline *Pipeline) LoadConf(configFile string) *Pipeline {

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("yamlFile.Get err  #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, pipeline)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return pipeline
}

func (pipeline *Pipeline) Run() {
	os.Mkdir("./.tmp", os.FileMode(0755))
	for stageIndex, stage := range pipeline.Stages {
		log.Printf("Stage: %s", stage.Name)

		for stepIndex, step := range stage.Steps {
			script := ".tmp/" + strconv.Itoa(stageIndex) + "_" + strconv.Itoa(stepIndex) + "_" + str.ReplaceAll(step.Name, " ", "_") + ".sh"
			err := ioutil.WriteFile(script, []byte(step.Run), 0755)
			if err != nil {
				log.Printf("Unable to write file: %v", err)
			}

			cmd := exec.Command("/bin/bash", "-c", script)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			log.Printf("Step: %s", step.Name)
			err = cmd.Run()

			if err != nil {
				log.Fatalf("Execute Shell:%s failed with error:%s", script, err.Error())
			}
		}
	}
}
