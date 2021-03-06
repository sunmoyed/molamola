package main

import (
	"flag"

	"github.com/sunmoyed/molamola/server/pkg/log"
	"github.com/sunmoyed/molamola/server/pkg/server"
)

var logger = log.DefaultLogger

func main() {
	var addr string
	var datadir string
	var webdir string

	flag.StringVar(&addr, "addr", ":4477", "Address to listen on")
	flag.StringVar(&datadir, "data", "data", "Data directory")
	flag.StringVar(&webdir, "webdir", "web/dist", "Web directory")

	flag.Parse()

	sv, svErr := server.NewServer(addr, datadir, webdir)
	if svErr != nil {
		logger.Fatal("new server error: ", svErr)
	}
	logger.Fatal("server run error: ", sv.Run())
}
