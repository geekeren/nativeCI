package main

import (
	"nativeCI/core/model"
)

func main() {
	var pipeline model.Pipeline
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	pipeline.LoadConf("example/nativeCI.yaml")
	pipeline.Run()
	//	c.JSON(200, gin.H{
	//		"message": pipeline,
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
