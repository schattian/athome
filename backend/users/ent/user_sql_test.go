package ent

import (
	"reflect"
	"sort"
	"testing"

	"github.com/athomecomar/storeql/name"
	"github.com/gedex/inflector"
	"github.com/google/go-cmp/cmp"
)

func TestUserSQLColumns(t *testing.T) {
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
		t.Errorf("User.SQLColumns() mismatch (-want +got): %s", diff)
	}
}

func TestUserSQLTable(t *testing.T) {
	u := User{}
	typeOf := reflect.TypeOf(u)
	want := inflector.Pluralize(name.ToSnakeCase(typeOf.Name()))
	got := u.SQLTable()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("User.SQLTable() mismatch (-want +got): %s", diff)
	}
}
