// 碎片下载代理
package main

import (
	"flag"
	"fmt"
	"os"
	PbDownMaxSize "test/grpcMaxSizeTest"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	name     string
	endpoint string
)

func main() {
	flag.StringVar(&name, "name", "大图片.jpg", "文件路径")
	flag.StringVar(&endpoint, "endpoint", "127.0.0.1:7575", "服务器地址")
	flag.Parse()

	maxSize := 20 * 1024 * 1024
	// 这里最关键,设置最大消息20Mb
	diaOpt := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxSize), grpc.MaxCallSendMsgSize(maxSize))
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure(), diaOpt)

	if err != nil {
		fmt.Printf("连接grpc服务器[%s]失败,%s\n", endpoint, err.Error())
		return
	}
	c := PbDownMaxSize.NewPbDownMaxSizeClient(conn)
	var req = new(PbDownMaxSize.ReqDown)
	req.Name = name
	reply, err := c.DownLoad(context.Background(), req)
	if err != nil {
		fmt.Printf("从grpc服务器[%s]下载文件[%s]失败,%s\n", endpoint, name, err.Error())
		return
	}
	err = SaveFile(reply.FileBuf, name)
	if err != nil {
		fmt.Printf("从grpc服务器[%s]下载文件[%s]写入失败,%s\n", endpoint, name, err.Error())
	} else {
		fmt.Printf("从grpc服务器[%s]下载文件[%s]写入成功\n", endpoint, name)
	}
}

func SaveFile(bytBuf []byte, strFileName string) error {
	file, err := os.OpenFile(strFileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("写文件[%s]失败,%s", strFileName, err.Error())
	}

	// 覆盖写入文件
	defer file.Close()
	if _, err = file.Write(bytBuf); err != nil {
		return fmt.Errorf("写文件[%s]失败,%s", strFileName, err.Error())
	}
	return nil
}
