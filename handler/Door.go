package handler

import (
	"context"
	arealogs "github.com/PonyWilliam/go-arealogs/proto/arealogs"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"strconv"
)
func DoorAdd (c *gin.Context){
	cl := arealogs.NewAreaLogsService("go.micro.service.arealogs",client.DefaultClient)
	pid := c.PostForm("pid")
	wid,_ := strconv.ParseInt(c.PostForm("wid"),10,64)
	aid,_ := strconv.ParseInt(c.PostForm("aid"),10,64)
	content := c.PostForm("content")
	rsp, _ := cl.AddLog(context.TODO(), &arealogs.ALog{AreaID: aid,PID: pid,WID: wid,Content: content})
	if !rsp.Result{
		c.JSON(200,gin.H{
			"code":500,
			"msg":rsp.Response,
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":rsp.Response,
	})
}
func DoorFindAll(c *gin.Context){
	cl := arealogs.NewAreaLogsService("go.micro.service.arealogs",client.DefaultClient)
	rsp,err := cl.FindAll(context.TODO(),&arealogs.Null{})
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
func DoorFindByAID(c *gin.Context)  {
	cl := arealogs.NewAreaLogsService("go.micro.service.arealogs",client.DefaultClient)
	aid,_ := strconv.ParseInt(c.Param("aid"),10,64)
	rsp,err := cl.FindByAID(context.TODO(),&arealogs.Area{Aid: aid})
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
func DoorFindByWID(c *gin.Context)  {
	cl := arealogs.NewAreaLogsService("go.micro.service.arealogs",client.DefaultClient)
	wid,_ := strconv.ParseInt(c.Param("wid"),10,64)
	rsp,err := cl.FindByWID(context.TODO(),&arealogs.Worker{Id: wid})
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
func DoorFindByID(c *gin.Context){
	cl := arealogs.NewAreaLogsService("go.micro.service.arealogs",client.DefaultClient)
	id,_ := strconv.ParseInt(c.Param("id"),10,64)
	rsp,err := cl.FindByID(context.TODO(),&arealogs.Id{Id: id})
	if err != nil{
		c.JSON(200,gin.H{
			"code":500,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":rsp,
	})
}
