package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"RU-Smart-Workspace/ru-smart-api/handlers"
	"RU-Smart-Workspace/ru-smart-api/handlers/studenth"
	"RU-Smart-Workspace/ru-smart-api/logger"

	"RU-Smart-Workspace/ru-smart-api/repositories/studentr"
	"RU-Smart-Workspace/ru-smart-api/services"
	"RU-Smart-Workspace/ru-smart-api/services/students"

	"RU-Smart-Workspace/ru-smart-api/handlers/public/mr30h"
	"RU-Smart-Workspace/ru-smart-api/repositories/public/mr30r"
	"RU-Smart-Workspace/ru-smart-api/services/public/mr30s"

	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"RU-Smart-Workspace/ru-smart-api/repositories"

	"RU-Smart-Workspace/ru-smart-api/handlers/masterhandlers"
	"RU-Smart-Workspace/ru-smart-api/repositories/masterrepo"
	"RU-Smart-Workspace/ru-smart-api/services/masterservice"

	"RU-Smart-Workspace/ru-smart-api/handlers/officerhandlers"
	"RU-Smart-Workspace/ru-smart-api/repositories/officerrepos"
	"RU-Smart-Workspace/ru-smart-api/services/officerservices"
)

func Setup(router *gin.Engine, oracle_db *sqlx.DB, oracle_db_dbg *sqlx.DB, redis_cache *redis.Client, mysql_db *sqlx.DB, mysql_db_stdapps *sqlx.DB, mysql_db_rotcs *sqlx.DB, oracleScholar_db *sqlx.DB) {

	jsonFileLogger, err := logger.NewJSONFileLogger("/logger/app.log")
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	router.Use(logger.ErrorLogger(jsonFileLogger))

	router.Use(middlewares.NewCorsAccessControl().CorsAccessControl())

	router.GET("/healthz", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "The service works normally...",
		})
	})

	officeAuth := router.Group("/officer")
	{
		officeRepo := officerrepos.NewOfficerRepo(oracle_db_dbg)
		officerService := officerservices.NewOfficerServices(officeRepo, redis_cache)
		officeHandler := officerhandlers.NewOfficerHandlers(officerService)

		officeAuth.POST("/authorization", officeHandler.Authentication)
		officeAuth.POST("/refresh-authentication", officeHandler.RefreshAuthentication)
	}

	googleAuth := router.Group("/google")
	{
		studentRepo := studentr.NewStudentRepo(oracle_db, oracle_db_dbg)
		studentService := students.NewStudentServices(studentRepo, redis_cache)
		studentHandler := studenth.NewStudentHandlers(studentService)

		googleAuth.POST("/authorization", middlewares.GoogleAuth, studentHandler.Authentication)
		googleAuth.POST("/authorization-test", studentHandler.AuthenticationTest)
		googleAuth.POST("/authorization-service", studentHandler.AuthenticationService)
		googleAuth.POST("/authorization-redirect", studentHandler.AuthenticationRedirect)

	}

	student := router.Group("/student")
	{
		studentRepo := studentr.NewStudentRepo(oracle_db, oracle_db_dbg)
		studentService := students.NewStudentServices(studentRepo, redis_cache)
		studentHandler := studenth.NewStudentHandlers(studentService)

		student.POST("/certificate", middlewares.Authorization(redis_cache), studentHandler.Certifiate)

		student.POST("/refresh-authentication", studentHandler.RefreshAuthentication)
		student.POST("/unauthorization", studentHandler.Unauthorization)
		student.POST("/exists-token", studentHandler.ExistsToken)
		student.GET("/checktoken/:token", studentHandler.CheckToken)
		student.GET("/profile/:std_code", middlewares.Authorization(redis_cache), studentHandler.GetStudentProfile)
		student.GET("/register", middlewares.Authorization(redis_cache), studentHandler.GetRegister)
		student.GET("/registers", middlewares.Authorization(redis_cache), studentHandler.GetRegisterAll)

		student.GET("/photoprofile", middlewares.Authorization(redis_cache), studentHandler.GetPhoto)
		student.GET("/photograduate", middlewares.Authorization(redis_cache), studentHandler.GetPhotoGraduate)
		student.GET("/photograduatesuccess/:id", studentHandler.GetPhotoGraduateSuccess)
		student.GET("/photo/:id", studentHandler.GetPhotoById)
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
		mr30.POST("/data/pagination", mr30Handler.GetMr30Pagination)
	}

	register := router.Group("/register")
	{

		registerRepo := repositories.NewRegisterRepo(oracle_db)
		registerService := services.NewRegisterServices(registerRepo, redis_cache)
		registerHandler := handlers.NewRegisterHandlers(registerService)

		register.GET("/yearsemesterlates", registerHandler.YearSemesterLates)

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

		ondemandRepo := repositories.NewOndemandRepo(mysql_db)
		ondemandService := services.NewOndemandServices(ondemandRepo, redis_cache)
		ondemandHandler := handlers.NewOndemandHandlers(ondemandService)

		ondemand.POST("/", ondemandHandler.GetOndemandAll)

		ondemand.POST("/subjectcode", ondemandHandler.GetOndemandSubjectCode)

	}

	rotcs := router.Group("/rotcs")
	{
		rotcsRepo := repositories.NewRotcsRepo(mysql_db_rotcs)
		rotcsService := services.NewRotcsServices(rotcsRepo, redis_cache)
		rotcsHandler := handlers.NewRotcsHandlers(rotcsService)
		rotcs.POST("/register", middlewares.Authorization(redis_cache), rotcsHandler.GetRotcsRegister)
		rotcs.POST("/extend", middlewares.Authorization(redis_cache), rotcsHandler.GetRotcsExtend)

	}

	insurance := router.Group("/insurance")
	{
		insuranceRepo := repositories.NewInsuranceRepo(mysql_db_stdapps)
		insuranceService := services.NewInsuranceServices(insuranceRepo, redis_cache)
		insuranceHandler := handlers.NewInsuranceHandlers(insuranceService)
		insurance.POST("/", middlewares.Authorization(redis_cache), insuranceHandler.GetInsuranceListAll)

	}

	scholarship := router.Group("/scholarship")
	{
		scholarShipRepo := repositories.NewScholarshipRepo(oracleScholar_db)
		scholarShipService := services.NewScholarShipServices(scholarShipRepo, redis_cache)
		scholarShipHandler := handlers.NewScholarShipHandlers(scholarShipService)

		scholarship.POST("/getScholarShip", scholarShipHandler.GetScholarshipAll)
	}

	event := router.Group("/event")
	{
		eventRepo := repositories.NewEventRepo(mysql_db_stdapps)
		eventService := services.NewEventServices(eventRepo, redis_cache)
		eventHandler := handlers.NewEventHandlers(eventService)
		event.POST("/", eventHandler.GetEventListAll)

	}

	master := router.Group("/master")
	{

		masterRepo := masterrepo.NewStudentRepo(oracle_db_dbg)
		masterService := masterservice.NewStudentServices(masterRepo, redis_cache)
		masterHandler := masterhandlers.NewStudentHandlers(masterService)

		studentMaster := master.Group("/student")
		studentMaster.GET("/profile", middlewares.Authorization(redis_cache), masterHandler.GetStudentProfile)
		studentMaster.GET("/success", middlewares.Authorization(redis_cache), masterHandler.GetStudentSuccess)
		studentMaster.GET("/successcheck/:id", masterHandler.GetStudentSuccessCheck)

		registerMaster := master.Group("/register")
		registerMaster.GET("/", middlewares.Authorization(redis_cache), masterHandler.GetRegisterAll)
		registerMaster.GET("/:year", middlewares.Authorization(redis_cache), masterHandler.GetRegisterByYear)

		gradeMaster := master.Group("/grade")
		gradeMaster.GET("/", middlewares.Authorization(redis_cache), masterHandler.GetGradeAll)
		gradeMaster.GET("/:year", middlewares.Authorization(redis_cache), masterHandler.GetGradeByYear)

	}

	PORT := viper.GetString("ruConnext.port")
	router.Run(PORT)

}

func errorLogger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Continue to the next middleware or route handler
		c.Next()
		// Check if any errors occurred during the request handling
		err := c.Errors.Last()
		if err != nil {
			// Log the error
			log.WithField("status", c.Writer.Status()).
				WithField("method", c.Request.Method).
				WithField("path", c.Request.URL.Path).
				Error(err.Err)
		}
	}
}
