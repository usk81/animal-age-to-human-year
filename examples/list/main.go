package main

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/usk81/animal-age-to-human-year/dog"
)

func main() {
	funcs := []func(float64) int{
		dog.Calc,
		dog.WrongCalc,
	}
	for _, f := range funcs {
		fmt.Println(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		for age := 0; age <= 20; age++ {
			fmt.Printf("dog age: %d / human year: %d\n", age, f(float64(age)))
		}
	}
}
