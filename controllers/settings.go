package controllers

import (
	"html/template"

	"github.com/nihilizm/openvpn-web-ui/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SettingsController struct {
	BaseController
}

func (c *SettingsController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.LoginPath())
		return
	}
	c.Data["breadcrumbs"] = &BreadCrumbs{
		Title: "Settings",
	}
}

func (c *SettingsController) Get() {
	c.TplName = "settings.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	settings := models.Settings{Profile: "default"}
	settings.Read("Profile")
	c.Data["Settings"] = &settings
}

func (c *SettingsController) Post() {
	c.TplName = "settings.html"

	flash := beego.NewFlash()
	settings := models.Settings{Profile: "default"}
	settings.Read("Profile")
	if err := c.ParseForm(&settings); err != nil {
		beego.Warning(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}
	c.Data["Settings"] = &settings

	o := orm.NewOrm()
	if _, err := o.Update(&settings); err != nil {
		flash.Error(err.Error())
	} else {
		flash.Success("Settings has been updated")
		models.GlobalCfg = settings
	}
	flash.Store(&c.Controller)
}
