package initialize

import (
	"fmt"
	"time"

	"github.com/twinbeard/goLearning/global"
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
	fmt.Println("helo")

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	s := fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname) // Format the connection string
	fmt.Println(s)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false, // Disable nested transactions => Giúp tăng tính nhất quán của dữ liệu như gây chậm
	}) // Open a database connection
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("MySQL initialization successful", zap.String("DB", m.Dbname))

	global.Mdb = db
	SetPool() // Set the maximum number of open and idle connections to the database
	// migrateTables() // create tables in the database
}

// SetPool() is used to set the number of connections to the database
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
	// AutoMigrate() will use blueprint in models.GoCrmUserV2 to creates tables in the database
	err := global.Mdb.AutoMigrate(
	// &po.User{}, &po.Role{}
	// &models.GoCrmUserV2{},
	)
	if err != nil {
		fmt.Println("AutoMigrate tables error", err)

	}
}
