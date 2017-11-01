package core

import (
	"os"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Environment struct {
	Env string `json:"env"`
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

var instance *Config

func InitConfig(file string) {
	f, err := os.Open(file)
	if err != nil {
		logs.Critical(err)
		return
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		logs.Critical(err)
		return
	}

	var environment Environment
	err = json.Unmarshal(data, &environment)
	if err != nil {
		//panic(err)
		logs.Critical(err)
		return
	}

	f, err = os.Open("./config/wechat_" + environment.Env + ".yml")
	if err != nil {
		logs.Critical(err)
		return
	}

	defer f.Close()

	data, err = ioutil.ReadAll(f)
	if err != nil {
		logs.Error(err)
		return
	}

	var config Config
	yaml.Unmarshal([]byte(data), &config)
	instance = &config
}

func Instance() *Config {
	if instance == nil {
		InitConfig(".env")
	}

	return instance
}
