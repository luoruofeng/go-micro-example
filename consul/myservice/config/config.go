package config

import (
	"fmt"

	consul "github.com/go-micro/plugins/v4/config/source/consul"

	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/file"
	"go-micro.dev/v4/util/log"
)

// defaultMysqlConfig mysql 配置
type MysqlConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Dbname       string `json:"dbname"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns string `json:"max_idle_conns"`
}

type LogConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:filename"`
	MaxSize    int    `yaml:max_size"`
	MaxAge     int    `yaml:max_age"`
	MaxBackips int    `yaml:max_backips"`
}

var CONFIG_PREFIX = "micro/config"
var CONFIG_FILE_PATH = "config.json"

var (
	Config   config.Config
	MysqlCnf MysqlConfig
	LogCnf   LogConfig
)

func Init(consul_addr string) error {
	fmt.Println("init consul config... addr:" + consul_addr)

	// Create a new Consul source
	consulSource := consul.NewSource(
		consul.WithAddress(consul_addr),
		consul.WithPrefix(CONFIG_PREFIX),
		//是否移除前缀，这里是设置为true，表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)

	// Create a config instance
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Println("Error consul config:", err)
		return err
	}
	conf.Load(consulSource)

	//把consul中设置好了的值设置到全局变量中
	//这个mysql是提前在consul的控制面板中配置好了的值。key是micro/config/mysql
	conf.Get("mysql").Scan(&MysqlCnf)
	conf.Get("log").Scan(&LogCnf)

	// Add a config file source
	// LoadConfig(CONFIG_FILE_PATH)

	// 侦听文件变动
	watcher, err := conf.Watch()
	if err != nil {
		log.Fatalf("[Init] 侦听consul配置中心 watcher异常，%s", err)
		panic(err)
	}
	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatalf("侦听consul配置中心 异常， %s", err)
				return
			}
			if err = conf.Load(consulSource); err != nil {
				panic(err)
			}
			log.Logf("consul配置中心有变化，%s", string(v.Bytes()))
		}
	}()

	Config = conf
	return nil
}

func LoadConfig(configPath string) error {
	// 加载配置文件
	if err := Config.Load(file.NewSource(
		file.WithPath(configPath),
	)); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ConfigGet(key string) string {
	// Get the value from config
	return Config.Get(key).String("")
}

func ConfigSet(key string, val interface{}) {
	Config.Set(val, key)
}
