package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/dtmimp"
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
		c.JSON(http.StatusConflict, &ResHTTP{Result: false, Err: err})
		return
	}
	log.Println(">> rm1-server transout handler success")
	c.JSON(http.StatusOK, &ResHTTP{Result: true})
}

func transOut(c *gin.Context) error {
	qs := c.Request.URL.Query()
	xaFunc := func(db *sql.DB, xa *dtmcli.Xa) error {
		var req ReqHTTP
		if err := c.ShouldBindJSON(&req); err != nil {
			return err
		}
		result, err := db.Exec("update dtm_busi.user_account set balance=balance-? where user_id=?", req.Amount, req.TransOutUserID)
		if err != nil {
			return err
		}
		affected, err := result.RowsAffected()
		log.Printf(">> rm1-server transout db success, result: %d, %v", affected, err)
		return nil
	}
	if err := dtmcli.XaLocalTransaction(qs, BasicDBConf, xaFunc); err != nil {
		return err
	}
	log.Println(">> rm1-server exec local transaction success")
	return nil
}

var BasicDBConf = dtmimp.DBConf{
	Driver:   "mysql",
	Host:     "localhost",
	Port:     3306,
	User:     "root",
	Password: "123456",
	// Db:       "",
	// Schema:   "",
}

type ReqHTTP struct {
	Amount         int `json:"amount"`
	TransOutUserID int `json:"trans_out_user_id"`
	TransInUserID  int `json:"trans_in_user_id"`
	// TransInResult  string `json:"trans_in_result"`
	// TransOutResult string `json:"trans_out_Result"`
}

type ResHTTP struct {
	Result bool        `json:"result"`
	Data   interface{} `json:"data"`
	Err    error       `json:"err"`
}
