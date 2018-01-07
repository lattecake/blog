package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
)

type Post struct {
	Id          int64     `orm:"auto;unique;column(id)"`
	Title       string    `orm:"index;size(255);column(title)"`
	Description string    `orm:"null;size(1024);column(description)"`
	Content     string    `orm:"size(65535);column(content);type(text)"`
	IsMarkdown  int       `orm:"null;default(0);column(is_rarkdown)"`
	Status      int       `orm:"default(0);column(status)"`
	ReadNum     int64     `orm:"default(0);column(read_num)"`
	Reviews     int64     `orm:"default(0);column(reviews)"`
	PushTime    time.Time `orm:"null;column(push_time);auto_now;type(datetime)"`
	CreatedAt   time.Time `orm:"null;column(created_at);auto_now_add;type(datetime)"`
	Action      int       `orm:"column(action)"`

	Images []*Image `orm:"-;reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Post))
}

// 表名
func (u *Post) TableName() string {
	return "posts"
}

// 引擎
func (u *Post) TableEngine() string {
	return "INNODB"
}

// 多字段索引
func (u *Post) TableIndex() [][]string {
	return [][]string{
		[]string{"Title"},
	}
}

// 多字段唯一键
func (u *Post) TableUnique() [][]string {
	return [][]string{
		[]string{},
	}
}

// GetPostById retrieves Post by Id. Returns error if
// Id doesn't exist
func GetPostById(id int64) (v *Post, err error) {
	o := orm.NewOrm()
	v = &Post{Id: id}
	if err = o.QueryTable(new(Post)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPost retrieves all Post matches certain condition. Returns empty list if
// no records exist
func GetAllPost(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Post))
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

	var l []Post
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

// UpdatePost updates Post by Id and returns error if
// the record to be updated doesn't exist
func UpdatePostById(m *Post) (err error) {
	o := orm.NewOrm()
	v := Post{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func PostPopular(num int) ([]*Post, error) {
	o := orm.NewOrm()

	user := new(Post)
	qs := o.QueryTable(user) // 返回 QuerySeter

	var posts []*Post

	_, err := qs.OrderBy("-ReadNum").Limit(num).All(&posts)
	if err != nil {
		return posts, err
	}

	return posts, nil

}
