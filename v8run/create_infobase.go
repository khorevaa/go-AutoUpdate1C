package v8run

import "github.com/khorevaa/go-AutoUpdate1C/v8run/types"

type CreateInfoBaseOptions struct {
	types.UserOptions

	disableStartupDialogs  bool
	disableStartupMessages bool
	visible                bool
}

func (d *CreateInfoBaseOptions) Command() string {
	return COMMAND_CREATEINFOBASE
}

func (d *CreateInfoBaseOptions) Check() bool {

	return true
}

func (d *CreateInfoBaseOptions) Values() (values types.UserOptions) {

	values = make(map[string]interface{})

	values.Append(d.UserOptions)

	values.SetOption("/DisableStartupDialogs", d.disableStartupDialogs)
	values.SetOption("/DisableStartupDialogs", d.disableStartupDialogs)
	values.SetOption("/Visible", d.visible)

	return values

}

func newDefaultCreateInfoBase() *CreateInfoBaseOptions {

	d := &CreateInfoBaseOptions{
		disableStartupDialogs:  true,
		disableStartupMessages: true,
		visible:                false,
	}

	return d
}
