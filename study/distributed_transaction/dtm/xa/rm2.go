package main

import (
	"database/sql"
	"go-app/study/distributed_transaction/dtm/common"
	"log"
	"net/http"
	"time"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func RunRM2Server() {
	app := gin.Default()
	app.POST("/api/busi/trans_in", transInHandler)
	go func() {
		if err := app.Run(":28082"); err != nil {
			panic(err)
		}
	}()
	time.Sleep(100 * time.Millisecond)
}

func transInHandler(c *gin.Context) {
	if err := transIn(c); err != nil {
		log.Printf(">> rm2-server transin handler failed, err: %v", err)
		c.JSON(http.StatusConflict, &common.ResHTTP{Result: false, Err: err})
		return
	}
	log.Println(">> rm2-server transin handler success")
	c.JSON(http.StatusOK, &common.ResHTTP{Result: true})
}

func transIn(c *gin.Context) error {
	qs := c.Request.URL.Query()
	xaFunc := func(db *sql.DB, xa *dtmcli.Xa) error {
		var req common.ReqHTTP
		if err := c.ShouldBindJSON(&req); err != nil {
			return err
		}
		result, err := db.Exec("update dtm_busi.user_account set balance=balance+? where user_id=?", req.Amount, req.TransInUserID)
		if err != nil {
			return err
		}
		affected, err := result.RowsAffected()
		log.Printf(">> rm2-server transin db success, result: %d, %v", affected, err)
		return nil
	}
	if err := dtmcli.XaLocalTransaction(qs, common.BasicDBConf, xaFunc); err != nil {
		return err
	}
	log.Println(">> rm2-server exec local transaction success")
	return nil
}
