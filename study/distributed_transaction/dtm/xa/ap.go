package main

import (
	"go-app/study/distributed_transaction/dtm/common"
	"log"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/go-resty/resty/v2"
)

func RunAPTx() {
	var (
		gid = common.NewGid()

		rm1TransOutUrl = "http://localhost:28081/api/busi/trans_out"
		rm1TransOutReq = &common.ReqHTTP{Amount: 30, TransOutUserID: 1}

		rm2TransInUrl = "http://localhost:28082/api/busi/trans_in"
		rm2TransInReq = &common.ReqHTTP{Amount: 30, TransInUserID: 2}
	)

	xaFunc := func(xa *dtmcli.Xa) (*resty.Response, error) {
		resp, err := xa.CallBranch(rm1TransOutReq, rm1TransOutUrl)
		if err != nil {
			return resp, err
		}
		return xa.CallBranch(rm2TransInReq, rm2TransInUrl)
	}
	if err := dtmcli.XaGlobalTransaction(common.DTMServer, gid, xaFunc); err != nil {
		log.Printf(">> create xa transaction failed, err: %v", err)
	}
	log.Printf(">> create xa transaction success, gid: %s", gid)
}
