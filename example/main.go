package main

import (
	"eager"
	"fmt"
	"time"
)

//Config config
type Config struct {
	Debug   bool
	Port    int
	TimeOut int
}

//config Config instance

func main() {
	eagerConfig := eager.NewConfig("配置文件路径")
	if err := eagerConfig.Parse(eager.TOML, nil); err != nil {
		fmt.Println("parse config error: ", err.Error())
	}
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		fmt.Println(eagerConfig.GetInt64("timeout"), eagerConfig.GetInt64("port"), eagerConfig.GetBool("debug"))
	}

	select {}
}
