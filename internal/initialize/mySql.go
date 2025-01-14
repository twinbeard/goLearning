package initialize

import (
	"fmt"
	"time"

	"github.com/twinbeard/goLearning/global"
	"github.com/twinbeard/goLearning/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	s := fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname) // Format the connection string
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false, // Disable nested transactions => Giúp tăng tính nhất quán của dữ liệu như gây chậm
	}) // Open a database connection
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("MySQL initialization successful", zap.String("DB", m.Dbname))

	global.Mdb = db
	SetPool()       // Set the maximum number of open and idle connections to the database
	migrateTables() // create tables in the database
}

// SetPool() sets the maximum number of open and idle connections to the database
func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()

	if err != nil {
		fmt.Printf("mysql error: %s", err)
	}

	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(&po.User{}, &po.Role{})
	if err != nil {
		fmt.Println("AutoMigrate tables error", err)

	}
}
