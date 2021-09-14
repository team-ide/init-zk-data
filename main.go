package main

import (
	"errors"
	"fmt"
)

var (
	zkService *ZKService
)

func main() {
	var err error
	zkService, err = CreateZKService(Config.Server)
	println("连接Zookeeper：", Config.Server)
	if err != nil {
		panic(err)
	}

	if Config.Data == nil {
		println("没有需要输入的数据")
		return
	}

	err = appendData("", Config.Data)
	if err != nil {
		panic(err)
	}
}

func appendData(parentPath string, data map[interface{}]interface{}) error {
	var err error
	var isExist bool
	for key_, _ := range data {
		key, _ := key_.(string)
		value, ok := data[key]
		if !ok {
			continue
		}
		if value == nil {
			continue
		}
		path := "/" + key
		if parentPath != "" {
			path = parentPath + path
		}
		subMap, ok := value.(map[interface{}]interface{})
		if ok {
			err = appendData(path, subMap)
			if err != nil {
				return err
			}
		} else {

			isExist, err = zkService.Exists(path)
			if err != nil {
				return err
			}

			if stringValue, ok := value.(string); ok {
				if isExist {
					println("写入路径：", path, "；已存在；覆盖值：", stringValue)
					err = zkService.SetData(path, []byte(stringValue))
					if err != nil {
						return err
					}
				} else {
					println("写入路径：", path, "；不存在；写入值：", stringValue)
					err = zkService.Create(path, []byte(stringValue), 0)
					if err != nil {
						return err
					}
				}
			} else {
				return errors.New(fmt.Sprint("不支持的值类型，路径：", path, "；值：", value))
			}
		}

	}
	return nil
}
