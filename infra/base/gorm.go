package base

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"moyutec.top/resk/infra"
	gormlogrus "moyutec.top/resk/infra/logrus"
	"time"
)

var orm *gorm.DB

func ORM() *gorm.DB {
	return orm
}

type GormStarter struct {
	infra.BaseStarter
}

func (g *GormStarter) Init(ctx infra.StarterContext) {
	props := ctx.Props()
	conn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		props.GetDefault("mysql.user", ""),
		props.GetDefault("mysql.password", ""),
		"tcp",
		props.GetDefault("mysql.host", "localhost:3306"),
		props.GetDefault("mysql.database", "db"),
		props.GetDefault("mysql.options.charset", "utf8mb4,utf8"),
		props.GetDefault("mysql.options.parseTime", "true"),
		props.GetDefault("mysql.options.loc", "Local"),
	)
	db, err := gorm.Open(props.GetDefault("mysql.driverName", "mysql"), conn)
	if err != nil {
		panic(err)
	}
	logrus.Info(db.DB().Ping())
	db.SetLogger(gormlogrus.NewUpperLogrusLogger())
	db.DB().SetConnMaxLifetime(props.GetDurationDefault("mysql.connMaxLifetime", time.Hour*1))
	db.DB().SetMaxIdleConns(props.GetIntDefault("mysql.maxIdleConns", 1))
	db.DB().SetMaxOpenConns(props.GetIntDefault("mysql.maxOpenConns", 3))
	orm = db
}
