package main

import (
	"go-app/study/distributed_transaction/dtm/common"
	"log"

	"github.com/dtm-labs/client/dtmcli"
)

func RunAPTx() {
	var (
		gid = common.NewGid()

		rm1TransOutUrl           = "http://localhost:28081/api/busi/trans_out"
		rm1TransOutCompensateUrl = "http://localhost:28081/api/busi/trans_out_compensate"
		rm1TransOutReq           = &common.ReqHTTP{Amount: 30, TransOutUserID: 1}

		rm2TransInUrl           = "http://localhost:28082/api/busi/trans_in"
		rm2TransInCompensateUrl = "http://localhost:28082/api/busi/trans_in_compensate"
		rm2TransInReq           = &common.ReqHTTP{Amount: 30, TransInUserID: 2}
	)

	saga := dtmcli.NewSaga(common.DTMServer, gid).
		Add(rm1TransOutUrl, rm1TransOutCompensateUrl, rm1TransOutReq).
		Add(rm2TransInUrl, rm2TransInCompensateUrl, rm2TransInReq)
	if err := saga.Submit(); err != nil {
		log.Printf(">> create saga transaction failed, err: %v", err)
	}
	log.Printf(">> create saga transaction success, gid: %s", gid)

}
