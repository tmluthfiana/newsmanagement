package controllers

import (
	. "newsmanagement/webapp/commons"
	. "newsmanagement/webapp/models"
	//"fmt"
	"github.com/eaciit/dbox"
	"github.com/eaciit/knot/knot.v1"
	tk "github.com/eaciit/toolkit"
	//"gopkg.in/mgo.v2/bson"
	//"time"
)

type LoginController struct {
	*BaseController
}

func (c *LoginController) Default(k *knot.WebContext) interface{} {
	c.LoadPartial(k, []string{})
	k.Config.NoLog = true

	k.Config.OutputType = knot.OutputTemplate
	k.Config.LayoutTemplate = "_bodyonly.html"
	return ""
}

func (c *LoginController) DoLogin(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputJson
	k.Config.NoLog = true

	p := struct {
		Username string
		Password string
	}{}

	e := k.GetPayload(&p)

	msg := "login failed"
	isLogged := false

	if e != nil {
		msg = "login failed"
	} else { //if p.Email == "eaciit" && p.Password == "Password.1"
		//=======	} else if p.Email == "circularedge" && p.Password == "newyork123" {
		//>>>>>>> .theirs		user:=&UserModel{}
		pass := GetMD5Hash(p.Password)
		user := &UserModel{}
		filter := tk.M{}
		filter.Set("where", dbox.And(dbox.Eq("Username", p.Username), dbox.Eq("Password", pass)))
		curr, _ := c.DB().Find(user, filter)

		defer curr.Close()
		e = curr.Fetch(user, 1, false)
		if e == nil {
			k.SetSession("username", user.Username)
			k.SetSession("userid", user.Id.Hex())
			k.SetSession("userrole", user.Role)
			k.SetSession("userphoto", user.Photo)
			k.SetSession("imageext", user.Ext)

			isLogged = true
			msg = ""
		}

	}

	return tk.M{}.Set("IsLogged", isLogged).Set("Message", msg)
}

func (c *LoginController) DoLogout(k *knot.WebContext) interface{} {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputNone
	k.SetSession("username", nil)
	//k.SetSession("Password", nil)
	k.SetSession("userid", nil)
	k.SetSession("userrole", nil)
	c.Redirect(k, "login", "default")

	return true
}
