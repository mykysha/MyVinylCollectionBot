package spreadsheetconverter

import (
	"bytes"
	"log"
	"strconv"

	"github.com/nndergunov/tgBot/bot/pkg/domain/entities"
	"github.com/xuri/excelize/v2"
)

func ConvertToExcel(albums []entities.Album) *bytes.Buffer {
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "Cover")
	f.SetCellValue("Sheet1", "B1", "Artist")
	f.SetCellValue("Sheet1", "C1", "Name")
	f.SetCellValue("Sheet1", "D1", "Genre")
	f.SetCellValue("Sheet1", "E1", "Label")
	f.SetCellValue("Sheet1", "F1", "Release year")
	f.SetCellValue("Sheet1", "G1", "Reissue year")
	f.SetCellValue("Sheet1", "H1", "Colored")

	for id, el := range albums {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(id+2), el.CoverID)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(id+2), el.Artist)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(id+2), el.Name)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(id+2), el.Genre)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(id+2), el.Label)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(id+2), el.ReleaseYear)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(id+2), el.ReissueYear)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(id+2), el.Coloured)
	}

	returnFile, err := f.WriteToBuffer()
	if err != nil {
		log.Println(err)
	}

	return returnFile
}
