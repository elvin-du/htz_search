package controllers

import (
	"htz_search/errors"
	"htz_search/models"

	"github.com/astaxie/beego"
)

type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	key := c.GetString("search_key")
	as, err := models.ArticleModel().Query(key)
	if nil != err && errors.E_ARTICLE_NOT_FOUND != err {
        beego.Error(err)
		c.Abort("500")
		return
	} /*else if errors.E_ARTICLE_NOT_FOUND == err {
		c.Data["notfound"] = true
	}*/

	c.Data["as"] = as

	c.TplNames = "result.html"
}

func (c *SearchController) GetByAuthor() {
	key := c.Ctx.Input.Param(":author")
	as, err := models.ArticleModel().InfoByAuthor(key)
	if nil != err && errors.E_ARTICLE_NOT_FOUND != err {
		beego.Error(err)
		c.Abort("500")
	} /*else if errors.E_ARTICLE_NOT_FOUND == err {
		c.Data["notfound"] = true
	}*/

	c.Data["as"] = as

	c.TplNames = "result.html"
}

func (c *SearchController) GetByKind() {
	key := c.Ctx.Input.Param(":kind")
	as, err := models.ArticleModel().InfoByKind(key)
	if nil != err && errors.E_ARTICLE_NOT_FOUND != err {
		beego.Error(err)
		c.Abort("500")
	} /*else if errors.E_ARTICLE_NOT_FOUND == err {
		c.Data["notfound"] = true
	}*/

	c.Data["as"] = as

	c.TplNames = "result.html"
}
