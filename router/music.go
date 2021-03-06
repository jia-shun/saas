package router

import (
	"github.com/gin-gonic/gin"
	"music-saas/api"
)

func InitMusicRouter(Router *gin.RouterGroup) {
	MusicRouter := Router.Group("music")
	{
		MusicRouter.GET(":id", api.GetMusicById)                      // 获取音乐详情
		MusicRouter.POST("list", api.GetMusic)                        // 音乐任务列表
		MusicRouter.POST("", api.CreateMusic)                         // 增加音乐任务
		MusicRouter.PUT("", api.UpdateMusic)                          // 编辑音乐任务
		MusicRouter.PUT("finish-status", api.UpdateMusicFinishStatus) // 编辑音乐完成状态
		MusicRouter.PUT("pay-status", api.UpdateMusicPayStatus)       // 编辑音乐支付状态
		MusicRouter.DELETE(":id", api.DeleteMusic)                    // 删除音乐任务
	}
}
