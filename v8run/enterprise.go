package v8run

type Enterprise struct {
	UserOptions

	disableSplash          bool
	disableStartupDialogs  bool
	disableStartupMessages bool
}

func (d Enterprise) Values() (values UserOptions) {

	values = make(map[string]interface{})

	values.Append(d.UserOptions)

	values.setOption("/DisableStartupDialogs", d.disableStartupDialogs)
	values.setOption("/DisableStartupDialogs", d.disableStartupDialogs)

	return values

}

func (d Enterprise) Command() string {
	return COMMAND_ENTERPRISE
}

func (d Enterprise) Check() bool {

	return true
}

func WithStartParams(params string) UserOption {
	return func(o Optioned) {
		o.setOption("/C", params)
	}
}

func NewEnterprise(opts ...UserOption) Enterprise {

	d := Enterprise{
		UserOptions: make(map[string]interface{}),
	}

	for _, opt := range opts {
		d.Option(opt)
	}

	return d
}

func newDefaultEnterprise() Enterprise {

	d := Enterprise{
		disableStartupDialogs:  true,
		disableStartupMessages: true,
		disableSplash:          true,
	}

	return d
}

///Execute <имя файла внешней обработки>
//— предназначен для запуска внешней обработки в режиме "1С:Предприятие"
// непосредственно после старта системы.
//
type ExecuteOptions struct {
	Enterprise
	File string
}

func (d ExecuteOptions) Values() (values UserOptions) {

	values = d.Enterprise.Values()
	values["/Execute"] = d.File

	return

}
