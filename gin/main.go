package main

import (
	"github.com/gin-gonic/gin"
)

func v1EndpointHandler(ctx *gin.Context) {
	ctx.String(200, "v1: %s %s", ctx.Request.Method, ctx.Request.URL.Path)
}

func v2EndpointHandler(ctx *gin.Context) {
	ctx.String(200, "v2: %s %s", ctx.Request.Method, ctx.Request.URL.Path)
}

func main() {
	router := gin.Default()

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	v1 := router.Group("/v1")
	v1.GET("/products", v1EndpointHandler)
	// Eg: /v1/products
	v1.GET("/products/:productId", v1EndpointHandler)
	v1.POST("/products", v1EndpointHandler)
	v1.PUT("/products/:productId", v1EndpointHandler)
	v1.DELETE("/products/:productId", v1EndpointHandler)

	v2 := router.Group("/v2")

	v2.GET("/products", v2EndpointHandler)
	v2.GET("/products/:productId", v2EndpointHandler)
	v2.POST("/products", v2EndpointHandler)
	v2.PUT("/products/:productId", v2EndpointHandler)
	v2.DELETE("/products/:productId", v2EndpointHandler)

	router.Run(":5000")

}
