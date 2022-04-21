package master
import "github.com/gin-gonic/gin"


func InitRouter(r *gin.Engine)  {
	r.GET("/job/save",HanleJobSave)

}


//保存任务端口
func HanleJobSave (c *gin.Context) {


}

