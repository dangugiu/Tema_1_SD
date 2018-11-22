package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  // connect to this socket bla bla
  conn, _ := net.Dial("tcp4", "127.0.0.1:8081")
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Nume?: ")
  text, _ := reader.ReadString('\n')
  fmt.Fprintf(conn, text + "\n")
  for { 
    // read in input from stdin
    fmt.Print("Text de trimis: ")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Mesaj de la server: "+message)
  }
}