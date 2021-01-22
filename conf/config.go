/**
 * @Author: hqd
 * @Description: 配置文件
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2021/1/17 23:06
 */

package conf

import (
	"time"

	"github.com/spf13/viper"
)

var (
	ServerConfig *ServerConf
	AppConfig    *AppConf
	MysqlConfig  *MysqlConf
)

type ServerConf struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppConf struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type MysqlConf struct {
	Host         string
	User         string
	Password     string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type Config struct {
	vip *viper.Viper
}

func NewConfig() (*Config, error) {
	vip := viper.New()
	vip.SetConfigName("config")
	vip.AddConfigPath("conf/")
	vip.SetConfigType("yaml")
	err := vip.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Config{vip: vip}, nil
}

func (c *Config) LoadConf(k string, v interface{}) error {
	err := c.vip.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
