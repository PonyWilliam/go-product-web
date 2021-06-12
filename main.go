package main

import (
	"github.com/PonyWilliam/go-ProductWeb/global"
	"github.com/PonyWilliam/go-ProductWeb/handler"
	"github.com/PonyWilliam/go-common"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"log"
	"net/http"
)

func main() {
	consul:= consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"1.116.62.214"}
	})
	router := gin.Default()
	service:= web.NewService(
		web.Name("go.micro.api.product"),
		web.Address(":9998"),
		web.Registry(consul),
		web.Handler(router),
	)
	_ = service.Init()
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go func(){
		err := http.ListenAndServe(":9092",hystrixStreamHandler)
		if err != nil{
			log.Fatal(err)
		}
	}()
	go common.PrometheusBoot("5008")
	v1 := router.Group("product")
	v1.GET("/rfid/:rfid",handler.GetProductByRFID)
	v1.GET("/id/:id",handler.GetProductByID)
	v1.GET("/name/:name",handler.GetProductByName)
	v1.GET("/area/:aid",handler.GetProductByArea)
	v1.GET("/worker/:wid",handler.GetProductByCustom)
	v1.GET("/bycategory/:cid",handler.JWTAuthMiddleware(),handler.GetProductByCategory)
	v1.GET("/",handler.GetProductAll)
	v1.PUT("/set/:id",handler.JWTAuthMiddleware(),handler.SetProductByID)
	v1.POST("/add",handler.JWTAuthMiddleware(),handler.CreateProduct)
	v1.GET("/test",handler.Test)
	v1.GET("/category",handler.JWTAuthMiddleware(),handler.FindCategories)
	v1.GET("/category/:id",handler.JWTAuthMiddleware(),handler.FindCategoryByID)
	v1.DELETE("/category/:id",handler.JWTAuthMiddleware(),handler.DeleteCategory)
	v1.POST("/category",handler.JWTAuthMiddleware(),handler.CreateCategory)
	v1.PUT("/category/:id",handler.JWTAuthMiddleware(),handler.UpdateCategory)
	v1.GET("/Area",handler.JWTAuthMiddleware(),handler.FindAreaAll)
	v1.DELETE("/Area/:id",handler.JWTAuthMiddleware(),handler.DelArea)
	v1.POST("/Area",handler.JWTAuthMiddleware(),handler.CreateArea)
	v1.PUT("/Area/:id",handler.JWTAuthMiddleware(),handler.UpdateArea)
	v1.DELETE("del/:id",handler.JWTAuthMiddleware(),handler.DelProduct)
	v1.GET("log/",handler.JWTAuthMiddleware(),handler.ProductLog)
	v1.GET("borrowlog/",handler.JWTAuthMiddleware(),handler.GetBorrowLog)
	v1.POST("door/add",handler.DoorMiddleWare(),handler.DoorAdd)
	v1.GET("door",handler.FindAreaAll)
	v1.GET("door/:id",handler.DoorFindByID)
	v1.GET("door/wid/:wid",handler.DoorFindByWID)
	v1.GET("door/aid/:aid",handler.DoorFindByAID)
	router.Use(Cors())
	_ = router.Run()
	router.Use(Cors())
	_ = router.Run()
	if err := global.SetupRedisDb();err!=nil{
		log.Fatal(err)
	}
	if err:=service.Run();err!=nil{
		log.Fatal(err)
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, token, Token")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}