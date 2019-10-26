package v8run

import "fmt"

type DesignerOptions func(designer Designer)

type UserOption func(Optioned)

type Optioned interface {
	setOption(key string, value interface{})
}

type CanUpdateDBCfg interface {
	SetUpdateDBCfg(UpdateDBCfgOptions)
}

type UserOptions map[string]interface{}

func (uo UserOptions) setOption(key string, value interface{}) {
	uo[key] = value
}

func (uo UserOptions) Append(uo2 UserOptions) {

	for k, v := range uo2 {
		uo[k] = k
	}

}

type Designer struct {
	UserOptions

	disableStartupDialogs  bool
	disableStartupMessages bool
	visible                bool
}

func (d Designer) Command() string {
	return COMMANE_DESIGNER
}

func (d Designer) Check() bool {

	return true
}

func (d Designer) Values() (values UserOptions) {

	values = make(map[string]interface{})

	values["/DisableStartupDialogs"] = d.disableStartupDialogs
	values["/DisableStartupDialogs"] = d.disableStartupDialogs
	values["/Visible"] = d.visible

	for k, v := range d.UserOptions {
		values[k] = v
	}

	return values

}

func processArgs(options UserOptions) (args []string) {

	for k, v := range options {

		switch v.(type) {

		case bool:

			val, _ := v.(bool)

			if val {
				args = append(args, k)
			}

		case string:

			val, _ := v.(string)

			if len(val) > 0 {
				args = append(args, fmt.Sprintf("%s %s", k, val))
			}

		case Optioned:

		default:

			continue

		}

	}

	return
}

func (d Designer) option(in interface{}) {

	opt, ok := in.(DesignerOptions)

	if !ok {
		return
	}

	d.Option(opt)
}

func NewDesigner(opts ...DesignerOptions) Designer {

	d := Designer{
		UserOptions: make(map[string]interface{}),
	}

	for _, opt := range opts {
		d.option(opt)
	}

	return d
}

func WithUnlockCode(uc string) UserOption {
	return func(o Optioned) {
		o.setOption("/UC", uc)
	}
}

func newDefaultDesigner() Designer {

	d := Designer{
		disableStartupDialogs:  true,
		disableStartupMessages: true,
		visible:                false,
	}

	return d
}

func (d Designer) Option(option DesignerOptions) {
	option(d)
}

func WithUpdateDBCfg(update UpdateDBCfgOptions) func(CanUpdateDBCfg) {
	return func(options CanUpdateDBCfg) {
		options.SetUpdateDBCfg(update)
	}
}

// loadcf
type LoadCfgOption func(LoadCfgOptions)

type LoadCfgOptions struct {
	Designer
	File        string
	Extension   string
	UpdateDBCfg UpdateDBCfgOptions
}

func (d LoadCfgOptions) Values() (values UserOptions) {

	values = d.Designer.Values()
	values["/LoadCfg"] = d.File
	values["-Extension"] = d.Extension

	values.Append(d.UpdateDBCfg.Values())

	return

}

func (d LoadCfgOptions) option(in interface{}) {

	switch in.(type) {

	case LoadCfgOption:
		opt, _ := in.(LoadCfgOption)
		d.Option(opt)
	case DesignerOptions:
		d.Designer.option(in)
	}
}

func (d LoadCfgOptions) Option(fn LoadCfgOption) {

	fn(d)

}

func (d LoadCfgOptions) SetUpdateDBCfg(updateDBCfg UpdateDBCfgOptions) {
	d.UpdateDBCfg = updateDBCfg
}

type DumpCfgOption func(DumpCfgOptions)

type DumpCfgOptions struct {
	Designer
	File      string
	Extension string
}

func (d DumpCfgOptions) Args() (args []string) {

	args = d.Designer.Args()

	args = append(args, fmt.Sprintf("/DumpCfg %s ", d.File))

	if len(d.Extension) > 0 {
		args = append(args, fmt.Sprintf("-Extension %s", d.Extension))

	}

	return

}

func (d DumpCfgOptions) Option(fn DumpCfgOption) {

	fn(d)

}

func (d DumpCfgOptions) option(in interface{}) {

	switch in.(type) {

	case DumpCfgOption:
		opt, _ := in.(DumpCfgOption)
		d.Option(opt)
	case DesignerOptions:
		d.Designer.option(in)
	}
}

type UpdateCfgOption func(UpdateCfgOptions)

type UpdateCfgOptions struct {
	Designer
	//<имя cf | cfu-файла>
	File string

	// <имя файла настроек> — содержит имя файла настроек объединения.
	Settings string

	// если в настройках есть объекты, не включенные в список обновляемых и отсутствующие в основной конфигурации,
	// на которые есть ссылки из объектов, включенных в список, то такие объекты также помечаются для обновления,
	// и выполняется попытка продолжить обновление.
	IncludeObjectsByUnresolvedRefs bool

	//— очищение ссылок на объекты, не включенные в список обновляемых.
	ClearUnresolvedRefs bool

	//— Если параметр используется, обновление будет выполнено несмотря на наличие предупреждений:
	//о применении настроек,
	//о дважды измененных свойствах, для которых не был выбран режим объединения,
	//об удаляемых объектах, на которые найдены ссылки в объектах, не участвующие в объединении.
	//Если параметр не используется, то в описанных случаях объединение будет прервано.
	Force bool

	//— вывести список всех дважды измененных свойств.
	DumpListOfTwiceChangedProperties bool

	UpdateDBCfg UpdateDBCfgOptions
}

func (d UpdateCfgOptions) Args() (args []string) {

	args = d.Designer.Args()

	args = append(args, fmt.Sprintf("/UpdateCfg %s ", d.File))

	if d.Force {
		args = append(args, "-Force")
	}

	if len(d.Settings) > 0 {
		args = append(args, fmt.Sprintf("-Settings %s", d.Settings))

	}

	return

}

func (d UpdateCfgOptions) Option(fn UpdateCfgOption) {

	fn(d)

}

func (d UpdateCfgOptions) SetUpdateDBCfg(updateDBCfg *UpdateDBCfgOptions) {
	d.UpdateDBCfg = updateDBCfg
}

func (d UpdateCfgOptions) option(in interface{}) {

	switch in.(type) {

	case UpdateCfgOption:
		opt, _ := in.(UpdateCfgOption)
		d.Option(opt)
	case DesignerOptions:
		d.Designer.option(in)
	}
}

type UpdateDBCfgOption func(UpdateDBCfgOptions)

///UpdateDBCfg [–Dynamic<Режим>] [-BackgroundStart] [-BackgroundCancel]
//[-BackgroundFinish [-Visible]] [-BackgroundSuspend] [-BackgroundResume]
//[-WarningsAsErrors] [-Server [-v1|-v2]][-Extension <имя расширения>]
type UpdateDBCfgOptions struct {
	Designer

	//-Dynamic<Режим> — признак использования динамического обновления. Режим может принимать следующие значения
	//-Dynamic+ — Значение параметра по умолчанию.
	// Сначала выполняется попытка динамического обновления, если она завершена неудачно, будет запущено фоновое обновление.
	//-Dynamic–  — Динамическое обновление запрещено.
	Dynamic bool

	//-BackgroundStart [-Dynamic<Режим>] — будет запущено фоновое обновление конфигурации,
	// текущий сеанс будет завершен. Если обновление уже выполняется, будет выдана ошибка.
	//-Dynamic+ — Значение параметра по умолчанию.
	// Сначала выполняется попытка динамического обновления, если она завершена неудачно,
	// будет запущено фоновое обновление.
	//-Dynamic–  — Динамическое обновление запрещено.
	BackgroundStart bool

	//-BackgroundCancel — отменяет запущенное фоновое обновление конфигурации базы данных.
	// Если фоновое обновление не запущено, будет выдана ошибка.
	BackgroundCancel bool

	//-BackgroundFinish — запущенное фоновое обновление конфигурации базы данных будет завершено:
	// при этом будет наложена монопольная блокировка и проведена финальная фаза обновления.
	// Если фоновое обновление конфигурации не запущено или переход к завершающей фазе обновления не возможен, будет выдана ошибка.
	// Возможно использование следующих параметров:
	//-Visible — На экран будет выведен диалоговое окно с кнопками Отмена, Повторить, Завершить сеансы и повторить.
	// В случае невозможности завершения фонового обновления, если данная опция не указана, выполнение обновления будет завершено с ошибкой..
	BackgroundFinish bool

	//-BackgroundResume — продолжает фоновое обновление конфигурации базы данных, приостановленное ранее.
	BackgroundResume bool

	//-BackgroundSuspend — приостанавливает фоновое обновление конфигурации на паузу.
	// Если фоновое обновление не запущено, будет выдана ошибка.
	BackgroundSuspend bool

	//-WarningsAsErrors —  все предупредительные сообщения будут трактоваться как ошибки.
	WarningsAsErrors bool

	//-Server — обновление будет выполняться на сервере (имеет смысл только на сервере).
	// Если параметр используется вместе с фоновым обновлением, то:
	//
	//Фаза актуализации всегда выполняется на сервере.
	//Фаза обработки и фаза принятия изменения могут выполняться как на клиенте, так и на сервере.
	//Допускается запуск фонового обновления на стороне клиента, а завершение - на стороне сервера, и наоборот.
	//Не используется 2-я версия механизма реструктуризации (игнорируется параметр -v2, если таковой указан).
	//Если не указана версия механизма реструктуризации (-v1 или -v2),
	// то будет использоваться механизм реструктуризации той версии, которая указана в файле conf.cfg.
	// В противном случае будет использована указанная версия механизма.
	// Если указана 2-я версия механизма реструктуризации, но использование этой версии конфликтует с другими параметрами
	// – будет использована 1-я версия.
	Server bool

	//-Extension <Имя расширения> — будет выполнено обновление расширения с указанным именем.
	// Если расширение успешно обработано возвращает код возврата 0,
	// в противном случае (если расширение с указанным именем не существует или в процессе работы произошли ошибки) — 1.
	Extension string
}

func (d UpdateDBCfgOptions) Args() (args []string) {

	args = d.Designer.Args()

	args = append(args, "/UpdateDBCfg")

	if d.Server {
		args = append(args, "-Server")
	}

	if d.WarningsAsErrors {
		args = append(args, "-WarningsAsErrors")
	}

	if len(d.Extension) > 0 {
		args = append(args, fmt.Sprintf("-Extension %s", d.Extension))

	}

	return

}

func (d UpdateDBCfgOptions) UpdateArgs() (args []string) {

	args = append(args, "/UpdateDBCfg")

	if d.Server {
		args = append(args, "-Server")
	}

	if d.WarningsAsErrors {
		args = append(args, "-WarningsAsErrors")
	}

	if len(d.Extension) > 0 {
		args = append(args, fmt.Sprintf("-Extension %s", d.Extension))

	}

	return

}

func (d UpdateDBCfgOptions) Option(fn UpdateDBCfgOption) {

	fn(d)

}

func (d UpdateDBCfgOptions) option(in interface{}) {

	switch in.(type) {

	case UpdateDBCfgOption:
		opt, _ := in.(UpdateDBCfgOption)
		d.Option(opt)
	case DesignerOptions:
		d.Designer.option(in)
	}
}

func processOptions(what Optioned, opts []interface{}) {

	for _, opt := range opts {
		what.option(opt)
	}

}
