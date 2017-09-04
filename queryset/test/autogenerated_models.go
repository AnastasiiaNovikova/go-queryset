package test

import (
	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set BlogQuerySet

// BlogQuerySet is an queryset type for Blog
type BlogQuerySet struct {
	db *gorm.DB
}

func NewBlogQuerySet(db *gorm.DB) BlogQuerySet {
	return BlogQuerySet{
		db: db,
	}
}

// IDEq is an autogenerated method
func (qs BlogQuerySet) IDEq(ID int) BlogQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id = ?", ID)
	})
	return qs
}

// IDGt is an autogenerated method
func (qs BlogQuerySet) IDGt(ID int) BlogQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id > ?", ID)
	})
	return qs
}

// IDGte is an autogenerated method
func (qs BlogQuerySet) IDGte(ID int) BlogQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id >= ?", ID)
	})
	return qs
}

// IDLt is an autogenerated method
func (qs BlogQuerySet) IDLt(ID int) BlogQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id < ?", ID)
	})
	return qs
}

// IDLte is an autogenerated method
func (qs BlogQuerySet) IDLte(ID int) BlogQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id <= ?", ID)
	})
	return qs
}

// IDNe is an autogenerated method
func (qs BlogQuerySet) IDNe(ID int) BlogQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id != ?", ID)
	})
	return qs
}

// Limit is an autogenerated method
func (qs BlogQuerySet) Limit(limit int) BlogQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Limit(limit)
	})
	return qs
}

// NameEq is an autogenerated method
func (qs BlogQuerySet) NameEq(name string) BlogQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("name = ?", name)
	})
	return qs
}

// NameNe is an autogenerated method
func (qs BlogQuerySet) NameNe(name string) BlogQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("name != ?", name)
	})
	return qs
}

// OrderByID is an autogenerated method
func (qs BlogQuerySet) OrderByID() BlogQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Order("ID")
	})
	return qs
}

// All is used to retieve slice of results
func (qs BlogQuerySet) All(ret *[]Blog) error {
	return qs.db.Find(ret).Error
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs BlogQuerySet) One(ret *Blog) error {
	return qs.db.First(ret).Error
}

// ===== END of query set BlogQuerySet

// ===== BEGIN of query set PostQuerySet

// PostQuerySet is an queryset type for Post
type PostQuerySet struct {
	db *gorm.DB
}

func NewPostQuerySet(db *gorm.DB) PostQuerySet {
	return PostQuerySet{
		db: db,
	}
}

// IDEq is an autogenerated method
func (qs PostQuerySet) IDEq(ID int) PostQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id = ?", ID)
	})
	return qs
}

// IDGt is an autogenerated method
func (qs PostQuerySet) IDGt(ID int) PostQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id > ?", ID)
	})
	return qs
}

// IDGte is an autogenerated method
func (qs PostQuerySet) IDGte(ID int) PostQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id >= ?", ID)
	})
	return qs
}

// IDLt is an autogenerated method
func (qs PostQuerySet) IDLt(ID int) PostQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id < ?", ID)
	})
	return qs
}

// IDLte is an autogenerated method
func (qs PostQuerySet) IDLte(ID int) PostQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id <= ?", ID)
	})
	return qs
}

// IDNe is an autogenerated method
func (qs PostQuerySet) IDNe(ID int) PostQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id != ?", ID)
	})
	return qs
}

// Limit is an autogenerated method
func (qs PostQuerySet) Limit(limit int) PostQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Limit(limit)
	})
	return qs
}

// OrderByID is an autogenerated method
func (qs PostQuerySet) OrderByID() PostQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Order("ID")
	})
	return qs
}

// PreloadBlog is an autogenerated method
func (qs PostQuerySet) PreloadBlog() PostQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Preload("Blog")
	})
	return qs
}

// TitleEq is an autogenerated method
func (qs PostQuerySet) TitleEq(title string) PostQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("title = ?", title)
	})
	return qs
}

// TitleNe is an autogenerated method
func (qs PostQuerySet) TitleNe(title string) PostQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("title != ?", title)
	})
	return qs
}

// All is used to retieve slice of results
func (qs PostQuerySet) All(ret *[]Post) error {
	return qs.db.Find(ret).Error
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs PostQuerySet) One(ret *Post) error {
	return qs.db.First(ret).Error
}

// ===== END of query set PostQuerySet

// ===== BEGIN of query set UserQuerySet

// UserQuerySet is an queryset type for User
type UserQuerySet struct {
	db *gorm.DB
}

func NewUserQuerySet(db *gorm.DB) UserQuerySet {
	return UserQuerySet{
		db: db,
	}
}

// IDEq is an autogenerated method
func (qs UserQuerySet) IDEq(ID int) UserQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id = ?", ID)
	})
	return qs
}

// IDGt is an autogenerated method
func (qs UserQuerySet) IDGt(ID int) UserQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id > ?", ID)
	})
	return qs
}

// IDGte is an autogenerated method
func (qs UserQuerySet) IDGte(ID int) UserQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id >= ?", ID)
	})
	return qs
}

// IDLt is an autogenerated method
func (qs UserQuerySet) IDLt(ID int) UserQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id < ?", ID)
	})
	return qs
}

// IDLte is an autogenerated method
func (qs UserQuerySet) IDLte(ID int) UserQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id <= ?", ID)
	})
	return qs
}

// IDNe is an autogenerated method
func (qs UserQuerySet) IDNe(ID int) UserQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("id != ?", ID)
	})
	return qs
}

// Limit is an autogenerated method
func (qs UserQuerySet) Limit(limit int) UserQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Limit(limit)
	})
	return qs
}

// NameEq is an autogenerated method
func (qs UserQuerySet) NameEq(name string) UserQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("name = ?", name)
	})
	return qs
}

// NameNe is an autogenerated method
func (qs UserQuerySet) NameNe(name string) UserQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Where("name != ?", name)
	})
	return qs
}

// OrderByID is an autogenerated method
func (qs UserQuerySet) OrderByID() UserQuerySet {
	qs.db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Order("ID")
	})
	return qs
}

// All is used to retieve slice of results
func (qs UserQuerySet) All(ret *[]User) error {
	return qs.db.Find(ret).Error
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs UserQuerySet) One(ret *User) error {
	return qs.db.First(ret).Error
}

// ===== END of query set UserQuerySet

// ===== END of all query sets
