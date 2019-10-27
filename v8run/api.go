package v8run

func LoadCfg(file string, opts ...UserOption) *LoadCfgOptions {

	command := &LoadCfgOptions{
		File:     file,
		Designer: newDefaultDesigner(),
	}

	processOptions(command, opts)

	return command

}

func LoadExtensionCfg(file, extension string, opts ...UserOption) *LoadCfgOptions {

	command := LoadCfg(file, opts...)
	command.Extension = extension

	return command

}

func DumpCfg(file string, opts ...UserOption) *DumpCfgOptions {

	command := &DumpCfgOptions{
		File:     file,
		Designer: newDefaultDesigner(),
	}

	processOptions(command, opts)

	return command

}

func DumpExtensionCfg(file, extension string, opts ...UserOption) *DumpCfgOptions {

	command := DumpCfg(file, opts...)
	command.Extension = extension
	return command

}

func UpdateDBCfg(server bool, Dynamic bool, opts ...UserOption) *UpdateDBCfgOptions {

	command := &UpdateDBCfgOptions{
		Designer: newDefaultDesigner(),
		Server:   server,
		Dynamic:  Dynamic,
	}

	processOptions(command, opts)

	return command

}

func UpdateDBExtensionCfg(extension string, server bool, Dynamic bool, opts ...UserOption) *UpdateDBCfgOptions {

	command := UpdateDBCfg(server, Dynamic, opts...)
	command.Extension = extension

	return command

}

func DumpIB(file string, opts ...UserOption) *DumpIBOptions {

	command := &DumpIBOptions{
		Designer: newDefaultDesigner(),
		File:     file,
	}

	processOptions(command, opts)
	return command
}

func RestoreIB(file string, opts ...UserOption) *RestoreIBOptions {

	command := &RestoreIBOptions{
		Designer: newDefaultDesigner(),
		File:     file,
	}

	processOptions(command, opts)

	return command
}

func Execute(file string, opts ...UserOption) *ExecuteOptions {

	command := &ExecuteOptions{
		Enterprise: newDefaultEnterprise(),
		File:       file,
	}

	processOptions(command, opts)

	return command
}
