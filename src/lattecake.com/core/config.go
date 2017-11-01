package core

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Environment struct {
	App        string `json:"app"`
	Env        string `json:"env"`
	ConfigPath string `json:"config_path"`
}

type Config struct {
	Env string `yaml:"env"`
	Wechat struct {
		AppId        string `yaml:"app_id"`
		AppSecret    string `yaml:"app_secret"`
		Token        string `yaml:"token"`
		EncodeAESKey string `yaml:"encode_AES_key"`
	}
	Mysql struct {
		Host     string `yaml:"host"`
		Port     int `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Charset  string `yaml:"utf8"`
		Database string `yaml:"lattecake-wechat"`
	}
	LogFileName string `yaml:"log_file_name"`
}

var instance *interface{}

func InitConfig(file string, config interface{}) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
		return
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
		return
	}

	var environment Environment
	err = json.Unmarshal(data, &environment)
	if err != nil {
		//panic(err)
		panic(err)
		return
	}

	f, err = os.Open("../" + environment.ConfigPath + environment.App + "_" + environment.Env + ".yml")
	if err != nil {
		panic(err)
		return
	}

	defer f.Close()

	data, err = ioutil.ReadAll(f)
	if err != nil {
		panic(err)
		return
	}

	//var config Config
	yaml.Unmarshal([]byte(data), &config)
	instance = &config
}

func Instance(envpath string, config interface{}) *interface{} {
	if instance == nil {
		InitConfig(envpath, config)
	}

	return instance
}
