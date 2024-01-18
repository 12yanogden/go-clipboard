package clipboard

import (
	"reflect"
	"testing"

	"golang.design/x/clipboard"
)

func TestPush(t *testing.T) {
	str := "Copy this stuff"
	expected := str
	actual := ""

	Push(str)

	actual = string(clipboard.Read(clipboard.FmtText))

	if expected != actual {
		t.Fatalf("\nExpected:\t'%s'\nActual:\t\t'%s'\n", expected, actual)
	}
}

func TestTop(t *testing.T) {
	str := "Copy this stuff"
	expected := str
	actual := ""

	initialize()

	clipboard.Write(clipboard.FmtText, []byte(str))

	actual = Top()

	if expected != actual {
		t.Fatalf("\nExpected:\t'%s'\nActual:\t\t'%s'\n", expected, actual)
	}
}

func TestPushBytes(t *testing.T) {
	bytes := []byte("Copy this stuff")
	expected := bytes
	var actual []byte

	PushBytes(bytes)

	actual = clipboard.Read(clipboard.FmtText)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t'%#v'\nActual:\t\t'%#v'\n", expected, actual)
	}
}

func TestTopBytes(t *testing.T) {
	bytes := []byte("Copy this stuff")
	expected := bytes
	var actual []byte

	initialize()

	clipboard.Write(clipboard.FmtText, bytes)

	actual = TopBytes()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t'%#v'\nActual:\t\t'%#v'\n", expected, actual)
	}
}
