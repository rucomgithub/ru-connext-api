package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/handlers/studenth"
	"RU-Smart-Workspace/ru-smart-api/repositories/studentr"
	"RU-Smart-Workspace/ru-smart-api/services"
	"RU-Smart-Workspace/ru-smart-api/services/students"

	"RU-Smart-Workspace/ru-smart-api/handlers/public/mr30h"
	"RU-Smart-Workspace/ru-smart-api/repositories/public/mr30r"
	"RU-Smart-Workspace/ru-smart-api/services/public/mr30s"

	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"RU-Smart-Workspace/ru-smart-api/repositories"

	
)

func Setup(router *gin.Engine, oracle_db *sqlx.DB, mysql_db *sqlx.DB, redis_cache *redis.Client) {

	router.Use(middlewares.NewCorsAccessControl().CorsAccessControl())

	router.GET("/healthz", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "The service works normally... kor jaa",
		})
	})

	googleAuth := router.Group("/google")
	{
		studentRepo := studentr.NewStudentRepo(oracle_db)
		studentService := students.NewStudentServices(studentRepo, redis_cache)
		studentHandler := studenth.NewStudentHandlers(studentService)

		googleAuth.POST("/authorization", middlewares.GoogleAuth, studentHandler.Authentication)
		googleAuth.POST("/authorization-test", studentHandler.AuthenticationTest)
		googleAuth.POST("/authorization-redirect", studentHandler.AuthenticationRedirect)

	}

	student := router.Group("/student")
	{
		studentRepo := studentr.NewStudentRepo(oracle_db)
		studentService := students.NewStudentServices(studentRepo, redis_cache)
		studentHandler := studenth.NewStudentHandlers(studentService)

		student.POST("/refresh-authentication", studentHandler.RefreshAuthentication)
		student.POST("/unauthorization", studentHandler.Unauthorization)
		student.POST("/exists-token", studentHandler.ExistsToken)
		student.GET("/profile/:std_code", middlewares.Authorization(redis_cache), studentHandler.GetStudentProfile)
		student.GET("/register", middlewares.Authorization(redis_cache), studentHandler.GetRegister)
		student.GET("/registers", middlewares.Authorization(redis_cache), studentHandler.GetRegisterAll)

		student.GET("/imageprofile", middlewares.Authorization(redis_cache), studentHandler.GetImageProfile)
		student.GET("/photoprofile", middlewares.Authorization(redis_cache), studentHandler.GetPhoto)
		student.GET("/photo/:id", studentHandler.GetPhotoById)
		student.GET("/photoaod", middlewares.Authorization(redis_cache), studentHandler.GetPhotoAOD)

		student.GET("/", studentHandler.GetStudentAll)
	}

	mr30 := router.Group("/mr30")
	{

		mr30Repo := mr30r.NewMr30Repo(oracle_db)
		mr30Service := mr30s.NewMr30Services(mr30Repo, redis_cache)
		mr30Handler := mr30h.NewMr30Handlers(mr30Service)

		// mr30.GET("/data", mr30Handler.GetMr30)
		mr30.POST("/year", mr30Handler.GetMr30Year)
		mr30.POST("/data", mr30Handler.GetMr30)
		mr30.GET("/data/search", mr30Handler.GetMr30Searching)
		mr30.GET("/data/pagination", mr30Handler.GetMr30Pagination)
	}

	register := router.Group("/register")
	{

		registerRepo := repositories.NewRegisterRepo(oracle_db)
		registerService := services.NewRegisterServices(registerRepo, redis_cache)
		registerHandler := handlers.NewRegisterHandlers(registerService)

		register.POST("/", middlewares.Authorization(redis_cache), registerHandler.Registers)

		register.GET("/:std_code/year", middlewares.Authorization(redis_cache), registerHandler.Years)
		register.GET("/:std_code/yearsemester", middlewares.Authorization(redis_cache), registerHandler.YearSemesters)
		register.POST("/:std_code/schedule", middlewares.Authorization(redis_cache), registerHandler.ScheduleYearSemesters)
		register.POST("/:std_code/schedulelatest", middlewares.Authorization(redis_cache), registerHandler.Schedules)
	}

	grade := router.Group("/grade")
	{

		gradeRepo := repositories.NewGradeRepo(oracle_db)
		gradeService := services.NewGradeServices(gradeRepo, redis_cache)
		gradeHandler := handlers.NewgradeHandlers(gradeService)

		grade.POST("/:std_code/year", middlewares.Authorization(redis_cache), gradeHandler.GradeYear)
		grade.POST("/:std_code", middlewares.Authorization(redis_cache), gradeHandler.Grades)
	}

	ondemand := router.Group("/ondemand")
	{

		ondemandRepo := repositories.NewOnDemandRepo(mysql_db)
		// ondemandService := services.NewOnDemandServices(ondemandRepo, redis_cache)
		// ondemandHandler := handlers.NewOnDeMandHandlers(ondemandService)


		// ondemand.POST("/", ondemandHandler.GetOnDemandAll)

	}

	PORT := viper.GetString("ruConnext.port")
	router.Run(PORT)

}