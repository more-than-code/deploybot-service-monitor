package task

import (
	"bytes"
	"net/http"

	"deploybot-service-monitor/util"

	"github.com/docker/docker/pkg/stdcopy"
	"github.com/gin-gonic/gin"
)

type Runner struct {
	cHelper *util.ContainerHelper
}

func NewRunner() *Runner {
	return &Runner{util.NewContainerHelper()}
}

func (s *Runner) ServiceLogHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Query("name")
		out, err := s.cHelper.LogContainer(ctx, name)

		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		var buf bytes.Buffer

		_, err = stdcopy.StdCopy(&buf, &buf, out)

		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.String(http.StatusOK, buf.String())
	}
}

func (r *Runner) HealthCheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
