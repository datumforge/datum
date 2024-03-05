package views

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image/png"
	"net/http"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

var secretBase32 = ""
var qrCodeBase64 = ""

func GenerateTOTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/totp.html"))
	context := map[string]interface{}{}

	data_action := r.FormValue("data_action")
	issuer := r.FormValue("issuer")
	accountName := r.FormValue("accountName")
	haveKey := r.FormValue("haveKey")

	if data_action == "GENERATE KEY" {
		fmt.Println("ISSUER: ", issuer)
		fmt.Println("ACCOUNT NAME: ", accountName)

		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      issuer,
			AccountName: accountName,
			Period:      30,
			SecretSize:  10,
			Algorithm:   otp.AlgorithmSHA256,
		})
		CheckErr(w, err)

		secretBase32 = key.Secret()
		qrCode, err := key.Image(200, 200)
		CheckErr(w, err)

		qrCodeBuffer := new(bytes.Buffer)
		err = png.Encode(qrCodeBuffer, qrCode)
		CheckErr(w, err)

		qrCodeBase64 = base64.StdEncoding.EncodeToString(qrCodeBuffer.Bytes())
		context["generateSecret"] = key.Secret()
		context["qrCode"] = qrCodeBase64
	}

	if data_action == "GENERATE TOTP" {
		totpCode, err := TOTPGenerator(secretBase32)
		CheckErr(w, err)

		context["generateTOTP"] = totpCode
		context["key"] = secretBase32
		context["qr"] = qrCodeBase64
	}

	if data_action == "HAVE A KEY" {
		totpCode, err := TOTPGenerator(haveKey)
		CheckErr(w, err)

		context["haveTOTP"] = totpCode
		context["genKey"] = haveKey
		context["haveKey"] = secretBase32
	}
	tmpl.Execute(w, context)
}

func TOTPGenerator(secret string) (string, error) {
	key, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return key, nil
}

func CheckErr(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
