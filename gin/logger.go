package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/xiaxin/logger"
)

func TraceMiddleware(logger *logger.Logger, tracer opentracing.Tracer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span := logger.NewTraceLogger("web", tracer)
		defer span.Finish()

		span.Info("request")
		ctx.Next()
		span.Info("response")
	}
}
