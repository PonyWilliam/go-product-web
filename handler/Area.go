package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PonyWilliam/go-ProductWeb/cache"
	area "github.com/PonyWilliam/go-area/proto"
	product "github.com/PonyWilliam/go-product/proto"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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
	cache.DelCache("area")
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
	cache.DelCache(fmt.Sprintf("area_%v",id))
	cache.DelCache("area")
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
	areas,err := cache.GetGlobalCache("area")
	if err == redis.Nil || err != nil{
		cl := area.NewAreaService("go.micro.service.area",client.DefaultClient)
		rsp,err := cl.FindAll(context.TODO(),&area.Request_NULL{})
		if err != nil{
			c.JSON(200,gin.H{
				"code":500,
				"msg":err.Error(),
			})
			return
		}
		_ = cache.SetGlobalCache("area", rsp.Infos)
		c.JSON(200,gin.H{
			"code":200,
			"data":rsp.Infos,
		})
		return
	}
	result := &product.Response_ProductInfos{}
	_ = json.Unmarshal([]byte(areas), &result.Infos)
	c.JSON(200,gin.H{
		"code":200,
		"data":result.Infos,
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
	cache.DelCache(fmt.Sprintf("area_%v",new_id))
	cache.DelCache("area")
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
	})
}