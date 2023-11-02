package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	HttpTraceId = "X-Trace-ID"
)

const (
	CtxRequestStart = "ctx_request_start"
	CtxTraceId      = "ctx_trace_id"
	CtxLoginUserId  = "ctx_login_user_id"
)

func SetRequestStartTime(c *gin.Context) {
	c.Set(CtxRequestStart, time.Now())
}

func GetRequestStartTime(c *gin.Context) time.Time {
	return c.GetTime(CtxRequestStart)
}

func SetTraceId(c *gin.Context) {
	value := c.GetHeader(HttpTraceId)
	if value == "" {
		value = uuid.NewString()
	}

	c.Set(CtxTraceId, value)
}

func GetTraceId(c *gin.Context) string {
	return c.GetString(CtxTraceId)
}

func SetLoginUserId(c *gin.Context, userId int) {
	c.Set(CtxLoginUserId, userId)
}

func GetLoginUserId(c *gin.Context) int {
	return c.GetInt(CtxLoginUserId)
}
