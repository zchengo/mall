package initialize

import (
	"fmt"
	"imall/api"
	"imall/global"
	"imall/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {

	engine := gin.Default()

	// 开启跨域
	engine.Use(middleware.Cors())

	// 静态资源请求映射
	engine.Static("/image", global.Config.Upload.SavePath)

	// 商城后台管理API
	web := engine.Group("/web")

	{
		// 用户登录
		web.GET("/captcha", api.GetWebUser().GetCaptcha)
		web.POST("/login", api.GetWebUser().UserLogin)

		// 开启JWT认证,以下接口需要认证成功才能访问
		web.Use(middleware.JwtAuth())

		// 文件上传
		web.POST("/upload", api.GetWebFileUpload().FileUpload)

		// 数据统计
		web.GET("/today/data", api.GetWebStatistics().GetOrderData)
		web.GET("/week/data", api.GetWebStatistics().GetWeekData)

		// 商品管理
		web.POST("/goods/create", api.GetWebGoods().CreateGoods)
		web.DELETE("/goods/delete", api.GetWebGoods().DeleteGoods)
		web.PUT("/goods/update", api.GetWebGoods().UpdateGoods)
		web.PUT("/goods/status/update", api.GetWebGoods().UpdateGoodsStatus)
		web.GET("/goods/list", api.GetWebGoods().GetGoodsList)

		// 订单管理
		web.DELETE("/order/delete", api.GetWebOrder().DeleteOrder)
		web.PUT("/order/update", api.GetWebOrder().UpdateOrder)
		web.GET("/order/list", api.GetWebOrder().GetOrderList)
		web.GET("/order/detail", api.GetWebOrder().GetOrderDetail)

		// 类目管理
		web.POST("/category/create", api.GetWebCategory().CreateCategory)
		web.DELETE("/category/delete", api.GetWebCategory().DeleteCategory)
		web.PUT("/category/update", api.GetWebCategory().UpdateCategory)
		web.GET("/category/list", api.GetWebCategory().GetCategoryList)
		web.GET("/category/option", api.GetWebCategory().GetCategoryOption)

		// 营销管理
		web.POST("/market/create", api.GetWebMarket().CreateMarket)
		web.DELETE("/market/delete", api.GetWebMarket().DeleteMarket)
		web.PUT("/market/update", api.GetWebMarket().UpdateMarket)
		web.PUT("/market/status/update", api.GetWebMarket().UpdateMarketStatus)
		web.GET("/market/list", api.GetWebMarket().GetMarketList)
		web.GET("/market/goods/list", api.GetWebMarket().GetMarketGoods)

		// 问题反馈
		web.POST("/feedback/send", api.GetWebFeedback().SendFeedback)
	}

	// 微信小程序API
	app := engine.Group("/app")

	{
		// 用户登录
		app.POST("/login", api.GetAppUser().UserLogin)

		// 钻石展位
		app.GET("/banner/list", api.GetAppMarket().GetBanners)

		// 商品
		app.GET("/goods/list", api.GetAppGoods().GetGoodsList)
		app.GET("/goods/info", api.GetAppGoods().GetGoodsInfo)
		app.GET("/goods/search", api.GetAppGoods().SearchGoods)

		// 购物车
		app.POST("/cart/add", api.GetAppCart().AddCart)
		app.DELETE("/cart/delete", api.GetAppCart().DeleteCart)
		app.DELETE("/cart/clear", api.GetAppCart().ClearCart)
		app.GET("/cart/info", api.GetAppCart().GetCartInfo)

		// 商品订单
		app.PUT("/order/update", api.GetAppOrder().UpdateOrder)
		app.POST("/order/submit", api.GetAppOrder().SubmitOrder)
		app.GET("/order/list", api.GetAppOrder().GetOrderList)

		// 收货地址管理
		app.POST("/address/save", api.GetAppAddress().SaveAddress)
		app.DELETE("/address/delete", api.GetAppAddress().DeleteAddress)
		app.GET("/address/list", api.GetAppAddress().GetAddressList)
		app.GET("/address/info", api.GetAppAddress().GetAddressInfo)

		// 商品分类
		app.GET("/category/option", api.GetAppCategory().GetCategoryOption)
	}

	// 启动、监听端口
	_ = engine.Run(fmt.Sprintf(":%s", global.Config.Server.Post))
}
