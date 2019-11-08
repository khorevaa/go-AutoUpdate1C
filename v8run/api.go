package v8run

import (
	"github.com/khorevaa/go-AutoUpdate1C/v8run/types"
	"io/ioutil"
)

func LoadCfg(file string, opts ...types.UserOption) *LoadCfgOptions {

	command := &LoadCfgOptions{
		File:     file,
		Designer: newDefaultDesigner(),
	}

	processOptions(command, opts)

	return command

}

func UpdateCfg(file string, force bool, opts ...types.UserOption) *UpdateCfgOptions {

	command := &UpdateCfgOptions{
		File:     file,
		Force:    force,
		Designer: newDefaultDesigner(),
	}

	processOptions(command, opts)

	return command

}

func LoadExtensionCfg(file, extension string, opts ...types.UserOption) *LoadCfgOptions {

	command := LoadCfg(file, opts...)
	command.Extension = extension

	return command

}

func DumpCfg(file string, opts ...types.UserOption) *DumpCfgOptions {

	command := &DumpCfgOptions{
		File:     file,
		Designer: newDefaultDesigner(),
	}

	processOptions(command, opts)

	return command

}

func DumpExtensionCfg(file, extension string, opts ...types.UserOption) *DumpCfgOptions {

	command := DumpCfg(file, opts...)
	command.Extension = extension
	return command

}

func UpdateDBCfg(server bool, Dynamic bool, opts ...types.UserOption) *UpdateDBCfgOptions {

	command := &UpdateDBCfgOptions{
		Designer: newDefaultDesigner(),
		Use:      true,
		Server:   server,
		Dynamic:  Dynamic,
	}

	processOptions(command, opts)

	return command

}

func UpdateDBExtensionCfg(extension string, server bool, Dynamic bool, opts ...types.UserOption) *UpdateDBCfgOptions {

	command := UpdateDBCfg(server, Dynamic, opts...)
	command.Extension = extension

	return command

}

func DumpIB(file string, opts ...types.UserOption) *DumpIBOptions {

	command := &DumpIBOptions{
		Designer: newDefaultDesigner(),
		File:     file,
	}

	processOptions(command, opts)
	return command
}

func RestoreIB(file string, opts ...types.UserOption) *RestoreIBOptions {

	command := &RestoreIBOptions{
		Designer: newDefaultDesigner(),
		File:     file,
	}

	processOptions(command, opts)

	return command
}

func CreateTempInfoBase(opts ...types.UserOption) *CreateInfoBaseOptions {

	return CreateInfoBase(opts...)

}

func CreateInfoBase(opts ...types.UserOption) *CreateInfoBaseOptions {

	command := newDefaultCreateInfoBase()

	processOptions(command, opts)

	return command

}

func Execute(file string, opts ...types.UserOption) *ExecuteOptions {

	command := &ExecuteOptions{
		Enterprise: newDefaultEnterprise(),
		File:       file,
	}

	processOptions(command, opts)

	return command
}

////////////////////////////////////////////////////////
// Доступные опции

func WithStartParams(params string) types.UserOption {
	return func(o types.Optioned) {
		o.SetOption("/C", params)
	}
}

func WithUnlockCode(uc string) types.UserOption {
	return func(o types.Optioned) {
		o.SetOption("/UC", uc)
	}
}

func WithUpdateDBCfg() types.UserOption {
	return func(o types.Optioned) {
		o.SetOption("/UpdateDBCfg", true)
	}
}

func WithUpdateDBCfgOptions(options *UpdateDBCfgOptions) types.UserOption {
	return func(o types.Optioned) {
		o.SetOption("/UpdateDBCfg", options)
	}
}

func WithExtension(ext string) types.UserOption {
	return func(o types.Optioned) {
		o.SetOption("-Extension", ext)
	}
}

func WithManagedApplication() types.UserOption {
	return func(o types.Optioned) {
		o.SetOption("/RunModeManagedApplication", true)
	}
}

func WithCredentials(user, password string) types.UserOption {
	return func(o types.Optioned) {

		if len(user) == 0 {
			return
		}

		o.SetOption("/U", user)

		if len(password) > 0 {
			o.SetOption("/P", user)
		}

	}
}

////////////////////////////////////////////////////////
// Create InfoBases

func NewFileIB(path string, opts ...types.UserOption) *FileInfoBase {

	ib := &FileInfoBase{
		baseInfoBase: baseInfoBase{},
		File:         path,
	}

	return ib
}

func NewTempIB(opts ...types.UserOption) *FileInfoBase {

	path, _ := ioutil.TempDir("", "1c_DB_")

	ib := NewFileIB(path, opts...)

	return ib
}

func NewServerIB(server, base string, opts ...types.UserOption) *ServerInfoBase {

	ib := &ServerInfoBase{
		baseInfoBase: baseInfoBase{},
		Srvr:         server,
		Ref:          base,
	}

	return ib
}

func NewWebServerIB(ref string, opts ...types.UserOption) *WSInfoBase {

	ib := &WSInfoBase{
		baseInfoBase: baseInfoBase{},
		Ref:          ref,
	}

	return ib
}
