package ent

import (
	"reflect"
	"sort"
	"testing"

	"github.com/athomecomar/storeql/name"
	"github.com/gedex/inflector"
	"github.com/google/go-cmp/cmp"
)

func TestOnboardingSQLColumns(t *testing.T) {
	u := User{}
	typeOf := reflect.TypeOf(u)
	var want []string
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		col := name.ToSnakeCase(field.Name)
		want = append(want, col)
	}
	sort.Strings(want)

	got := u.SQLColumns()
	sort.Strings(got)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Onboarding.SQLColumns() mismatch (-want +got): %s", diff)
	}
}

func TestOnboardingsSQLTable(t *testing.T) {
	u := Onboarding{}
	typeOf := reflect.TypeOf(u)
	want := inflector.Pluralize(name.ToSnakeCase(typeOf.Name()))
	got := u.SQLTable()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Onboarding.SQLTable() mismatch (-want +got): %s", diff)
	}
}
