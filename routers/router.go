package routers

import (
	"github.com/astaxie/beego"
	"htz_search/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/search", &controllers.SearchController{})
	beego.Router("/search/kind/:kind", &controllers.SearchController{},"get:GetByKind")
	beego.Router("/search/author/:author", &controllers.SearchController{},"get:GetByAuthor")
	beego.Router("/article/:id", &controllers.ArticleController{})
}
