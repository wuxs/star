package main

import (
	"flag"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"github.com/wuxs/star/pkg/star"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	connType := flag.String("t", "serial", "connect type: serial / tcp / file")
	port := flag.String("p", "/dev/ttyS6", "serial port name or tcp address or file path")
	flag.Parse()

	var conn io.ReadWriteCloser
	switch *connType {
	case "serial":
		options := serial.OpenOptions{
			PortName:        *port,
			BaudRate:        9600,
			DataBits:        8,
			StopBits:        1,
			MinimumReadSize: 4,
		}
		s, err := serial.Open(options)
		conn = s
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	case "tcp":
		s, err := net.Dial("tcp", *port)
		conn = s
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	case "file":
		s, err := os.OpenFile(*port, os.O_CREATE|os.O_WRONLY, 644)
		conn = s
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	default:
		fmt.Println("unknown type")
		return
	}
	defer conn.Close()
	s := star.NewStar(conn)
	s.Init().
		SpecifyAlignment(star.Center).
		SpecifyBold().
		Print("车辆通行费\n").
		CancelBold().
		SpecifyAlignment(star.Left).
		Print("\n").
		Print(fmt.Sprintf("          发票代码: %s\n", "1234567890")).
		Print(fmt.Sprintf("          发票代码: %s\n", "1234567890")).
		Print("\n").
		Print(fmt.Sprintf("    车类: %s\n", "客一型")).
		Print(fmt.Sprintf("    车号: %s\n", "京A12345")).
		Print(fmt.Sprintf("    入口: %s\n", "西出口")).
		Print(fmt.Sprintf("    出口: %s\n", "东出口口")).
		Print(fmt.Sprintf("    日期: " + time.Now().Format("2006-01-01 15:04") + "\n")).
		FeedPaperLines(5).
		CutFull()
	_, err := s.Flush()
	if err != nil {
		fmt.Println(err.Error())
	}
}
