package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
)

/*
监听配置文件
*/

func main() {
	watcher, err := fsnotify.NewWatcher() //创建监听者
	if err != nil {
		panic(err)
	}
	defer watcher.Close()
	done := make(chan struct{})
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events: //监听文件或文件夹的变化
				if !ok {
					return
				}
				fmt.Println(event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println(event.Name)
				}
			case err, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println(err)
			}
		}
	}()
	path := "src/package/fsnotify/test"
	err = watcher.Add(path) //监听此路径
	if err != nil {
		panic(err)
	}
	<-done
}
