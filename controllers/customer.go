package controllers

import (
	"blog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type CustomerController struct {
	BaseController
}

// @router /customer/home [get]
func (this *CustomerController) Index()  {
	this.Layout = "layout/frontend/single.tpl"
	this.TplName = "customer.tpl"
}

// @router /customer/setting [get,post]
func (this *CustomerController) Setting() {
	if this.Ctx.Input.Method() == "GET" {
		this.Data["xsrf_token"] = this.XSRFToken()
		this.Layout = "layout/frontend/single.tpl"
		this.TplName = "customer_setting.tpl"
	} else {
		flash := beego.NewFlash()
		id, _ := this.GetInt64("customer_id")
		name := this.GetString("name")
		oldPwd := this.GetString("old_password")
		newPwd := this.GetString("password")
		cnewPwd := this.GetString("password_confirmation")

		params := make(map[string]string)
		if name != "" {
			params["name"] = name
		}
		if newPwd == cnewPwd {
			if oldPwd != "" {
				params["old_password"] = oldPwd
				params["new_password"] = newPwd
			}
		} else {
			flash.Warning("两次密码不一致，请重试！")
			this.Redirect("/customer/setting", 302)
		}
		err := models.UpdateCustomer(id, params)
		if err == nil {
			this.Redirect("/customer/setting", 302)
		} else {
			flash.Warning("更新客户数据失败，请重试！")
			this.Redirect("/customer/setting", 302)
		}
	}
}

// @router /customer/upload [post]
func (this *CustomerController) Upload() {
	// 获取上传文件
	f, h, err := this.GetFile("customer_icon")
	if err == nil {
		// 关闭文件
		f.Close()
	} else {
		// 获取错误则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		this.ServeJSON()
		return
	}
	// 设置保存目录
	dirPath := "./static/uploads/avatar/"
	// 设置保存文件名
	FileName := h.Filename
	// 将文件保存到服务器中
	err = this.SaveToFile("customer_icon", fmt.Sprintf("%s/%s", dirPath, FileName))
	if err != nil {
		// 出错则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		this.ServeJSON()
		return
	}

	// 更新客户头像
	id, _ := this.GetInt64("id")
	params := make(map[string]string)
	params["icon"] = "avatar/" + FileName
	models.UpdateCustomer(id, params)

	this.Data["json"] = map[string]interface{}{"success": 1, "message": "avatar/" + FileName}
	this.ServeJSON()
}

// @router /customer/login [get,post]
func (this *CustomerController) Login() {
	if this.Ctx.Input.Method() == "GET" {
		this.Data["xsrf_token"] = this.XSRFToken()
		this.Layout = "layout/frontend/single_no_head.tpl"
		this.TplName = "login.tpl"
	} else {
		flash := beego.NewFlash()
		email := this.GetString("email")
		password := this.GetString("password")
		customer, err := models.CustomerLogin(email, password)
		if err == nil {
			this.SetSession("userLogin", 1)
			this.SetSession("userId", customer.Id)
			this.Redirect("/customer/home", 302)
		} else {
			flash.Error("登陆失败，请重试!")
			this.Redirect("/customer/login", 302)
		}
	}
}

// @router /customer/logout [get]
func (this *CustomerController) Logout() {
	this.DelSession("userLogin")
	this.DelSession("userInfo")
	this.Redirect("/customer/login", 302)
}

// @router /customer/register [get,post]
func (this *CustomerController) Register() {
	if this.Ctx.Input.Method() == "GET" {
		this.Data["xsrf_token"] = this.XSRFToken()
		this.Layout = "layout/frontend/single_no_head.tpl"
		this.TplName = "register.tpl"
	} else {
		flash := beego.NewFlash()
		customer := &models.Customer{}
		if err := this.ParseForm(customer); err != nil {
			return
		}

		if customer.Password != customer.Repassword {
			flash.Warning("两次密码输入不一致")
			flash.Store(&this.Controller)
			return
		}

		if id, err := models.InsertCustomer(customer); err == nil {
			this.SetSession("userLogin", 1)
			this.SetSession("userId", id)
			this.Redirect("/customer/home", 302)
		} else {
			flash.Error("注册失败，请重试！")
			this.Redirect("/customer/register", 302)
		}
	}
}
