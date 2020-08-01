package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuiyi1997/go-gin-api/app/controller/jaeger_conn"
	"github.com/zhuiyi1997/go-gin-api/app/controller/product"
	"github.com/zhuiyi1997/go-gin-api/app/controller/test"
	"github.com/zhuiyi1997/go-gin-api/app/route/middleware/exception"
	"github.com/zhuiyi1997/go-gin-api/app/route/middleware/jaeger"
	"github.com/zhuiyi1997/go-gin-api/app/route/middleware/logger"
	"github.com/zhuiyi1997/go-gin-api/app/util/response"
	"github.com/zhuiyi1997/go-gin-api/app/controller/auth"
)

func SetupRouter(engine *gin.Engine) {

	//设置路由中间件
	engine.Use(logger.SetUp(), exception.SetUp(), jaeger.SetUp())

	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404,"请求方法不存在", nil)
	})

	engine.GET("/ping", func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(1,"pong", nil)
	})

	// 测试链路追踪
	engine.GET("/jaeger_test", jaeger_conn.JaegerTest)

	//@todo 记录请求超时的路由

	ProductRouter := engine.Group("/product")
	{
		// 新增产品
		ProductRouter.POST("", product.Add)

		// 更新产品
		ProductRouter.PUT("/:id", product.Edit)

		// 删除产品
		ProductRouter.DELETE("/:id", product.Delete)

		// 获取产品详情
		ProductRouter.GET("/:id", product.Detail)
	}

	// 登录注册
	LoginRoute := engine.Group("/auth")
	{
		// 发送验证码
		LoginRoute.POST("/send_sms", auth.SendSms)
		//　登录
		LoginRoute.POST("/login", auth.Login)
	}

	// 测试加密性能
	TestRouter := engine.Group("/test")
	{
		// 测试 MD5 组合 的性能
		TestRouter.GET("/md5", test.Md5Test)

		// 测试 AES 的性能
		TestRouter.GET("/aes", test.AesTest)

		// 测试 RSA 的性能
		TestRouter.GET("/rsa", test.RsaTest)
	}
}
