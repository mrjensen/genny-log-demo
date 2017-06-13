package log

import (
	"os"

	"github.com/go-kit/kit/log"
)

var logger = log.NewJSONLogger(os.Stderr)

func Debug(args ...interface{}) {
	logger.Log("Debug: ", args)
}
