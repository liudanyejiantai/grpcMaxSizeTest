// grpc将单个4M放开为20M
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	PbDownMaxSize "test/grpcMaxSizeTest"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// 获得grpc的客户端IP地址和端口信息
func GetClietInfo(ctx context.Context) (string, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("[getClinetIP] invoke FromContext() failed")
	}
	if pr.Addr == net.Addr(nil) {
		return "", fmt.Errorf("[getClientIP] peer.Addr is nil")
	}
	return pr.Addr.String(), nil
}

func main() {
	hostAndPort := ":7575"
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.ListenTCP("tcp", serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	maxSize := 1024 * 1024 * 20
	// 这里是最关键的
	s := grpc.NewServer(grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize))
	PbDownMaxSize.RegisterPbDownMaxSizeServer(s, &server{})
	s.Serve(listener)
}

type server struct{}

// 实现proto文件中的接口
func (s *server) DownLoad(ctx context.Context, in *PbDownMaxSize.ReqDown) (*PbDownMaxSize.ResDown, error) {
	clientStr, err := GetClietInfo(ctx)
	if err != nil {
		fmt.Printf("获得grpc客户端地址错误,%s\n", err.Error())
	}

	var reply = new(PbDownMaxSize.ResDown)
	var msg string

	if reply.FileBuf, err = ReadFileBuffer(in.Name); err != nil {
		fmt.Printf("%s读取文件[%s]失败，原因:%s\n", clientStr, in.Name, err.Error())
		reply.IRet, reply.RetMsg = -1, msg
		return reply, err
	}
	reply.IRet, reply.RetMsg = 0, "success"
	fmt.Printf("[%s]下载文件[%s]成功\n", clientStr, in.Name)

	return reply, nil
}

// 读取文件
func ReadFileBuffer(strFileName string) ([]byte, error) {
	file, err := os.OpenFile(strFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeType)
	if err != nil {
		return nil, fmt.Errorf("读取文件[%s]失败,原因,%s", strFileName, err.Error())
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return buf, fmt.Errorf("打开文件[%s]失败,原因,%s", strFileName, err.Error())
	}
	return buf, err
}
