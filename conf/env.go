package conf

import (
	"log"

	"github.com/go-ini/ini"
)

type Environment struct {
	Storage       string
	Key           string
}
var EnvironmentSetting = &Environment{}

var cfg *ini.File

/*
File path - enviroment file ends with .ini
*/
func Setup(filePath string) {
	var err error
	
	cfg, err = ini.Load(filePath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse '%s': %v", filePath, err)
	}

	mapTo("environment", EnvironmentSetting)

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}