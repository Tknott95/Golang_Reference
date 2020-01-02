package analytics

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	ErrInvalidID    = errors.New("Invalid ID")
	ErrInvalidEmail = errors.New("Invalid email")
)

func CreateLogger(filename string) *log.Logger {
	file, err := os.OpenFile("./Analytics/"+filename+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) /*  Better number for prod app ( this reads and writes) */
	if err != nil {
		panic(err)
	}
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}

func Time(logger *log.Logger, next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		elapsed := time.Since(start)
		logger.Println(elapsed)
	})
}

/*  Recovers server if any coroutine panicks. This will allow for it to not crash it where it usually does */
func Recover(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				switch err {
				case ErrInvalidEmail:
					http.Error(w, ErrInvalidEmail.Error(), http.StatusUnauthorized)
				case ErrInvalidID:
					http.Error(w, ErrInvalidID.Error(), http.StatusUnauthorized)
				default:
					http.Error(w, "Unknown error, recovered from panic", http.StatusInternalServerError)
				}
			}
		}()
		next.ServeHTTP(w, req)
	})
}
