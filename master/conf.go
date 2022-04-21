package master

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Port string `json:"port"`
}

var Conf Config

//解析配置文件
func InitConfig(filename string) error {
	context,err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	Conf  = Config{}
	err = json.Unmarshal(context,&Conf)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

