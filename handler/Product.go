package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PonyWilliam/go-ProductWeb/cache"
	borrowlog "github.com/PonyWilliam/go-borrow-logs/proto"
	borrow "github.com/PonyWilliam/go-borrow/proto"
	"github.com/PonyWilliam/go-common"
	product "github.com/PonyWilliam/go-product/proto"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/v2/client"
	"strconv"
)

//获取产品接口均不需要验证信息，只有update方法需要
func GetProductByRFID(c *gin.Context){
	fmt.Println("test")
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
	fmt.Println(res)
	c.JSON(200,res)
}
func GetProductByID(c *gin.Context){
	ID := c.Param("id")
	new_ID,_ := strconv.ParseInt(ID,10,64)
	pro,err := cache.GetGlobalCache(fmt.Sprintf("product_%v",new_ID))
	if err == redis.Nil || err != nil{
		cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
		res,err := cl.FindProductByID(context.TODO(),&product.Request_ProductID{Id: new_ID})
		if err!=nil{
			c.JSON(200,gin.H{
				"code":500,
				"msg":err.Error(),
			})
			return
		}
		_ = cache.SetGlobalCache(fmt.Sprintf("product_%v",new_ID),res)
		c.JSON(200,res)
		return
	}
	result := &product.Response_ProductInfo{}
	_ = json.Unmarshal([]byte(pro), &result)
	c.JSON(200,result)
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
	c.JSON(200,res.Infos)

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
	c.JSON(200,gin.H{
		"code":200,
		"data":res.Infos,
	})

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
	c.JSON(200,gin.H{
		"code":200,
		"data":res.Infos,
	})

}
func GetProductAll(c *gin.Context){
	//做redis缓存
	res,err :=  cache.GetGlobalCache("product")
	if err == redis.Nil || err != nil{
		cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
		res,err := cl.FindAll(context.TODO(),&product.Request_Null{})
		if err!=nil{
			c.JSON(200,gin.H{
				"code":500,
				"msg":err.Error(),
			})
			return
		}
		//存入缓存
		err = cache.SetGlobalCache("product",res)
		if err != nil{
			fmt.Println(err.Error())
		}
		c.JSON(200,gin.H{
			"code":200,
			"data":res.Infos,
		})
	}else{
		result := &product.Response_ProductInfos{}
		_ = json.Unmarshal([]byte(res), &result)
		c.JSON(200,gin.H{
			"code":200,
			"data":result.Infos,
		})
	}
}
func GetProductByCategory(c *gin.Context){
	id,_ := strconv.ParseInt(c.Param("cid"),10,64)
	cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
	rsp,err := cl.FindProductByCategory(context.TODO(),&product.Request_ProductCategory{Cid: id})
	if err!=nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":rsp.Infos,
	})

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
		if important == "true" || important == "1"{
			rsp.IsImportant = true
		}else{
			rsp.IsImportant = false
		}
	}
	if is:=c.PostForm("is");is!=""{
		if is == "true" || is == "1"{
			rsp.ProductIs = true
		}else{
			rsp.ProductIs = false
		}
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
	cache.DelCache("product")
	cache.DelCache(fmt.Sprintf("product_%v",new_id))
	c.JSON(200,gin.H{
		"code":200,
		"msg":rsp2.Message,
	})
}
func DelProduct(c *gin.Context){
	res,_ := c.Get("username")
	if res != "admin"{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"您无权访问",
		})
		return
	}
	id,_ := strconv.ParseInt(c.Param("id"),10,64)
	cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
	rsp, err := cl.DelProduct(context.TODO(), &product.Request_ProductID{Id: id})
	if err!=nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	cache.DelCache("product")
	cache.DelCache(fmt.Sprintf("product_%v",id))
	c.JSON(200,gin.H{
		"code":200,
		"msg":rsp.Message,
	})
}
func ProductLog(c *gin.Context){
	res,_ := c.Get("username")
	if res != "admin"{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"您无权访问",
		})
		return
	}
	cl := borrow.NewBorrowService("go.micro.service.borrow",client.DefaultClient)
	rsp,err := cl.FindBorrowAll(context.TODO(),&borrow.Null_Request{})
	if err != nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":rsp.Logs,
	})
}
func CreateProduct(c *gin.Context){
	res,_ := c.Get("username")
	if res != "admin"{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"您无权访问",
		})
		return
	}
	var is,important bool
	cl := product.NewProductService("go.micro.service.product",client.DefaultClient)
	name := c.PostForm("name")
	category,_ := strconv.ParseInt(c.PostForm("category"),10,64)
	area,_ := strconv.ParseInt(c.PostForm("area"),10,64)
	temp,_ := strconv.ParseInt(c.PostForm("is"),10,64)
	if temp == 0{
		is = false
	}else{
		is = true
	}
	temp,_ = strconv.ParseInt(c.PostForm("important"),10,64)
	if temp == 0{
		important = false
	}else{
		important = true
	}
	level,_ := strconv.ParseInt(c.PostForm("level"),10,64)
	rfid := c.PostForm("rfid")
	description := c.PostForm("description")
	rsp,err := cl.AddProduct(context.TODO(),&product.Request_ProductInfo{
		ProductName: name,
		ProductLocation: " ",
		ProductIs: is,
		ImageId: 0,
		IsImportant: important,
		ProductDescription: description,
		ProductBelongArea: area,
		ProductBelongCategory: category,
		ProductLevel: level,
		ProductRfid: rfid,
	})
	if err != nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	cache.DelCache("product")
	c.JSON(200,gin.H{
		"code":200,
		"msg":rsp.Message,
	})
}
func GetBorrowLog(c *gin.Context){
	res,_ := c.Get("username")
	if res != "admin"{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"您无权访问",
		})
		return
	}
	resu,err := cache.GetGlobalCache("logs")
	if err == redis.Nil || err != nil{
		cl := borrowlog.NewBorrowLogsService("go.micro.service.borrowlog",client.DefaultClient)
		rsp,err := cl.FindAll(context.TODO(),&borrowlog.Req_Null{})
		if err != nil{
			c.JSON(200,gin.H{
				"code":500,
				"msg":err.Error(),
			})
			return
		}
		_ = cache.SetGlobalCache("logs", rsp.Logs)
		c.JSON(200,gin.H{
			"code":200,
			"data":rsp.Logs,
		})
		return
	}
	result := &borrowlog.RspLogs{}
	_ = json.Unmarshal([]byte(resu), &result.Logs)
	c.JSON(200,gin.H{
		"code":200,
		"data":result.Logs,
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


