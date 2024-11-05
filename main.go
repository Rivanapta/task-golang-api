package main

import (
	"net/http"
	"strconv"

	"go-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Group route untuk magic
	magic := r.Group("/magic")
	{
		magic.GET("/sum", func(c *gin.Context) {
			n, _ := strconv.Atoi(c.DefaultQuery("n", "0"))
			c.JSON(http.StatusOK, gin.H{"result": utils.MagicSum(n)})
		})

		magic.POST("/pow", func(c *gin.Context) {
			var data struct {
				Num int `json:"n"`
			}
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"result": utils.MagicPow(data.Num)})
		})

		magic.GET("/odd", func(c *gin.Context) {
			n, _ := strconv.Atoi(c.DefaultQuery("n", "0"))
			c.JSON(http.StatusOK, gin.H{"result": utils.MagicOdd(n)})
		})

		magic.POST("/grade", func(c *gin.Context) {
			var data struct {
				Num int `json:"n"`
			}
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"result": utils.MagicGrade(data.Num)})
		})

		magic.GET("/name", func(c *gin.Context) {
			n, _ := strconv.Atoi(c.DefaultQuery("n", "1"))
			c.JSON(http.StatusOK, gin.H{"result": utils.MagicName(n)})
		})

		magic.POST("/tria", func(c *gin.Context) {
			var data struct {
				Num int `json:"n"`
			}
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"result": utils.MagicTria(data.Num)})
		})
	}

	// Group route untuk account
	account := r.Group("/account")
	{
		account.POST("/create", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Success"})
		})
		account.GET("/read", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": map[string]interface{}{}})
		})
		account.PUT("/update", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Success"})
		})
		account.DELETE("/delete", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Success"})
		})
		account.GET("/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
		})
		account.GET("/:username", func(c *gin.Context) {
			username := c.Param("username")
			c.JSON(http.StatusOK, gin.H{"data": "Hi, my name is " + username})
		})
	}

	// Route untuk auth login
	r.POST("/auth/login", func(c *gin.Context) {
		var data struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		isAlphanumeric := func(s string) bool {
			for _, c := range s {
				if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
					return false
				}
			}
			return true
		}

		if isAlphanumeric(data.Username) {
			if _, err := strconv.Atoi(data.Password); err == nil {
				c.JSON(http.StatusOK, gin.H{"message": "Login success"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
		}
	})

	r.Run(":8080") // Listen and serve on port 8080
}
