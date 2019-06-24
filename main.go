package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/t-tomalak/logrus-easy-formatter"
	"os"
)

func main() {
	filename := "app.log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	logr := &logrus.Logger{
		Out: os.Stderr,
        // Level: logrus.WarnLevel,
        Level: logrus.DebugLevel,
        Formatter: &easy.Formatter{
            TimestampFormat: "2006-01-02 15:04:05",
            LogFormat:       "%time% - [%lvl%]:   %msg%\n",
        },
    }

	
	logr.SetOutput(f)

	logr.Debug("Just boring development stuff.")
	logr.Info("This is a nice info.")
	logr.Warning("Alert! Alert!")
	logr.Error("Something very bad has happened!")
}
