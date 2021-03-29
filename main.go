package main

import (
	"github.com/PonyWilliam/go-ProductWeb/handler"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
)
func main(){
	consul:= consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1"}
	})
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	service:= web.NewService(
		web.Name("go.micro.api.product"),
		web.Address("0.0.0.0:8100"),
		web.Registry(consul),
		web.Handler(router),
	)
	_ = service.Init()
	v1 := router.Group("product")
	v1.GET("/rfid/:rfid",handler.GetProductByRFID)
	v1.GET("/id/:id",handler.GetProductByID)
	v1.GET("/name/:name",handler.GetProductByName)
	v1.GET("/area/:aid",handler.GetProductByArea)
	v1.GET("/worker/:wid",handler.GetProductByCustom)
	v1.GET("/",handler.GetProductAll)
	v1.POST("/:id",handler.SetProductByID)
	v1.GET("/test",handler.Test)
	v2 := router.Group("category")
	v2.GET("/",handler.JWTAuthMiddleware(),handler.FindCategories)
	v2.GET("/:id",handler.JWTAuthMiddleware(),handler.FindCategoryByID)
	v2.DELETE("/:id",handler.JWTAuthMiddleware(),handler.DeleteCategory)
	_ = router.Run()
	if err:=service.Run();err!=nil{
		log.Fatal(err.Error())
	}
}
