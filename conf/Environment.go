package conf

import (
	"database/sql"
	"github.com/gorilla/sessions"
	"gopkg.in/mailgun/mailgun-go.v1"
)

var Db *sql.DB
var Store *sessions.CookieStore
var MG mailgun.Mailgun
var EncryptionKey []byte
