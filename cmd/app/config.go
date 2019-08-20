package app

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/glog"
	yaml "gopkg.in/yaml.v2"
)

func loadConfig(configPath string) (config *AppConfig, err error) {
	config = &AppConfig{}
	if configPath == "" {
		return nil, fmt.Errorf("configPath is empty")
	}

	if err = LoadYamlFile(configPath, config); err != nil {
		return nil, fmt.Errorf("load api config file :%s error:%s", configPath, err)
	}

	if err := checkConfig(config); err != nil {
		return nil, err
	}

	return
}

func LoadYamlFile(filePath string, ac interface{}) (err error) {
	fileContext, err := ioutil.ReadFile(filePath)
	if err != nil {
		glog.Errorf("loadConfig context from:%s error:%s", filePath, err.Error())
		return err
	}

	err = yaml.Unmarshal(fileContext, ac)
	if err != nil {
		glog.Errorf("parse yaml config from:%s error:%s", filePath, err.Error())
		return err
	}

	return
}

func checkConfig(config *AppConfig) error {
	if config.Port == 0 {
		glog.Errorf("app config Port is empty!")
		return fmt.Errorf("app config Port is empty")
	}

	return nil
}
