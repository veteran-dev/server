package initialize

import (
	"net/http"

	swaggerFiles "github.com/swaggo/files"

	"github.com/5asp/gin-vue-admin/server/docs"
	"github.com/5asp/gin-vue-admin/server/global"
	"github.com/5asp/gin-vue-admin/server/middleware"
	"github.com/5asp/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {

	if global.GVA_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router := gin.New()
	Router.Use(gin.Recovery())
	if global.GVA_CONFIG.System.Env != "public" {
		Router.Use(gin.Logger())
	}

	InstallPlugin(Router)
	systemRouter := router.RouterGroupApp.System

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath))

	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	WebGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	// WebGroup.Use(middleware.JWTAuth())
	{
		webRouter := router.RouterGroupApp.Web
		webRouter.InitWebRouter(WebGroup)
	}

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{

		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)
	}

	GeneralGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	GeneralGroup.Use(middleware.JWTAuth())
	{
		generalRouter := router.RouterGroupApp.General
		generalRouter.InitGeneralRouter(GeneralGroup)
	}

	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitAutoCodeRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)
		systemRouter.InitSysExportTemplateRouter(PrivateGroup)

	}
	{
		cityRouter := router.RouterGroupApp.City
		cityRouter.InitCityDataRouter(PrivateGroup)
	}
	{
		orderRouter := router.RouterGroupApp.Order
		orderRouter.InitOrderRouter(PrivateGroup)
	}
	{
		carCombinationRouter := router.RouterGroupApp.CarCombination
		carCombinationRouter.InitCarCombinationRouter(PrivateGroup)
	}
	{
		cityCarCombinationRouter := router.RouterGroupApp.CityCarCombination
		cityCarCombinationRouter.InitCityCarCombinationRouter(PrivateGroup)
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
