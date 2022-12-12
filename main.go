package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go_forum/common/initdo"
	"go_forum/common/setUp/config"
	"go_forum/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var c int = 1

// @title forum项目接口文档
// @version v1
// @description go开发的论坛
// @termsOfService github.com/20gu00

// @contact.name cjq
// @contact.url github.com/20gu00

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	var confFile string
	flag.StringVar(&confFile, "conf", "", "配置文件")
	flag.Parse()
	//读取配置文件,加载配置文件需要时间如果用goroutine方式去加载最好主goroutine阻塞一会,不然那拿到的配置值为空
	if err := config.ConfRead(confFile); err != nil {
		fmt.Printf("读取配置文件失败, err:%v\n", err)
		panic(err)
	}

	ch := make(chan int)
	go func() {
		initdo.InitDO(ch)
	}()
	r := router.InitRouter()

	server := http.Server{
		Addr:           fmt.Sprintf(":%d", config.Conf.Port),
		Handler:        r,
		ReadTimeout:    time.Duration(config.Conf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Conf.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << config.Conf.MaxHeader,
	}

	go func() {
		zap.L().Info("[Info]",
			zap.String("程序名称", viper.GetString("app_name")),
			zap.String("程序版本", viper.GetString("version")),
			zap.Int("server port", viper.GetInt("app_port")),
		)
		fmt.Println("[Info] server port:", viper.GetInt("app_port"))
		if err := server.ListenAndServe(); err != nil { //阻塞
			zap.L().Info("[Info] web server启动失败", zap.Error(err))
		}

	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	ch <- c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("server不正常退出,shutdown", zap.Error(err))
	}

	zap.L().Info("server退出了")
}
