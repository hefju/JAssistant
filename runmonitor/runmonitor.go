//用于检测目标设备运行情况, 如果超过一定时间没有状态报告发上来, 就会发送警示信息
package runmonitor
import (
    "time"
    "fmt"
    "github.com/hefju/JAssistant/models"
)

func Start(){
    go func(){
        tk:=time.NewTicker(time.Hour)
        for t := range tk.C {

          count:=  models. GetStatusCount(t)
            if count<5{
                fmt.Println("一个小时内状态报告小于5, 需要警报. SentEmail")
            }


           // fmt.Println(t)
        }
    }()
}
