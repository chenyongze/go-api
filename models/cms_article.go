package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CmsArticle struct {
	Id              int    `orm:"column(id);auto" json:"id"`
	UserId          uint64 `orm:"column(user_id)" description:"发表者用户id" json:"user_id"`
	ArticleStatus   int8   `orm:"column(article_status)" description:"状态;1:已发布;0:未发布;-1:删除"`
	IsTop           uint8  `orm:"column(is_top)" description:"是否置顶;1:置顶;0:不置顶"`
	ArticleHits     uint64 `orm:"column(article_hits)" description:"浏览数"`
	CreateTime      uint   `orm:"column(create_time)" description:"创建时间"`
	UpdateTime      uint   `orm:"column(update_time)" description:"更新时间"`
	PublishedTime   uint   `orm:"column(published_time)" description:"发布时间"`
	DeleteTime      uint   `orm:"column(delete_time)" description:"删除时间"`
	ArticleTitle    string `orm:"column(article_title);size(100)" description:"article标题"`
	ArticleKeywords string `orm:"column(article_keywords);size(150)" description:"seo keywords"`
	ArticleContent  string `orm:"column(article_content);null" description:"文章内容"`
	Thumbnail       string `orm:"column(thumbnail);size(255)" description:"缩略图"`
	ArticleDesc     string `orm:"column(article_desc);size(500)" description:"文章描述"`
	Author          string `orm:"column(author);size(100)" description:"作者"`
}

func (t *CmsArticle) TableName() string {
	return "cms_article"
}

func init() {
	orm.RegisterModel(new(CmsArticle))
}

// AddCmsArticle insert a new CmsArticle into database and returns
// last inserted Id on success.
func AddCmsArticle(m *CmsArticle) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCmsArticleById retrieves CmsArticle by Id. Returns error if
// Id doesn't exist
func GetCmsArticleById(id int) (v *CmsArticle, err error) {
	o := orm.NewOrm()
	o.Using("tangka")
	v = &CmsArticle{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCmsArticle retrieves all CmsArticle matches certain condition. Returns empty list if
// no records exist
func GetAllCmsArticle(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CmsArticle))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []CmsArticle
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateCmsArticle updates CmsArticle by Id and returns error if
// the record to be updated doesn't exist
func UpdateCmsArticleById(m *CmsArticle) (err error) {
	o := orm.NewOrm()
	v := CmsArticle{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCmsArticle deletes CmsArticle by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCmsArticle(id int) (err error) {
	o := orm.NewOrm()
	v := CmsArticle{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CmsArticle{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
