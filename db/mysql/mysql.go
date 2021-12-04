package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type MySqlDB struct {
	Database        string
	UserName        string
	Password        string
	Host            string
	Port            string
	Protocol        string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime int64
	SourceName      string
}

func defaultConfig() *MySqlDB {
	m := &MySqlDB{
		UserName:        "root",
		Host:            "localhost",
		Port:            "3306",
		Protocol:        "tcp",
		MaxOpenConn:     100,
		MaxIdleConn:     10,
		ConnMaxLifetime: 10 * 60,
	}

	return m
}

func NewMysqlDB(database, user, password, host, port, protocol string, maxOpenConn, maxIdleConn int, ConnMaxLifetime int64) *MySqlDB {
	defaultConf := defaultConfig()

	if maxOpenConn == 0 {
		maxOpenConn = defaultConf.MaxOpenConn
		maxIdleConn = defaultConf.MaxIdleConn
		ConnMaxLifetime = defaultConf.ConnMaxLifetime
	}

	return &MySqlDB{
		Database:        database,
		UserName:        user,
		Password:        password,
		Host:            host,
		Port:            port,
		Protocol:        protocol,
		MaxOpenConn:     maxOpenConn,
		MaxIdleConn:     maxIdleConn,
		ConnMaxLifetime: ConnMaxLifetime,
	}
}

func (m *MySqlDB) Connect() *sqlx.DB {
	m.SourceName = fmt.Sprintf("%s:%s@%s(%s:%s)/%s", m.UserName, m.Password, m.Protocol, m.Host, m.Port, m.Database)

	pool, err := sqlx.Open("mysql", m.SourceName)
	if err != nil {
		panic(err)
	}

	pool.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
	pool.SetMaxIdleConns(m.MaxIdleConn)
	pool.SetMaxOpenConns(m.MaxOpenConn)

	if err = pool.Ping(); err != nil {
		log.Panicf("[Connect] Ping err:%s", err)
	}

	return pool
}
