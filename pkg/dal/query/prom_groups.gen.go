// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"prometheus-manager/pkg/dal/model"
)

func newPromGroup(db *gorm.DB, opts ...gen.DOOption) promGroup {
	_promGroup := promGroup{}

	_promGroup.promGroupDo.UseDB(db, opts...)
	_promGroup.promGroupDo.UseModel(&model.PromGroup{})

	tableName := _promGroup.promGroupDo.TableName()
	_promGroup.ALL = field.NewAsterisk(tableName)
	_promGroup.ID = field.NewInt32(tableName, "id")
	_promGroup.Name = field.NewString(tableName, "name")
	_promGroup.StrategyCount = field.NewInt64(tableName, "strategy_count")
	_promGroup.Status = field.NewInt32(tableName, "status")
	_promGroup.Remark = field.NewString(tableName, "remark")
	_promGroup.CreatedAt = field.NewTime(tableName, "created_at")
	_promGroup.UpdatedAt = field.NewTime(tableName, "updated_at")
	_promGroup.DeletedAt = field.NewField(tableName, "deleted_at")
	_promGroup.PromStrategies = promGroupHasManyPromStrategies{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("PromStrategies", "model.PromStrategy"),
		AlarmPages: struct {
			field.RelationField
			PromStrategies struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("PromStrategies.AlarmPages", "model.PromAlarmPage"),
			PromStrategies: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("PromStrategies.AlarmPages.PromStrategies", "model.PromStrategy"),
			},
		},
		Categories: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("PromStrategies.Categories", "model.PromDict"),
		},
		AlertLevel: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("PromStrategies.AlertLevel", "model.PromDict"),
		},
		GroupInfo: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("PromStrategies.GroupInfo", "model.PromGroup"),
		},
	}

	_promGroup.Categories = promGroupHasManyCategories{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Categories", "model.PromDict"),
	}

	_promGroup.fillFieldMap()

	return _promGroup
}

type promGroup struct {
	promGroupDo

	ALL            field.Asterisk
	ID             field.Int32
	Name           field.String // 规则组名称
	StrategyCount  field.Int64  // 规则数量
	Status         field.Int32  // 启用状态1:启用;2禁用
	Remark         field.String // 描述信息
	CreatedAt      field.Time   // 创建时间
	UpdatedAt      field.Time   // 更新时间
	DeletedAt      field.Field  // 删除时间
	PromStrategies promGroupHasManyPromStrategies

	Categories promGroupHasManyCategories

	fieldMap map[string]field.Expr
}

func (p promGroup) Table(newTableName string) *promGroup {
	p.promGroupDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p promGroup) As(alias string) *promGroup {
	p.promGroupDo.DO = *(p.promGroupDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *promGroup) updateTableName(table string) *promGroup {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.Name = field.NewString(table, "name")
	p.StrategyCount = field.NewInt64(table, "strategy_count")
	p.Status = field.NewInt32(table, "status")
	p.Remark = field.NewString(table, "remark")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *promGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *promGroup) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 10)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["strategy_count"] = p.StrategyCount
	p.fieldMap["status"] = p.Status
	p.fieldMap["remark"] = p.Remark
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt

}

func (p promGroup) clone(db *gorm.DB) promGroup {
	p.promGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p promGroup) replaceDB(db *gorm.DB) promGroup {
	p.promGroupDo.ReplaceDB(db)
	return p
}

type promGroupHasManyPromStrategies struct {
	db *gorm.DB

	field.RelationField

	AlarmPages struct {
		field.RelationField
		PromStrategies struct {
			field.RelationField
		}
	}
	Categories struct {
		field.RelationField
	}
	AlertLevel struct {
		field.RelationField
	}
	GroupInfo struct {
		field.RelationField
	}
}

func (a promGroupHasManyPromStrategies) Where(conds ...field.Expr) *promGroupHasManyPromStrategies {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a promGroupHasManyPromStrategies) WithContext(ctx context.Context) *promGroupHasManyPromStrategies {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a promGroupHasManyPromStrategies) Session(session *gorm.Session) *promGroupHasManyPromStrategies {
	a.db = a.db.Session(session)
	return &a
}

func (a promGroupHasManyPromStrategies) Model(m *model.PromGroup) *promGroupHasManyPromStrategiesTx {
	return &promGroupHasManyPromStrategiesTx{a.db.Model(m).Association(a.Name())}
}

type promGroupHasManyPromStrategiesTx struct{ tx *gorm.Association }

func (a promGroupHasManyPromStrategiesTx) Find() (result []*model.PromStrategy, err error) {
	return result, a.tx.Find(&result)
}

func (a promGroupHasManyPromStrategiesTx) Append(values ...*model.PromStrategy) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a promGroupHasManyPromStrategiesTx) Replace(values ...*model.PromStrategy) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a promGroupHasManyPromStrategiesTx) Delete(values ...*model.PromStrategy) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a promGroupHasManyPromStrategiesTx) Clear() error {
	return a.tx.Clear()
}

func (a promGroupHasManyPromStrategiesTx) Count() int64 {
	return a.tx.Count()
}

type promGroupHasManyCategories struct {
	db *gorm.DB

	field.RelationField
}

func (a promGroupHasManyCategories) Where(conds ...field.Expr) *promGroupHasManyCategories {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a promGroupHasManyCategories) WithContext(ctx context.Context) *promGroupHasManyCategories {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a promGroupHasManyCategories) Session(session *gorm.Session) *promGroupHasManyCategories {
	a.db = a.db.Session(session)
	return &a
}

func (a promGroupHasManyCategories) Model(m *model.PromGroup) *promGroupHasManyCategoriesTx {
	return &promGroupHasManyCategoriesTx{a.db.Model(m).Association(a.Name())}
}

type promGroupHasManyCategoriesTx struct{ tx *gorm.Association }

func (a promGroupHasManyCategoriesTx) Find() (result []*model.PromDict, err error) {
	return result, a.tx.Find(&result)
}

func (a promGroupHasManyCategoriesTx) Append(values ...*model.PromDict) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a promGroupHasManyCategoriesTx) Replace(values ...*model.PromDict) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a promGroupHasManyCategoriesTx) Delete(values ...*model.PromDict) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a promGroupHasManyCategoriesTx) Clear() error {
	return a.tx.Clear()
}

func (a promGroupHasManyCategoriesTx) Count() int64 {
	return a.tx.Count()
}

type promGroupDo struct{ gen.DO }

type IPromGroupDo interface {
	gen.SubQuery
	Debug() IPromGroupDo
	WithContext(ctx context.Context) IPromGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPromGroupDo
	WriteDB() IPromGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPromGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPromGroupDo
	Not(conds ...gen.Condition) IPromGroupDo
	Or(conds ...gen.Condition) IPromGroupDo
	Select(conds ...field.Expr) IPromGroupDo
	Where(conds ...gen.Condition) IPromGroupDo
	Order(conds ...field.Expr) IPromGroupDo
	Distinct(cols ...field.Expr) IPromGroupDo
	Omit(cols ...field.Expr) IPromGroupDo
	Join(table schema.Tabler, on ...field.Expr) IPromGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPromGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPromGroupDo
	Group(cols ...field.Expr) IPromGroupDo
	Having(conds ...gen.Condition) IPromGroupDo
	Limit(limit int) IPromGroupDo
	Offset(offset int) IPromGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPromGroupDo
	Unscoped() IPromGroupDo
	Create(values ...*model.PromGroup) error
	CreateInBatches(values []*model.PromGroup, batchSize int) error
	Save(values ...*model.PromGroup) error
	First() (*model.PromGroup, error)
	Take() (*model.PromGroup, error)
	Last() (*model.PromGroup, error)
	Find() ([]*model.PromGroup, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromGroup, err error)
	FindInBatches(result *[]*model.PromGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PromGroup) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPromGroupDo
	Assign(attrs ...field.AssignExpr) IPromGroupDo
	Joins(fields ...field.RelationField) IPromGroupDo
	Preload(fields ...field.RelationField) IPromGroupDo
	FirstOrInit() (*model.PromGroup, error)
	FirstOrCreate() (*model.PromGroup, error)
	FindByPage(offset int, limit int) (result []*model.PromGroup, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPromGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	SaFindById(ctx context.Context, id int32) (result *model.PromGroup, err error)
}

// select * from @@table where id = @id
func (p promGroupDo) SaFindById(ctx context.Context, id int32) (result *model.PromGroup, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("select * from prom_groups where id = ? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p promGroupDo) Debug() IPromGroupDo {
	return p.withDO(p.DO.Debug())
}

func (p promGroupDo) WithContext(ctx context.Context) IPromGroupDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p promGroupDo) ReadDB() IPromGroupDo {
	return p.Clauses(dbresolver.Read)
}

func (p promGroupDo) WriteDB() IPromGroupDo {
	return p.Clauses(dbresolver.Write)
}

func (p promGroupDo) Session(config *gorm.Session) IPromGroupDo {
	return p.withDO(p.DO.Session(config))
}

func (p promGroupDo) Clauses(conds ...clause.Expression) IPromGroupDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p promGroupDo) Returning(value interface{}, columns ...string) IPromGroupDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p promGroupDo) Not(conds ...gen.Condition) IPromGroupDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p promGroupDo) Or(conds ...gen.Condition) IPromGroupDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p promGroupDo) Select(conds ...field.Expr) IPromGroupDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p promGroupDo) Where(conds ...gen.Condition) IPromGroupDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p promGroupDo) Order(conds ...field.Expr) IPromGroupDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p promGroupDo) Distinct(cols ...field.Expr) IPromGroupDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p promGroupDo) Omit(cols ...field.Expr) IPromGroupDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p promGroupDo) Join(table schema.Tabler, on ...field.Expr) IPromGroupDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p promGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPromGroupDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p promGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) IPromGroupDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p promGroupDo) Group(cols ...field.Expr) IPromGroupDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p promGroupDo) Having(conds ...gen.Condition) IPromGroupDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p promGroupDo) Limit(limit int) IPromGroupDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p promGroupDo) Offset(offset int) IPromGroupDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p promGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPromGroupDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p promGroupDo) Unscoped() IPromGroupDo {
	return p.withDO(p.DO.Unscoped())
}

func (p promGroupDo) Create(values ...*model.PromGroup) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p promGroupDo) CreateInBatches(values []*model.PromGroup, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p promGroupDo) Save(values ...*model.PromGroup) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p promGroupDo) First() (*model.PromGroup, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromGroup), nil
	}
}

func (p promGroupDo) Take() (*model.PromGroup, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromGroup), nil
	}
}

func (p promGroupDo) Last() (*model.PromGroup, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromGroup), nil
	}
}

func (p promGroupDo) Find() ([]*model.PromGroup, error) {
	result, err := p.DO.Find()
	return result.([]*model.PromGroup), err
}

func (p promGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromGroup, err error) {
	buf := make([]*model.PromGroup, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p promGroupDo) FindInBatches(result *[]*model.PromGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p promGroupDo) Attrs(attrs ...field.AssignExpr) IPromGroupDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p promGroupDo) Assign(attrs ...field.AssignExpr) IPromGroupDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p promGroupDo) Joins(fields ...field.RelationField) IPromGroupDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p promGroupDo) Preload(fields ...field.RelationField) IPromGroupDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p promGroupDo) FirstOrInit() (*model.PromGroup, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromGroup), nil
	}
}

func (p promGroupDo) FirstOrCreate() (*model.PromGroup, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromGroup), nil
	}
}

func (p promGroupDo) FindByPage(offset int, limit int) (result []*model.PromGroup, count int64, err error) {
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

func (p promGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p promGroupDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p promGroupDo) Delete(models ...*model.PromGroup) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *promGroupDo) withDO(do gen.Dao) *promGroupDo {
	p.DO = *do.(*gen.DO)
	return p
}