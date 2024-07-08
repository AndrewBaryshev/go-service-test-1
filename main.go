package main

import (
	"github.com/gin-gonic/gin"
)

var savedData []map[string]interface{}

func main() {
    router := gin.Default()

    router.POST("/save_data", func(context *gin.Context) {
        var data map[string]interface{}
        if err := context.ShouldBindJSON(&data); err != nil {
            context.JSON(400, gin.H{"error": err.Error()})
            return
        }
        savedData = append(savedData, data)
        context.JSON(200, gin.H{"message": "Data saved successfully!"})
    })

    router.GET("/get_data", func(context *gin.Context) {
        context.JSON(200, gin.H{"saved_data": savedData})
    })

    router.Run(":8080")
}
