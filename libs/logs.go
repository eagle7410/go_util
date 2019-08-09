package lib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const logfile = "development.log"

var logFileName string

func OpenDailyRotateLogFile() () {
	if err := MustDir("./logs", 0777); err != nil {
		log.Fatalf("Mkdir logs error %v", err)
	}

	now  := time.Now()
	next := time.Date(now.Year(),now.Month(),now.Day() + 1,0,0,0 , 0, now.Location())

	diff := next.Sub(now);

	fmt.Printf("diff is %v ", diff)

	logFileName = now.Format("2006-01-02") + "_" + logfile

	fmt.Printf("Logging to file %v\n", logFileName)

	lf, err := os.OpenFile("./logs/"+logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

	if err != nil {
		log.Fatalf("OpenLogfile: os.OpenFile: %s", err)
	}

	log.SetOutput(lf)

	go func() {
		time.Sleep(diff)
		OpenLogFile()
	}()
}

func OpenLogFile() {
	t := time.Now()
	oldFileName := logFileName

	logFileName = t.Format("2006-01-02") + "_" + logfile

	if err := MustDir("./logs", 0777); err != nil {
		log.Fatalf("Mkdir logs error %v", err)
	}

	if len(oldFileName) == 0 || oldFileName != logFileName {
		fmt.Printf("Logging to file %v\n", logFileName)

		lf, err := os.OpenFile("./logs/"+logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatalf("OpenLogfile: os.OpenFile: %s", err)
		}

		log.SetOutput(lf)
	}
}

func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SetCors(w, r)

		if r.Method == "OPTIONS" {
			return
		}

		OpenLogFile()

		if *Env.IsDev {
			log.Printf("req %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
			fmt.Printf("req %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		}

		handler.ServeHTTP(w, r)
	})
}

func LogAppRun(port string) {
	mask := "\n=== App run At (" + time.Now().Format("2006-01-02T15:04:05") + ") in http://localhost" + port + " ===\n"
	fmt.Printf(mask)
	log.Printf(mask)
}

func LogFatalf(format string, v ...interface{}) {
	format = "[ERR|FATAL]|" + format + "\n"
	fmt.Printf(format, v...)
	log.Fatalf(format, v...)
}

func Logf(format string, v ...interface{}) {
	format = format + "\n"
	fmt.Printf(format, v...)
	log.Printf(format, v...)
}

func LogEF(format string, v ...interface{}) {
	fmt.Printf("[0;31m"+format+"[39m\n", v...)
	log.Printf("[ERR]|"+format+"\n", v...)
}
