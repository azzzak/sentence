package sentence

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func initPlural(lang, delim string) func(int, reflect.Value) (string, error) {
	var pluralFn pluralFunc
	formsNum := formsNum[lang]

	switch lang {
	case LangEnglish:
		pluralFn = pluralEnglish
	default:
		pluralFn = pluralRussian
	}

	return func(num int, item reflect.Value) (string, error) {
		item = indirectInterface(item)
		if !item.IsValid() {
			return "", fmt.Errorf("untyped nil")
		}

		var forms []string
		switch item.Kind() {
		case reflect.String:
			forms = strings.Split(item.String(), delim)
		case reflect.Slice:
			var ok bool
			if forms, ok = item.Interface().([]string); !ok {
				return "", fmt.Errorf("can't slice item of type %s", item.Type())
			}
		default:
			return "", fmt.Errorf("can't slice item of type %s", item.Type())
		}

		if len(forms) != formsNum {
			return "", fmt.Errorf("lang '%s' want %d forms of word, you pass %d", lang, formsNum, len(forms))
		}

		return pluralFn(num, forms), nil
	}
}

func initPluraln(fn func(int, reflect.Value) (string, error)) func(num int, item reflect.Value) (string, error) {
	return func(num int, item reflect.Value) (string, error) {
		str, err := fn(num, item)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s %s", strconv.Itoa(num), str), nil
	}
}

func initAny(delim string) func(reflect.Value) (string, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return func(item reflect.Value) (string, error) {
		item = indirectInterface(item)
		if !item.IsValid() {
			return "", fmt.Errorf("untyped nil")
		}

		var options []string
		switch item.Kind() {
		case reflect.String:
			options = strings.Split(item.String(), delim)
		case reflect.Slice:
			var ok bool
			if options, ok = item.Interface().([]string); !ok {
				return "", fmt.Errorf("can't slice item of type %s", item.Type())
			}
		default:
			return "", fmt.Errorf("can't slice item of type %s", item.Type())
		}

		return options[r.Intn(len(options))], nil
	}
}

func initAnyf() func(items ...reflect.Value) (string, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return func(items ...reflect.Value) (string, error) {
		res := make([]string, len(items))
		for i, item := range items {
			res[i] = fmt.Sprintf("%v", item)
		}

		return res[r.Intn(len(res))], nil
	}
}

func indirectInterface(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Interface {
		return v
	}
	if v.IsNil() {
		return reflect.Value{}
	}
	return v.Elem()
}
