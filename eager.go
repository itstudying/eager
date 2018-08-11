package eager

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"log"
	"sync"
)

//Enum the config type
const (
	_    ConfigType = iota //int enum type
	TOML                   //toml
)

//ConfigType config type constant int
type ConfigType int

// Config struct is details info
type config struct {
	Path        string
	MonitorTime time.Duration
	modTime     time.Time
	storge      interface{}
	parseFunc   func(path string, config interface{}) error
	rwLock      sync.RWMutex
}

// NewConfig method is return Config instance
func NewConfig(path string) *config {

	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			panic("file is not exist")
		}

		if fileInfo.IsDir() {
			panic("is dir")
		}
		panic(fmt.Errorf("get file stat error: %s", err.Error()))
	}

	return &config{Path: path, modTime: fileInfo.ModTime(), MonitorTime: time.Second * 3}
}

//Parse method load config，
//If config is not empty, the configuration is parsed into config
//If config is empty, it is parsed into config according to configType, You'll be able to use the get... method
func (c *config) Parse(configType ConfigType, config interface{}) error {

	var parser Parser

	switch configType {
	case TOML:
		parser = &TOMLParser{}
	}
	if config != nil {
		if reflect.ValueOf(config).Kind() != reflect.Ptr {
			return ErrNotPtr
		}
		if err := parser.ParseConfig(c.Path, config); err != nil {
			return err
		}
		c.storge, c.parseFunc = config, parser.ParseConfig
	} else {
		m := make(map[string]interface{})
		if err := parser.ParseConfig(c.Path, &m); err != nil {
			return err
		}
		c.storge, c.parseFunc = &m, parser.ParseConfig
	}

	//start goroutine to monitor change
	go c.monitorChange()
	return nil
}

//monitorChange monitor file Change
func (c *config) monitorChange() {

	ticker := time.NewTicker(c.MonitorTime)

	for range ticker.C {

		func() {
			fileInfo, err := os.Stat(c.Path)
			if err != nil {
				if os.IsNotExist(err) {
					log.Println(ErrNotExist)
				}

				if fileInfo.IsDir() {
					log.Println(ErrNotFile)
				}
				log.Println("get file stat error: ", err)
				return
			}

			if fileInfo.ModTime().Equal(c.modTime) {
				return
			}

			c.rwLock.Lock()
			defer c.rwLock.Unlock()
			if err := c.parseFunc(c.Path, c.storge); err == nil {
				c.modTime = fileInfo.ModTime()
			} else {
				log.Println("parse error: ", err.Error())
			}
		}()
	}

}
