package game

import (
	"fmt"
	"time"

	"github.com/responserms/response/response"
	"github.com/tidwall/redcon"
)

func ListenAndServe(addr string, core *response.Core) error {
	return redcon.ListenAndServe(
		addr, func(conn redcon.Conn, cmd redcon.Command) {
			handler := conn.Context().(*handler)

			mux := redcon.NewServeMux()
			mux.HandleFunc("ping", handler.ping)
			mux.HandleFunc("quit", handler.quit)
			mux.HandleFunc("join", handler.join)
			//mux.HandleFunc("drop", handler.drop)
			mux.HandleFunc("print", handler.print)
			mux.HandleFunc("detach", handler.detach)

			mux.ServeRESP(conn, cmd)
		}, handleConnect(core), handleDisconnect(core))
}

func handleConnect(core *response.Core) func(redcon.Conn) bool {
	return func(conn redcon.Conn) bool {
		handler := &handler{
			core:    core,
			players: map[string]time.Time{},
		}

		fmt.Println("new conn?")
		conn.SetContext(handler)

		return true
	}
}

func handleDisconnect(core *response.Core) func(redcon.Conn, error) {
	return func(conn redcon.Conn, err error) {
		fmt.Println("detach")
	}
}
