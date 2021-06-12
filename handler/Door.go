package handler

import (
	"context"
	arealogs "github.com/PonyWilliam/go-arealogs/proto/arealogs"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"strconv"
	"strings"
)
var cl arealogs.AreaLogsService
func init(){
	cl = arealogs.NewAreaLogsService("go.micro.service.arealogs",client.DefaultClient)
}
func DoorAdd (c *gin.Context){
	var pids []int64
	pid := c.PostForm("pid")
	wid,_ := strconv.ParseInt(c.PostForm("wid"),10,64)
	aid,_ := strconv.ParseInt(c.PostForm("aid"),10,64)
	content := c.PostForm("content")
	temp := strings.Split(",",pid)
	for _,v := range temp{
		temp2,err := strconv.ParseInt(v,10,64)
		if err != nil{
			c.JSON(200,gin.H{
				"code":500,
				"msg":"pid有问题",
			})
			return
		}
		pids = append(pids,temp2)
	}
	rsp, _ := cl.AddLog(context.TODO(), &arealogs.ALog{AreaID: aid,PID: pids,WID: wid,Content: content})
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
func DoorFindByAID(c *gin.Context)  {
	aid,_ := strconv.ParseInt(c.PostForm("aid"),10,64)
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
	wid,_ := strconv.ParseInt(c.PostForm("aid"),10,64)
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
	id,_ := strconv.ParseInt(c.PostForm("id"),10,64)
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