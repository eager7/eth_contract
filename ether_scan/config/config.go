package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var Cfg *Config

type EthOpt struct {
	Address    string `json:"address"`
	ApiAddress string `json:"api_address"`
}

type DbOpt struct {
	Address     string `json:"address"`
	User        string `json:"user"`
	Password    string `json:"password"`
	DbName      string `json:"db_name"`
	MaxOpenConn int    `json:"max_open_conn"`
	MaxIdleConn int    `json:"max_idle_conn"`
}

type Config struct {
	Path   string
	Worker int    `json:"worker"`
	EthOpt EthOpt `json:"eth_opt"`
	DbOpt  DbOpt  `json:"db_opt"`
}

func Initialize() (*Config, error) {
	var BaseDir string
	if flag.Lookup("test.v") == nil {
		BaseDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
		BaseDir = strings.Replace(BaseDir, "\\", "/", -1)
	} else {
		BaseDir = "/tmp"
	}
	Cfg = &Config{Path: BaseDir}
	if err := Cfg.ReadConfigFile(BaseDir + "/eth_contract.toml"); err != nil {
		return nil, err
	}

	Cfg.readFileConfig()
	fmt.Println("Initialize config module success:", Cfg)
	return Cfg, nil
}

//读取配置文件，如果文件不存在，则创建一个文件
func (c *Config) ReadConfigFile(file string) error {
	viper.SetConfigFile(file)

	c.initDefaultConfig()

	if _, err := os.Stat(viper.ConfigFileUsed()); os.IsNotExist(err) {
		if _, err := os.Create(viper.ConfigFileUsed()); err != nil {
			fmt.Println("create file err:", err)
			return err
		}
		fmt.Println("write config into file")
		if err := viper.WriteConfig(); err != nil {
			return err
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

//此处设置默认配置，如果没有在此处设置，当读不到值时，会使用零值
func (c *Config) initDefaultConfig() {
	viper.SetDefault("worker", 100)

	viper.SetDefault("eth.address", "http://47.52.157.31:8585")
	viper.SetDefault("eth.api_address", "https://etherscan.io/address/")

	viper.SetDefault("db.address", "127.0.0.1:3306")
	viper.SetDefault("db.user", "root")
	viper.SetDefault("db.password", "plainchant")
	viper.SetDefault("db.db_name", "eth_contract")
	viper.SetDefault("db.max_open_conn", 2000)
	viper.SetDefault("db.max_idle_conn", 1000)
}

//如果文件中配置被更改，此处读取会覆盖默认配置参数
func (c *Config) readFileConfig() {
	c.Worker = viper.GetInt("worker")

	c.EthOpt.Address = viper.GetString("eth.address")
	c.EthOpt.ApiAddress = viper.GetString("eth.api_address")

	c.DbOpt.Address = viper.GetString("db.address")
	c.DbOpt.User = viper.GetString("db.user")
	c.DbOpt.Password = viper.GetString("db.password")
	c.DbOpt.DbName = viper.GetString("db.db_name")
	c.DbOpt.MaxOpenConn = viper.GetInt("db.max_open_conn")
	c.DbOpt.MaxIdleConn = viper.GetInt("db.max_idle_conn")
}
