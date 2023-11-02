package conf

import (
	"os"

	"gopkg.in/yaml.v2"
)

type MysqlConfig struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	Database        string `yaml:"database"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	MaxIdleConn     int    `yaml:"max_idle_conn"`
	MaxOpenConn     int    `yaml:"max_open_conn"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

type LogConfig struct {
	FilePath string `yaml:"file_path"`
	Level    string `yaml:"level"`
}

type Configure struct {
	Mysql MysqlConfig `yaml:"mysql"`
	Log   LogConfig   `yaml:"log"`
}

var Conf *Configure

func Init(cfgPath string) error {
	content, err := os.ReadFile(cfgPath)
	if err != nil {
		return err
	}
	Conf = &Configure{}
	err = yaml.Unmarshal(content, Conf)
	if err != nil {
		return err
	}

	return nil
}
