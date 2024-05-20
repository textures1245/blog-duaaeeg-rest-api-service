package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"time"

	"github.com/textures1245/BlogDuaaeeg-backend/pkg/error/entity"
	"github.com/xuri/excelize/v2"
)

type Excel[T interface{}] struct {
	Data []*T
}

type ExcelData struct {
	FileName   string
	FileBuffer bytes.Buffer
}

func (e Excel[T]) ExportData() (*ExcelData, error) {
	genTimestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("file_%s.xlsx", genTimestamp)
	f := excelize.NewFile()
	index, err := f.NewSheet(fileName)
	if err != nil {
		return nil, &entity.CError{
			Err:        err,
			StatusCode: 500,
		}
	}

	v := reflect.ValueOf(*e.Data[0])

	// set first row as column name from key struct
	for j := 0; j < v.NumField(); j++ {
		val := v.Type().Field(j).Name

		cellName := fmt.Sprintf("%c%d", 'A'+j, 1)
		f.SetCellValue(fileName, cellName, val)
	}

	for i, dat := range e.Data {
		// TODO: Refactor this to skip the pointer check if the value is a inside a struct not an pointer

		// increment i for skipped the first row
		if i == 0 {
			i++
		}

		v := reflect.ValueOf(dat).Elem()

		// Loop over the fields of the product
		for j := 0; j < v.NumField(); j++ {
			// Get the interface{} value
			val := v.Field(j).Interface()

			// Check if the value is a pointer
			rv := reflect.ValueOf(val)
			if rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Slice || rv.Kind() == reflect.Map || rv.Kind() == reflect.Chan || rv.Kind() == reflect.Func || rv.Kind() == reflect.Interface {
				// If the value is a pointer and not nil, dereference it
				if !rv.IsNil() {
					rv = rv.Elem()
					val = rv.Interface()
				} else {
					// If the value is a nil pointer, set val to nil
					val = nil
				}
			}

			// Set the value of the cell in the Excel file
			cellName := fmt.Sprintf("%c%d", 'A'+j, i+1)
			f.SetCellValue(fileName, cellName, val)
		}
	}
	f.SetActiveSheet(index)

	var buf bytes.Buffer
	if errOnW := f.Write(&buf); errOnW != nil {
		return nil, &entity.CError{
			Err:        errOnW,
			StatusCode: 500,
		}
	}

	return &ExcelData{
		FileName:   fileName,
		FileBuffer: buf,
	}, nil
}
