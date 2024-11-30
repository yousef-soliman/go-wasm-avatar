package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"syscall/js"

	"github.com/o1egl/govatar"
)

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("handleSubmitForm", js.FuncOf(HanldeSubmitForm))
	<-make(chan struct{})
}

func logError(msg string) {
	js.Global().Get("document").
		Call("getElementById", "status").
		Set("innerText", msg)
}

func GenerateAvatar(gender govatar.Gender, username string) (image.Image, error) {
	img, err := govatar.GenerateForUsername(gender, username)
	if err != nil {
		logError(err.Error())
	}
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)
	imgBytes := buf.Bytes()
	dst := js.Global().Get("Uint8Array").New(len(imgBytes))
	js.CopyBytesToJS(dst, imgBytes)
	js.Global().Call("displayImage", dst)
	return img, err
}

func HanldeSubmitForm(this js.Value, args []js.Value) any {
	var typeGender govatar.Gender
	gender := args[0].String()
	if gender == "male" {
		typeGender = govatar.MALE
	} else {
		typeGender = govatar.FEMALE
	}
	username := args[1].String()
	fmt.Printf("username %s\n", username)
	fmt.Printf("gender %s\n", gender)
	_, err := GenerateAvatar(typeGender, username)
	if err != nil {
		return false
	}
	return true
}
