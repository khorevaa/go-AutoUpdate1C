package v8run

func LoadCfg(file string, opts ...interface{}) *LoadCfgOptions {

	command := &LoadCfgOptions{
		File:     file,
		Designer: newDefaultDesigner(),
	}

	processOptions(command, opts)

	return command

}

func LoadExtensionCfg(file, extension string, opts ...interface{}) *LoadCfgOptions {

	extensionFunc := func(options LoadCfgOptions) {
		options.Extension = extension
	}

	opts = append(opts, extensionFunc)
	command := LoadCfg(file, opts...)

	return command

}

func DumpCfg(file string, opts ...interface{}) *DumpCfgOptions {

	command := &DumpCfgOptions{
		File:     file,
		Designer: newDefaultDesigner(),
	}

	processOptions(command, opts)

	return command

}

func DumpExtensionCfg(file, extension string, opts ...interface{}) *DumpCfgOptions {

	extensionFunc := func(options DumpCfgOptions) {
		options.Extension = extension
	}

	opts = append(opts, extensionFunc)

	command := DumpCfg(file, opts...)

	return command

}
