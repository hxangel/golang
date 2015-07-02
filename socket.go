package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"bytes"
	"encoding/binary"
	"iconv"
)

/*
typedef struct _Data {
    UINT32 data1;
    UINT32 data2;
} Data;
*/

var msg string = "[\"USD\",\"USDJPY\"]"
var msg_l = len(msg)
var str string

func main() {
	buf := make([]byte, 16)
	//	mbuf := make([]byte, 7)
	binary.BigEndian.PutUint16(buf[1:], 1)
	binary.BigEndian.PutUint32(buf[8:], 1001)
	binary.BigEndian.PutUint32(buf[12:], uint32(msg_l))
	//	binary.Write(mbuf,[]byte(msg))
	//	binary.LittleEndian.PutUint16(buf[2:], 1)
	//	binary.LittleEndian.PutUint32(buf[8:], 1001)
	//	binary.LittleEndian.PutUint32(buf[12:], 7)
	fmt.Println(buf)
	var data = make([]byte, 1024)
	var rheader = make([]byte, 16)
	var (
		host   = "61.145.163.218"
//		host   = "192.168.0.165"
		port   = "6688"
		remote = host + ":" + port
	)
	con, err := net.Dial("tcp", remote)
	defer con.Close()
	if err != nil {
		fmt.Println("Server not found.")
		os.Exit(-1)
	}
	fmt.Println("Connection OK.")

	//	send data
	in, err := con.Write(buf)
	if err != nil {
		fmt.Printf("Error when send to server: %d\n", in)
		os.Exit(0)
	}
	sm, err := con.Write([]byte(msg))
	if err != nil {
		fmt.Printf("Error when send to server: %d\n", sm)
		os.Exit(0)
	}

	//	read header
	ll, err := con.Read(rheader)
	if err != nil {
		fmt.Printf("Error when read from server.\n")
		os.Exit(0)
	}
	fmt.Printf("%d\n",ll)
	aa, err := ConverttoUint(rheader[12:ll])
	if err != nil {
		fmt.Printf("Convert failed\n")
		os.Exit(0)
	}
	//	ret := read_int32(data[ll-4:ll]);
	fmt.Println(aa)
	var res string
	//	get data
	for {
		length, err := con.Read(data)
		if err != nil {
			fmt.Printf("Error when read from server.\n")
			os.Exit(0)
		}
		str = string(data[0:length])
		res +=ConvertEncoding(str)

	}
	fmt.Println(res)
}


func ConvertEncoding(str string)(ret string){
	cd, err := iconv.Open("utf-8", "gbk")
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()
	ret  = cd.ConvString(str)
	return  ret
}

func Converttoint(data []byte) (uint32, error) {
	v, err := strconv.ParseUint(string(data), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

func read_int32(data []byte) (ret uint32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &ret)
	return
}
func ConverttoUint(data []byte) (uint32, error) {
	var v uint32
	err := binary.Read(bytes.NewReader(data), binary.BigEndian, &v)
	if err != nil {
		return 0, err
	}
	return v, nil
}
