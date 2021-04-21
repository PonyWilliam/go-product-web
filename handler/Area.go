package handler

import (
	"context"
	area "github.com/PonyWilliam/go-area/proto"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"strconv"
)

func CreateArea(c *gin.Context){
	user,ok := c.Get("username")
	if ok == false{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"无法读取到用户信息",
		})
		return
	}
	if user!= "admin"{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"请使用管理员账号登陆",
		})
		return
	}
	name := c.PostForm("name")
	description := c.PostForm("description")
	cl := area.NewAreaService("go.micro.service.area",client.DefaultClient)
	_, err := cl.CreateArea(context.TODO(),&area.Request_Add_Area{Name: name,Description: description})
	if err != nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
	})
}
func DelArea(c *gin.Context){
	user,ok := c.Get("username")
	if ok == false{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"无法读取到用户信息",
		})
		return
	}
	if user!= "admin"{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"请使用管理员账号登陆",
		})
		return
	}
	id,_ := strconv.ParseInt(c.Param("id"),10,64)
	cl := area.NewAreaService("go.micro.service.area",client.DefaultClient)
	_,err := cl.DelArea(context.TODO(),&area.Request_AreaID{Id: id})
	if err != nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
	})
}
func FindAreaAll(c *gin.Context){
	user,ok := c.Get("username")
	if ok == false{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"无法读取到用户信息",
		})
		return
	}
	if user!= "admin"{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"请使用管理员账号登陆",
		})
		return
	}
	cl := area.NewAreaService("go.micro.service.area",client.DefaultClient)
	rsp,err := cl.FindAll(context.TODO(),&area.Request_NULL{})
	if err != nil{
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
func UpdateArea(c *gin.Context){
	user,ok := c.Get("username")
	if ok == false{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"无法读取到用户信息",
		})
		return
	}
	if user!= "admin"{
		c.JSON(200,gin.H{
			"code":500,
			"msg":"请使用管理员账号登陆",
		})
		return
	}
	id := c.Param("id")
	new_id,err := strconv.ParseInt(id,10,64)
	if id == "" || err != nil {
		c.JSON(200,gin.H{
			"code":500,
			"msg":"无法解析的id",
		})
		return
	}
	name := c.PostForm("name")
	description := c.PostForm("description")
	cl := area.NewAreaService("go.micro.service.area",client.DefaultClient)
	_, err = cl.UpdateArea(context.TODO(),&area.Request_Update_Area{Id:new_id,Name: name,Description: description})
	if err != nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
	})
}