package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
)

type Image struct {
	Id                 int64     `orm:"unique;column(id)"`
	ImageName          string    `orm:"size(255);column(image_name)"`
	Extension          string    `orm:"null;size(32);column(extension)"`
	PostId             int64     `orm:"column(post_id);default(0)"`
	ImagePath          string    `orm:"null;size(255);column(image_path)"`
	RealPath           string    `orm:"null;size(255);column(real_path)"`
	ImageTime          time.Time `orm:"null;column(image_time);auto_now_add;type(datetime)"`
	ImageStatus        int       `orm:"null;default(0);column(image_status)"`
	ImageSize          string    `orm:"null;default(0);column(image_size)"`
	Md5                string    `orm:"null;size(255);default(0);column(md5)"`
	ClientOriginalName string    `orm:"null;size(255);default(0);column(client_original_mame)"`

	Post *Post `orm:"-"`
}

func init() {
	orm.RegisterModel(new(Image))
}

// 表名
func (u *Image) TableName() string {
	return "images"
}

// 引擎
func (u *Image) TableEngine() string {
	return "INNODB"
}

// 多字段索引
func (u *Image) TableIndex() [][]string {
	return [][]string{
		[]string{"ImageName"},
	}
}

// 多字段唯一键
func (u *Image) TableUnique() [][]string {
	return [][]string{
		[]string{},
	}
}

// AddImage insert a new Image into database and returns
// last inserted Id on success.
func AddImage(m *Image) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetImageByPostId(postId int64) (v []*Image, err error) {
	o := orm.NewOrm()

	v = []*Image{}
	if _, err = o.QueryTable(new(Image)).Filter("PostId", postId).Limit(1).All(&v); err == nil {
		return
	}
	return v, err
}

func GetImageByPostIds(postIds []int64) (v []*Image, err error) {
	o := orm.NewOrm()
	v = []*Image{}
	if _, err = o.QueryTable(new(Image)).Filter("PostId__in", postIds).GroupBy("PostId").All(&v); err == nil {
		return
	}
	return v, err
}

// GetImageById retrieves Image by Id. Returns error if
// Id doesn't exist
func GetImageById(id int64) (v *Image, err error) {
	o := orm.NewOrm()
	v = &Image{Id: id}
	if err = o.QueryTable(new(Image)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllImage retrieves all Image matches certain condition. Returns empty list if
// no records exist
func GetAllImage(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Image))
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

	var l []Image
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

// UpdateImage updates Image by Id and returns error if
// the record to be updated doesn't exist
func UpdateImageById(m *Image) (err error) {
	o := orm.NewOrm()
	v := Image{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteImage deletes Image by Id and returns error if
// the record to be deleted doesn't exist
func DeleteImage(id int64) (err error) {
	o := orm.NewOrm()
	v := Image{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Image{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetImageByMd5(name string) (v *Image, err error) {
	o := orm.NewOrm()
	v = &Image{Md5: name}
	if err = o.QueryTable(new(Image)).Filter("Md5", name).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

//
func ExistsImageByMd5(name string) (ok bool, err error) {
	o := orm.NewOrm()

	v := Image{}
	if err = o.QueryTable(new(Image)).Filter("Md5", name).One(&v); err == nil {
		return true, nil
	}
	return false, err
}
