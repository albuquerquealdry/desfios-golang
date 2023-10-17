package main

import {
	"github.com/badoux/checkmail"
	"fmt"
}

func main() {
    err := checkmail.ValidateFormat("ç$€§/az@gmail.com")
    if err != nil {
        fmt.Println(err)
    }

	

}
