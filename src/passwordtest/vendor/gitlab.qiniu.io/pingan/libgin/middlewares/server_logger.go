package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/color"
	"gitlab.qiniu.io/pingan/libbase/loggers"
	"gitlab.qiniu.io/pingan/libserver/colorful"
)

type GetLogger func(*gin.Context) loggers.Logger

func ServerLogger(getLogger GetLogger) gin.HandlerFunc {
	return func(context *gin.Context) {
		logger := getLogger(context)
		colorizer := colorful.StdoutColor
		path := context.Request.URL.RequestURI()
		clientIP := context.ClientIP()
		method := context.Request.Method

		if authorization := context.Request.Header.Get("Authorization"); authorization != "" {
			logger.Infof(
				"[Start] %s %s | %s | %s",
				colorfulMethod(method, colorizer), path, clientIP, authorization)
		} else {
			logger.Infof(
				"[Start] %s %s | %s",
				colorfulMethod(method, colorizer), path, clientIP)
		}

		var (
			body []byte
			err  error
		)

		if body, err = ioutil.ReadAll(context.Request.Body); err != nil {
			logger.Errorf("Failed to read request body for server logging: %s", err)
			context.Abort()
			return
		} else if err = context.Request.Body.Close(); err != nil {
			logger.Errorf("Failed to close request body: %s", err)
			context.Abort()
			return
		}

		context.Request.Body = ioutil.NopCloser(bytes.NewReader(body)) // Put request body back
		if len(body) > 0 {
			logger.Infof("  Body: %s", body)
		}

		start := time.Now()
		context.Next()
		end := time.Now()
		latency := end.Sub(start)

		statusCode := context.Writer.Status()
		logger.Infof(
			"[Performed] %s | %v",
			colorfulStatus(statusCode, colorizer),
			latency)

		comment := context.Errors.ByType(gin.ErrorTypePrivate).String()
		if len(comment) > 0 {
			logger.Infof("  Comment: %s", comment)
		}
	}
}

func colorfulStatus(code int, colorizer *color.Color) string {
	str := fmt.Sprintf(" %3d ", code)
	switch {
	case code >= 200 && code < 300:
		return colorizer.GreenBg(str)
	case code >= 300 && code < 400:
		return colorizer.WhiteBg(str)
	case code >= 400 && code < 500:
		return colorizer.YellowBg(str)
	default:
		return colorizer.RedBg(str)
	}
}

func colorfulMethod(method string, colorizer *color.Color) string {
	var colorized string

	toBeColorized := fmt.Sprintf(" %s ", method)
	switch method {
	case "GET":
		colorized = colorizer.BlueBg(toBeColorized)
	case "POST":
		colorized = colorizer.CyanBg(toBeColorized)
	case "PUT":
		colorized = colorizer.YellowBg(toBeColorized)
	case "DELETE":
		colorized = colorizer.RedBg(toBeColorized)
	case "PATCH":
		colorized = colorizer.GreenBg(toBeColorized)
	case "HEAD":
		colorized = colorizer.MagentaBg(toBeColorized)
	default:
		colorized = colorizer.WhiteBg(toBeColorized)
	}
	return fmt.Sprintf("%-7s", colorized)
}
