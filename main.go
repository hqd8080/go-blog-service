/**
 * @Author: hqd
 * @Description: main
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2021/1/17 15:13
 */
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hqd8080/go-blog-service/conf"
	"github.com/hqd8080/go-blog-service/model"
	"github.com/hqd8080/go-blog-service/routers"
)

func setupSetting() error {
	config, err := conf.NewConfig()
	if err != nil {
		return err
	}
	err = config.LoadConf("Server", &conf.ServerConfig)
	if err != nil {
		return err
	}
	err = config.LoadConf("App", &conf.AppConfig)
	if err != nil {
		return err
	}
	err = config.LoadConf("Mysql", &conf.MysqlConfig)
	if err != nil {
		return err
	}
	conf.ServerConfig.ReadTimeout *= time.Second
	conf.ServerConfig.WriteTimeout *= time.Second
	return nil
}

func setupDatabase() error {
	err := model.NewDBConnection(conf.MysqlConfig)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		panic(err)
	}
	err = setupDatabase()
	if err != nil {
		panic(err)
	}
}

func main() {
	gin.SetMode(conf.ServerConfig.RunMode)
	router := routers.NewRouter()
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", conf.ServerConfig.HttpPort),
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       conf.ServerConfig.ReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      conf.ServerConfig.WriteTimeout,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}

	log.Println("server exiting")
}
