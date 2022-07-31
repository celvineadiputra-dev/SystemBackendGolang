package Middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

//func LogMiddleware() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		f, err := os.OpenFile("logs/loggingMiddleware-"+time.Now().Format("01-02-2006")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//		if err != nil {
//			log.Fatalf("error opening file: %v", err)
//		}
//		defer f.Close()
//
//		log.SetOutput(f)
//		//ByteBody, _ := ioutil.ReadAll(context.Request.Body)
//		//body, _ := ioutil.ReadAll(context.Request.Body)
//		log.Printf("IP [%s], Access Route : [%s%s] Request : \n",
//			context.ClientIP(),
//			context.Request.Host,
//			context.Request.RequestURI)
//	}
//}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// read request
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

		//read response
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Request.Body = rdr2
		c.Next()

		f, err := os.OpenFile("logs/loggingMiddleware-"+time.Now().Format("01-02-2006")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)

		containRequest := readBody(rdr1)
		if strings.Contains(c.Request.RequestURI, "login") || strings.Contains(c.Request.RequestURI, "register") {
			containRequest = "-"
		}

		log.Printf("IP [%s], Access Route : [%s%s] Request : %s Response : %s\n",
			c.ClientIP(),
			c.Request.Host,
			c.Request.RequestURI,
			containRequest,
			blw.body.String())
	}
}
