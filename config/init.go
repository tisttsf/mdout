package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/JabinGP/mdout/log"
	"github.com/JabinGP/mdout/static"
	"github.com/JabinGP/mdout/tool"
)

func readConfig() {
	if _, err := toml.DecodeFile(static.ConfigFileFullName, &Obj); err != nil {
		log.Errorln("读取配置文件失败！", err)
		panic(err)
	}
}

// ShowConfig 输出读取到的配置文件
func ShowConfig() {
	confBts, err := json.Marshal(Obj)
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("---这是你的配置文件参数---")
	log.Debugln(string(confBts))
	log.Debugln("--------------------------")
}

// InitConfigFile 初始化配置文件
func InitConfigFile() {
	if !tool.IsExists(static.ConfigFileFullName) {
		log.Infoln("配置文件 " + static.ConfigFileFullName + " 不存在，创建中...")
		confBytes := DownloadConfig("v1")

		if !tool.IsExists(static.ConfigFolderFullName) {
			log.Infoln("配置文件夹 " + static.ConfigFolderFullName + " 不存在，创建中...")
			err := os.Mkdir(static.ConfigFolderFullName, os.ModePerm)
			if err != nil {
				log.Errorf("创建文件夹 " + static.ConfigFolderFullName + " 失败!\n")
				panic(err)
			}
			log.Infoln("创建文件夹 " + static.ConfigFolderFullName + " 成功!\n")
		}

		err := ioutil.WriteFile(static.ConfigFileFullName, confBytes, 0777)
		if err != nil {
			log.Errorf("创建配置文件 " + static.ConfigFileFullName + " 失败，请重新尝试或者手动创建！")
			panic(err)
		}
		log.Infoln("创建配置文件 " + static.ConfigFileFullName + " 成功！")
	}
}
