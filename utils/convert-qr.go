package utils

import (
	typesLinks "convertqr/types"
	"fmt"
	"strconv"
	"time"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func randomValue() string {
	currentTime := time.Now()
	milliseconds := currentTime.UnixNano() / int64(time.Millisecond)

	return strconv.Itoa(int(milliseconds))
}
func QRConverter(links []typesLinks.Link) {

	for _, link := range links {
		qrc, err := qrcode.New(link.Link)

		if err != nil {
			fmt.Printf("could not generate QRCode: %v", err)
			return
		}

		// generate a time to concact as filename
		fileName := fmt.Sprintf("qr-code-%s", randomValue())
		writeFile, err := standard.New(fileName)
		if err != nil {
			fmt.Printf("Cnat Standart: %v", err)
			return
		}

		// save file
		if err = qrc.Save(writeFile); err != nil {
			fmt.Printf("Cant Create the Image: %v", err)
		}

	}
}
