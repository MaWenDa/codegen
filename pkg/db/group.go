package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// Client 继承gorm.DB
type Client struct {
	// go没有继承的概念，可以通过嵌套结构体来实现代码复用和组合
	*gorm.DB
}

// Group 数据库客户端Group
type Group struct {
	name   string
	master *Client
	slaves []*Client
	next   uint64
	total  uint64
	mu     sync.RWMutex
}

// Master 获取主库客户端
func (g *Group) Master() *Client {
	return g.master
}

// Slave 获取从库客户端
func (g *Group) Slave() *Client {

	// 获取客户端时不允许修改
	g.mu.RLock()
	defer g.mu.RUnlock()

	// 没有从库返回主库实例
	if g.total == 0 {
		return g.master
	}

	// 轮询slaves客户端实例
	next := atomic.AddUint64(&g.next, 1)
	return g.slaves[next%g.total]
}

// openDB 打开mysql链接
func openDB(dsn string) (*Client, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	// 创建mysql链接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("open mysql addr [%s] err: %v", dsn, err)
	}

	// 是否开启debug
	if debug, _ := strconv.ParseBool(os.Getenv("DEBUG")); debug {
		db.Logger.LogMode(logger.Silent)
	}

	return &Client{DB: db}, nil
}

// NewGroup 获取数据库链接组
func NewGroup(name string, cfg Config) (*Group, error) {

	var err error

	// 数据库组实例
	g := &Group{name: name}

	// 主库实例
	g.master, err = openDB(cfg.Master.Dsn)
	if err != nil {
		return nil, err
	}

	for _, slaveCfg := range cfg.Slaves {

		// 从库实例
		slave, err := openDB(slaveCfg.Dsn)
		if err != nil {
			return nil, err
		}

		// 从库集合
		g.slaves = append(g.slaves, slave)

		// 从库数量
		g.total++
	}

	return g, err
}

// ConnMysql 链接mysql
// username 	用户名
// password 	密码
// dataSourceName  数据源链接
func ConnMysql(username, password, dataSource string) *gorm.DB {

	if !TestConnMysql(username, password, dataSource) {
		panic("数据库链接失败")
	}
	// 拼接dsn，并增加链接超时时间，默认为2s
	dsn := username + ":" + password + dataSource

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 连接数据库，
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	// 是否开启debug
	if debug, _ := strconv.ParseBool(os.Getenv("DEBUG")); debug {
		db.Logger.LogMode(logger.Silent)
	}
	return db

}

// ConnMysql 链接测试
// username 	用户名
// password 	密码
// dataSourceName  数据源链接
func TestConnMysql(username, password, dataSource string) bool {

	// 拼接dsn，并增加链接超时时间，默认为2s
	dsn := username + ":" + password + dataSource + "&timeout=2s"

	// 连接数据库，
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return false
	}

	return true

}
