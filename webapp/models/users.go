package models

import (
	"github.com/eaciit/orm"
	"gopkg.in/mgo.v2/bson"
	//"time"
)

type UserModel struct {
	orm.ModelBase `bson:"-",json:"-"`
	Id            bson.ObjectId `bson:"_id"`
	Username      string        `bson:"Username"`
	Password      string        `bson:"Password"`
	Role          string        `bson:"Role"`
	Photo         string        `bson:"Photo"`
	Ext           string        `bson:"Ext"`
	IsHas         bool
}

func (e *UserModel) RecordID() interface{} {
	return e.Id
}

func NewUser() *UserModel {
	e := new(UserModel)
	e.Id = bson.NewObjectId()
	return e
}

func (m *UserModel) TableName() string {
	return "user"
}
