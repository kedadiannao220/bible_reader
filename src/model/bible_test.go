package model

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConvertBookID(t *testing.T) {
	//Init()
	//fmt.Println(ConvertBookID("1"))

	int, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(int)
}

