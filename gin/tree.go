package main

import "github.com/gin-gonic/gin"

/**
  engine.trees.tree.root
**/
func main() {
	router := gin.Default()

	router.GET("/a", func(ctx *gin.Context) {})
	router.GET("/a1", func(ctx *gin.Context) {})
	router.GET("/a2", func(ctx *gin.Context) {})
	router.GET("/a22", func(ctx *gin.Context) {})
	router.GET("/a3", func(ctx *gin.Context) {})

	router.Run()
}
