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

		rm1TransOutTryUrl     = "http://localhost:28081/api/busi/trans_out_try"
		rm1TransOutConfirmUrl = "http://localhost:28081/api/busi/trans_out_confirm"
		rm1TransOutCancelUrl  = "http://localhost:28081/api/busi/trans_out_cancel"
		rm1TransOutReq        = &common.ReqHTTP{Amount: 30, TransOutUserID: 1}

		rm2TransInTryUrl     = "http://localhost:28082/api/busi/trans_in_try"
		rm2TransInConfirmUrl = "http://localhost:28082/api/busi/trans_in_confirm"
		rm2TransInCancelUrl  = "http://localhost:28082/api/busi/trans_in_cancel"
		rm2TransInReq        = &common.ReqHTTP{Amount: 30, TransInUserID: 2}
	)

	tccFunc := func(tcc *dtmcli.Tcc) (*resty.Response, error) {
		resp, err := tcc.CallBranch(rm1TransOutReq, rm1TransOutTryUrl, rm1TransOutConfirmUrl, rm1TransOutCancelUrl)
		if err != nil {
			return resp, err
		}
		return tcc.CallBranch(rm2TransInReq, rm2TransInTryUrl, rm2TransInConfirmUrl, rm2TransInCancelUrl)
	}
	if err := dtmcli.TccGlobalTransaction(common.DTMServer, gid, tccFunc); err != nil {
		log.Printf(">> create tcc transaction failed, err: %v", err)
	}
	log.Printf(">> create tcc transaction success, gid: %s", gid)
}
