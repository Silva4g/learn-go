package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// int, int8, int16, int32, int64
	// unit (unsigned int) ele so pode armazenar valores positivos diferente do int que armazena valores negativos
	// uint, uint8, uint16, uint32, uint64
	//var numero int16 = 1000000000
	//fmt.Println(numero)
	var numero0 int
	fmt.Println(numero0)
	var numero int64 = 1000000000
	fmt.Println(numero)
	var numero2 int = -1000000000 // igual numero2 := 100000
	fmt.Println(numero2)
	var numero3 uint32 = 100000
	fmt.Println(numero3)

	//alias
	var numero4 rune = 123456 //int32 = rune
	fmt.Println(numero4)
	var numero5 byte = 123 //uint8 = byte
	fmt.Println(numero5)

	//float32, float64
	var numeroReal float32
	fmt.Println(numeroReal)
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)
	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)
	numeroReal3 := 12345.6789
	fmt.Println(numeroReal3)

	var str string = "textao"
	fmt.Println(str)
	str2 := "texitinho"
	fmt.Println(str2)
	char := '%' // uma string com aspas simples ele pega o valor da tabela ascii, se colocar 2 valores ele nao compila
	fmt.Println(char)
	var texto string
	fmt.Println(texto)

	var boleano1 bool
	fmt.Println(boleano1)
	var boleano2 bool = true
	fmt.Println(boleano2)

	var erro error
	fmt.Println(erro)
	var erro1 error = errors.New("Erro Interno")
	fmt.Println(erro1)

	// Brincando com cor no terminal
	fmt.Println(Color(erro1, Red, Invert))
}

var (
	// Colors and font options via ANSI escape codes
	Reset     = "\033[0m"
	Black     = "\033[30m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	Gray      = "\033[37m"
	White     = "\033[97m"
	Bold      = "\033[1m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Invert    = "\033[7m"
)

func Color(input interface{}, color ...string) string {
	var s string
	c := ""
	for i := range color {
		c = c + color[i]
	}
	switch v := input.(type) {
	case error:
		s = c + v.Error()
	case int:
		s = c + strconv.Itoa(v) + Reset
	case bool:
		s = c + strconv.FormatBool(v) + Reset
	case []string:
		s = c + strings.Join(v, ", ") + Reset
	case string:
		s = c + v + Reset
	default:
		fmt.Printf("Unsupported type provided to Color func - %T\n", v)
	}
	return s
}
