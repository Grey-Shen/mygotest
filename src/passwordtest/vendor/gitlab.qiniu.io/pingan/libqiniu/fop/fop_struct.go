package fop

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

const fopTagName = "fop"

type pair struct {
	key, value string
}

type fopCommand struct {
	commandName string
	mainOption  string
	options     []pair
}

func (cmd fopCommand) String() string {
	result := cmd.commandName
	if cmd.mainOption != "" {
		result += "/" + cmd.mainOption
	}
	for _, pair := range cmd.options {
		result += "/" + pair.key
		if pair.value != "" {
			result += "/" + pair.value
		}
	}
	return result
}

type NoValue struct{}
type EmptyMainOption struct{}
type EmptyTypeName struct{}

func (_ *EmptyMainOption) Error() string {
	return "The main option must not be empty"
}

func (_ *EmptyTypeName) Error() string {
	return "The value must have a name"
}

type encodable interface {
	Encode() string
}

func convertToFopCommand(value interface{}) (cmd fopCommand, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch errMsg := r.(type) {
			case string:
				err = errors.New(errMsg)
			case error:
				err = errMsg
			}
		}
	}()

	val := reflect.ValueOf(value)
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		fieldType := val.Type().Field(i)
		if fopOptionName, ok := fieldType.Tag.Lookup(fopTagName); ok {
			if fopOptionName == "" {
				fopOptionName = fieldType.Name
			}
			if valStr, present := convertToString(fieldValue); present {
				cmd.options = append(cmd.options, pair{key: fopOptionName, value: valStr})
			}
		}
	}
	if len(cmd.options) > 0 {
		cmd.commandName = cmd.options[0].key
		cmd.mainOption = cmd.options[0].value
		cmd.options = cmd.options[1:]
	}
	return cmd, nil
}

func convertToString(value reflect.Value) (string, bool) {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := value.Int()
		return strconv.FormatInt(val, 10), val != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := value.Uint()
		return strconv.FormatUint(val, 10), val != 0
	case reflect.Float32:
		val := value.Float()
		return strconv.FormatFloat(val, 'f', -1, 32), val != 0.0
	case reflect.Float64:
		val := value.Float()
		return strconv.FormatFloat(val, 'f', -1, 64), val != 0.0
	case reflect.Bool:
		if value.Bool() {
			return "1", true
		} else {
			return "0", false
		}
	case reflect.String:
		val := value.String()
		return val, val != ""
	default:
		if value.Type().ConvertibleTo(reflect.TypeOf(NoValue{})) {
			return "", true
		} else if value.IsValid() && value.CanInterface() {
			if interf := value.Interface(); interf == nil {
				return "", false
			} else if value.Kind() == reflect.Ptr && value.IsNil() {
				return "", false
			} else {
				switch val := interf.(type) {
				case encodable:
					str := val.Encode()
					return str, str != ""
				case error:
					str := val.Error()
					return str, str != ""
				case fmt.Stringer:
					str := val.String()
					return str, str != ""
				}
			}
			if value.Kind() == reflect.Ptr {
				return convertToString(value.Elem())
			}
		}
		str := value.String()
		return str, str != ""
	}
}
