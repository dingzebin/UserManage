package util

import "github.com/alexflint/go-arg"

type Arguments struct {
	Db string `arg:"required"`
}

var Args Arguments

func Parse() {
	Args = Arguments{}
	arg.MustParse(&Args)
}
