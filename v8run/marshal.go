package v8run

import (
	"errors"
	"fmt"
	"github.com/khorevaa/go-AutoUpdate1C/v8run/tags"
	"reflect"
	"strconv"
	"time"
)

const TAG_NAMESPACE = "v8"
const COMMAND_FIELD_NAME = "command"

type Marshaler interface {
	MarshalV8() (string, error)
}

func v8Marshal(object interface{}) []string {

	var fieldsList []string

	rType := reflect.TypeOf(object).Elem()
	fieldsCount := rType.NumField()

	v := reflect.ValueOf(object)

	for i := 0; i < fieldsCount; i++ {
		field := rType.Field(i)

		fieldInfo := tags.GetFieldTagInfo(field)

		if fieldInfo == nil {
			continue
		}

		if field.Name == COMMAND_FIELD_NAME {
			fieldsList = append(fieldsList, fieldInfo.Name)
			continue
		}

		iface := reflect.Indirect(v).FieldByName(field.Name).Interface()

		// if the field is a pointer to a struct, follow the pointer then create fieldinfo for each field
		if field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct {
			// unless it implements marshalText or marshalCSV. Structs that implement this
			// should result in one iface and not have their fields exposed
			if fieldInfo.Inherit {
				fieldsList = append(fieldsList, v8Marshal(iface)...)
				continue
			}
		}
		// if the field is a struct, create a fieldInfo for each of its fields
		if field.Type.Kind() == reflect.Struct {
			// unless it implements marshalText or marshalCSV. Structs that implement this
			// should result in one iface and not have their fields exposed
			if fieldInfo.Inherit {
				fieldsList = append(fieldsList, v8Marshal(iface)...)
				continue
			}
		}

		switch m := iface.(type) {

		case Marshaler:
			v, err := m.MarshalV8()
			if err != nil {
				_ = fmt.Errorf("error marshal type: %s", err)
				continue
			}

			err = checkField(v, fieldInfo)
			if err != nil {
				continue // TODO Error
			}

			fieldArg := fieldInfo.Name + " " + v

			if fieldInfo.Argument {
				fieldArg = v
			}

			fieldsList = append(fieldsList, fieldArg)

		case time.Time, *time.Time:
			// Although time.Time implements TextMarshaler,
			// we don't want to treat it as a string for YAML
			// purposes because YAML has special support for
			// timestamps.

		case string:

			v := iface.(string)

			err := checkField(v, fieldInfo)
			if err != nil {
				continue // TODO Error
			}

			fieldArg := fieldInfo.Name + " " + v
			fieldsList = append(fieldsList, fieldArg)

		case bool:

			fieldArg := fieldInfo.Name

			if !fieldInfo.Inherit {

			}

			fieldsList = append(fieldsList, fieldArg)

		case int, int32, int64:

			fieldArg := fieldInfo.Name + " " + strconv.FormatInt(iface.(int64), 10)
			fieldsList = append(fieldsList, fieldArg)

		case nil:
			continue
		}

	}
	return fieldsList

}

func checkField(value string, tagInfo *tags.FieldTagInfo) error {

	if tagInfo.Argument && len(value) == 0 {
		return errors.New("need value")
	}

	if tagInfo.Optional && len(value) == 0 {
		return nil
	}

	if len(value) == 0 {
		return errors.New("need value")
	}

	return nil
}
