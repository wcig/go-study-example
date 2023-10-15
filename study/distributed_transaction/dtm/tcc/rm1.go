package main

import (
	"database/sql"
	"go-app/study/distributed_transaction/dtm/common"
	"log"
	"net/http"
	"time"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	rm1DB *sql.DB
)

func RunRM1Server() {
	rm1DB = InitDB()
	InitRM1App()
	time.Sleep(100 * time.Millisecond)
}

func InitRM1App() {
	app := gin.Default()
	app.POST("/api/busi/trans_out_try", transOutTryHandler)
	app.POST("/api/busi/trans_out_confirm", transOutConfirmHandler)
	app.POST("/api/busi/trans_out_cancel", transOutCancelHandler)
	go func() {
		if err := app.Run(":28081"); err != nil {
			panic(err)
		}
	}()
}

func transOutTryHandler(c *gin.Context) {
	if err := transOutTry(c); err != nil {
		log.Printf(">> rm1-server transout try handler failed, err: %v", err)
		c.JSON(http.StatusConflict, &common.ResHTTP{Result: false, Err: err})
		return
	}
	log.Println(">> rm1-server transout try handler success")
	c.JSON(http.StatusOK, &common.ResHTTP{Result: true})
}

func transOutConfirmHandler(c *gin.Context) {
	if err := transOutConfirm(c); err != nil {
		log.Printf(">> rm1-server transout confirm handler failed, err: %v", err)
		c.JSON(http.StatusConflict, &common.ResHTTP{Result: false, Err: err})
		return
	}
	log.Println(">> rm1-server transout confirm handler success")
	c.JSON(http.StatusOK, &common.ResHTTP{Result: true})
}

func transOutCancelHandler(c *gin.Context) {
	if err := transOutCancel(c); err != nil {
		log.Printf(">> rm1-server transout cancel handler failed, err: %v", err)
		c.JSON(http.StatusConflict, &common.ResHTTP{Result: false, Err: err})
		return
	}
	log.Println(">> rm1-server transout cancel handler success")
	c.JSON(http.StatusOK, &common.ResHTTP{Result: true})
}

func transOutTry(c *gin.Context) error {
	// return errors.New("my transout try test error")

	req, err := ParseReq(c)
	if err != nil {
		return err
	}
	barrier := MustBarrierFromGin(c)
	busiFunc := func(tx *sql.Tx) error {
		driver := common.BasicDBConf.Driver
		sqlFormat := "update dtm_busi.user_account set trading_balance=trading_balance+? where user_id=? and trading_balance + ? + balance >= 0"
		affected, err2 := dtmimp.DBExec(driver, rm1DB, sqlFormat, -req.Amount, req.TransOutUserID, -req.Amount)
		if err2 != nil {
			return err2
		}
		log.Printf(">> rm1-server transout try db success, result: %d", affected)
		return nil
	}
	return barrier.CallWithDB(rm1DB, busiFunc)
}

func transOutConfirm(c *gin.Context) error {
	req, err := ParseReq(c)
	if err != nil {
		return err
	}
	barrier := MustBarrierFromGin(c)
	busiFunc := func(tx *sql.Tx) error {
		driver := common.BasicDBConf.Driver
		sqlFormat := "update dtm_busi.user_account set trading_balance=trading_balance-?, balance=balance+?  where user_id=?"
		affected, err2 := dtmimp.DBExec(driver, rm1DB, sqlFormat, -req.Amount, -req.Amount, req.TransOutUserID)
		if err2 != nil {
			return err2
		}
		log.Printf(">> rm1-server transout confirm db success, result: %d", affected)
		return nil
	}
	return barrier.CallWithDB(rm1DB, busiFunc)
}

func transOutCancel(c *gin.Context) error {
	req, err := ParseReq(c)
	if err != nil {
		return err
	}
	barrier := MustBarrierFromGin(c)
	busiFunc := func(tx *sql.Tx) error {
		driver := common.BasicDBConf.Driver
		sqlFormat := "update dtm_busi.user_account set trading_balance=trading_balance+? where user_id=? and trading_balance + ? + balance >= 0"
		affected, err2 := dtmimp.DBExec(driver, rm1DB, sqlFormat, req.Amount, req.TransOutUserID, req.Amount)
		if err2 != nil {
			return err2
		}
		log.Printf(">> rm1-server transout cancel db success, result: %d", affected)
		return nil
	}
	return barrier.CallWithDB(rm1DB, busiFunc)
}

func MustBarrierFromGin(c *gin.Context) *dtmcli.BranchBarrier {
	ti, err := dtmcli.BarrierFromQuery(c.Request.URL.Query())
	if err != nil {
		log.Fatalf(">> MustBarrierFromGin err: %v", err)
	}
	return ti
}

func ParseReq(c *gin.Context) (*common.ReqHTTP, error) {
	var req common.ReqHTTP
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func InitDB() *sql.DB {
	sqlDB, err := dtmimp.PooledDB(common.BasicDBConf)
	if err != nil {
		log.Fatalf(">> init db err: %v", err)
	}
	return sqlDB
}
