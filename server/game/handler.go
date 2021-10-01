package game

import (
	"fmt"
	"github.com/responserms/response/response"
	"github.com/tidwall/redcon"
	"time"
)

type handler struct {
	core    *response.Core
	players map[string]time.Time
}

func (h *handler) ping(conn redcon.Conn, cmd redcon.Command) {
	conn.WriteString("PONG")
}

func (h *handler) quit(conn redcon.Conn, cmd redcon.Command) {
	conn.WriteString("OK")
	conn.Close()
}

func (h *handler) join(conn redcon.Conn, cmd redcon.Command) {
	fmt.Println(string(cmd.Args[1]))

	h.players[string(cmd.Args[1])] = time.Now()
	conn.WriteString("OK")
}

func (h *handler) print(conn redcon.Conn, cmd redcon.Command) {
	conn.WriteArray(len(h.players))

	for k, v := range h.players {
		conn.WriteString(fmt.Sprintf("%s: %s", k, v.String()))
	}
}

func (h *handler) detach(conn redcon.Conn, cmd redcon.Command) {
	d := conn.Detach()
	go func(c redcon.DetachedConn) {
		defer c.Close()

		c.WriteString("OK")
		c.Flush()
	}(d)
}
