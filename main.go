package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main(){
	fmt.Println("Listening on port 6379")
	l,err:=net.Listen("tcp",":6379")
	if err!=nil{
		fmt.Println(err)
		return
	}
	conn,err:=l.Accept()
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		// buf:=make([]byte,1024) //buf := []byte{0, 0, 0, ..., 0} // 1024 times
		// value,err:=resp.Read()
		resp:=NewResp(conn)
		value,err=resp.Read()
		if err!=nil {
			// if err==io.EOF {
			// 	break //io.EOF: this means the client closed the connection, break the loop and stop
			// }
			fmt.Println(err)
			return
		}
		_=value
		// fmt.Println(value)
		// conn.Write([]byte("+OK\r\n"))
		writer:=NewWriter(conn)
		writer.Write(Value{typ: "string", str: "OK"})
	}

}
