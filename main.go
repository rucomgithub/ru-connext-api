package main

import (
	"RU-Smart-Workspace/ru-smart-api/databases"
	"RU-Smart-Workspace/ru-smart-api/environments"
	"RU-Smart-Workspace/ru-smart-api/routers"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
)

var oracle_db *sqlx.DB

var mysql_db *sqlx.DB
var mysql_db_rotcs *sqlx.DB

var redis_cache *redis.Client

var oracleScholar_db *sqlx.DB

func init() {
	//logger.LoggerInit()
	environments.TimeZoneInit()
	environments.EnvironmentInit()
	oracle_db = databases.NewDatabases().OracleInit()
	redis_cache = databases.NewDatabases().RedisInint()
	mysql_db = databases.NewDatabases().MysqlInit()
	mysql_db_rotcs = databases.NewDatabases().MysqlInitRotcs()
	oracleScholar_db = databases.NewDatabases().OracleScholarShipInit()
}

func main() {
	defer oracle_db.Close()
	defer redis_cache.Close()
	defer mysql_db.Close()
	defer mysql_db_rotcs.Close()
	defer oracleScholar_db.Close()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routers.Setup(router, oracle_db, redis_cache, mysql_db, mysql_db_rotcs, oracleScholar_db)
}
