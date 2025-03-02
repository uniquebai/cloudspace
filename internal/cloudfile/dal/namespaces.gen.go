// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/yanguiyuan/cloudspace/internal/cloudfile/model"
)

func newNamespace(db *gorm.DB, opts ...gen.DOOption) namespace {
	_namespace := namespace{}

	_namespace.namespaceDo.UseDB(db, opts...)
	_namespace.namespaceDo.UseModel(&model.Namespace{})

	tableName := _namespace.namespaceDo.TableName()
	_namespace.ALL = field.NewAsterisk(tableName)
	_namespace.ID = field.NewInt64(tableName, "id")
	_namespace.Name = field.NewString(tableName, "name")
	_namespace.CreateTime = field.NewTime(tableName, "create_time")
	_namespace.UpdateTime = field.NewTime(tableName, "update_time")

	_namespace.fillFieldMap()

	return _namespace
}

type namespace struct {
	namespaceDo namespaceDo

	ALL        field.Asterisk
	ID         field.Int64
	Name       field.String
	CreateTime field.Time
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (n namespace) Table(newTableName string) *namespace {
	n.namespaceDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n namespace) As(alias string) *namespace {
	n.namespaceDo.DO = *(n.namespaceDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *namespace) updateTableName(table string) *namespace {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewInt64(table, "id")
	n.Name = field.NewString(table, "name")
	n.CreateTime = field.NewTime(table, "create_time")
	n.UpdateTime = field.NewTime(table, "update_time")

	n.fillFieldMap()

	return n
}

func (n *namespace) WithContext(ctx context.Context) INamespaceDo {
	return n.namespaceDo.WithContext(ctx)
}

func (n namespace) TableName() string { return n.namespaceDo.TableName() }

func (n namespace) Alias() string { return n.namespaceDo.Alias() }

func (n namespace) Columns(cols ...field.Expr) gen.Columns { return n.namespaceDo.Columns(cols...) }

func (n *namespace) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *namespace) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 4)
	n.fieldMap["id"] = n.ID
	n.fieldMap["name"] = n.Name
	n.fieldMap["create_time"] = n.CreateTime
	n.fieldMap["update_time"] = n.UpdateTime
}

func (n namespace) clone(db *gorm.DB) namespace {
	n.namespaceDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n namespace) replaceDB(db *gorm.DB) namespace {
	n.namespaceDo.ReplaceDB(db)
	return n
}

type namespaceDo struct{ gen.DO }

type INamespaceDo interface {
	gen.SubQuery
	Debug() INamespaceDo
	WithContext(ctx context.Context) INamespaceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() INamespaceDo
	WriteDB() INamespaceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) INamespaceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INamespaceDo
	Not(conds ...gen.Condition) INamespaceDo
	Or(conds ...gen.Condition) INamespaceDo
	Select(conds ...field.Expr) INamespaceDo
	Where(conds ...gen.Condition) INamespaceDo
	Order(conds ...field.Expr) INamespaceDo
	Distinct(cols ...field.Expr) INamespaceDo
	Omit(cols ...field.Expr) INamespaceDo
	Join(table schema.Tabler, on ...field.Expr) INamespaceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INamespaceDo
	RightJoin(table schema.Tabler, on ...field.Expr) INamespaceDo
	Group(cols ...field.Expr) INamespaceDo
	Having(conds ...gen.Condition) INamespaceDo
	Limit(limit int) INamespaceDo
	Offset(offset int) INamespaceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INamespaceDo
	Unscoped() INamespaceDo
	Create(values ...*model.Namespace) error
	CreateInBatches(values []*model.Namespace, batchSize int) error
	Save(values ...*model.Namespace) error
	First() (*model.Namespace, error)
	Take() (*model.Namespace, error)
	Last() (*model.Namespace, error)
	Find() ([]*model.Namespace, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Namespace, err error)
	FindInBatches(result *[]*model.Namespace, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Namespace) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INamespaceDo
	Assign(attrs ...field.AssignExpr) INamespaceDo
	Joins(fields ...field.RelationField) INamespaceDo
	Preload(fields ...field.RelationField) INamespaceDo
	FirstOrInit() (*model.Namespace, error)
	FirstOrCreate() (*model.Namespace, error)
	FindByPage(offset int, limit int) (result []*model.Namespace, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INamespaceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n namespaceDo) Debug() INamespaceDo {
	return n.withDO(n.DO.Debug())
}

func (n namespaceDo) WithContext(ctx context.Context) INamespaceDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n namespaceDo) ReadDB() INamespaceDo {
	return n.Clauses(dbresolver.Read)
}

func (n namespaceDo) WriteDB() INamespaceDo {
	return n.Clauses(dbresolver.Write)
}

func (n namespaceDo) Session(config *gorm.Session) INamespaceDo {
	return n.withDO(n.DO.Session(config))
}

func (n namespaceDo) Clauses(conds ...clause.Expression) INamespaceDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n namespaceDo) Returning(value interface{}, columns ...string) INamespaceDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n namespaceDo) Not(conds ...gen.Condition) INamespaceDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n namespaceDo) Or(conds ...gen.Condition) INamespaceDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n namespaceDo) Select(conds ...field.Expr) INamespaceDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n namespaceDo) Where(conds ...gen.Condition) INamespaceDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n namespaceDo) Order(conds ...field.Expr) INamespaceDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n namespaceDo) Distinct(cols ...field.Expr) INamespaceDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n namespaceDo) Omit(cols ...field.Expr) INamespaceDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n namespaceDo) Join(table schema.Tabler, on ...field.Expr) INamespaceDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n namespaceDo) LeftJoin(table schema.Tabler, on ...field.Expr) INamespaceDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n namespaceDo) RightJoin(table schema.Tabler, on ...field.Expr) INamespaceDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n namespaceDo) Group(cols ...field.Expr) INamespaceDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n namespaceDo) Having(conds ...gen.Condition) INamespaceDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n namespaceDo) Limit(limit int) INamespaceDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n namespaceDo) Offset(offset int) INamespaceDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n namespaceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INamespaceDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n namespaceDo) Unscoped() INamespaceDo {
	return n.withDO(n.DO.Unscoped())
}

func (n namespaceDo) Create(values ...*model.Namespace) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n namespaceDo) CreateInBatches(values []*model.Namespace, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n namespaceDo) Save(values ...*model.Namespace) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n namespaceDo) First() (*model.Namespace, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Namespace), nil
	}
}

func (n namespaceDo) Take() (*model.Namespace, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Namespace), nil
	}
}

func (n namespaceDo) Last() (*model.Namespace, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Namespace), nil
	}
}

func (n namespaceDo) Find() ([]*model.Namespace, error) {
	result, err := n.DO.Find()
	return result.([]*model.Namespace), err
}

func (n namespaceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Namespace, err error) {
	buf := make([]*model.Namespace, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n namespaceDo) FindInBatches(result *[]*model.Namespace, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n namespaceDo) Attrs(attrs ...field.AssignExpr) INamespaceDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n namespaceDo) Assign(attrs ...field.AssignExpr) INamespaceDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n namespaceDo) Joins(fields ...field.RelationField) INamespaceDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n namespaceDo) Preload(fields ...field.RelationField) INamespaceDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n namespaceDo) FirstOrInit() (*model.Namespace, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Namespace), nil
	}
}

func (n namespaceDo) FirstOrCreate() (*model.Namespace, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Namespace), nil
	}
}

func (n namespaceDo) FindByPage(offset int, limit int) (result []*model.Namespace, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n namespaceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n namespaceDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n namespaceDo) Delete(models ...*model.Namespace) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *namespaceDo) withDO(do gen.Dao) *namespaceDo {
	n.DO = *do.(*gen.DO)
	return n
}
