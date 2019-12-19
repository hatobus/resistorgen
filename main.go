package p

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/hatobus/resistorgen/resistor"
	"github.com/nlopes/slack"
)

type PostParam struct {
	Token          string `json:"token"`
	Channels       string `json:"channels"`
	Filename       string `json:"filename"`
	InitialComment string `json:"initial_comment"`
	Title          string `json:"title"`
}

func ResistorResolver(w http.ResponseWriter, r *http.Request) {
	api := slack.New(os.Getenv("ACCESS_TOKEN"))

	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !s.ValidateToken(os.Getenv("VERIFICATION_TOKEN")) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/resistor":
		resistor := s.Text
		fname, err := generatepic(resistor)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		file, err := os.Open(fname)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// upload picture
		params := slack.FileUploadParameters{
			File: fname,
			// Content:        "picture",
			Reader:         file,
			Filename:       fname,
			Title:          resistor,
			InitialComment: fmt.Sprintf("%sの抵抗の色です", resistor),
			Channels:       []string{s.ChannelName},
		}

		_, err = api.UploadFile(params)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// params := &slack.Msg{
		// 	Channel: s.ChannelName,
		// 	Text:    fmt.Printf("%sのカラーコードはこちらです", resistor),
		// }

	default:
	}

	w.WriteHeader(http.StatusOK)
	return

}

func generatepic(value string) (string, error) {
	colors, _ := resistor.GenerateColor(value)

	width := 200
	height := 100

	resistorimage := image.NewRGBA(image.Rect(0, 0, width, height))
	fillcolor := color.RGBA{255, 192, 203, 255}

	draw.Draw(resistorimage, resistorimage.Bounds(), &image.Uniform{fillcolor}, image.ZP, draw.Src)

	firstband := image.Rect(24, 0, 44, height)
	draw.Draw(resistorimage, firstband, &image.Uniform{colors.FirstBand}, image.ZP, draw.Src)

	secondband := image.Rect(68, 0, 88, height)
	draw.Draw(resistorimage, secondband, &image.Uniform{colors.SecondBand}, image.ZP, draw.Over)

	thirdband := image.Rect(112, 0, 132, height)
	draw.Draw(resistorimage, thirdband, &image.Uniform{colors.ThirdBand}, image.ZP, draw.Over)

	fourthband := image.Rect(156, 0, 176, height)
	draw.Draw(resistorimage, fourthband, &image.Uniform{colors.Tolerance}, image.ZP, draw.Over)

	fname := fmt.Sprintf("/tmp/resistor%s.png", value)

	f, _ := os.Create(fname)
	png.Encode(f, resistorimage)

	return fname, nil

}
