package core

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Env              string `yaml:"env"`
	UploadImageToken string `yaml:"upload_image_token"`
	UploadImageDir   string `yaml:"upload_image_dir"`
	UploadImageUrl   string `yaml:"upload_image_url"`

	PostArticle struct {
		Security struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}
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

type Environment struct {
	App        string `json:"app"`
	Env        string `json:"env"`
	ConfigPath string `json:"config_path"`
}

var instance *Config

func InitConfig(file string) {
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

	var config Config
	yaml.Unmarshal([]byte(data), &config)
	instance = &config
}

func Instance(envpath string) *Config {
	if instance == nil {
		InitConfig(envpath)
	}

	return instance
}
