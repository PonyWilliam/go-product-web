package handler

import (
	"context"
	product "github.com/PonyWilliam/go-product/proto"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"strconv"
)

//获取产品接口均不需要验证信息，只有update方法需要
func GetProductByRFID(c *gin.Context){
	RFID := c.Param("rfid")
	cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
	res,err := cl.FindProductByRFID(context.TODO(),&product.Request_ProductRFID{Rfid: RFID})
	if err!=nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,res)
}
func GetProductByID(c *gin.Context){
	ID := c.Param("id")
	new_ID,_ := strconv.ParseInt(ID,10,64)
	cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
	res,err := cl.FindProductByID(context.TODO(),&product.Request_ProductID{Id: new_ID})
	if err!=nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,res)
}
func GetProductByName(c *gin.Context){
	name := c.Param("name")
	cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
	res,err := cl.FindProductByName(context.TODO(),&product.Request_ProductName{Name: name})
	if err!=nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,res)

}
func GetProductByArea(c *gin.Context){
	aid := c.Param("aid")
	new_aid,_ := strconv.ParseInt(aid,10,64)
	cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
	res,err := cl.FindProductByArea(context.TODO(),&product.Request_ProductArea{Aid: new_aid})
	if err!=nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,res)

}
func GetProductByCustom(c *gin.Context){
	wid := c.Param("wid")
	new_wid,_ := strconv.ParseInt(wid,10,64)
	cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
	res,err := cl.FindProductByCustom(context.TODO(),&product.Request_ProductCustom{Wid:new_wid})
	if err!=nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,res)

}
func GetProductAll(c *gin.Context){
	cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
	res,err := cl.FindAll(context.TODO(),&product.Request_Null{})
	if err!=nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,res)
}
//func GetProductByCategory(c *gin.Context){
//	cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
//	res,err := cl.F
//	if err!=nil{
//		c.JSON(200,gin.H{
//			"code":500,
//			"msg":err.Error(),
//		})
//		return
//	}
//	c.JSON(200,res)
//
//}
