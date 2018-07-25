package dsvc_test

import (
	"bytes"
	"testing"

	. "github.com/Wing924/dsvc"
)

func TestDsvc(t *testing.T) {
	testData := []struct {
		Input  string
		Expect string
	}{
		{
			"", "",
		}, {
			"\n", "\n",
		}, {
			`abc`,
			"abc",
		}, {
			"abc\n",
			"abc\n",
		}, {
			"abc 123",
			"abc\t123",
		}, {
			`abc "123 456"`,
			"abc\t123 456",
		}, {
			"abc \"123\t456\"",
			"abc\t123\\t456",
		}, {
			`abc "123\"456"`,
			"abc\t123\"t456",
		}, {
			`abc "123\" 456"`,
			"abc\t123\" t456",
		}, {
			`abc "123""456"`,
			"abc\t123\"t456",
		}, {
			"abc def\n123\n456",
			"abc\tdef\n123\n456",
		}, {
			`127.0.0.1 user-identifier frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326`,
			"127.0.0.1\tuser-identifier\tfrank\t[10/Oct/2000:13:55:36\t-0700]\tGET /apache_pb.gif HTTP/1.0\t200\t2326",
		},
	}

	for _, tt := range testData {
		in := bytes.NewBufferString(tt.Input)
		r := NewConverter(in)
		out := &bytes.Buffer{}
		r.WriteTo(out)
		actual := out.String()
		if actual != tt.Expect {
			t.Errorf("\nInput:  %#v\nExpect: %#v\nActual: %#v", tt.Input, tt.Expect, actual)
		}
	}
}
