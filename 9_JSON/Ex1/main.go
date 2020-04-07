package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type StatCodes []struct {
	Code    string `json:"Code"`
	Descrip int    `json:"Descrip"`
}


func main() {
	var data StatCodes

	//   ` is the character on the ~ key
	rcvd := `[{"Code":"StatusOK","Descrip":200},{"Code":"StatusMovedPermanently","Descrip":301},{"Code":"StatusFound","Descrip":302},{"Code":"StatusSeeOther","Descrip":303},{"Code":"StatusTemporaryRedirect","Descrip":307},{"Code":"StatusBadRequest","Descrip":400},{"Code":"StatusUnauthorized","Descrip":401},{"Code":"StatusPaymentRequired","Descrip":402},{"Code":"StatusForbidden","Descrip":403},{"Code":"StatusNotFound","Descrip":404},{"Code":"StatusMethodNotAllowed","Descrip":405},{"Code":"StatusTeapot","Descrip":418},{"Code":"StatusInternalServerError","Descrip":500}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)

	for _, v := range data{
		fmt.Printf("%s %d\n", v.Code, v.Descrip)
	}
}
