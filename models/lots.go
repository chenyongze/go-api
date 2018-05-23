package models

import "github.com/astaxie/beego/orm"

type Lots struct {
	//Lot_id int `pk:"auto"`
	Lot_id int `orm:"pk"`
	Lot_name string `json:"lot_name"`
	Lot_body string `json:"lot_body"`
	Lot_sn  string `json:"lot_sn"`
	Lot_image string `json:"lot_image"`
}

func (self *Lots) TableName() string {
	return TableSLName("sale_lots")
}


func LotsGetById(lot_id int) ([]*Lots, error) {

	list := make([]*Lots, 0)
	o := orm.NewOrm()
	o.Using("tangka")
	//var lots [] Lots
	o.Raw("select * from sl_sale_lots where lot_id = ?",lot_id).QueryRows(&list)
	return list, nil
	////lots := new(Lots)
	//o := orm.NewOrm()
	//o.Using("tangka")
	//lots := Lots{Lot_id:lot_id}
	//err := o.QueryTable(TableSLName("sale_lots")).Filter("Lot_id", lot_id).One(lots)
	//if err != nil {
	//	return nil, err
	//}
	//return lots, nil
}