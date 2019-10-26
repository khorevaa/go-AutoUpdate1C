package v8run

import (
	"fmt"
	"strconv"
)

type credentials struct {
	User     string
	Password string
}

func (c credentials) authString() string {

	var auth string

	auth += "/N" + c.User

	if len(c.Password) > 0 {
		auth += "/P" + c.Password
	}

	return auth

}

type Infobase struct {
	cs   connectString
	auth credentials
	zn   []string
}

func (ib *Infobase) Args() []string {

	var args []string
	args = append(args, ib.cs.connectString())
	args = append(args, ib.auth.authString())
	return args
}

func (ib *Infobase) ConnectString() string {
	return fmt.Sprintf("%s %s", ib.cs.connectString(), ib.auth.authString())
}

type serverConnectString struct {
	Base     string
	Host     string `long:"host" description:"The host to connect to." default:"127.0.0.1"`
	Port     int16  `long:"port" description:"The port to connect to." default:"5432"`
	Protocol string
}

func (s serverConnectString) connectString() string {

	var serverAddr string

	if len(s.Protocol) > 0 {
		serverAddr += s.Protocol + "://"
	}

	serverAddr += s.Host

	if s.Port != 0 {
		serverAddr += ":" + strconv.FormatInt(int64(s.Port), 10)

	}

	return fmt.Sprintf("/S %s\\\"%s\"", serverAddr, s.Base)
}

type connectString interface {
	connectString() string
}

type fileConnectString string

func (f fileConnectString) connectString() string {
	return fmt.Sprintf("/F \"%s\"", string(f))
}

type wsConnectString string

func (ws wsConnectString) connectString() string {
	return fmt.Sprintf("/WS \"%s\"", string(ws))
}
