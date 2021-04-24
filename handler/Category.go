package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PonyWilliam/go-ProductWeb/cache"
	category "github.com/PonyWilliam/go-category/proto"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/v2/client"
	"strconv"
)
func DeleteCategory(c *gin.Context){
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
	if id == ""{
		c.JSON(200,gin.H{
			"code":200,
			"msg":"参数非法",
		})
		return
	}
	new_id,_ := strconv.ParseInt(id,10,64)
	cl := category.NewCategoryService("go.micro.services.category",client.DefaultClient)
	rsp,_ := cl.DeleteCategory(context.TODO(),&category.Delete_Category_Request{CategoryId: new_id})
	cache.DelCache("category")
	cache.DelCache(fmt.Sprintf("category_%v",new_id))
	c.JSON(200,gin.H{
		"code":200,
		"msg":rsp.Message,
	})
}
func FindCategoryByID(c *gin.Context){
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
	if id == ""{
		c.JSON(200,gin.H{
			"code":200,
			"msg":"参数非法",
		})
		return
	}
	new_id,_ := strconv.ParseInt(id,10,64)
	res,err := cache.GetGlobalCache(fmt.Sprintf("category_%v",new_id))
	if err != nil || err == redis.Nil{
		cl := category.NewCategoryService("go.micro.services.category",client.DefaultClient)
		rsp,_ := cl.FindCategoryById(context.TODO(),&category.FindCateGoryById_Request{Id: new_id})
		_ = cache.SetGlobalCache(fmt.Sprintf("category_%v",new_id),rsp)
		c.JSON(200,gin.H{
			"code":200,
			"data":rsp,
		})
		return
	}
	result := &category.Category_Response{}
	_ = json.Unmarshal([]byte(res), &result)
	c.JSON(200,gin.H{
		"code":200,
		"data":result,
	})
}
func FindCategoriesByName(c *gin.Context){
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
	name := c.Param("name")
	if name == ""{
		c.JSON(200,gin.H{
			"code":200,
			"msg":"参数非法",
		})
		return
	}
	cl := category.NewCategoryService("go.micro.services.category",client.DefaultClient)
	rsp,_ := cl.FindCategoryByName(context.TODO(),&category.Find_CategoryByName_Request{Name: name})
	c.JSON(200,gin.H{
		"code":200,
		"data":rsp,
	})
}
func FindCategories(c *gin.Context){
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
	res,err := cache.GetGlobalCache("category")
	if err == redis.Nil || err != nil{
		cl := category.NewCategoryService("go.micro.services.category",client.DefaultClient)
		rsp,_ := cl.FindAllCategory(context.TODO(),&category.Find_All_Request{})
		_ = cache.SetGlobalCache("category", rsp)
		c.JSON(200,gin.H{
			"code":200,
			"data":rsp,
		})
		return
	}
	result := &category.Find_All_Response{}
	_ = json.Unmarshal([]byte(res), &result)
	c.JSON(200,gin.H{
		"code":200,
		"data":result,
	})
}
func CreateCategory(c *gin.Context){
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
	cl := category.NewCategoryService("go.micro.services.category",client.DefaultClient)
	rsp,_ := cl.CreateCategory(context.TODO(),&category.Create_Category_Request{CategoryName: name,CategoryDescription: description})
	cache.DelCache("category")
	c.JSON(200,gin.H{
		"code":200,
		"msg":rsp.Message,
	})
}
func UpdateCategory(c *gin.Context){
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
	cl := category.NewCategoryService("go.micro.services.category",client.DefaultClient)
	rsp,_ := cl.UpdateCategory(context.TODO(),&category.Update_Category_Request{CategoryId:new_id,CategoryName: name,CategoryDescription: description})
	cache.DelCache("category")
	cache.DelCache(fmt.Sprintf("category_%v",new_id))
	c.JSON(200,gin.H{
		"code":200,
		"msg":rsp.Message,
	})
}