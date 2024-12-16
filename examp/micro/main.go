package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Flyingmn/ebase"
	_ "github.com/Flyingmn/ebase"
	"github.com/gin-gonic/gin"
)

func getExecutableDir() (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	executableDir := filepath.Dir(executablePath)
	return executableDir, nil
}

// 使用go run main.go  启动测试服务
func main() {
	path, _ := os.Getwd()
	ebase.SetProjectPath(path)
	//增加心跳推送
	ebase.SetHeartbeatPush(func() error {
		log.Println("HeartbeatPush")
		return nil
	})

	//增加服务注册
	ebase.SetRegfunc(func() error {
		log.Println("Regfunc")
		return nil
	})
	ebase.Init()
	eb := ebase.GetEbInstance()
	r, err := eb.GetHttpServer()
	//gin 库
	if err != nil {
		log.Panicln(err.Error())
	}

	r.GET("/ping", func(context *gin.Context) {
		value, exists := context.Get("EbaseRequestID")
		if exists {
			fmt.Println("requestID:", value)
		}
	})

	eb.Run()
}
