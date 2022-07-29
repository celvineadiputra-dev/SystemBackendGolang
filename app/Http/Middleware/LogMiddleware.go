package Middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func LogMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		f, err := os.OpenFile("logs/loggingMiddleware-"+time.Now().Format("01-02-2006")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		//ByteBody, _ := ioutil.ReadAll(context.Request.Body)
		log.Printf("IP [%s], Access Route : [%s%s]\n",
			context.ClientIP(),
			context.Request.Host,
			context.Request.URL.Path)
	}
}
