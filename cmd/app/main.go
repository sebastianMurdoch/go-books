package main

import (
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
)

var newRelicApp *newrelic.Application

func main() {
	r := gin.Default()
	r.GET("/ping", HandlerFunc)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func init()  {
	var err error
	newRelicApp, err = newrelic.NewApplication(
		newrelic.ConfigAppName("newRelicApp"),
		newrelic.ConfigLicense("0cfe72a4df09c8d5723a142b8040134281b8NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		panic(err)
	}
}

func HandlerFunc(c *gin.Context)  {
	txn := newRelicApp.StartTransaction("status")
	defer txn.End()
	c.String(http.StatusOK, "pong")
}

