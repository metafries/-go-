package _mfConfig

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ConfigFields map[string]reflect.Value

func (cf ConfigFields) Add(n, v, t string) error {
	switch t {
	case "STRING":
		cf[n] = reflect.ValueOf(v)
	case "INTEGER":
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		cf[n] = reflect.ValueOf(i)
	case "FLOAT":
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		cf[n] = reflect.ValueOf(f)
	case "BOOL":
		b, err := strconv.ParseBool(v)
		if err != nil {
			return err
		}
		cf[n] = reflect.ValueOf(b)
	}
	return nil
}

func MarshalCustomConfig(v reflect.Value, fileName string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occured:", r)
		}
	}()
	if !v.CanSet() {
		return errors.New("[ERROR] Value passed not settable")
	}
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	fields := make(ConfigFields) // make(map[string]reflect.Value)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Processing line", line)
		if strings.Count(line, "|") != 1 || strings.Count(line, ";") != 1 {
			continue
		}
		args := strings.Split(line, "|")
		valuetype := strings.Split(args[1], ";")
		n, v, t := strings.TrimSpace(args[0]), strings.TrimSpace(valuetype[0]), strings.ToUpper(strings.TrimSpace(valuetype[1]))
		fields.Add(n, v, t)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	vt := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldType := vt.Field(i)
		fieldValue := v.Field(i)
		name := fieldType.Tag.Get("name")
		if name == "" {
			name = fieldType.Name
		}
		if v, ok := fields[name]; ok {
			fieldValue.Set(v)
		}
	}
	return nil
}
