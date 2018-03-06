package gorm

import (
	"fmt"

	"anla.io/taizhou-fe-api/config"
	// is mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// gorm mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbConfig  = config.Config.DB
	mysqlConn *gorm.DB
	err       error
)

// initialize database
func init() {
	if dbConfig.Driver == "mysql" {
		setupMysqlConn()
	}
}

// setupMysqlConn: setup mysql database connection using the configuration from config.yml
func setupMysqlConn() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	fmt.Println(connectionString)

	mysqlConn, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	err = mysqlConn.DB().Ping()
	if err != nil {
		panic(err)
	}
	mysqlConn.LogMode(true)
	// mysqlConn.DB().SetMaxIdleConns(mysql.MaxIdleConns)
}

// MysqlConn is: return mysql connection from gorm ORM
func MysqlConn() *gorm.DB {
	return mysqlConn
}
