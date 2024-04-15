package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var Db *gorm.DB

func init() {
	cfg, _ := LoadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		cfg.MySQL.Username,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.Name,
	)
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbl_", // 表前缀
			SingularTable: true,   // 禁用表名复数
		},
		NowFunc: func() time.Time {
			return time.Now().UTC().Truncate(time.Second)
		},
		//Logger: logger.Default.LogMode(logger.Info),
	}
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		log.Fatal(err)
	}
}
