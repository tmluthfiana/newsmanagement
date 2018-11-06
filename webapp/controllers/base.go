package controllers

import (
	"bufio"
	//. "eaciit/jde/webapp/models"
	"fmt"
	"github.com/eaciit/dbox"
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/orm"
	"github.com/leekchan/accounting"
	//tk "github.com/eaciit/toolkit"
	//"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	wd = func() string {
		d, _ := os.Getwd()
		return d + "/"
	}()
)

type IBaseController interface {
	// not implemented anything yet
}

type BaseController struct {
	base IBaseController
	Ctx  *orm.DataContext
}

type PageInfo struct {
	Username     string
	UserId       string
	UserRole     string
	PageTitle    string
	SelectedMenu string
	Breadcrumbs  map[string]string
}

func (b *BaseController) CloseDb() {
	if b.Ctx != nil {
		b.Ctx.Close()
	}
}

func (b *BaseController) ReadConfig() map[string]string {
	ret := make(map[string]string)
	file, err := os.Open(wd + "conf/app.conf")
	if err == nil {
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, _, e := reader.ReadLine()
			if e != nil {
				break
			}

			sval := strings.Split(string(line), "=")
			ret[sval[0]] = sval[1]
		}
	} else {
		fmt.Println(err.Error())
	}

	return ret
}

func (b *BaseController) SetDb(conn dbox.IConnection) error {
	b.CloseDb()
	b.Ctx = orm.New(conn)
	return nil
}

func (b *BaseController) DB() *orm.DataContext {
	return b.Ctx
}

func (b *BaseController) SetToBaseDB() {
	config := b.ReadConfig()
	ci := &dbox.ConnectionInfo{config["host"], config["database"], config["username"], config["password"], nil}
	newConnection, e := dbox.NewConnection("mongo", ci)

	if e != nil {
		fmt.Println("Gagal Koneksi", e.Error())
		fmt.Println("Isi config")
		fmt.Println(config)
		return

	}
	e = newConnection.Connect()
	if e != nil {
		fmt.Println("Gagal Koneksi 2", e.Error())
		fmt.Println("Isi config")
		fmt.Println(config)
	}
	b.SetDb(newConnection)
}

func (b *BaseController) SetToDB(ci dbox.ConnectionInfo) {
	newConnection, e := dbox.NewConnection("mongo", &ci)

	if e != nil {
		fmt.Println("Gagal Koneksi", e.Error())
		fmt.Println("Isi config")
		return

	}
	e = newConnection.Connect()
	if e != nil {
		fmt.Println("Gagal Koneksi 2", e.Error())
		fmt.Println("Isi config")

	}
	b.SetDb(newConnection)
}

func (b *BaseController) LoadBase(k *knot.WebContext) {
	k.Config.NoLog = true
	b.IsAuthenticate(k)
}

func (b *BaseController) IsAuthenticate(k *knot.WebContext) {
	if k.Session("userid") == nil {
		fmt.Println("Redirecting")
		b.Redirect(k, "login", "default")
	}
	return
}
func (b *BaseController) AllowRoles(roles []string, k *knot.WebContext) {
	for _, i := range roles {
		if i == k.Session("userrole") {
			return
		}
	}
	k.Config.ViewName = "notallowed.html"
}
func (b *BaseController) LoadPartial(k *knot.WebContext, tpls []string) {
	defaultTpls := []string{"_loader.html", "_loader_alt.html", "_profile.html", "_profile_simple.html"}
	if len(tpls) > 0 {
		defaultTpls = append(defaultTpls, tpls...)
	}
	k.Config.IncludeFiles = defaultTpls
}

func (b *BaseController) WriteLog(msg interface{}) {
	log.Printf("%#v\n\r", msg)
	return
}

func (b *BaseController) Redirect(k *knot.WebContext, controller string, action string) {
	log.Println("invalid session , redirecting to " + controller + "/" + action)
	http.Redirect(k.Writer, k.Request, "/"+controller+"/"+action, http.StatusTemporaryRedirect)
}

func (b *BaseController) ConvertJdeDate(k *knot.WebContext, jdedate string) string {
	century := 1900
	year := 0
	daynum := 0
	newfuture := ""
	if jdedate != "0" {
		if string(jdedate[0]) == "0" {
			century = century
			ye, e := strconv.Atoi(string(jdedate[1]) + string(jdedate[2]))
			if e != nil {
				fmt.Println(e.Error())
			}
			year = ye
			day, err := strconv.Atoi(string(jdedate[3]) + string(jdedate[4]) + string(jdedate[5]))
			if err != nil {
				fmt.Println(err.Error())
			}
			daynum = day
			year = century + year
			initial := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
			future := time.Hour * 24 * (time.Duration(daynum) - 1)
			newfuture = (initial.Add(future)).String()
			//fmt.Println(newfuture)
		} else {
			cen, er := strconv.Atoi(string(jdedate[0]))
			if er != nil {
				fmt.Println(er.Error())
			}
			century = century + (100 * cen)
			ye, e := strconv.Atoi(string(jdedate[1]) + string(jdedate[2]))
			if e != nil {
				fmt.Println(e.Error())
			}
			year = ye
			day, err := strconv.Atoi(string(jdedate[3]) + string(jdedate[4]) + string(jdedate[5]))
			if err != nil {
				fmt.Println(err.Error())
			}
			daynum = day
			year = century + year
			initial := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
			future := time.Hour * 24 * (time.Duration(daynum) - 1)
			newfuture = (initial.Add(future)).String()
			//fmt.Println(newfuture)
		}
	}

	return newfuture
}

func (b *BaseController) Currency(originalValue int, simbol string, precision int) string {
	ac := accounting.Accounting{Symbol: simbol, Precision: precision}
	result := ac.FormatMoney(originalValue)

	return result
}

func (b *BaseController) Substr(s string) string {
	result := ""

	datePart := []string{"", "", ""}
	datePart1 := []string{"", "", ""}
	dates := strings.Split(s, " ")
	if len(dates) > 1 {
		datePart = strings.Split(dates[0], "-")
		if len(datePart) > 2 {
			datePart1[0] = datePart[2]
			datePart1[1] = datePart[1]
			datePart1[2] = datePart[0]

			result = strings.Join(datePart1, "-")
		}
	}

	return result
}

/*
func (b *BaseController) GetUserSession(k *knot.WebContext) tk.M {
	currentuser := new(UserModel)

	b.SetToBaseDB()
	b.DB().GetById(currentuser, bson.ObjectIdHex(k.Session("userid").(string)))

	organizations := []OrganizationModel{}
	currentorg := OrganizationModel{}

	if len(currentuser.Groups) > 0 {
		for _, i := range currentuser.Groups {
			neworg := &OrganizationModel{}
			b.DB().GetById(neworg, bson.ObjectIdHex(i))
			organizations = append(organizations, *neworg)
			if neworg.Id.Hex() == currentuser.CurrentOrganization {
				currentorg = *neworg
				fmt.Println("Current org=",currentorg.Id.Hex())
			}
		}
	}

	config := b.ReadConfig()

	if currentuser.CurrentOrganization != "" {
		ci := &dbox.ConnectionInfo{config["host"], strings.Replace(currentorg.Name, " ", "_", -1), "", "", nil}
		b.SetToDB(*ci)
	}

	return (tk.M{}).Set("User", currentuser).Set("Organization", organizations).Set("CurrentOrg", currentorg)
}*/
