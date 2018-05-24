/**********************************************
** @Des: This file ...
** @Author: yongze.chen
** @Date:   2017-09-08 00:18:02
** @Last Modified by:   yongze.chen
** @Last Modified time: 2017-09-16 17:26:48
***********************************************/

package models

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func Init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	dbname_tangka := beego.AppConfig.String("db.sl.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	fmt.Println(dsn)
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	dsn_tangka := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname_tangka + "?charset=utf8"
	fmt.Println(dsn_tangka)
	if timezone != "" {
		dsn_tangka = dsn_tangka + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterDataBase("tangka", "mysql", dsn_tangka)
	orm.RegisterModel(new(Auth), new(Role), new(RoleAuth), new(Admin),
		new(Group), new(Env), new(Code), new(ApiSource),
		new(Lots),
		new(ApiDetail), new(ApiPublic), new(Template))

	//dev 环境 开启 orm sql
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}

func TableSLName(name string) string {
	return beego.AppConfig.String("db.sl.prefix") + name
}
