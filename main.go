package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {

	newUuid := CreateUUID()
	fmt.Println(a.first)

}

func CreateUUID()string {
	uuidWithHyphen := uuid.New()
	// fmt.Println(uuidWithHyphen)
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	// fmt.Println(uuid)
	return uuid
}
