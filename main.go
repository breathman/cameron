package main

import (
	"bytes"
	"fmt"
	"github.com/aofei/cameron"
	"image/jpeg"
	"os"
)

func main() {
	f, err := os.OpenFile("avatar.jpeg", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	buf := bytes.Buffer{}
	err = jpeg.Encode(
		&buf,
		cameron.Identicon(
			[]byte("Donald Trump"),
			600,
			50,
		),
		&jpeg.Options{
			Quality: 100,
		},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = f.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
