package service

import (
	"github.com/go-xorm/xorm"
	"go-im/model"
)

var DbEngine *xorm.Engine

func init () {
	DbEngine = model.InitDatabase()
}
