package common

import (
	"fmt"
	"os"
	"reflect"
	"slices"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
)

type CommandVariant int

const (
	Dos CommandVariant = iota
	Posix
	PowerShell
)

type Arguments struct {
	Name           *string `dos:"n" posix:"name" posix_short:"n" poweshell:"Name" poweshell_short:"n"`
	Duration       *uint16 `dos:"d" posix:"duration" posix_short:"d" poweshell:"Duration" poweshell_short:"d"`
	Pineapple      *bool   `dos:"p" posix:"pineapple" posix_negative:"no-pineapple" posix_short:"p" posix_short_negative:"P" poweshell:"Pineapple" poweshell_short:"p"`
	CheckSuperuser *bool   `dos:"s" posix:"check-superuser" posix_negative:"no-check-superuser" posix_short:"s" posix_short_negative:"S" poweshell:"CheckSuperuser" poweshell_short:"s"`
	ExitCode       *uint16 `dos:"e" posix:"exit-code" posix_short:"e" poweshell:"ExitCode" poweshell_short:"e"`
	Help           *bool   `dos:"?" posix:"help" posix_short:"h" poweshell:"Help" poweshell_short:"h"`
	Version        *bool   `dos:"v" posix:"version" posix_short:"v" poweshell:"Version" poweshell_short:"v"`
}

type MappingQuery struct {
	Long     bool
	Field    reflect.StructField
	Negative bool
}

func PrintLicense() {
	fmt.Println(`
Copyright (c) 2026 Reinhart Previano Koentjoro <reinhart1010.id> & Citra Manggala Dirgantara <shiftine.sh>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.`)
}

func GetValidArgs(variant CommandVariant) (longArgs []string, shortArgs []string, mapping map[string]MappingQuery) {
	longArgs = []string{}
	shortArgs = []string{}
	mapping = map[string]MappingQuery{}

	args := reflect.ValueOf(Arguments{})
	typ := args.Type()
	for i := 0; i < args.NumField(); i++ {
		field := typ.Field(i)

		longCandidates := []string{}

		switch variant {
		case Posix:
			longCandidates = []string{
				field.Tag.Get("posix"),
				field.Tag.Get("posix_negative"),
			}
		case PowerShell:
			longCandidates = []string{
				field.Tag.Get("powershell"),
			}
		}

		for i, insert := range longCandidates {
			if len(insert) > 0 {
				longArgs = append(longArgs, insert)
				mapping[insert] = MappingQuery{
					Long:     true,
					Field:    field,
					Negative: i > 0,
				}
			}
		}

		shortCandidates := []string{}

		switch variant {
		case Dos:
			shortCandidates = []string{
				field.Tag.Get("dos"),
			}
		case Posix:
			shortCandidates = []string{
				field.Tag.Get("posix_short"),
				field.Tag.Get("posix_short_negative"),
			}
		case PowerShell:
			shortCandidates = []string{
				field.Tag.Get("powershell_short"),
			}
		}

		for _, insert := range shortCandidates {
			if len(insert) > 0 {
				shortArgs = append(shortArgs, insert)
				mapping[insert] = MappingQuery{
					Long:     false,
					Field:    field,
					Negative: i > 0,
				}
			}
		}
	}
	return longArgs, shortArgs, mapping
}

func ParseArguments(parsed map[string]string, mapping map[string]MappingQuery) Arguments {
	args := Arguments{}
	argsRf := reflect.ValueOf(args)
	for key, rawVal := range parsed {
		var val interface{}
		if field, ok := mapping[key]; ok {
			if field.Field.Type == reflect.TypeFor[bool]() {
				if len(rawVal) > 0 {
					valTemp, err := strconv.ParseBool(rawVal)
					if err != nil {
						log.Fatalf("Parameter `%s` has a non-Boolean value: %s", key, rawVal)
					}
					val = valTemp
				} else {
					// Use hints from the created mapping
					val = !field.Negative
				}
			} else if slices.Contains([]reflect.Type{
				reflect.TypeFor[int](),
				reflect.TypeFor[int8](),
				reflect.TypeFor[int16](),
				reflect.TypeFor[int32](),
				reflect.TypeFor[int64](),
			}, field.Field.Type) {
				valTemp, err := strconv.ParseInt(rawVal, 10, 0)
				if err != nil {
					log.Fatalf("Parameter `%s` has a non-integer value: %s", key, rawVal)
				}

				switch field.Field.Type {
				case reflect.TypeFor[int8]():
					val = int8(valTemp)
				case reflect.TypeFor[int16]():
					val = int16(valTemp)
				case reflect.TypeFor[int32]():
					val = int32(valTemp)
				case reflect.TypeFor[int64]():
					val = int64(valTemp)
				default:
					val = valTemp
				}
			} else if slices.Contains([]reflect.Type{
				reflect.TypeFor[uint](),
				reflect.TypeFor[uint8](),
				reflect.TypeFor[uint16](),
				reflect.TypeFor[uint32](),
				reflect.TypeFor[uint64](),
			}, field.Field.Type) {
				valTemp, err := strconv.ParseUint(rawVal, 10, 0)
				if err != nil {
					log.Fatalf("Parameter `%s` has a non-unsigned integer value: %s", key, rawVal)
				}

				switch field.Field.Type {
				case reflect.TypeFor[uint8]():
					val = uint8(valTemp)
				case reflect.TypeFor[uint16]():
					val = uint16(valTemp)
				case reflect.TypeFor[uint32]():
					val = uint32(valTemp)
				case reflect.TypeFor[uint64]():
					val = uint64(valTemp)
				default:
					val = valTemp
				}
			} else if field.Field.Type == reflect.TypeFor[float32]() {
				valTemp, err := strconv.ParseFloat(rawVal, 32)
				if err != nil {
					log.Fatalf("Parameter `%s` has a non 32-bit float value: %s", key, rawVal)
				}
				val = float32(valTemp)
			} else if field.Field.Type == reflect.TypeFor[float64]() {
				valTemp, err := strconv.ParseFloat(rawVal, 64)
				if err != nil {
					log.Fatalf("Parameter `%s` has a non 64-bit float value: %s", key, rawVal)
				}
				val = valTemp
			} else if field.Field.Type == reflect.TypeFor[complex64]() {
				valTemp, err := strconv.ParseComplex(rawVal, 64)
				if err != nil {
					log.Fatalf("Parameter `%s` has a non 64-bit complex value: %s", key, rawVal)
				}

				val = complex64(valTemp)
			} else if field.Field.Type == reflect.TypeFor[complex128]() {
				valTemp, err := strconv.ParseComplex(rawVal, 128)
				if err != nil {
					log.Fatalf("Parameter `%s` has a non 128-bit complex value: %s", key, rawVal)
				}
				val = valTemp
			} else {
				val = rawVal
			}
		} else {
			log.Fatalf("INTERNAL ERROR: Mapping data for field %s is not found", key)
		}

		fv := argsRf.FieldByName(key)
		if !fv.IsValid() {
			log.Fatalf("INTERNAL ERROR: Argument struct field %s is not valid", key)
		} else if !fv.CanSet() {
			log.Fatalf("INTERNAL ERROR: Argument struct field %s cannot be set", key)
		}
		valOf := reflect.ValueOf(val)
		if fv.Kind() != valOf.Kind() {
			log.Fatalf("INTERNAL ERROR: Argument struct field %s has type mismatch", key)
		}
		fv.Set(valOf)
	}
	return args
}

func MakePizza(contents string, args Arguments) uint16 {
	startTime := time.Now()
	log.Info("Preparing the pizza", "name", args.Name, "contents", contents)

	if args.CheckSuperuser != nil && *args.CheckSuperuser {
		if os.Geteuid() != 0 {
			log.Fatal("You are not (#- ). This pizza shall be made by (#- ).", "name", args.Name, "contents", contents)
		} else {
			log.Info("Hi, (#- )!")
		}
	}

	if args.Pineapple == nil || *args.Pineapple {
		log.Info("Adding some pineapples to the pizza", args.Name, "contents", contents)
	}

	if args.Duration != nil {
		remainingTime := time.Duration(float64(*args.Duration)-time.Since(startTime).Seconds()) * time.Second
		if remainingTime > 0 {
			log.Info(fmt.Sprintf("Baking for %d seconds...", remainingTime), "name", args.Name, "contents", contents)
			time.Sleep(remainingTime)
		}
	}

	exitCode := uint16(0)
	if args.ExitCode != nil {
		exitCode = *args.ExitCode
	}

	log.Info(fmt.Sprintf("Finished baking pizza with exit code %d", exitCode), "name", args.Name, "contents", contents)
	return exitCode
}
