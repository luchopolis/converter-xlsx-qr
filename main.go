package main

import (
	"convertqr/service/files"
	typesLinks "convertqr/types"
	"convertqr/utils"
	"log"
	"slices"
)

func main() {
	columnsExpected := []string{"links"}
	rows, error := files.ReaderXlsx("./links.xlsx")

	if error != nil {
		log.Fatal("Could not find links.xlsx file")
	}

	linksToConvert := []typesLinks.Link{}

	for row := range rows {
		for _, cell := range row.Cells {
			checkExist := slices.IndexFunc(columnsExpected, func(c string) bool { return c == cell.Value })
			if checkExist < 0 {
				linksToConvert = append(linksToConvert, typesLinks.Link{
					Link: cell.Value,
				})
			}
		}
	}

	utils.QRConverter(linksToConvert)
}
