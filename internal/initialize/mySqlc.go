package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/twinbeard/goLearning/global"
	"go.uber.org/zap"
)

//* Dùng SQLC thay cho gorm

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlC() {
	m := global.Config.Mysql

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	s := fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname) // Format the connection string
	fmt.Println(s)
	db, err := sql.Open("mysql", s) // Open a database connection
	checkErrorPanicC(err, "InitMysql initialization error")
	global.Logger.Info("MySQL initialization successful", zap.String("DB", m.Dbname))
	global.Mdbc = db
	// SetPoolC() // Set the maximum number of open and idle connections to the database
	// migrateTablesC() // create tables in the database
}

// SetPool() is used to set the number of connections to the database
func SetPoolC() {
	m := global.Config.Mysql
	sqlDb := global.Mdbc

	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

// func migrateTablesC() {
// 	// AutoMigrate() will use blueprint in models.GoCrmUserV2 to creates tables in the database
// 	err := global.Mdb.AutoMigrate(
// 		// &po.User{}, &po.Role{}
// 		&models.GoCrmUserV2{},
// 	)
// 	if err != nil {
// 		fmt.Println("AutoMigrate tables error", err)

// 	}
// }
