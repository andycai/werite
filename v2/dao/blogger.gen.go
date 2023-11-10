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

func newBlogger(db *gorm.DB, opts ...gen.DOOption) blogger {
	_blogger := blogger{}

	_blogger.bloggerDo.UseDB(db, opts...)
	_blogger.bloggerDo.UseModel(&model.Blogger{})

	tableName := _blogger.bloggerDo.TableName()
	_blogger.ALL = field.NewAsterisk(tableName)
	_blogger.BlogName = field.NewString(tableName, "blog_name")
	_blogger.SubTitle = field.NewString(tableName, "sub_title")
	_blogger.BeiAn = field.NewString(tableName, "bei_an")
	_blogger.BTitle = field.NewString(tableName, "b_title")
	_blogger.Copyright = field.NewString(tableName, "copyright")
	_blogger.SeriesSay = field.NewString(tableName, "series_say")
	_blogger.ArchivesSay = field.NewString(tableName, "archives_say")

	_blogger.fillFieldMap()

	return _blogger
}

type blogger struct {
	bloggerDo

	ALL         field.Asterisk
	BlogName    field.String
	SubTitle    field.String
	BeiAn       field.String
	BTitle      field.String
	Copyright   field.String
	SeriesSay   field.String
	ArchivesSay field.String

	fieldMap map[string]field.Expr
}

func (b blogger) Table(newTableName string) *blogger {
	b.bloggerDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b blogger) As(alias string) *blogger {
	b.bloggerDo.DO = *(b.bloggerDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *blogger) updateTableName(table string) *blogger {
	b.ALL = field.NewAsterisk(table)
	b.BlogName = field.NewString(table, "blog_name")
	b.SubTitle = field.NewString(table, "sub_title")
	b.BeiAn = field.NewString(table, "bei_an")
	b.BTitle = field.NewString(table, "b_title")
	b.Copyright = field.NewString(table, "copyright")
	b.SeriesSay = field.NewString(table, "series_say")
	b.ArchivesSay = field.NewString(table, "archives_say")

	b.fillFieldMap()

	return b
}

func (b *blogger) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *blogger) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 7)
	b.fieldMap["blog_name"] = b.BlogName
	b.fieldMap["sub_title"] = b.SubTitle
	b.fieldMap["bei_an"] = b.BeiAn
	b.fieldMap["b_title"] = b.BTitle
	b.fieldMap["copyright"] = b.Copyright
	b.fieldMap["series_say"] = b.SeriesSay
	b.fieldMap["archives_say"] = b.ArchivesSay
}

func (b blogger) clone(db *gorm.DB) blogger {
	b.bloggerDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b blogger) replaceDB(db *gorm.DB) blogger {
	b.bloggerDo.ReplaceDB(db)
	return b
}

type bloggerDo struct{ gen.DO }

type IBloggerDo interface {
	gen.SubQuery
	Debug() IBloggerDo
	WithContext(ctx context.Context) IBloggerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBloggerDo
	WriteDB() IBloggerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBloggerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBloggerDo
	Not(conds ...gen.Condition) IBloggerDo
	Or(conds ...gen.Condition) IBloggerDo
	Select(conds ...field.Expr) IBloggerDo
	Where(conds ...gen.Condition) IBloggerDo
	Order(conds ...field.Expr) IBloggerDo
	Distinct(cols ...field.Expr) IBloggerDo
	Omit(cols ...field.Expr) IBloggerDo
	Join(table schema.Tabler, on ...field.Expr) IBloggerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBloggerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBloggerDo
	Group(cols ...field.Expr) IBloggerDo
	Having(conds ...gen.Condition) IBloggerDo
	Limit(limit int) IBloggerDo
	Offset(offset int) IBloggerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBloggerDo
	Unscoped() IBloggerDo
	Create(values ...*model.Blogger) error
	CreateInBatches(values []*model.Blogger, batchSize int) error
	Save(values ...*model.Blogger) error
	First() (*model.Blogger, error)
	Take() (*model.Blogger, error)
	Last() (*model.Blogger, error)
	Find() ([]*model.Blogger, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Blogger, err error)
	FindInBatches(result *[]*model.Blogger, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Blogger) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBloggerDo
	Assign(attrs ...field.AssignExpr) IBloggerDo
	Joins(fields ...field.RelationField) IBloggerDo
	Preload(fields ...field.RelationField) IBloggerDo
	FirstOrInit() (*model.Blogger, error)
	FirstOrCreate() (*model.Blogger, error)
	FindByPage(offset int, limit int) (result []*model.Blogger, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBloggerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b bloggerDo) Debug() IBloggerDo {
	return b.withDO(b.DO.Debug())
}

func (b bloggerDo) WithContext(ctx context.Context) IBloggerDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bloggerDo) ReadDB() IBloggerDo {
	return b.Clauses(dbresolver.Read)
}

func (b bloggerDo) WriteDB() IBloggerDo {
	return b.Clauses(dbresolver.Write)
}

func (b bloggerDo) Session(config *gorm.Session) IBloggerDo {
	return b.withDO(b.DO.Session(config))
}

func (b bloggerDo) Clauses(conds ...clause.Expression) IBloggerDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bloggerDo) Returning(value interface{}, columns ...string) IBloggerDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bloggerDo) Not(conds ...gen.Condition) IBloggerDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bloggerDo) Or(conds ...gen.Condition) IBloggerDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bloggerDo) Select(conds ...field.Expr) IBloggerDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bloggerDo) Where(conds ...gen.Condition) IBloggerDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bloggerDo) Order(conds ...field.Expr) IBloggerDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bloggerDo) Distinct(cols ...field.Expr) IBloggerDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bloggerDo) Omit(cols ...field.Expr) IBloggerDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bloggerDo) Join(table schema.Tabler, on ...field.Expr) IBloggerDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bloggerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBloggerDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bloggerDo) RightJoin(table schema.Tabler, on ...field.Expr) IBloggerDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bloggerDo) Group(cols ...field.Expr) IBloggerDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bloggerDo) Having(conds ...gen.Condition) IBloggerDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bloggerDo) Limit(limit int) IBloggerDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bloggerDo) Offset(offset int) IBloggerDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bloggerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBloggerDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bloggerDo) Unscoped() IBloggerDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bloggerDo) Create(values ...*model.Blogger) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bloggerDo) CreateInBatches(values []*model.Blogger, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bloggerDo) Save(values ...*model.Blogger) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bloggerDo) First() (*model.Blogger, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Blogger), nil
	}
}

func (b bloggerDo) Take() (*model.Blogger, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Blogger), nil
	}
}

func (b bloggerDo) Last() (*model.Blogger, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Blogger), nil
	}
}

func (b bloggerDo) Find() ([]*model.Blogger, error) {
	result, err := b.DO.Find()
	return result.([]*model.Blogger), err
}

func (b bloggerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Blogger, err error) {
	buf := make([]*model.Blogger, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bloggerDo) FindInBatches(result *[]*model.Blogger, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bloggerDo) Attrs(attrs ...field.AssignExpr) IBloggerDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bloggerDo) Assign(attrs ...field.AssignExpr) IBloggerDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bloggerDo) Joins(fields ...field.RelationField) IBloggerDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bloggerDo) Preload(fields ...field.RelationField) IBloggerDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bloggerDo) FirstOrInit() (*model.Blogger, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Blogger), nil
	}
}

func (b bloggerDo) FirstOrCreate() (*model.Blogger, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Blogger), nil
	}
}

func (b bloggerDo) FindByPage(offset int, limit int) (result []*model.Blogger, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bloggerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bloggerDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bloggerDo) Delete(models ...*model.Blogger) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bloggerDo) withDO(do gen.Dao) *bloggerDo {
	b.DO = *do.(*gen.DO)
	return b
}
