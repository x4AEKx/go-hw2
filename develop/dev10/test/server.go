package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, _ := net.Listen("tcp", ":8081")

	// Открываем порт
	conn, _ := ln.Accept()

	// Запускаем цикл
	for {
		// Будем прослушивать все сообщения разделенные \n
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// Распечатываем полученое сообщение
		fmt.Print("Message Received:", string(message))
		// Процесс выборки для полученной строки
		newmessage := "Message from server:" + strings.ToUpper(message)
		// Отправить новую строку обратно клиенту
		conn.Write([]byte(newmessage + "\n"))
	}

}
