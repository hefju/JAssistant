package models
import (
    "github.com/go-xorm/xorm"
    "log"
    "github.com/go-xorm/core"
    "github.com/hefju/JAssistant/myconfig"
    _"github.com/mattn/go-sqlite3"
    "fmt"
    "time"
)

var engine *xorm.Engine

func init() {
    fmt.Println("model init()")
    var err error
    //	engine, err = xorm.NewEngine("odbc", "driver={SQL Server};Server=192.168.1.200; Database=charge; uid=sa; pwd=123;")
  //  engine, err = xorm.NewEngine("odbc", "driver={SQL Server};server=.;database=charge;integrated security=SSPI;")
    engine, err =  xorm.NewEngine("sqlite3", myconfig.AppRootPath+"/test.db")

    if err != nil {
        log.Fatalln("xorm create error", err)
    }
    //engine.ShowSQL = true
    engine.SetMapper(core.SameMapper{})
    // engine.CreateTables(new(tp_charge_billing))
    err = engine.Sync2(new(StatusReport)) //, new(Group))
    if err != nil {
        log.Fatalln("xorm sync error", err)
    }
}


//插入单条数据, 测试用的
func InsertStatus(status StatusReport)int64{
    affected, err := engine.Omit("Id").Insert(status)
    if err!=nil{
        log.Fatalln("insert StatusReport",err)
    }
    return affected
}


//报告类
type StatusReport struct  {
    Id int64
    From string      //发送人
    FromTime time.Time  //发送的时间
    Title string //标题(分类: 健康,统计的,)
    Content string //详细内容
    ServerTime time.Time //服务器时间
}

