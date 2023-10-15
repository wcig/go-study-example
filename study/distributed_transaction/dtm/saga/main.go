package main

import "go-app/study/distributed_transaction/dtm/common"

func main() {
	// dtm client logger
	common.EnableDtmcliLogger()

	// run rm1, rm2
	RunRM1Server()
	RunRM2Server()

	// run ap
	RunAPTx()

	// wait exit
	common.WaitExit()
}
