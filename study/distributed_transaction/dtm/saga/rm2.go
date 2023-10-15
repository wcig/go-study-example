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
	rm2DB *sql.DB
)

func RunRM2Server() {
	rm2DB = InitDB()
	InitRM2App()
	time.Sleep(100 * time.Millisecond)
}

func InitRM2App() {
	app := gin.Default()
	app.POST("/api/busi/trans_in", transInHandler)
	app.POST("/api/busi/trans_in_compensate", transInCompensageHandler)
	go func() {
		if err := app.Run(":28082"); err != nil {
			panic(err)
		}
	}()
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

func transInCompensageHandler(c *gin.Context) {
	if err := transInCompensage(c); err != nil {
		log.Printf(">> rm2-server transin compensage handler failed, err: %v", err)
		c.JSON(http.StatusConflict, &common.ResHTTP{Result: false, Err: err})
		return
	}
	log.Println(">> rm2-server transin compensage handler success")
	c.JSON(http.StatusOK, &common.ResHTTP{Result: true})
}

func transIn(c *gin.Context) error {
	// return errors.New("my transin test error")

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
		affected, err2 := dtmimp.DBExec(common.BasicDBConf.Driver, rm2DB, sqlFormat, req.Amount, req.TransInUserID)
		if err2 != nil {
			return err
		}
		log.Printf(">> rm2-server transin db success, result: %d", affected)
		return nil
	}
	return barrier.CallWithDB(rm2DB, busiCall)
}

func transInCompensage(c *gin.Context) error {
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
		affected, err2 := dtmimp.DBExec(common.BasicDBConf.Driver, rm2DB, sqlFormat, req.Amount, req.TransInUserID)
		if err2 != nil {
			return err
		}
		log.Printf(">> rm2-server transin compensage db success, result: %d", affected)
		return nil
	}
	return barrier.CallWithDB(rm2DB, busiCall)
}
