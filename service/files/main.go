package files

import (
	xlsxUtil "github.com/thedatashed/xlsxreader"
)

func ReaderXlsx(path string) (chan xlsxUtil.Row, error) {
	var rows chan xlsxUtil.Row

	file, error := xlsxUtil.OpenFile(path)

	if error != nil {
		return rows, error
	}

	rows = file.ReadRows(file.Sheets[0])
	return rows, nil
}
