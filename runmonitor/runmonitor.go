//用于检测目标设备运行情况, 如果超过一定时间没有状态报告发上来, 就会发送警示信息
package runmonitor
import (
    "time"
    "fmt"
)

func Start(){
    go func(){
        tk:=time.NewTicker(time.Second)
        for t := range tk.C {

            //
            fmt.Println(t)
        }
    }()
}
