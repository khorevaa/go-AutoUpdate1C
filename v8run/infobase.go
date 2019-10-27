package v8run

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
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
	UserOptions
	Path connectString
	Auth credentials
	ZN   []string
	UC   string
}

func NewInfobase(path connectString, opts ...UserOption) Infobase {

	ib := Infobase{
		Path: path,
	}

	return ib
}

func NewFileInfobase(path string, opts ...UserOption) Infobase {

	ib := NewInfobase(fileConnectString(path), opts...)

	return ib
}

func NewServerInfobase(server, base string, opts ...UserOption) Infobase {

	serverUrl, _ := url.Parse(server)
	port64, _ := strconv.ParseInt(serverUrl.Port(), 64, 10)
	port := int16(port64)

	if port == 0 {
		port = DEFAULT_1SSERVER_PORT
	}

	path := serverConnectString{
		Base:     base,
		Host:     serverUrl.Hostname(),
		Port:     port,
		Protocol: serverUrl.Scheme,
	}

	ib := NewInfobase(path, opts...)

	return ib
}

func (ib *Infobase) userOptionsString() string {

	return strings.Join(processArgs(ib.UserOptions), " ")
}

func (ib *Infobase) ConnectString() string {
	return fmt.Sprintf("%s %s", ib.Path.connectString(), ib.Auth.authString())
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
