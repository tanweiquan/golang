package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Database struct {
	Sql   Mysql `yaml:"mysql"`
	Nosql Redis `yaml:"redis"`
	Rpcok []Rpc `yaml:"rpc"`
}
type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type Rpc struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

var Data = new(Database)

func main() {
	b, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("读取配置文件失败，错误为：", err)
		return
	}
	err = yaml.Unmarshal(b, Data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Data.Sql.Host)
	fmt.Println(Data.Sql.Dbname)
	fmt.Println(Data.Rpcok[0])
	fmt.Println(Data.Rpcok[1])
}
