package server

import (
	"net"
	"runtime"
	"fmt"
	"time"
	"io"
	"bytes"
)

func Server(){
	// 创建一个监听器
	listener,err := net.Listen("tcp","127.0.0.1:8085")
	SimplePanic(err)
	defer func(){
		err:=listener.Close()
		SimplePanic(err)
		}()

	//监听器的地址
	LogPrint("监听器地址:",listener.Addr())

	for{
		conn,err :=listener.Accept()
		//观察一下同一个客户端同一个请求是不是同一个conn实例
		fmt.Println("连接实例",conn)
		SimplePanic(err)

		//连接源的地址
		LogPrint("收到客户端地址是: :" ,conn.RemoteAddr())

		go handleCon(conn)

	}




}

func handleCon(conn net.Conn){
	defer conn.Close()
	for {
		//设置尾期
		err := conn.SetDeadline(time.Now().Add(10 * time.Second))
		SimplePanic(err)
		strReq, err2 := read(conn)
		if err2!=nil{
			if err == io.EOF{
				LogPrint("服务端读完了请求")
			}else{
				SimplePanic(err)
			}
			break
		}
		LogPrint("收到请求数据#:",strReq)


		_,err = write(conn,"你好，收到消息了")
		SimplePanic(err)


	}
}

func write(conn net.Conn,content string) (int,error){
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte('#')
	return conn.Write(buffer.Bytes())
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

func SimplePanic(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		LogPrint(file, line, err)
		runtime.Goexit()
	}
}

func LogPrint(a ... interface{}){
	fmt.Println(a...)
}