package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xiaxin/logger"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var config *logger.Config

func init() {
	config = &logger.Config{
		Name: "gin",
		JaegerConfig: &jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				Type:  "const",
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LocalAgentHostPort: "127.0.0.1:6831",
				LogSpans:           true,
			},
		},
	}
}

func main() {

	logger := config.Build()

	defer logger.Sync()

	cfg := config.ToJaegerConfiguration()

	tracer, closer, err := logger.NewTracer(cfg)
	if err != nil {
		fmt.Errorf("Init failed: %v\n", err)
		return
	}

	defer closer.Close()

	// web
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(TraceMiddleware(logger, tracer))

	r.SetTrustedProxies([]string{"localhost"})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
