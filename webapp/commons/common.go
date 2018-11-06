package commons

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/eaciit/gomail"
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"os"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GenerateRandomString(size int) string {
	y := ""
	for i := 0; i < size; i++ {
		y += string(letters[rand.Intn(len(letters))])
	}
	return y
}

func SendEmail(from, address, subject, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", address)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	d := gomail.NewPlainDialer("smtp.mandrillapp.com", 587, "greyroot00@gmail.com", "kOhVHNe49GA5arhldyko2w")
	//d := gomail.NewPlainDialer("smtp.gmail.com", 587, "aryadhira25@gmail.com", "asusultrabook")
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func UploadHandler(r *knot.WebContext, filename, dstpath string, replacename string) (error, string) {
	file, handler, err := r.Request.FormFile(filename)
	if err != nil {
		return err, ""
	}
	defer file.Close()

	dstSource := dstpath + toolkit.PathSeparator + replacename
	f, err := os.OpenFile(dstSource, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err, ""
	}
	defer f.Close()
	io.Copy(f, file)

	return nil, handler.Filename
}

func UploadHandlerCopy(r *knot.WebContext, filename, dstpath string) (error, string) {
	file, handler, err := r.Request.FormFile(filename)
	if err != nil {
		return err, ""
	}
	defer file.Close()

	dstSource := dstpath + toolkit.PathSeparator + handler.Filename
	f, err := os.OpenFile(dstSource, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err, ""
	}
	defer f.Close()
	io.Copy(f, file)

	return nil, handler.Filename
}

func ResizeImg(x int, y int, sourcefile string, destinationfile string) {
	// Read input file.
	// f, err := os.Open("assets/photos/"+ k.Request.FormValue("userid") +"/photo." + k.Request.FormValue("filetype"))
	f, err := os.Open(sourcefile)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	src, _, err := image.Decode(f)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Scale down by a factor of 2.
	sb := src.Bounds()
	// dst := image.NewRGBA(image.Rect(0, 0, 64, 64))
	dst := image.NewRGBA(image.Rect(0, 0, x, y))
	draw.BiLinear.Scale(dst, dst.Bounds(), src, sb, draw.Over, nil)

	// Write output file.
	// if f, err = os.Create("assets/photos/"+ k.Request.FormValue("userid") +"/photo_64." + k.Request.FormValue("filetype")); err != nil {
	if f, err = os.Create(destinationfile); err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	var opt jpeg.Options
	opt.Quality = 80
	if err := jpeg.Encode(f, dst, &opt); err != nil {
		if err := png.Encode(f, dst); err != nil {
			fmt.Println(err.Error())
		}
	}
}
