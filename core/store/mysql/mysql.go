package mysql

import (
	"fmt"
	"os"

	"go-backend-template/core/logx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Orm struct {
		Host     string
		Port     string
		User     string
		Pass     string
		Database string
		Debug    bool

		Conf *gorm.Config
		*gorm.DB
	}
	Option func(r *Orm)
)

func (r *Orm) GetOrm() *gorm.DB {
	return r.DB
}

func (r *Orm) OrmConnectionUpdate(conf OrmConf) *Orm {
	orm, err := NewMysqlOrm(conf)
	if err != nil {
		return r
	}
	return orm
}

func MustNewMysqlOrm(conf OrmConf, opts ...Option) *Orm {
	orm, err := NewMysqlOrm(conf, opts...)
	if err != nil {
		logx.Errorw("fail to load mysql orm", zap.Field{Key: "err", Type: zapcore.StringType, String: err.Error()})
		os.Exit(1)
	}
	return orm
}

func NewMysqlOrm(conf OrmConf, opts ...Option) (*Orm, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}
	opts = append([]Option{WithAddr(conf.Host, conf.Port)}, opts...)
	opts = append([]Option{WithAuth(conf.User, conf.Pass)}, opts...)
	opts = append([]Option{WithDBName(conf.Database)}, opts...)
	opts = append([]Option{WithDebug(conf.Debug)}, opts...)

	return newOrm(opts...)
}

func WithGormConf(conf *gorm.Config) Option {
	return func(r *Orm) {
		r.Conf = conf
	}
}

func WithDebug(debug bool) Option {
	return func(r *Orm) {
		r.Debug = debug
	}
}

func WithAddr(host, port string) Option {
	return func(r *Orm) {
		r.Host = host
		r.Port = port
	}
}

func WithAuth(user, pass string) Option {
	return func(r *Orm) {
		r.Pass = pass
		r.User = user
	}
}

func WithDBName(db string) Option {
	return func(r *Orm) {
		r.Database = db
	}
}

func newOrm(opts ...Option) (*Orm, error) {
	m := &Orm{}
	for _, opt := range opts {
		opt(m)
	}
	conf := m.Conf
	if conf == nil {
		conf = &gorm.Config{}
	}
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User, m.Pass, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(dsn), conf)
	if m.Debug {
		db = db.Debug()
	}
	m.DB = db
	return m, err
}
