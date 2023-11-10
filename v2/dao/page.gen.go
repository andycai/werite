// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/andycai/werite/v2/model"
)

func newPage(db *gorm.DB, opts ...gen.DOOption) page {
	_page := page{}

	_page.pageDo.UseDB(db, opts...)
	_page.pageDo.UseModel(&model.Page{})

	tableName := _page.pageDo.TableName()
	_page.ALL = field.NewAsterisk(tableName)
	_page.ID = field.NewInt32(tableName, "id")
	_page.UserID = field.NewInt32(tableName, "user_id")
	_page.Author = field.NewString(tableName, "author")
	_page.Slug = field.NewString(tableName, "slug")
	_page.Title = field.NewString(tableName, "title")
	_page.Count_ = field.NewInt32(tableName, "count")
	_page.Content = field.NewString(tableName, "content")
	_page.Tags = field.NewString(tableName, "tags")
	_page.IsDraft = field.NewInt32(tableName, "is_draft")
	_page.DeletedAt = field.NewField(tableName, "deleted_at")
	_page.UpdatedAt = field.NewTime(tableName, "updated_at")
	_page.CreatedAt = field.NewTime(tableName, "created_at")

	_page.fillFieldMap()

	return _page
}

type page struct {
	pageDo

	ALL       field.Asterisk
	ID        field.Int32
	UserID    field.Int32
	Author    field.String
	Slug      field.String
	Title     field.String
	Count_    field.Int32
	Content   field.String
	Tags      field.String
	IsDraft   field.Int32
	DeletedAt field.Field
	UpdatedAt field.Time
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (p page) Table(newTableName string) *page {
	p.pageDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p page) As(alias string) *page {
	p.pageDo.DO = *(p.pageDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *page) updateTableName(table string) *page {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.UserID = field.NewInt32(table, "user_id")
	p.Author = field.NewString(table, "author")
	p.Slug = field.NewString(table, "slug")
	p.Title = field.NewString(table, "title")
	p.Count_ = field.NewInt32(table, "count")
	p.Content = field.NewString(table, "content")
	p.Tags = field.NewString(table, "tags")
	p.IsDraft = field.NewInt32(table, "is_draft")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.CreatedAt = field.NewTime(table, "created_at")

	p.fillFieldMap()

	return p
}

func (p *page) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *page) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["author"] = p.Author
	p.fieldMap["slug"] = p.Slug
	p.fieldMap["title"] = p.Title
	p.fieldMap["count"] = p.Count_
	p.fieldMap["content"] = p.Content
	p.fieldMap["tags"] = p.Tags
	p.fieldMap["is_draft"] = p.IsDraft
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["created_at"] = p.CreatedAt
}

func (p page) clone(db *gorm.DB) page {
	p.pageDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p page) replaceDB(db *gorm.DB) page {
	p.pageDo.ReplaceDB(db)
	return p
}

type pageDo struct{ gen.DO }

type IPageDo interface {
	gen.SubQuery
	Debug() IPageDo
	WithContext(ctx context.Context) IPageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPageDo
	WriteDB() IPageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPageDo
	Not(conds ...gen.Condition) IPageDo
	Or(conds ...gen.Condition) IPageDo
	Select(conds ...field.Expr) IPageDo
	Where(conds ...gen.Condition) IPageDo
	Order(conds ...field.Expr) IPageDo
	Distinct(cols ...field.Expr) IPageDo
	Omit(cols ...field.Expr) IPageDo
	Join(table schema.Tabler, on ...field.Expr) IPageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPageDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPageDo
	Group(cols ...field.Expr) IPageDo
	Having(conds ...gen.Condition) IPageDo
	Limit(limit int) IPageDo
	Offset(offset int) IPageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPageDo
	Unscoped() IPageDo
	Create(values ...*model.Page) error
	CreateInBatches(values []*model.Page, batchSize int) error
	Save(values ...*model.Page) error
	First() (*model.Page, error)
	Take() (*model.Page, error)
	Last() (*model.Page, error)
	Find() ([]*model.Page, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Page, err error)
	FindInBatches(result *[]*model.Page, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Page) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPageDo
	Assign(attrs ...field.AssignExpr) IPageDo
	Joins(fields ...field.RelationField) IPageDo
	Preload(fields ...field.RelationField) IPageDo
	FirstOrInit() (*model.Page, error)
	FirstOrCreate() (*model.Page, error)
	FindByPage(offset int, limit int) (result []*model.Page, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pageDo) Debug() IPageDo {
	return p.withDO(p.DO.Debug())
}

func (p pageDo) WithContext(ctx context.Context) IPageDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pageDo) ReadDB() IPageDo {
	return p.Clauses(dbresolver.Read)
}

func (p pageDo) WriteDB() IPageDo {
	return p.Clauses(dbresolver.Write)
}

func (p pageDo) Session(config *gorm.Session) IPageDo {
	return p.withDO(p.DO.Session(config))
}

func (p pageDo) Clauses(conds ...clause.Expression) IPageDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pageDo) Returning(value interface{}, columns ...string) IPageDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pageDo) Not(conds ...gen.Condition) IPageDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pageDo) Or(conds ...gen.Condition) IPageDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pageDo) Select(conds ...field.Expr) IPageDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pageDo) Where(conds ...gen.Condition) IPageDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pageDo) Order(conds ...field.Expr) IPageDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pageDo) Distinct(cols ...field.Expr) IPageDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pageDo) Omit(cols ...field.Expr) IPageDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pageDo) Join(table schema.Tabler, on ...field.Expr) IPageDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pageDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPageDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pageDo) RightJoin(table schema.Tabler, on ...field.Expr) IPageDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pageDo) Group(cols ...field.Expr) IPageDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pageDo) Having(conds ...gen.Condition) IPageDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pageDo) Limit(limit int) IPageDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pageDo) Offset(offset int) IPageDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPageDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pageDo) Unscoped() IPageDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pageDo) Create(values ...*model.Page) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pageDo) CreateInBatches(values []*model.Page, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pageDo) Save(values ...*model.Page) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pageDo) First() (*model.Page, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Page), nil
	}
}

func (p pageDo) Take() (*model.Page, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Page), nil
	}
}

func (p pageDo) Last() (*model.Page, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Page), nil
	}
}

func (p pageDo) Find() ([]*model.Page, error) {
	result, err := p.DO.Find()
	return result.([]*model.Page), err
}

func (p pageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Page, err error) {
	buf := make([]*model.Page, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pageDo) FindInBatches(result *[]*model.Page, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pageDo) Attrs(attrs ...field.AssignExpr) IPageDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pageDo) Assign(attrs ...field.AssignExpr) IPageDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pageDo) Joins(fields ...field.RelationField) IPageDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pageDo) Preload(fields ...field.RelationField) IPageDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pageDo) FirstOrInit() (*model.Page, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Page), nil
	}
}

func (p pageDo) FirstOrCreate() (*model.Page, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Page), nil
	}
}

func (p pageDo) FindByPage(offset int, limit int) (result []*model.Page, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pageDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pageDo) Delete(models ...*model.Page) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pageDo) withDO(do gen.Dao) *pageDo {
	p.DO = *do.(*gen.DO)
	return p
}
