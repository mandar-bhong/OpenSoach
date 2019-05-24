// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websockethelper

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	//"opensoach.com/core/logger"
	//	"text/template"
)

/* var addr = flag.String("addr", ":8080", "http service address") */
var addr *string

//var homeTempl = template.Must(template.ParseFiles("home.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found r", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//homeTempl.Execute(w, r.Host)
}

func Init(wss *WebsocketInitHelperStruct) error {

	serverWebSocketPort := fmt.Sprintf(":%d", wss.WebSocketPort)
	addr = flag.String("addr", serverWebSocketPort, "http service address")

	fmt.Printf("Address %s", addr)
	WebsocketInitHelperData = wss

	flag.Parse()
	go h.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)

	var serverStartErr error
	go func() {
		serverStartErr = http.ListenAndServe(*addr, nil)

	}()

	time.Sleep(time.Second * 2)

	if serverStartErr != nil {
		fmt.Printf("Error occured %s", serverStartErr.Error())
	}

	return serverStartErr
}

func CloseConnection(conn *WebSocketConnection) {
	if conn.ws == nil {
		//logger.Log(logger.CORESERVER, logger.WARNING, "WebSocketHelper conn.ws is nil")
		return
	}

	conn.ws.Close()
}
