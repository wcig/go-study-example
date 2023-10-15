package main

import (
	"database/sql"
	"go-app/study/distributed_transaction/dtm/common"
	"log"
	"net/http"
	"time"

	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	rm2DB *sql.DB
)

func RunRM2Server() {
	rm2DB = InitDB()
	InitRM2App()
	time.Sleep(100 * time.Millisecond)
}

func InitRM2App() {
	app := gin.Default()
	app.POST("/api/busi/trans_in_try", transInTryHandler)
	app.POST("/api/busi/trans_in_confirm", transInConfirmHandler)
	app.POST("/api/busi/trans_in_cancel", transInCancelHandler)
	go func() {
		if err := app.Run(":28082"); err != nil {
			panic(err)
		}
	}()
}

func transInTryHandler(c *gin.Context) {
	if err := transInTry(c); err != nil {
		log.Printf(">> rm1-server transin try handler failed, err: %v", err)
		c.JSON(http.StatusConflict, &common.ResHTTP{Result: false, Err: err})
		return
	}
	log.Println(">> rm1-server transin try handler success")
	c.JSON(http.StatusOK, &common.ResHTTP{Result: true})
}

func transInConfirmHandler(c *gin.Context) {
	if err := transInConfirm(c); err != nil {
		log.Printf(">> rm2-server transin confirm handler failed, err: %v", err)
		c.JSON(http.StatusConflict, &common.ResHTTP{Result: false, Err: err})
		return
	}
	log.Println(">> rm2-server transin confirm handler success")
	c.JSON(http.StatusOK, &common.ResHTTP{Result: true})
}

func transInCancelHandler(c *gin.Context) {
	if err := transInCancel(c); err != nil {
		log.Printf(">> rm2-server transin cancel handler failed, err: %v", err)
		c.JSON(http.StatusConflict, &common.ResHTTP{Result: false, Err: err})
		return
	}
	log.Println(">> rm2-server transin cancel handler success")
	c.JSON(http.StatusOK, &common.ResHTTP{Result: true})
}

func transInTry(c *gin.Context) error {
	// return errors.New("my transin try test error")

	req, err := ParseReq(c)
	if err != nil {
		return err
	}
	barrier := MustBarrierFromGin(c)
	busiFunc := func(tx *sql.Tx) error {
		driver := common.BasicDBConf.Driver
		sqlFormat := "update dtm_busi.user_account set trading_balance=trading_balance+? where user_id=? and trading_balance + ? + balance >= 0"
		affected, err2 := dtmimp.DBExec(driver, rm2DB, sqlFormat, req.Amount, req.TransInUserID, req.Amount)
		if err2 != nil {
			return err2
		}
		log.Printf(">> rm2-server transin try db success, result: %d", affected)
		return nil
	}
	return barrier.CallWithDB(rm2DB, busiFunc)
}

func transInConfirm(c *gin.Context) error {
	req, err := ParseReq(c)
	if err != nil {
		return err
	}
	barrier := MustBarrierFromGin(c)
	busiFunc := func(tx *sql.Tx) error {
		driver := common.BasicDBConf.Driver
		sqlFormat := "update dtm_busi.user_account set trading_balance=trading_balance-?, balance=balance+?  where user_id=?"
		affected, err2 := dtmimp.DBExec(driver, rm2DB, sqlFormat, req.Amount, req.Amount, req.TransInUserID)
		if err2 != nil {
			return err2
		}
		log.Printf(">> rm2-server transin confirm db success, result: %d", affected)
		return nil
	}
	return barrier.CallWithDB(rm2DB, busiFunc)
}

func transInCancel(c *gin.Context) error {
	req, err := ParseReq(c)
	if err != nil {
		return err
	}
	barrier := MustBarrierFromGin(c)
	busiFunc := func(tx *sql.Tx) error {
		driver := common.BasicDBConf.Driver
		sqlFormat := "update dtm_busi.user_account set trading_balance=trading_balance+? where user_id=? and trading_balance + ? + balance >= 0"
		affected, err2 := dtmimp.DBExec(driver, rm2DB, sqlFormat, -req.Amount, req.TransInUserID, -req.Amount)
		if err2 != nil {
			return err2
		}
		log.Printf(">> rm2-server transin cancel db success, result: %d", affected)
		return nil
	}
	return barrier.CallWithDB(rm2DB, busiFunc)
}
