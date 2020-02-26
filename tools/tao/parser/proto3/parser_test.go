package proto3

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/repr"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	var parser = participle.MustBuild(&Proto{}, participle.UseLookahead(2))
	proto := &Proto{}
	r, err := os.Open("testdata/comment.proto")
	if err != nil {
		t.Fatal(err)
	}
	err = parser.Parse(r, proto)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range proto.Entries {
		if entry.Message != nil {

		} else if entry.Service != nil {

		}
	}
	repr.Println(proto, repr.Hide(&lexer.Position{}))
}
