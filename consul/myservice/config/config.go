package config

import (
	"fmt"

	consul "github.com/go-micro/plugins/v4/config/source/consul"

	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/file"
	"go-micro.dev/v4/util/log"
)

type mysqlConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Dbname       string `json:"dbname"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns string `json:"max_idle_conns"`
}

type prometheusConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

var CONFIG_PREFIX = "micro/config"
var CONFIG_FILE_PATH = "config/prometheus.json"

var (
	cnf           config.Config
	MysqlCnf      mysqlConfig
	PrometheusCnf prometheusConfig
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

	conf, err := config.NewConfig()
	if err != nil {
		fmt.Println("Error consul config:", err)
		return err
	}

	//加载source前需要在consul创建micro/config/mysql
	err = conf.Init(config.WithSource(consulSource))

	if err != nil {
		fmt.Println("Error consul init:", err)
		return err
	}

	//把consul中设置好了的值设置到全局变量中
	//提前在consul的控制面板中配置好了的值。key是micro/config/mysql
	conf.Get("mysql").Scan(&MysqlCnf)

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

	cnf = conf

	// 另一种把数据加载到config中的方法，是不通过consul已经配置好的kv，而是直接把一个配置文件加载到config中
	LoadConfig(CONFIG_FILE_PATH)
	conf.Get("prometheus").Scan(&PrometheusCnf)

	//可以查看目前有config中的kv
	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	fmt.Println(ConfigMap())
	// }()

	return nil
}

func LoadConfig(configPath string) error {
	// 加载配置文件
	fileSource := file.NewSource(file.WithPath(configPath))
	if err := cnf.Load(fileSource); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ConfigGet(key string) string {
	// Get the value from config
	return cnf.Get(key).String("")
}

func ConfigMap() map[string]interface{} {
	return cnf.Map()
}

func ConfigSet(key string, val interface{}) {
	cnf.Set(val, key)
}
