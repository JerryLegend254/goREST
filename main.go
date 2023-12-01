package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SportsPlanModel struct{
	Id int `json:"id" binding:"required"`
	SportName string `json:"sport_name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

var SportsPlans []SportsPlanModel

func main(){
	server := gin.Default()
	

	server.GET("/sportsplans", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"data": SportsPlans,
		})
	})

	server.POST("/sportsplans", func (ctx *gin.Context)  {
		var data SportsPlanModel
		err := ctx.ShouldBind(&data)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": fmt.Sprintf("%v", err),
			})
		} else{
			SportsPlans = append(SportsPlans, data)
			ctx.JSON(http.StatusCreated, gin.H{"msg": "New sportplan was added!"})
		}
	})

	server.Run()
}