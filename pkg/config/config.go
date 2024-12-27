package config

import (
	"encoding/json"
	"os"
)

var config configT

type configT struct {

    MuninnHostIp string

}

func Get() configT {

    if config.MuninnHostIp == "" {

        homedir, _ := os.UserHomeDir()
        validateDir(homedir + "/.config/muninn")

        configpath := homedir + "/.config/muninn/config.json"
        file, _ := os.ReadFile(configpath)

        json.Unmarshal(file, &config)

    }

    return config

}

func GetMuninnHost() string {
    return config.MuninnHostIp
}

func validateDir(path string) {

    _, err := os.ReadDir(path)
    if err != nil {
        os.Mkdir(path, 0666)
    }

} 
