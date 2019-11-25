package test

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	u "go_util/libs"
	"testing"
	"time"
)

type transfer struct {
	Message string `json:"message"`
	Phrase  string `json:"phrase"`
}

var data transfer = transfer{
	Message: "Message",
	Phrase:  "Test",
}

var secret []byte = []byte("12345qwerty")
var tokenString string
var err error

func TestJWT(t *testing.T) {
	Convey("Jwt pack", t, func() {
		tokenString, err = u.JwtPackStruct(&data, &secret, time.Minute*2)
		So(err, ShouldBeNil)
		So(len(tokenString), ShouldBeGreaterThan, 0)
		fmt.Printf("token string is %v", tokenString)
	})

	Convey("Jwt unpack", t, func() {
		arrByte, err := u.JwtUnpackStruct(&tokenString, &secret)
		So(err, ShouldBeNil)

		dt := transfer{}
		err = json.Unmarshal(arrByte, &dt)
		So(err, ShouldBeNil)
		So(dt.Message, ShouldEqual, data.Message)
		So(dt.Phrase, ShouldEqual, data.Phrase)
	})
}

func BenchmarkJwtPack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = u.JwtPackStruct(&data, &secret, time.Minute*2)
	}
}
