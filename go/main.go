package main

import (
    "net/http"
    "encoding/json"
    "log"
)

type Ping struct {
    Status int `json:"status"`
    Result string `json:"result"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    ping := Ping{http.StatusOK, "ok"} // 変更する
    // ping := Ping{http.StatusOK, "hello"} //　変更した
    // docker.htmlで受け取るjsonが変わって表示も変わる
    res, _ := json.Marshal(ping)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(res)
}


func main() {
    var httpServer http.Server
    http.HandleFunc("/", rootHandler)
    httpServer.Addr = ":8080"
    log.Println(httpServer.ListenAndServe())
}