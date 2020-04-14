package main

import (
	"feedProject/http/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	route.Load(router)

	_ = router.Run(":8085")
}
