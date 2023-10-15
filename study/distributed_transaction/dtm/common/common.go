package common

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dtm-labs/client/dtmcli/dtmimp"

	"github.com/dtm-labs/client/dtmcli/logger"
	uuid "github.com/satori/go.uuid"
)

const (
	DTMServer = "http://localhost:36789/api/dtmsvr"
)

var BasicDBConf = dtmimp.DBConf{
	Driver:   "mysql",
	Host:     "localhost",
	Port:     3306,
	User:     "root",
	Password: "123456",
}

type ReqHTTP struct {
	Amount         int `json:"amount"`
	TransOutUserID int `json:"trans_out_user_id"`
	TransInUserID  int `json:"trans_in_user_id"`
}

type ResHTTP struct {
	Result bool        `json:"result"`
	Data   interface{} `json:"data"`
	Err    error       `json:"err"`
}

func NewGid() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func EnableDtmcliLogger() {
	logger.InitLog("debug")
}

func WaitExit() {
	// Ctrl+C exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("quit (%v)\n", <-sig)
}
