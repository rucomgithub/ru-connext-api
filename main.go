package main

import (
	"RU-Smart-Workspace/ru-smart-api/databases"
	"RU-Smart-Workspace/ru-smart-api/environments"
	"RU-Smart-Workspace/ru-smart-api/routers"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
)

var oracle_db *sqlx.DB
var mysql_db *sql.DB
var redis_cache *redis.Client

func init() {
	//logger.LoggerInit()
	environments.TimeZoneInit()
	environments.EnvironmentInit()
	oracle_db = databases.NewDatabases().OracleInit()
	mysql_db = databases.NewDatabases().MySQLInit()
	redis_cache = databases.NewDatabases().RedisInint()
}

func main() {

	defer oracle_db.Close()
	defer redis_cache.Close()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routers.Setup(router, oracle_db, mysql_db, redis_cache)
}
