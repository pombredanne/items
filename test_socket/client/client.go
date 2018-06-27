package client

import (
	"net"
	"runtime"
	"fmt"
	"time"
	"bytes"
)

func Client() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8085", 20*time.Second)
	simplePanic(err)
	defer func() {
		simplePanic(err)
	}()

	fmt.Println(conn.LocalAddr(), "发送给", conn.RemoteAddr())

	time.Sleep(5 * time.Second)

	//接收数据数据
	go func() error {
		for {
			str, _ := read(conn)
			fmt.Println(str)
		}
		return nil
	}()
	for i := 0; i < 5; i++ {
		_, err = write(conn, "你好，我是客户端，我要吃屎")
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}

func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content + "#")
	return conn.Write(buffer.Bytes())
}

func simplePanic(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file, line, err)
		runtime.Goexit()
	}
}

func read(conn net.Conn) (string,error){
	readBytes :=make([]byte,1)
	var buffer bytes.Buffer
	for{
		_,err := conn.Read(readBytes)
		if err !=nil {
			return "",err
		}
		readByte := readBytes[0]
		if readByte == '#'{
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(),nil
}