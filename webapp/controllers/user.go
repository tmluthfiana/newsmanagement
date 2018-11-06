package controllers

import (
	"fmt"
	"github.com/eaciit/dbox"
	"github.com/eaciit/knot/knot.v1"
	tk "github.com/eaciit/toolkit"
	. "gopkg.in/mgo.v2/bson"
	. "newsmanagement/webapp/commons"
	. "newsmanagement/webapp/models"
	"os"
)

type UserController struct {
	*BaseController
}

func (d *UserController) Default(k *knot.WebContext) interface{} {
	d.LoadBase(k)
	k.Config.NoLog = false
	d.IsAuthenticate(k)

	k.Config.OutputType = knot.OutputTemplate

	infos := PageInfo{}

	if k.Session("username") != nil {
		infos.Username = k.Session("username").(string)
		infos.UserId = k.Session("userid").(string)
		infos.UserRole = k.Session("userrole").(string)
	}

	infos.PageTitle = "User Management"
	infos.SelectedMenu = "user"
	infos.Breadcrumbs = make(map[string]string, 0)

	d.LoadPartial(k, []string{})
	return infos
}
func (d *UserController) Getuser(k *knot.WebContext) interface{} {
	//d.LoadBase(k)
	k.Config.OutputType = knot.OutputJson
	users := []UserModel{}
	curr, _ := d.DB().Find(&UserModel{}, tk.M{})
	curr.Fetch(&users, 0, false)
	defer curr.Close()
	return users
}
func (d *UserController) Create(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputJson
	user := UserModel{}
	e := k.GetPayload(&user)
	msg := "create failed"
	isSuccess := false
	if e != nil {
		fmt.Println(user)
		msg = "Login failed" + e.Error()
	} else {
		user.Password = GetMD5Hash(user.Password)
		user.Id = NewObjectId()
		e = d.DB().Save(&user)
		if e == nil {
			isSuccess = true
			msg = "Success"

		} else {
			msg = e.Error()
		}
	}
	return tk.M{}.Set("Success", isSuccess).Set("msg", msg)
}

func (d *UserController) Delete(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputJson
	userTemp := UserModel{}
	user := UserModel{}

	_ = k.GetPayload(&userTemp)

	msg := "delete success"
	isSuccess := true
	var err error

	d.DB().GetById(&user, userTemp.Id)

	os.RemoveAll("assets/photos/" + user.Username)

	e := d.DB().Delete(&user)

	if e != nil {
		msg = "delete unsuccess"
		err = e
		isSuccess = false
	}

	return tk.M{}.Set("Success", isSuccess).Set("msg", msg).Set("error", err)
}

func (d *UserController) Update(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputJson
	userTemp := UserModel{}

	e := k.GetPayload(&userTemp)

	if userTemp.IsHas {
		userTemp.Password = GetMD5Hash(userTemp.Password)
	}

	msg := "update success"
	isSuccess := true
	var err error

	e = d.DB().Save(&userTemp)

	if e != nil {
		msg = "update unsuccess"
		err = e
		isSuccess = false
	}

	return tk.M{}.Set("Success", isSuccess).Set("msg", msg).Set("error", err)
}

func (d *UserController) CekByUsername(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputJson
	userTemp := struct {
		Username string
	}{}
	msg := ""
	isSuccess := false
	user := &UserModel{}

	e := k.GetPayload(&userTemp)

	filter := tk.M{}
	filter.Set("where", dbox.Eq("Username", userTemp.Username))
	curr, _ := d.DB().Find(user, filter)

	defer curr.Close()
	e = curr.Fetch(user, 1, true)

	if e != nil {
		isSuccess = true
		msg = "Username Valid"
	}

	return tk.M{}.Set("Success", isSuccess).Set("msg", msg)
}

func (c *UserController) GetUserData(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputJson

	Uname := k.Session("username")
	ress := new(UserModel)
	query := tk.M{}.Set("where", dbox.Eq("Username", Uname))
	csr, _ := c.DB().Find(ress, query)
	res := make([]UserModel, 0)
	csr.Fetch(&res, 0, false)
	defer csr.Close()

	return (tk.M{}).Set("data", res)
}

func (m *UserController) SaveImageProfile(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputJson

	os.RemoveAll("assets/photos/" + k.Request.FormValue("username"))
	os.MkdirAll("assets/photos/"+k.Request.FormValue("username"), 0777)

	e, _ := UploadHandler(k, "userfile", "assets/photos/"+k.Request.FormValue("username")+"/", "photo."+k.Request.FormValue("filetype"))

	if e != nil {
		return e.Error()
	}

	sourcefile := "assets/photos/" + k.Request.FormValue("username") + "/photo." + k.Request.FormValue("filetype")
	destinationfile := "assets/photos/" + k.Request.FormValue("username") + "/photo_32." + k.Request.FormValue("filetype")
	ResizeImg(32, 32, sourcefile, destinationfile)

	destinationfile = "assets/photos/" + k.Request.FormValue("username") + "/photo_64." + k.Request.FormValue("filetype")
	ResizeImg(64, 64, sourcefile, destinationfile)

	destinationfile = "assets/photos/" + k.Request.FormValue("username") + "/photo_128." + k.Request.FormValue("filetype")
	ResizeImg(128, 128, sourcefile, destinationfile)

	os.Remove(sourcefile)

	return (tk.M{}).Set("pathPhoto", "photos/"+k.Request.FormValue("username")+"/photo_128."+k.Request.FormValue("filetype"))
}

func (m *UserController) GetUserProfile(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputJson

	path := ""
	if k.Session("userphoto") != nil && k.Session("imageext") != nil {
		imagePath := k.Session("userphoto").(string)
		imageExt := k.Session("imageext").(string)

		path = imagePath + "/photo_128." + imageExt
	}

	return (tk.M{}).Set("PathImage", path).Set("Username", k.Session("username")).Set("UserRole", k.Session("userrole"))
}
