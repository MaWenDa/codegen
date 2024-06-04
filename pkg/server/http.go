package server

import (
	"codegen/pkg/config"
	"codegen/pkg/server/middleware"
	"context"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func NewHttp(options ...HttpOption) *Http {

	// 默认设置
	viper.SetDefault("app.trusted_proxies", []string{"10.0.0.0/8", "172.16.0.0/12"})

	// 初始化gin
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	if config.Debug() {
		gin.SetMode(gin.DebugMode)
		g.Use(gin.Logger())
	}

	g.SetTrustedProxies(viper.GetStringSlice("app.trusted_proxies"))
	// 如果前面不挡nginx的话，腾讯云clb需要应用自己开启gzip
	if viper.GetBool("http.gzip") {
		g.Use(gzip.Gzip(gzip.BestSpeed))
	}

	// 默认中间件
	g.Use(middleware.RequestID(), middleware.Recovery())

	// 开启维护模式，维护模式请求全部禁止
	if viper.GetBool("http.maintenance") {
		g.Use(middleware.UnderMaintenance())
	}

	// 健康检查
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	srv := &Http{
		Server: http.Server{
			Handler: g,
		},
		gin: g,
	}
	for _, f := range options {
		f(srv)
	}

	return srv
}

type Http struct {
	gin *gin.Engine
	http.Server
}

func (h *Http) GracefulStart(ctx context.Context) {
	go func() {
		log.Printf("Http server listen at %s", h.Addr)
		if err := h.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Http server forced to shutdown:", err)
		}
	}()

	select {
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := h.Shutdown(ctx); err != nil {
			log.Fatal("Http server forced to shutdown:", err)
		}
		log.Println("Http server stopped")

		// todo:: 这里并没有平滑重启pprof 和 gops，目前看没必要
	}
}
