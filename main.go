package main

import (
	"fmt"
	"log"
	"RU-Smart-Workspace/ru-smart-api/databases"
	"RU-Smart-Workspace/ru-smart-api/infrastructure/db"
	"RU-Smart-Workspace/ru-smart-api/environments"
	"RU-Smart-Workspace/ru-smart-api/routers"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var oracle_db *sqlx.DB
var oracle_db_dbg *sqlx.DB
var mysql_db *sqlx.DB
var mysql_db_rotcs *sqlx.DB
var mysql_db_stdapps *sqlx.DB

var redis_cache *redis.Client

var oracleScholar_db *sqlx.DB

var clientID string

func init() {
	//logger.LoggerInit()
	environments.TimeZoneInit()
	environments.EnvironmentInit()
	oracle_db = databases.NewDatabases().OracleInit()
	oracle_db_dbg = databases.NewDatabases().OracleDBGInit()
	redis_cache = databases.NewDatabases().RedisInint()
	mysql_db = databases.NewDatabases().MysqlInit()
	mysql_db_rotcs = databases.NewDatabases().MysqlInitRotcs()
	mysql_db_stdapps = databases.NewDatabases().MysqlInitStdApps()
	oracleScholar_db = databases.NewDatabases().OracleScholarShipInit()
}

func main() {
	dns := fmt.Sprintf("%v", viper.GetString("db.connectionDBG"))
	clientID := fmt.Sprintf("%v", viper.GetString("google_client_id"))
	database, err := db.NewOracleDB(dns)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	defer oracle_db.Close()
	defer redis_cache.Close() 
	defer mysql_db.Close()
	defer mysql_db_rotcs.Close()
	defer oracleScholar_db.Close()
	gin.SetMode(gin.ReleaseMode) 
	router := gin.Default()
	routers.Setup(router, oracle_db,oracle_db_dbg, redis_cache, mysql_db, mysql_db_stdapps, mysql_db_rotcs, oracleScholar_db, database , clientID)
}
