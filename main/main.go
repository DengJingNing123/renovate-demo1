package main

import (
	_ "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化空的服务器
	app := gin.New()

	log.Printf("启动服务器在 http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", app).Error())
}
