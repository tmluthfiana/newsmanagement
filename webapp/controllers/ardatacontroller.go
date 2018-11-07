package controllers

import (
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/mongo"
	"github.com/eaciit/knot/knot.v1"
	tk "github.com/eaciit/toolkit"
	// "gopkg.in/mgo.v2/bson"
	//"encoding/json"
	// "strconv"
	//"strings"
)

func (m *DataBrowserController) ArData(k *knot.WebContext) interface{} {
	m.LoadBase(k)

	k.Config.OutputType = knot.OutputTemplate

	infos := PageInfo{}
	infos.Username = k.Session("username").(string)
	infos.UserId = k.Session("userid").(string)
	infos.UserRole = k.Session("userrole").(string)

	infos.PageTitle = "News Management"
	infos.SelectedMenu = "databrowser"
	infos.Breadcrumbs = make(map[string]string, 0)

	m.WriteLog(infos.UserRole)

	m.LoadPartial(k, []string{})

	return infos
}

func (c *DataBrowserController) GetNews(k *knot.WebContext) interface{} {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputJson

	param := struct {
		Status string
		Topic  string
	}{}
	_ = k.GetPayload(&param)

	filter := []*dbox.Filter{}
	if param.Status != "" {
		filter = append(filter, dbox.Eq("news_status", param.Status))
	}

	if param.Topic != "" {
		filter = append(filter, dbox.Eq("news_topic.topic", param.Topic))
	}

	cursor, e := c.DB().Connection.NewQuery().
		From("news").
		Where(filter...).
		Cursor(nil)

	if e != nil {
		tk.Println(e.Error())
	}

	defer cursor.Close()

	results := []tk.M{}
	e = cursor.Fetch(&results, 0, false)
	if e != nil {
		tk.Println(e.Error())
	}

	return (tk.M{}).Set("data", results).Set("total", len(results))
}

func (c *DataBrowserController) EditNews(k *knot.WebContext) interface{} {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputJson

	param := struct {
		Id string
	}{}
	_ = k.GetPayload(&param)

	filter := []*dbox.Filter{}
	if param.Id != "" {
		filter = append(filter, dbox.Eq("_id", param.Id))
	}

	cursor, e := c.DB().Connection.NewQuery().
		From("news").
		Where(filter...).
		Cursor(nil)

	if e != nil {
		tk.Println(e.Error())
	}

	defer cursor.Close()

	results := []tk.M{}
	e = cursor.Fetch(&results, 0, false)
	if e != nil {
		tk.Println(e.Error())
	}

	return (tk.M{}).Set("data", results)
}

func (c *DataBrowserController) SaveNews(k *knot.WebContext) interface{} {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputJson

	payload := tk.M{}
	_ = k.GetPayload(&payload)

	id := payload.GetString("news_title")

	payload.Set("_id", id)

	ext := c.DB().Connection.NewQuery().
		From("news").
		SetConfig("multiexec", true).
		Save()

	defer ext.Close()
	er := ext.Exec(tk.M{}.Set("data", payload))
	if er != nil {
		return (tk.M{}).Set("data", nil).Set("msg", "save unsuccess")
	}

	return (tk.M{}).Set("data", payload).Set("msg", "save success")
}

func (c *DataBrowserController) DeleteNews(k *knot.WebContext) interface{} {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputJson

	param := struct {
		Id string
	}{}
	_ = k.GetPayload(&param)

	filter := []*dbox.Filter{}
	if param.Id != "" {
		filter = append(filter, dbox.Eq("_id", param.Id))
	}

	ex := c.DB().Connection.NewQuery().
		Where(dbox.And(dbox.In("_id", param.Id))).
		From("news").
		Delete()
	defer ex.Close()

	exQ := ex.Exec(nil)
	if exQ != nil {
		return (tk.M{}).Set("msg", "delete unsuccess")
	}

	return (tk.M{}).Set("msg", "delete success")
}

func (c *DataBrowserController) GetTopic(k *knot.WebContext) interface{} {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputJson

	cursor, e := c.DB().Connection.NewQuery().
		From("topic").
		Cursor(nil)

	if e != nil {
		tk.Println(e.Error())
	}

	defer cursor.Close()

	results := []tk.M{}
	e = cursor.Fetch(&results, 0, false)
	if e != nil {
		tk.Println(e.Error())
	}

	return (tk.M{}).Set("data", results).Set("total", len(results))
}

func (c *DataBrowserController) EditTopic(k *knot.WebContext) interface{} {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputJson

	param := struct {
		Id string
	}{}
	_ = k.GetPayload(&param)

	filter := []*dbox.Filter{}
	if param.Id != "" {
		filter = append(filter, dbox.Eq("_id", param.Id))
	}

	cursor, e := c.DB().Connection.NewQuery().
		From("topic").
		Where(filter...).
		Cursor(nil)

	if e != nil {
		tk.Println(e.Error())
	}

	defer cursor.Close()

	results := []tk.M{}
	e = cursor.Fetch(&results, 0, false)
	if e != nil {
		tk.Println(e.Error())
	}

	return (tk.M{}).Set("data", results)
}

func (c *DataBrowserController) SaveTopic(k *knot.WebContext) interface{} {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputJson

	payload := tk.M{}
	_ = k.GetPayload(&payload)

	ext := c.DB().Connection.NewQuery().
		From("topic").
		SetConfig("multiexec", true).
		Save()

	defer ext.Close()
	er := ext.Exec(tk.M{}.Set("data", payload))
	if er != nil {
		return (tk.M{}).Set("data", nil).Set("msg", "save unsuccess")
	}

	return (tk.M{}).Set("data", payload).Set("msg", "save success")
}

func (c *DataBrowserController) DeleteTopic(k *knot.WebContext) interface{} {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputJson

	param := struct {
		Id string
	}{}
	_ = k.GetPayload(&param)

	filter := []*dbox.Filter{}
	if param.Id != "" {
		filter = append(filter, dbox.Eq("_id", param.Id))
	}

	ex := c.DB().Connection.NewQuery().
		Where(dbox.And(dbox.In("_id", param.Id))).
		From("topic").
		Delete()
	defer ex.Close()

	exQ := ex.Exec(nil)
	if exQ != nil {
		return (tk.M{}).Set("msg", "delete unsuccess")
	}

	return (tk.M{}).Set("msg", "delete success")
}

func (c *DataBrowserController) GetStatus(k *knot.WebContext) interface{} {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputJson

	cursor, e := c.DB().Connection.NewQuery().
		From("news").
		Group("news_status").
		Cursor(nil)

	if e != nil {
		tk.Println(e.Error())
	}

	defer cursor.Close()

	results := []tk.M{}
	e = cursor.Fetch(&results, 0, false)
	if e != nil {
		tk.Println(e.Error())
	}

	return (tk.M{}).Set("data", results).Set("total", len(results))
}
