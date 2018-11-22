package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"unicode"
	"strconv"
	"io/ioutil"
	"os"
)

var numarMaximDeConexiuni int
var numarCurentDeConexiuni int = 0

func handleConnection(c net.Conn) {
	adresaClient := c.RemoteAddr().String()
	nume, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
			fmt.Println(err)
			return
	}
	fmt.Printf("S-a conectat %s aka %s\n", adresaClient, nume)
	for {
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
					fmt.Println(err)
					return
			}

			temp := strings.TrimSpace(string(netData))
			if temp == "STOP" {
					numarCurentDeConexiuni --
					fmt.Printf("S-a deconectat %s\n", nume)
					break
			}

			cifra1 := 0
			cifra2 := 0
			var raspuns strings.Builder

			for _, caracter := range temp {
				if unicode.IsNumber(caracter) {
					if cifra1 != 0 {
						cifra2 = int(caracter - '0')
						cifra1 = cifra1 * 10
					} else {
						cifra1 = int(caracter - '0')
					}
				} else if unicode.IsLetter(caracter) {
					suma := cifra1 + cifra2
					if suma <= 20 {
						for i := 0; i < suma; i ++ {
							raspuns.WriteString(string(caracter))
						}
						cifra1 = 0
						cifra2 = 0
					}
				} else {
					raspuns.WriteString(string(caracter))
				}
			}
			fmt.Printf("Am sa trimit %s la %s\n", raspuns.String(), adresaClient)
			raspuns.WriteString("\n")
			c.Write([]byte(raspuns.String()))
			

	}
	c.Close()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	fmt.Printf("Acest server a pornit!\nSper ca te simti fericit!\n")
	l, err := net.Listen("tcp4", ":8081")
	if err != nil {
			fmt.Println(err)
			return
	}
	defer l.Close()

	pwd, _ := os.Getwd()
	b, err := ioutil.ReadFile(pwd + "/config.txt")
	check(err)
	numarMaximDeConexiuni, _ = strconv.Atoi(string(b))
	fmt.Println("Numarul maxim de clienti care se pot conecta concurent:", numarMaximDeConexiuni)

	for {
		if (numarCurentDeConexiuni < numarMaximDeConexiuni) {
			numarCurentDeConexiuni ++
			c, err := l.Accept()
			if err != nil {
					fmt.Println(err)
					return
			}
			go handleConnection(c)
		}
	}
}