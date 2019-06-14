package controllers

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
	"github.com/webonise/csv_upload/pkg/framework"
)

func (s *Srv) RenderUploadView(w *framework.Response, r *framework.Request) {
	tmplList := []string{
		"./web/views/upload.html",
		"./web/layouts/flash.html",
	}
	res, err := s.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		s.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	w.RenderHTML(res)
}

func (s *Srv) UploadFile(w *framework.Response, r *framework.Request) {

	var returnMessage string

	file, handler, err := r.FormFile("file")
	if err != nil {
		returnMessage = "Error while retrieving uploaded file"
	}
	defer file.Close()
	//Extracting the extension of file
	fslice := strings.Split(handler.Filename, ".")
	fileExtension := fslice[len(fslice)-1]
	//Handling CSV and XLSX file extension
	switch fileExtension {
	case "csv": //Handle file with extension CSV
		reader := csv.NewReader(file)
		record, _ := reader.ReadAll()
		for _, line := range record[1:] {
			err := s.Service.CheckEmployeesExist(line)
			if err != nil {
				//fmt.Println("Error while processing CSV extension", err)
				returnMessage = "Error while processing CSV extension"
			}
		}
		defer file.Close()

	case "xlsx": //Handle file with extension XLSX
		f, err := os.OpenFile("./web/assets/files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			//fmt.Println("Error while trying to open a file", err)
			returnMessage = "Error while trying to open a file"
		}
		defer f.Close()
		//Copy the file content
		io.Copy(f, file)

		xlFile, _ := xlsx.OpenFile("./web/assets/files/" + handler.Filename)

		for _, sheet := range xlFile.Sheets {
			//ignore first row because it contain Header
			for _, row := range sheet.Rows[1:] {
				emp := make([]string, 0)
				for _, cell := range row.Cells {
					text := cell.String()
					emp = append(emp, text)
				}
				err := s.Service.CheckEmployeesExist(emp)
				if err != nil {
					//fmt.Println("Error while processing XLSX extension", err)
					returnMessage = "Error while processing XLSX extension"
				}
			}
		}

	default:
		//if extension not supported
		returnMessage = "Upload file with " + fileExtension + " extension not allowed"
	}

	returnStatus := false
	if returnMessage == "" {
		returnMessage = "File uploaded and information update sucessfully"
		returnStatus = true

	}
	w.Message(returnMessage)
	w.SetSuccess(returnStatus)
}
