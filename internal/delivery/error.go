package delivery

import (
	"fmt"
	"log"
	"net/http"

	"grTracker/internal/models"
)

func ErrorCheck(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ErrorHandler(page http.ResponseWriter, status int, errText error) {
	var errT models.Error
	page.WriteHeader(status)
	errT.ErrorCode = fmt.Sprintf("%d %s\n", status, http.StatusText(status))
	errT.ErrorText = fmt.Sprintf("%s", errText)
	if err := tmpl.ExecuteTemplate(page, "error.html", errT); err != nil {
		ErrorCheck(err)
	}
}
