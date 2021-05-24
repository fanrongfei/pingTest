package pkg

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func GoPing(remoteServer string,c,w int) {
	addr, err := net.ResolveUDPAddr("udp", remoteServer)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err1 := net.DialUDP("udp", nil, addr)
	if err1 != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	msg:=[]byte("hello")
	successNum,failNum:=0,0
	var costTimeArr =make([]int64,0,c)
	fmt.Printf("正在 Ping %s 具有 %d 字节的数据:\n",strings.Split(remoteServer,":")[0],len(msg))
	for i:=0;i<c;i++ {
		_, err = conn.Write(msg)
		if err != nil {
			failNum++
			continue
		}
		tb := time.Now()
		data := make([]byte, 4)
		conn.SetDeadline(time.Now().Add(time.Second*time.Duration(w)))
		_, err = conn.Read(data)
		if err != nil {
			failNum++
			continue
		}
		successNum++
		costTime := time.Now().Sub(tb).Microseconds()
		costTimeArr = append(costTimeArr, costTime)
		fmt.Printf("来自%s 的回复: 时间 = %v μs\n",strings.Split(remoteServer,":")[0], costTime)
	}
	miniTime,maxTime,avgTime:=decodeTime(costTimeArr)
	fmt.Printf("%s 的Ping 统计信息:\n数据包: 已发送 = %d,已接受 = %d,丢失 =%d <%.2f%%丢失>，往返行程的估计时间<以微妙为单位>:\n最短 = %d μs,最长 = %d μs,平均 = %d μs",
		strings.Split(remoteServer,":")[0],c,successNum,failNum,float32(failNum)/float32(c)*100,miniTime,maxTime,avgTime)
	os.Exit(0)
}
func decodeTime(timeArr []int64)(miniTime,maxTime,avgTime int64){
	if len(timeArr)==0{
		return
	}
	var totalTime int64
	for i:=0;i<len(timeArr);i++{
		totalTime=totalTime+timeArr[i]
		if i==0{
			miniTime=timeArr[i]
			maxTime=timeArr[i]
			continue
		}
		if timeArr[i]<=miniTime{
			miniTime=timeArr[i]
		}
		if timeArr[i]>=maxTime{
			maxTime=timeArr[i]
		}
	}
	avgTime=totalTime/int64(len(timeArr))
	return
}
