package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/utils"
)

func GetSession(ctx *gin.Context) {
	// 初始化错误返回的map
	resp := make(map[string]string)
	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ctx.JSON(http.StatusOK, resp)
}
