package network

import (
	"github.com/gorilla/sessions"
	"github.com/super-link-manager/utils"
	"net/http"
)

var Key = "FLASH-SESSION"
var Store = sessions.NewCookieStore([]byte(Key))

func Set(w http.ResponseWriter, r *http.Request, message string, messageType string) {
	session, err := Store.Get(r, Key)
	utils.CheckErr(err)

	session.AddFlash(message, "message")
	session.AddFlash(messageType, "message")
	session.Save(r, w)
}

func Get(w http.ResponseWriter, r *http.Request) *Message {
	session, err := Store.Get(r, Key)
	utils.CheckErr(err)

	fm := session.Flashes("message")
	if fm == nil {
		return nil
	}
	session.Save(r, w)

	return &Message{Text: fm[0].(string), Type: fm[1].(string)}
}

type Message struct {
	Text string
	Type string
}