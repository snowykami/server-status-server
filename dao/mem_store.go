package dao

import (
	"encoding/gob"
	"github.com/LiteyukiStudio/go-logger/log"
	"os"
	"time"
)

var memStore = make(map[string]interface{})

func Save(key string, value interface{}) {
	memStore[key] = value
}

func Delete(key string) error {
	if _, ok := memStore[key]; ok {
		delete(memStore, key)
		return nil
	} else {
		return os.ErrNotExist
	}
}

func Get(key string) interface{} {
	if v, ok := memStore[key]; ok {
		return v
	}
	return nil
}

func GetAll() map[string]interface{} {
	return memStore
}

// init 初始化,从磁盘加载数据到内存
func init() {
	gob.Register(Report{})
	file, err := os.Open("data.gob")
	if err != nil {
		log.Error("Open file error: ", err)
	}
	log.Info("Open file success")
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&memStore)
	if err != nil {
		log.Error("Decode error: ", err)
	}

	// 启动定期持久化goroutine
	go func() {
		for {
			err := persist()
			if err != nil {
				log.Error("Persist error: ", err)
			}
			log.Info("Persist success")
			time.Sleep(30 * time.Second)
		}
	}()
}

// persist 持久化,定期将内存中的数据持久化到磁盘
func persist() error {
	file, err := os.OpenFile("data.gob", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Error("Open file error: ", err)
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(memStore)
	if err != nil {
		log.Error("Encode error: ", err)
		return err
	}
	return nil
}
