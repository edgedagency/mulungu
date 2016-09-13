package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//HTTPException function used to broadcast an execption
func HTTPException(w http.ResponseWriter, r *http.Request, response *Response) {
	response.Code = http.StatusInternalServerError

	//check client accepted header to respond with the appopriate response
	b, err := json.Marshal(response)

	if err != nil {
		fmt.Println("Failed to process reponse")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(b)
}
