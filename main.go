package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/donnie4w/go-logger/logger"
    "net/http"
    "time"
    "github.com/hefju/JAssistant/myconfig"
    "github.com/hefju/JAssistant/models"
    "github.com/hefju/JAssistant/runmonitor"
)
func main(){
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
    router.GET("/", func(c *gin.Context) { //测试，获取数据表信息
        // log.Println("visit homepage\r\n")
        logger.Debug("visit homepage")
        c.String(http.StatusOK, "Welcome to "+myconfig.MyName+"...")
    })


    //获取服务器时间
    router.GET("/time", func(c *gin.Context) {
        c.String(http.StatusOK, time.Now().Format("2006-01-02 15:04:05"))
    })

    router.GET("/index", func(c *gin.Context) {
        obj := gin.H{"title": "JAssistant"}
        c.HTML(http.StatusOK, "index.html", obj)
    })

    router.POST("/report", receiveReport)       //批量上传数据

    runmonitor.Start()//运行状态监视器

    router.Run(":8085")
    fmt.Println("end")
}

//批量插入数据, 正式使用
func receiveReport(c *gin.Context) {
    var json models.StatusReport
    c.Bind(&json)
    json.ServerTime=time.Now().Unix() //设置服务器本地时间
    count := models.InsertStatus(json)
    logger.Debug("upload插入结果:", count)
    result:="JAssistant 操作失败"
    if count>0{
        result="JAssistant 操作成功"
    }
    c.String(http.StatusOK, result) //返回结果, 插入了多少条记录
}