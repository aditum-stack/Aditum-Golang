package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type EmailInfo struct {
	Id               int64
	EmailId          string `orm:"size(128)"`
	EmailTitle       string `orm:"size(128)"`
	EmailContent     string `orm:"size(128)"`
	SenderName       string `orm:"size(128)"`
	SenderAddress    string `orm:"size(128)"`
	RecipientName    string `orm:"size(128)"`
	RecipientAddress string `orm:"size(128)"`
	EmailCreateTime  string `orm:"size(128)"`
	EmailIsDeleted   int
}

func init() {
	orm.RegisterModel(new(EmailInfo))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:wangshihao@tcp(47.100.236.64:3306)/golang-email?charset=utf8")
	orm.SetMaxIdleConns("default", 1000)
	orm.SetMaxOpenConns("default", 2000)
}

// AddEmailInfo insert a new EmailInfo into database and returns
// last inserted Id on success.
func AddEmailInfo(m *EmailInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEmailInfoById retrieves EmailInfo by Id. Returns error if
// Id doesn't exist
func GetEmailInfoById(id int64) (v *EmailInfo, err error) {
	o := orm.NewOrm()
	v = &EmailInfo{Id: id}
	if err = o.QueryTable(new(EmailInfo)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEmailInfo retrieves all EmailInfo matches certain condition. Returns empty list if
// no records exist
func GetAllEmailInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EmailInfo))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []EmailInfo
	qs = qs.OrderBy(sortFields...).RelatedSel()
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

// UpdateEmailInfo updates EmailInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateEmailInfoById(m *EmailInfo) (err error) {
	o := orm.NewOrm()
	v := EmailInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEmailInfo deletes EmailInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEmailInfo(id int64) (err error) {
	o := orm.NewOrm()
	v := EmailInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EmailInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
