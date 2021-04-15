package api

import (
	"fmt"
	"main/Interceptor"
	"main/model"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", Interceptor.JwtVerify, getProduct)
		productAPI.POST("/product", createProduct)
	}
}

func getProduct(c *gin.Context) {
	username := c.GetString("jwt_username")
	c.JSON(http.StatusOK, gin.H{"result": "get Product", "username": username})
}

func createProduct(c *gin.Context) {
	product := model.Product{}
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	image, _ := c.FormFile("image")
	product.Image = image.Filename

	runningDir, _ := os.Getwd()
	filepath := fmt.Sprintf("%s/uploaded/images/%s", runningDir, image.Filename)
	c.SaveUploadedFile(image, filepath)

	c.JSON(http.StatusOK, gin.H{"result": "create Product", "product": product})
}
