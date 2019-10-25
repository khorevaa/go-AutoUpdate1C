package v8run

type Designer struct {
	baseRunner
	infobase               infobase
	unlockCode             string
	disableStartupDialogs  bool
	disableStartupMessages bool
	visible                bool

	userArgs map[string]string
}

func (d Designer) command() string {
	return commandDesigner
}

func (d Designer) check() bool {

	return true
}

func (d Designer) args() (args []string) {

	return

}
