package utils

import (
	"encoding/json"
	"net/http"
)


type Envelope map[string]any


func WriteJSON(w http.ResponseWriter, statusCode int, responseData Envelope)error{
	js, err := json.MarshalIndent(responseData, "", " ")
	if err != nil {
		return err
	}
	 js = append(js, '\n')
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(statusCode)
	 w.Write(js)
	 return nil
}