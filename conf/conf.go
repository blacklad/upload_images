package conf

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const default_config_yaml = "./conf_example.yaml"

type OssConfig struct {
	Endpoint string `yaml:"endpoint"`
	Key      string `yaml:"key"`
	Secret   string `yaml:"secret"`
}

type Config struct {
	OssConfig   OssConfig `yaml:"oss"`
	OssBasePath string    `yaml:"ossBasePath"`
	OssDomain   string    `yaml:"ossDomain"`
}

func GetConf(path string) (*Config, error) {
	if len(path) == 0 {
		//path = default_config_yaml
		return nil, errors.New("配置文件不存在：[~/conf.yaml]");
	}
	return getConf(path)
}

//并转换成conf对象
func getConf(path string) (*Config, error) {
	conf := &Config{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
