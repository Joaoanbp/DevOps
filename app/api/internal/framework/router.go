package framework

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController struct {
	Router   *gin.RouterGroup
	Database *gorm.DB
}

type HandlerFunc func(*gin.Context)

func (this *BaseController) Use(handlers ...gin.HandlerFunc) *BaseController {
	return &BaseController{
		Router:   this.Router.Group("", handlers...),
		Database: this.Database,
	}
}

func (this *BaseController) Group(path string) *BaseController {
	return &BaseController{
		Router:   this.Router.Group(path),
		Database: this.Database,
	}
}

func (this *BaseController) Register(path string, controller any) {
	group := this.Router.Group(path)

	if c, ok := controller.(interface{ Get(*gin.Context) }); ok {
		group.GET("", c.Get)
	}
	if c, ok := controller.(interface{ Post(*gin.Context) }); ok {
		group.POST("", c.Post)
	}
	if c, ok := controller.(interface{ Put(*gin.Context) }); ok {
		group.PUT("", c.Put)
	}
	if c, ok := controller.(interface{ Delete(*gin.Context) }); ok {
		group.DELETE("", c.Delete)
	}
}

func Register(base *BaseController, path string, controller any) {
	base.Register(path, controller)
}
