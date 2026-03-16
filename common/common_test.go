package common

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestArgumentsForClash(t *testing.T) {
	dosFlags := []string{}
	posixFlags := []string{}
	psFlags := []string{}

	args := reflect.ValueOf(Arguments{})
	typ := args.Type()
	for i := 0; i < args.NumField(); i++ {
		field := typ.Field(i)

		// Get and assert the DOS flag
		dosOrig := field.Tag.Get("dos")
		dos := strings.ToLower(dosOrig)
		if len(dos) == 0 {
			t.Errorf("The field %s does not have the `dos` attribute.", field.Name)
		} else if strings.HasPrefix(dos, "-") || strings.HasPrefix(dos, "/") {
			t.Errorf("The field %s's DOS flag must not begin with the argument delimiter themselves (`-`, `--`, or `/`).", field.Name)
		} else if strings.Contains(dos, ":") || strings.Contains(dos, "=") {
			t.Errorf("The field %s's DOS flag must not contain the value delimiter (`:` or `=`).", field.Name)
		} else if slices.Contains(dosFlags, dos) {
			t.Errorf("The field %s uses a DOS flag `/%s` that has been used in another argument.", field.Name, dosOrig)
		} else {
			dosFlags = append(dosFlags, dos)
		}

		// Get and assert the POSIX flag(s)
		posix := field.Tag.Get("posix")
		if len(posix) == 0 {
			t.Errorf("The field %s does not have the `posix` attribute.", field.Name)
		} else if posix != strings.ToLower(posix) {
			t.Errorf("The field %s's POSIX long flag `--%s` must be written in lowercase.", field.Name, posix)
		} else if strings.HasPrefix(posix, "-") || strings.HasPrefix(posix, "/") {
			t.Errorf("The field %s's POSIX long flag must not begin with the argument delimiter themselves (`-`, `--`, or `/`).", field.Name)
		} else if strings.Contains(posix, ":") || strings.Contains(posix, "=") {
			t.Errorf("The field %s's POSIX long flag must not contain the value delimiter (`:` or `=`).", field.Name)
		} else if slices.Contains(posixFlags, posix) {
			t.Errorf("The field %s uses a POSIX long flag `--%s` that has been used in another argument.", field.Name, posix)
		} else {
			posixFlags = append(posixFlags, posix)
		}

		posixNegative := field.Tag.Get("posix_negative")
		if (field.Type != reflect.TypeFor[bool]() && field.Type != reflect.TypeFor[*bool]()) && len(posixNegative) > 0 {
			t.Errorf("The field %s does not accept boolean values, therefore `posix_negative` must not be set.", field.Name)
		} else if (field.Type == reflect.TypeFor[bool]() || field.Type == reflect.TypeFor[*bool]()) && len(posixNegative) == 0 {
			t.Errorf("The field %s accepts boolean values, but `posix_negative` was not set.", field.Name)
		}

		if len(posixNegative) > 0 {
			if strings.HasPrefix(posixNegative, "-") || strings.HasPrefix(posixNegative, "/") {
				t.Errorf("The field %s's POSIX long negative flag must not begin with the argument delimiter themselves (`-`, `--`, or `/`).", field.Name)
			} else if strings.Contains(posixNegative, ":") || strings.Contains(posixNegative, "=") {
				t.Errorf("The field %s's POSIX long negative flag must not contain the value delimiter (`:` or `=`).", field.Name)
			} else if posixNegative != strings.ToLower(posixNegative) {
				t.Errorf("The field %s's POSIX long negative flag `--%s` must be written in lowercase.", field.Name, posixNegative)
			} else if slices.Contains(posixFlags, posixNegative) {
				t.Errorf("The field %s uses a POSIX long negative flag `--%s` that has been used in another argument.", field.Name, posixNegative)
			} else {
				posixFlags = append(posixFlags, posixNegative)
			}
		}

		posixShort := field.Tag.Get("posix_short")
		if len(posixShort) > 0 {
			if strings.HasPrefix(posixShort, "-") || strings.HasPrefix(posixShort, "/") {
				t.Errorf("The field %s's POSIX short flag must not begin with the argument delimiter themselves (`-`, `--`, or `/`).", field.Name)
			} else if strings.Contains(posixShort, ":") || strings.Contains(posixShort, "=") {
				t.Errorf("The field %s's POSIX short flag must not contain the value delimiter (`:` or `=`).", field.Name)
			} else if slices.Contains(posixFlags, posixShort) {
				t.Errorf("The field %s uses a POSIX short flag `-%s` that has been used in another argument.", field.Name, posixShort)
			} else {
				posixFlags = append(posixFlags, posixShort)
			}
		}

		posixShortNegative := field.Tag.Get("posix_short_negative")
		if (field.Type != reflect.TypeFor[bool]() && field.Type != reflect.TypeFor[*bool]()) && len(posixShortNegative) > 0 {
			t.Errorf("The field %s does not accept boolean values, therefore `posix_short_negative` must not be set.", field.Name)
		} else if (field.Type == reflect.TypeFor[bool]() || field.Type == reflect.TypeFor[*bool]()) && len(posixShort) > 0 {
			if len(posixShortNegative) == 0 {
				t.Errorf("The field %s accepts boolean values, but `posix_short_negative` was not set.", field.Name)
			} else if strings.HasPrefix(posixShortNegative, "-") || strings.HasPrefix(posixShortNegative, "/") {
				t.Errorf("The field %s's POSIX short negative flag must not begin with the argument delimiter themselves (`-`, `--`, or `/`).", field.Name)
			} else if strings.Contains(posixShortNegative, ":") || strings.Contains(posixShortNegative, "=") {
				t.Errorf("The field %s's POSIX short negative flag must not contain the value delimiter (`:` or `=`).", field.Name)
			} else if slices.Contains(posixFlags, posixShortNegative) {
				t.Errorf("The field %s uses a POSIX short negative flag `-%s` that has been used in another argument.", field.Name, posixShortNegative)
			} else {
				posixFlags = append(posixFlags, posixShortNegative)
			}
		}

		// Get and assert the PowerShell flag(s)
		psOrig := field.Tag.Get("ps")
		ps := strings.ToLower(field.Tag.Get("ps"))
		if len(ps) == 0 {
			t.Errorf("The field %s does not have the `ps` attribute.", field.Name)
		} else if strings.HasPrefix(ps, "-") || strings.HasPrefix(ps, "/") {
			t.Errorf("The field %s's PowerShell long flag must not begin with the argument delimiter themselves (`-`, `--`, or `/`).", field.Name)
		} else if strings.Contains(ps, ":") || strings.Contains(ps, "=") {
			t.Errorf("The field %s's PowerShell long flag must not contain the value delimiter (`:` or `=`).", field.Name)
		} else if slices.Contains(psFlags, ps) {
			t.Errorf("The field %s uses a PowerShell long flag `-%s` that has been used in another argument.", field.Name, psOrig)
		} else {
			psFlags = append(psFlags, ps)
		}

		psShortOrig := field.Tag.Get("ps_short")
		psShort := strings.ToLower(psShortOrig)
		if len(psShort) > 0 {
			if strings.HasPrefix(psShort, "-") || strings.HasPrefix(psShort, "/") {
				t.Errorf("The field %s's PowerShell short flag must not begin with the argument delimiter themselves (`-`, `--`, or `/`).", field.Name)
			} else if strings.Contains(psShort, ":") || strings.Contains(psShort, "=") {
				t.Errorf("The field %s's PowerShell short flag must not contain the value delimiter (`:` or `=`).", field.Name)
			} else if slices.Contains(psFlags, psShort) {
				t.Errorf("The field %s uses a PowerShell short flag `-%s` that has been used in another argument.", field.Name, psShortOrig)
			} else {
				psFlags = append(psFlags, psShort)
			}
		}
	}
}
