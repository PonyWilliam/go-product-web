package handler

import (
	"context"
	"github.com/PonyWilliam/go-common"
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
func SetProductByID(c *gin.Context){
	cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
	id := c.Param("id")
	if id == ""{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"无法识别id",
		})
	}
	new_id,_ := strconv.ParseInt(id,10,64)
	rsp,err := cl.FindProductByID(context.TODO(),&product.Request_ProductID{Id: new_id})
	if rsp == nil || err!=nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"无法读取商品信息",
		})
		return
	}
	if name:=c.PostForm("name");name!=""{
		rsp.ProductName = name
	}
	if description:=c.PostForm("description");description!=""{
		rsp.ProductDescription = description
	}
	if level:=c.PostForm("level");level!=""{
		new_level,_ := strconv.ParseInt(level,10,64)
		rsp.ProductLevel = new_level
	}
	if category:=c.PostForm("category");category!=""{
		new_category,_ := strconv.ParseInt(category,10,64)
		rsp.ProductBelongCategory = new_category
	}
	if important:=c.PostForm("important");important!=""{
		rsp.IsImportant = true
	}
	if is:=c.PostForm("is");is!=""{
		rsp.ProductIs = true
	}
	if wid:=c.PostForm("wid");wid!=""{
		new_wid,_ := strconv.ParseInt(wid,10,64)
		rsp.ProductBelongCustom = new_wid
	}

	if aid:=c.PostForm("aid");aid!=""{
		new_aid,_ := strconv.ParseInt(aid,10,64)
		rsp.ProductBelongArea = new_aid
	}
	if location:=c.PostForm("location");location!=""{
		rsp.ProductLocation = location
	}
	if rfid:=c.PostForm("rfid");rfid!=""{
		rsp.ProductRfid = rfid
	}
	if imageid:=c.PostForm("imageid");imageid!=""{
		new_imageid,_ := strconv.ParseInt(imageid,10,64)
		rsp.ImageId = new_imageid
	}
	//rsp2,err := cl.ChangeProduct(context.TODO(),&product.Request_ProductInfo{
	//	Id: new_id,
	//	ProductName: rsp.ProductName,
	//	ProductDescription: rsp.ProductDescription,
	//	ProductLevel: rsp.ProductLevel,
	//	ProductBelongCategory: rsp.ProductBelongCategory,
	//	IsImportant: rsp.IsImportant
	//})
	temp := &product.Request_ProductInfo{}
	_ = common.SwapTo(rsp,temp)
	rsp2,err := cl.ChangeProduct(context.TODO(),temp)
	if err!=nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":rsp2.Message,
	})
}
 func Test(c *gin.Context){
	 c.String(200,"success")
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
