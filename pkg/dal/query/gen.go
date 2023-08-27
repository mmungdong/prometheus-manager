// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                    = new(Query)
	PromAlarmHistory     *promAlarmHistory
	PromAlarmPage        *promAlarmPage
	PromAlarmPageHistory *promAlarmPageHistory
	PromDict             *promDict
	PromGroup            *promGroup
	PromStrategy         *promStrategy
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	PromAlarmHistory = &Q.PromAlarmHistory
	PromAlarmPage = &Q.PromAlarmPage
	PromAlarmPageHistory = &Q.PromAlarmPageHistory
	PromDict = &Q.PromDict
	PromGroup = &Q.PromGroup
	PromStrategy = &Q.PromStrategy
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                   db,
		PromAlarmHistory:     newPromAlarmHistory(db, opts...),
		PromAlarmPage:        newPromAlarmPage(db, opts...),
		PromAlarmPageHistory: newPromAlarmPageHistory(db, opts...),
		PromDict:             newPromDict(db, opts...),
		PromGroup:            newPromGroup(db, opts...),
		PromStrategy:         newPromStrategy(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	PromAlarmHistory     promAlarmHistory
	PromAlarmPage        promAlarmPage
	PromAlarmPageHistory promAlarmPageHistory
	PromDict             promDict
	PromGroup            promGroup
	PromStrategy         promStrategy
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		PromAlarmHistory:     q.PromAlarmHistory.clone(db),
		PromAlarmPage:        q.PromAlarmPage.clone(db),
		PromAlarmPageHistory: q.PromAlarmPageHistory.clone(db),
		PromDict:             q.PromDict.clone(db),
		PromGroup:            q.PromGroup.clone(db),
		PromStrategy:         q.PromStrategy.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		PromAlarmHistory:     q.PromAlarmHistory.replaceDB(db),
		PromAlarmPage:        q.PromAlarmPage.replaceDB(db),
		PromAlarmPageHistory: q.PromAlarmPageHistory.replaceDB(db),
		PromDict:             q.PromDict.replaceDB(db),
		PromGroup:            q.PromGroup.replaceDB(db),
		PromStrategy:         q.PromStrategy.replaceDB(db),
	}
}

type queryCtx struct {
	PromAlarmHistory     IPromAlarmHistoryDo
	PromAlarmPage        IPromAlarmPageDo
	PromAlarmPageHistory IPromAlarmPageHistoryDo
	PromDict             IPromDictDo
	PromGroup            IPromGroupDo
	PromStrategy         IPromStrategyDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		PromAlarmHistory:     q.PromAlarmHistory.WithContext(ctx),
		PromAlarmPage:        q.PromAlarmPage.WithContext(ctx),
		PromAlarmPageHistory: q.PromAlarmPageHistory.WithContext(ctx),
		PromDict:             q.PromDict.WithContext(ctx),
		PromGroup:            q.PromGroup.WithContext(ctx),
		PromStrategy:         q.PromStrategy.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
