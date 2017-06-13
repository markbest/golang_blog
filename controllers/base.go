package controllers

import (
	"github.com/astaxie/beego"
	"blog/models"
)

type BaseController struct {
	beego.Controller
	isLogin bool
}

func (this *BaseController) Prepare() {
	//前台登陆信息
	var loginCustomer models.Customer
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.isLogin = false
	} else {
		this.isLogin = true
		loginCustomer = models.GetCustomerInfo(this.GetSession("userId"))
	}

	//分类列表
	allCategory := models.GetCategoryList()

	this.Data["current_url"] = this.Ctx.Request.RequestURI
	this.Data["isLogin"] = this.isLogin
	this.Data["loginCustomer"] = loginCustomer
	this.Data["category"] = allCategory
	beego.Info(this.Ctx.Request.RequestURI)
}

type AdminBaseController struct {
	beego.Controller
	isAdminLogin bool
}

func (this *AdminBaseController) Prepare() {
	//后台登录信息
	var loginUser models.User
	admin_userLogin := this.GetSession("admin_userLogin")
	if admin_userLogin == nil {
		this.isAdminLogin = false
	} else {
		this.isAdminLogin = true
		loginUser = models.GetUserInfo(this.GetSession("admin_userId"))
	}

	this.Data["current_url"] = this.Ctx.Request.RequestURI
	this.Data["isAdminLogin"] = this.isAdminLogin
	this.Data["loginUser"] = loginUser
}

