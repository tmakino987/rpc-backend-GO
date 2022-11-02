package util

import (
	"database/sql"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB interface {
	Close() error
	Create(value interface{}) DB
	First(dest interface{}, conds ...interface{}) DB
	FirstOrCreate(dest interface{}, conds ...interface{}) DB
	Last(dest interface{}, conds ...interface{}) DB
	Find(dest interface{}, conds ...interface{}) DB
	Update(target interface{}, value interface{}) DB
	Delete(value interface{}, conds ...interface{}) DB
	Where(query interface{}, args ...interface{}) DB
	Joins(query string, args ...interface{}) DB
	Select(query interface{}, args ...interface{}) DB
	Table(name string, args ...interface{}) DB
	Group(name string) DB
	Having(query interface{}, args ...interface{}) DB
	Raw(sql string, values ...interface{}) DB
	Scan(dest interface{}) DB
	Preload(query string, args ...interface{}) DB
	Order(value interface{}) DB
	Rows() (*Rows, error)
	Exec(sql string, values ...interface{}) DB
	ScanRows(rows *Rows, dest interface{}) error
	Begin() DB
	Rollback() DB
	Commit() DB
	RowsAffected() int64
	Distinct(args ...interface{}) DB
	Limit(limit int) DB
	Count(count *int64) DB
	Error() error
}

type db struct {
	Conn *gorm.DB
}

// NewDB DBと接続する
func NewDB(dsn string) DB {
	conn := newGormDB(dsn)
	return &db{Conn: conn}
}

func newGormDB(dsn string) *gorm.DB {
	
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Info("DB接続失敗", err)
		panic(err)
	}

	return db
}

func (d *db) Close() error {
	connDB, _ := d.Conn.DB()
	return connDB.Close()
}

func (d *db) Create(value interface{}) DB {
	return &db{Conn: d.Conn.Create(value)}
}

func (d *db) First(dest interface{}, conds ...interface{}) DB {
	return &db{Conn: d.Conn.First(dest, conds...)}
}

func (d *db) FirstOrCreate(dest interface{}, conds ...interface{}) DB {
	return &db{Conn: d.Conn.FirstOrCreate(dest, conds...)}
}

func (d *db) Last(dest interface{}, conds ...interface{}) DB {
	return &db{Conn: d.Conn.Last(dest, conds...)}
}

func (d *db) Find(dest interface{}, conds ...interface{}) DB {
	return &db{Conn: d.Conn.Find(dest, conds...)}
}

func (d *db) Update(target interface{}, value interface{}) DB {
	return &db{Conn: d.Conn.Model(target).Updates(value)}
}

func (d *db) Delete(value interface{}, conds ...interface{}) DB {
	return &db{Conn: d.Conn.Delete(value, conds...)}
}

func (d *db) Where(query interface{}, args ...interface{}) DB {
	return &db{Conn: d.Conn.Where(query, args...)}
}

func (d *db) Joins(query string, args ...interface{}) DB {
	return &db{Conn: d.Conn.Joins(query, args...)}
}

func (d *db) Select(query interface{}, args ...interface{}) DB {
	return &db{Conn: d.Conn.Select(query, args...)}
}

func (d *db) Table(name string, args ...interface{}) DB {
	return &db{Conn: d.Conn.Table(name, args...)}
}

func (d *db) Group(name string) DB {
	return &db{Conn: d.Conn.Group(name)}
}

func (d *db) Having(query interface{}, args ...interface{}) DB {
	return &db{Conn: d.Conn.Having(query, args...)}
}

func (d *db) Raw(sql string, values ...interface{}) DB {
	return &db{Conn: d.Conn.Raw(sql, values...)}
}

func (d *db) Exec(sql string, values ...interface{}) DB {
	return &db{Conn: d.Conn.Exec(sql, values...)}
}

func (d *db) Scan(dest interface{}) DB {
	return &db{Conn: d.Conn.Scan(dest)}
}

func (d *db) Preload(query string, args ...interface{}) DB {
	return &db{Conn: d.Conn.Preload(query, args...)}
}

func (d *db) Order(value interface{}) DB {
	return &db{Conn: d.Conn.Order(value)}
}

func (d *db) Begin() DB {
	return &db{d.Conn.Begin()}
}

func (d *db) Rollback() DB {
	return &db{d.Conn.Rollback()}
}

func (d *db) Commit() DB {
	return &db{d.Conn.Commit()}
}

func (d *db) RowsAffected() int64 {
	return d.Conn.RowsAffected
}

func (d *db) Distinct(args ...interface{}) DB {
	return &db{Conn: d.Conn.Distinct(args...)}
}

func (d *db) Count(count *int64) DB {
	return &db{Conn: d.Conn.Count(count)}
}

func (d *db) Limit(limit int) DB {
	return &db{Conn: d.Conn.Limit(limit)}
}

func (d *db) Error() error {
	return d.Conn.Error
}

type IRows interface {
	Next() bool
}

type Rows struct {
	IRows
	*SqlRows
}

type SqlRows struct {
	*sql.Rows
}

func (rs *SqlRows) Next() bool {
	return rs.Rows.Next()
}

func (d *db) Rows() (*Rows, error) {
	rows, err := d.Conn.Rows()
	srs := &SqlRows{rows}
	rs := &Rows{IRows: srs, SqlRows: srs}
	return rs, err
}

func (d *db) ScanRows(rows *Rows, dest interface{}) error {
	return d.Conn.ScanRows(rows.Rows, dest)
}
