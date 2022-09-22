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
	CheckOrigin: func(r *http.Request) bool {
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
			nowTime:= <- neTicker.C
			info:= &map[string]string{
				"周一":"1",
				"周二":"2",
				"周三":"3",
				"周四":"4",
				"周五":"5",
				"周六":"6",
				"周日":time.Unix(nowTime.Unix(), 0).Format("2006-01-02 15:04:05"),
			}
			//res,_ :=http.Get("https://quotedata.cnfin.com/quote/v1/kline?localDate=1663233097386&get_type=offset&prod_code=510310.SS&candle_period=6&candle_mode=1&fields=open_px%2Chigh_px%2Clow_px%2Cclose_px%2Cbusiness_amount%2Cbusiness_balance%2Cturnover_ratio%2Cpreclose_px%2Cspecial_close_px%2Csettlement%2Camount&data_count=1")
			/*res,_ :=http.Get("https://quotedata.cnfin.com/quote/v1/kline?localDate=1663656817701&get_type=offset&prod_code=562800.SS&candle_period=1&candle_mode=1&fields=open_px%2Chigh_px%2Clow_px%2Cclose_px%2Cbusiness_amount%2Cbusiness_balance%2Cturnover_ratio%2Cpreclose_px%2Cspecial_close_px%2Csettlement%2Camount&data_count=400&date=20220913&min_time=1501")
			defer res.Body.Close()
			body,_ := ioutil.ReadAll(res.Body)
			var actress model.ETFHistoryList
			err := json.Unmarshal(body,&actress)
			if err != nil {
				fmt.Println(err)
			}*/
			//fmt.Print(actress)
			//response.OkWithDetailed(actress,"获取到了大批数据",c)
			ws.WriteJSON(*info)
			//ws.WriteJSON(actress)
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
	global.GVA_VP = lib.Viper()       // 初始化Viper
	global.GVA_LOG = lib.Zap()        //初始化zap
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	//initialize.Timer()
	initialize.Redis()
	if global.GVA_DB != nil {
		//initialize.MysqlTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	go Socket()
	GinRouterStart()
}
func Socket() {
	http.HandleFunc("/ping", ping)
	// 设置监听的端口
	err := http.ListenAndServe(global.GVA_CONFIG.ServerInfo.SocketIp, nil)
	if err != nil {
		fmt.Print(err)
	}
}
func GinRouterStart() {
	gin.SetMode(gin.ReleaseMode)
	router := initialize.Routers()
	//开启服务
	if err := router.Run(global.GVA_CONFIG.ServerInfo.Port); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}


