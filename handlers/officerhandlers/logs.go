package officerhandlers

import (
	"context"
	//"database/sql"
	"net/http"
	//"strconv"
	//"fmt"
	"time"

	"github.com/gin-gonic/gin"
	//"encoding/json"
)

type (
	Log struct {
		Id       string `json:"Id"`
		Code     int64  `json:"Code" validate:"required"`
		Module   string `json:"Module" validate:"required"`
		Task     string `json:"Task" validate:"required"`
		Username string `json:"Username" validate:"required"`
		Created  string `json:"Created"`
		Modified string `json:"Modified"`
	}

	LogForm struct {
		Code   int64  `json:"Code" validate:"required"`
		Module string `json:"Module" validate:"required"`
	}
)

func (h *officerHandlers) CreateLogs(c *gin.Context) {
	var (
		logs []Log
		log  Log
	)
	form := new(Log)

	if err := c.Bind(form); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	tx, err := h.oracle_db_dbg.BeginTx(ctx, nil)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	_, err = tx.ExecContext(ctx, "insert into egrad_logs (ID,CODE,MODULE,TASK,USERNAME,CREATED,MODIFIED) values (EGRAD_LOGS_SEQ.NEXTVAL,:1,:2,:3,:4,sysdate,sysdate)", form.Code, form.Module, form.Task, form.Username)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	err = tx.Commit()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	sql := `select * from egrad_logs where code = :1 and module = :2 order by id desc `

	rows, err := h.oracle_db_dbg.Query(sql, form.Code, form.Module)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&log.Id, &log.Code, &log.Module, &log.Task, &log.Username,
			&log.Created, &log.Modified)

		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	if len(logs) < 1 {
		c.IndentedJSON(http.StatusBadRequest, map[string]string{"message": "ไม่พบข้อมูล."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, logs)

}

func (h *officerHandlers) FindLogs(c *gin.Context) {
	var (
		logs []Log
		log  Log
	)
	form := new(LogForm)

	if err := c.Bind(form); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	// err := c.ShouldBindJSON(&form)
	// if err != nil {
	// 	ErrValidateRequest(c)
	// 	c.Abort()
	// 	return
	// }

	sql := `select * from egrad_logs where code = :1 and module = :2 order by id desc `

	rows, err := h.oracle_db_dbg.Query(sql, form.Code, form.Module)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&log.Id, &log.Code, &log.Module, &log.Task, &log.Username,
			&log.Created, &log.Modified)

		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	if len(logs) < 1 {
		c.IndentedJSON(http.StatusBadRequest, map[string]string{"message": "ไม่พบข้อมูล."})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, logs)

}
