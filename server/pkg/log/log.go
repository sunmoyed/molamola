package log

import (
	stdlog "log"
	"os"
)

var DefaultLogger *stdlog.Logger = stdlog.New(os.Stdout, "", stdlog.Lshortfile|stdlog.Ldate|stdlog.Ltime)
