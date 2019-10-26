package v8run

type Enterprise struct {
	unlockCode             string
	startParams            string
	execute                string
	disableStartupDialogs  bool
	disableStartupMessages bool
	out                    string
	dumpResult             string

	userArgs map[string]string
}
