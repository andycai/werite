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

func newSerie(db *gorm.DB, opts ...gen.DOOption) serie {
	_serie := serie{}

	_serie.serieDo.UseDB(db, opts...)
	_serie.serieDo.UseModel(&model.Serie{})

	tableName := _serie.serieDo.TableName()
	_serie.ALL = field.NewAsterisk(tableName)
	_serie.ID = field.NewInt32(tableName, "id")
	_serie.Slug = field.NewString(tableName, "slug")
	_serie.Name = field.NewString(tableName, "name")
	_serie.Desc = field.NewString(tableName, "desc")
	_serie.CreatedAt = field.NewTime(tableName, "created_at")

	_serie.fillFieldMap()

	return _serie
}

type serie struct {
	serieDo

	ALL       field.Asterisk
	ID        field.Int32
	Slug      field.String
	Name      field.String
	Desc      field.String
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (s serie) Table(newTableName string) *serie {
	s.serieDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s serie) As(alias string) *serie {
	s.serieDo.DO = *(s.serieDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *serie) updateTableName(table string) *serie {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.Slug = field.NewString(table, "slug")
	s.Name = field.NewString(table, "name")
	s.Desc = field.NewString(table, "desc")
	s.CreatedAt = field.NewTime(table, "created_at")

	s.fillFieldMap()

	return s
}

func (s *serie) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *serie) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["id"] = s.ID
	s.fieldMap["slug"] = s.Slug
	s.fieldMap["name"] = s.Name
	s.fieldMap["desc"] = s.Desc
	s.fieldMap["created_at"] = s.CreatedAt
}

func (s serie) clone(db *gorm.DB) serie {
	s.serieDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s serie) replaceDB(db *gorm.DB) serie {
	s.serieDo.ReplaceDB(db)
	return s
}

type serieDo struct{ gen.DO }

type ISerieDo interface {
	gen.SubQuery
	Debug() ISerieDo
	WithContext(ctx context.Context) ISerieDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISerieDo
	WriteDB() ISerieDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISerieDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISerieDo
	Not(conds ...gen.Condition) ISerieDo
	Or(conds ...gen.Condition) ISerieDo
	Select(conds ...field.Expr) ISerieDo
	Where(conds ...gen.Condition) ISerieDo
	Order(conds ...field.Expr) ISerieDo
	Distinct(cols ...field.Expr) ISerieDo
	Omit(cols ...field.Expr) ISerieDo
	Join(table schema.Tabler, on ...field.Expr) ISerieDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISerieDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISerieDo
	Group(cols ...field.Expr) ISerieDo
	Having(conds ...gen.Condition) ISerieDo
	Limit(limit int) ISerieDo
	Offset(offset int) ISerieDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISerieDo
	Unscoped() ISerieDo
	Create(values ...*model.Serie) error
	CreateInBatches(values []*model.Serie, batchSize int) error
	Save(values ...*model.Serie) error
	First() (*model.Serie, error)
	Take() (*model.Serie, error)
	Last() (*model.Serie, error)
	Find() ([]*model.Serie, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Serie, err error)
	FindInBatches(result *[]*model.Serie, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Serie) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISerieDo
	Assign(attrs ...field.AssignExpr) ISerieDo
	Joins(fields ...field.RelationField) ISerieDo
	Preload(fields ...field.RelationField) ISerieDo
	FirstOrInit() (*model.Serie, error)
	FirstOrCreate() (*model.Serie, error)
	FindByPage(offset int, limit int) (result []*model.Serie, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISerieDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s serieDo) Debug() ISerieDo {
	return s.withDO(s.DO.Debug())
}

func (s serieDo) WithContext(ctx context.Context) ISerieDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s serieDo) ReadDB() ISerieDo {
	return s.Clauses(dbresolver.Read)
}

func (s serieDo) WriteDB() ISerieDo {
	return s.Clauses(dbresolver.Write)
}

func (s serieDo) Session(config *gorm.Session) ISerieDo {
	return s.withDO(s.DO.Session(config))
}

func (s serieDo) Clauses(conds ...clause.Expression) ISerieDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s serieDo) Returning(value interface{}, columns ...string) ISerieDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s serieDo) Not(conds ...gen.Condition) ISerieDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s serieDo) Or(conds ...gen.Condition) ISerieDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s serieDo) Select(conds ...field.Expr) ISerieDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s serieDo) Where(conds ...gen.Condition) ISerieDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s serieDo) Order(conds ...field.Expr) ISerieDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s serieDo) Distinct(cols ...field.Expr) ISerieDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s serieDo) Omit(cols ...field.Expr) ISerieDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s serieDo) Join(table schema.Tabler, on ...field.Expr) ISerieDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s serieDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISerieDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s serieDo) RightJoin(table schema.Tabler, on ...field.Expr) ISerieDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s serieDo) Group(cols ...field.Expr) ISerieDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s serieDo) Having(conds ...gen.Condition) ISerieDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s serieDo) Limit(limit int) ISerieDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s serieDo) Offset(offset int) ISerieDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s serieDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISerieDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s serieDo) Unscoped() ISerieDo {
	return s.withDO(s.DO.Unscoped())
}

func (s serieDo) Create(values ...*model.Serie) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s serieDo) CreateInBatches(values []*model.Serie, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s serieDo) Save(values ...*model.Serie) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s serieDo) First() (*model.Serie, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Serie), nil
	}
}

func (s serieDo) Take() (*model.Serie, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Serie), nil
	}
}

func (s serieDo) Last() (*model.Serie, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Serie), nil
	}
}

func (s serieDo) Find() ([]*model.Serie, error) {
	result, err := s.DO.Find()
	return result.([]*model.Serie), err
}

func (s serieDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Serie, err error) {
	buf := make([]*model.Serie, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s serieDo) FindInBatches(result *[]*model.Serie, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s serieDo) Attrs(attrs ...field.AssignExpr) ISerieDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s serieDo) Assign(attrs ...field.AssignExpr) ISerieDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s serieDo) Joins(fields ...field.RelationField) ISerieDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s serieDo) Preload(fields ...field.RelationField) ISerieDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s serieDo) FirstOrInit() (*model.Serie, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Serie), nil
	}
}

func (s serieDo) FirstOrCreate() (*model.Serie, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Serie), nil
	}
}

func (s serieDo) FindByPage(offset int, limit int) (result []*model.Serie, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s serieDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s serieDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s serieDo) Delete(models ...*model.Serie) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *serieDo) withDO(do gen.Dao) *serieDo {
	s.DO = *do.(*gen.DO)
	return s
}
