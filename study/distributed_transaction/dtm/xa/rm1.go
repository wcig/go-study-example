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

func RunRM1Server() {
	app := gin.Default()
	app.POST("/api/busi/trans_out", transOutHandler)
	go func() {
		if err := app.Run(":28081"); err != nil {
			panic(err)
		}
	}()
	time.Sleep(100 * time.Millisecond)
}

func transOutHandler(c *gin.Context) {
	if err := transOut(c); err != nil {
		log.Printf(">> rm1-server transout handler failed, err: %v", err)
		c.JSON(http.StatusConflict, &common.ResHTTP{Result: false, Err: err})
		return
	}
	log.Println(">> rm1-server transout handler success")
	c.JSON(http.StatusOK, &common.ResHTTP{Result: true})
}

func transOut(c *gin.Context) error {
	qs := c.Request.URL.Query()
	xaFunc := func(db *sql.DB, xa *dtmcli.Xa) error {
		var req common.ReqHTTP
		if err := c.ShouldBindJSON(&req); err != nil {
			return err
		}
		sqlFormat := "update dtm_busi.user_account set balance=balance-? where user_id=?"
		result, err := db.Exec(sqlFormat, req.Amount, req.TransOutUserID)
		if err != nil {
			return err
		}
		affected, err := result.RowsAffected()
		log.Printf(">> rm1-server transout db success, result: %d, %v", affected, err)
		return nil
	}
	if err := dtmcli.XaLocalTransaction(qs, common.BasicDBConf, xaFunc); err != nil {
		return err
	}
	log.Println(">> rm1-server exec local transaction success")
	return nil
}
