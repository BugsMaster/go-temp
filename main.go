package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"temp/global"
	"temp/initialize"
	"temp/lib"
	"time"
)
var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

//webSocket请求ping 返回pong
func ping(w http.ResponseWriter , r *http.Request) {
	neTicker:=time.NewTicker(time.Second*3)
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(w,r,nil)
	if err != nil {
		return
	}
	defer ws.Close()
	ws.SetCloseHandler(func(code int, text string) error {
		fmt.Println(code, text)
		return nil
	})
	go func() {
		for {
			nowTime:= <-neTicker.C
			ws.WriteMessage(websocket.TextMessage, []byte(time.Unix(nowTime.Unix(), 0).Format("2006-01-02 15:04:05")+"  终于下班了"))
		}
	}()
	//接受前端发送的消息
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		log.Printf("recv: %s", message)
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func main() {
	global.GVA_VP = lib.Viper() // 初始化Viper
	global.GVA_LOG = lib.Zap()//初始化zap
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	//initialize.Timer()
	if global.GVA_DB != nil {
		//initialize.MysqlTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	go Socket()
	GinRouterStart()
}
func Socket(){
	http.HandleFunc("/ping",ping)
	http.ListenAndServe("0.0.0.0:10104",nil)
}
func GinRouterStart() {
	gin.SetMode(gin.ReleaseMode)
	router:=initialize.Routers()
	//开启服务
	if err := router.Run(":10103"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
//type Msg struct {
//	UserId    string   `json:"userId"`
//	Text      string   `json:"text"`
//	State     string   `json:"state"`
//	Namespace string   `json:"namespace"`
//	Rooms     []string `json:"rooms"`
//}
//gin.SetMode(gin.ReleaseMode)
//router := gin.Default()
//router.GET("/ping", ping)
//if err := router.Run(":10104"); err != nil {
//	fmt.Println("startup service failed, err:%v\n", err)
//}

//runtime.Goexit()
//runtime.Gosched()
//neTicker:=time.NewTicker(time.Second)
//go func() {
//	//for {
//		nowTime:= <-neTicker.C
//		fmt.Println(nowTime)
//	//}
//	for i:=0;i<9;i++{
//	}
//}()
//type User struct {
//	//gorm.Model
//}