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
	app.POST("/api/busi/trans_out", transOutHandler)
	app.POST("/api/busi/trans_out_compensate", transOutCompensageHandler)
	go func() {
		if err := app.Run(":28081"); err != nil {
			panic(err)
		}
	}()
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

func transOutCompensageHandler(c *gin.Context) {
	if err := transOutCompensage(c); err != nil {
		log.Printf(">> rm1-server transout compensage handler failed, err: %v", err)
		c.JSON(http.StatusConflict, &common.ResHTTP{Result: false, Err: err})
		return
	}
	log.Println(">> rm1-server transout compensage handler success")
	c.JSON(http.StatusOK, &common.ResHTTP{Result: true})
}

func transOut(c *gin.Context) error {
	// return errors.New("my transout test error")

	var (
		req     *common.ReqHTTP
		barrier *dtmcli.BranchBarrier
		err     error
	)

	if req, err = ParseReq(c); err != nil {
		return err
	}
	barrier = MustBarrierFromGin(c)
	busiCall := func(tx *sql.Tx) error {
		sqlFormat := "update dtm_busi.user_account set balance = balance - ? where user_id = ?"
		affected, err2 := dtmimp.DBExec(common.BasicDBConf.Driver, rm1DB, sqlFormat, req.Amount, req.TransOutUserID)
		if err2 != nil {
			return err
		}
		log.Printf(">> rm1-server transout db success, result: %d", affected)
		return nil
	}
	return barrier.CallWithDB(rm1DB, busiCall)
}

func transOutCompensage(c *gin.Context) error {
	var (
		req     *common.ReqHTTP
		barrier *dtmcli.BranchBarrier
		err     error
	)

	if req, err = ParseReq(c); err != nil {
		return err
	}
	barrier = MustBarrierFromGin(c)
	busiCall := func(tx *sql.Tx) error {
		sqlFormat := "update dtm_busi.user_account set balance = balance + ? where user_id = ?"
		affected, err2 := dtmimp.DBExec(common.BasicDBConf.Driver, rm1DB, sqlFormat, req.Amount, req.TransOutUserID)
		if err2 != nil {
			return err
		}
		log.Printf(">> rm1-server transout compensage db success, result: %d", affected)
		return nil
	}
	return barrier.CallWithDB(rm1DB, busiCall)
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
