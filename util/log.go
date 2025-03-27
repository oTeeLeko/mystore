package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func LogActivity(ctx *gin.Context) {
	now := time.Now()
	year := now.Format("2006")
	month := now.Format("Jan")
	fullDate := now.Format("2006-01-02")

	logDir := filepath.Join("./logs", year, month)
	logFile := filepath.Join(logDir, fullDate+".txt")

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Println("Error creating log directory:", err)
		return
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	ip := ctx.ClientIP()

	var params string
	if len(ctx.Request.URL.Query()) > 0 {
		params += "Query Params: " + ctx.Request.URL.Query().Encode() + " "
	}
	if len(ctx.Request.PostForm) > 0 {
		params += "Form Params: " + ctx.Request.PostForm.Encode() + " "
	}

	var body string
	if ctx.Request.Body != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(ctx.Request.Body)
		body = buf.String()
		ctx.Request.Body = ioutil.NopCloser(buf)
	}

	status := ctx.Writer.Status()

	logMessage := fmt.Sprintf("[%s] %s - Status: %d Method: %s Path: %s",
		now.Format("2006-01-02 15:04:05"), ip, status, method, path)

	if len(params) > 0 {
		logMessage += fmt.Sprintf(" Params: %s", params)
	}

	if len(body) > 0 {
		logMessage += fmt.Sprintf(" Body: %s", body)
	}

	logger := log.New(file, "", 0)
	logger.Println(logMessage)
}
