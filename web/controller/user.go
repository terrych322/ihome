package controller

import (
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"image/color"
	"image/png"
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

func GetImageCd(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	// 生成图片验证码
	cap := captcha.New()
	cap.SetFont("./config/comic.ttf")
	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})
	img, str := cap.Create(4, captcha.NUM)
	png.Encode(ctx.Writer, img)
	fmt.Println("str=", str)
	fmt.Println("uuid=", uuid)
}
