package mysql

import (
	"fmt"
	"github.com/lemonlzy/vegetableBlog/internal/app"
	"github.com/lemonlzy/vegetableBlog/internal/server/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

func Init(cfg *conf.DBConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	app.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{Colorful: false}), // 禁用控制台颜色，防止乱码
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.TablePrefix, // 表名前缀，`User`表为`v_user`
			SingularTable: true,
		}})

	if err != nil {
		log.Printf("DB init err: %v\n", err)
		return
	}

	// 根据Models结构体初始化数据库表
	if cfg.MigrateTable {
		if err = app.DB.AutoMigrate(getModels()...); err != nil {
			return
		}
	}

	sqlDB, err := app.DB.DB()
	if err != nil {
		log.Printf("DB init err: %v\n", err)
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	return
}

func getModels() []interface{} {
	return []interface{}{
		&app.User{},
		&app.Tag{},
		&app.Article{},
	}
}
