package main

import (
	"feedProject/http/route"
	"feedProject/pkg/task"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	route.Load(router)

	task.Start()

	_ = router.Run(":8085")
}
