package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func JSON(w http.ResponseWriter, stat int, mes string) {
	json.NewEncoder(w).Encode(mes)
	w.WriteHeader(stat)
}
