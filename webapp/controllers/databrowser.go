package controllers

import (
	//. "eaciit/jde/webapp/models"
	"github.com/eaciit/knot/knot.v1"
	//tk "github.com/eaciit/toolkit"
	//"fmt"
)

type DataBrowserController struct {
	*BaseController
}

func (d *DataBrowserController) Default(k *knot.WebContext) interface{} {
	d.LoadBase(k)
	k.Config.NoLog = false
	d.IsAuthenticate(k)
	
	k.Config.OutputType = knot.OutputTemplate
	return ""
}
