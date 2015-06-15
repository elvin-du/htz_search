package controllers

import (
	"htz_search/errors"
	"htz_search/models"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	id := c.Ctx.Input.Param(":id")
	a, err := models.ArticleModel().InfoById(id)
	if nil != err && errors.E_ARTICLE_NOT_FOUND != err {
		beego.Error(err)
		c.Abort("500")
	} /*else if errors.E_ARTICLE_NOT_FOUND == err {
		c.Data["notfound"] = true
	}*/

	c.Data["A"] = a

	c.TplNames = "article.html"
}
