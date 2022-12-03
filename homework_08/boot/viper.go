package boot

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"homework_08/app/global"
	"os"
)

const (
	configEnv  = "ZHIHU_CONFIG_PATH"
	configFile = "manifest/config/config.yaml"
)

func ViperSetup(path ...string) {
	var configPath string

	//获取配置文件位置
	//优先级：参数>命令行>环境变量>默认值
	configPath = path[0]
	if configPath == "" {
		//命令行
		flag.StringVar(&configPath, "c", "", "use -c to set config path")
		flag.Parse()

		if configPath == "" {

			if configPath = os.Getenv(configEnv); configPath == "" {
				//环境变量
			} else {
				//默认路径
				configPath = configFile
			}
		}
	}

	fmt.Printf("get config path: %s\n", configPath)

	//创建viper实例,设置参数,读取配置文件
	v := viper.New()
	v.SetConfigFile(configPath) //配置信息位置
	v.SetConfigType("yaml")     //配置信息的存储格式
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("get config file failed, err: %v", err))
	}

	//将配置信息转换成结构体，放在global包里
	if err = v.Unmarshal(global.Config); err != nil {
		//对配置文件进行语法分析，并将信息存到结构体中，便于go语言调用
		panic(err)
	}
}
