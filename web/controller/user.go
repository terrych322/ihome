package controller

import (
	"fmt"
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

func GetImageCd(ctx *gin.Context)  {
	uuid := ctx.Param("uuid")
	fmt.Println("uuid=", uuid)
}
