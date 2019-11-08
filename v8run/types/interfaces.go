package types

type InfoBase interface {
	Path() string
	ShortConnectString() string
	IBConnectionString() (string, error)
	CreateString() (string, error)
}

type Command interface {
	Command() string
	Values() (values UserOptions)
	Check() bool
}

type Optioned interface {
	SetOption(key string, value interface{})
	Values() (values UserOptions)
	Option(opt interface{})
}
