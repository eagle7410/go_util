package test

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"go_util/libs"
	"strings"
	"testing"
)

func TestTools(t *testing.T) {
	Convey("Gene password", t, func() {
		opt := lib.PasswordGeneOptions{}
		opt.Digits = "1"

		pass := string(lib.GenePassword(&opt))

		fmt.Printf("\n pass is %v \n", pass)

		So(len(pass), ShouldEqual, lib.PwdGeneCharsMinDefault)
		opt.Length = lib.PwdGeneCharsMin


		pass = string(lib.GenePassword(&opt))
		So(len(pass), ShouldEqual, lib.PwdGeneCharsMin)

		res := struct{Up, Down, Digit, Special int} {}

		fmt.Printf("\n pass is %+v \n", pass)

		for _, val := range pass {

			charStr := string(val)

			if strings.Index(opt.Up, charStr) > -1 {
				res.Up++
				continue;
			}

			if strings.Index(opt.Down, charStr) > -1 {
				res.Down++
				continue;
			}

			if strings.Index(opt.Digits, charStr) > -1 {
				So(charStr, ShouldEqual, "1")
				res.Digit++
				continue;
			}

			if strings.Index(opt.Specials, charStr) > -1 {
				res.Special++
				continue;
			}
		}

		So(res.Down,    ShouldEqual, 1)
		So(res.Up,      ShouldEqual, 1)
		So(res.Special, ShouldEqual, 1)
		So(res.Digit,   ShouldEqual, 1)

		fmt.Printf("\n res is %+v \n", res)
	})

	Convey("Use password (hash, salt)", t, func() {
		testPass := "1"

		drive := lib.PasswordConverter{}

		err := drive.GeneSalt()
		So(err, ShouldBeNil)

		drive.GeneHashFromPassword(&testPass)

		salt, hash := drive.GetSaltHash()

		So(len(salt), ShouldEqual, lib.PwdConverterSaltLen)
		So(len(hash), ShouldEqual, lib.PwdConverterKeyLen)

		drive2 := lib.PasswordConverter{}

		drive2.SetSaltPass(&salt, &hash)
		isValid := drive2.IsValidPass(&testPass)

		So(isValid, ShouldEqual, true)

		b64Salt, b64Hash := drive.GetSaltHashBase64Encode()

		err = drive2.SetSaltPassBase64Decode(&b64Salt, &b64Hash)
		So(err, ShouldBeNil)

		isValid = drive2.IsValidPass(&testPass)
		So(isValid, ShouldEqual, true)

		fmt.Printf("\n Salt %v , hash %v \n", b64Salt, b64Hash)
	})
}
