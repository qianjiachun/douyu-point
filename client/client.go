package client

import (
	"bytes"
	"douyu-point/common"
	"encoding/binary"
	"io"
	"net"

	"time"
)

const ADDR  = "danmuproxy.douyu.com:8601"

type DouyuClient struct {
	Rid string
	conn net.Conn
}

func (client *DouyuClient) Connect(callback func(data string)) {
	var err error
	client.conn, err = net.Dial("tcp", ADDR)

	if err != nil {
		println("connect to server failed!")
		return
	}
	_, err = client.conn.Write(buildDouyuPkg("type@=loginreq/roomid@=" + client.Rid + "/"))
	_, err = client.conn.Write(buildDouyuPkg("type@=joingroup/rid@=" + client.Rid + "/gid@=-9999/"))
	common.CheckErr(err)
	go client.heartBeat()
	go client.recv(callback)
}

func (client *DouyuClient) recv(callback func(data string)) {
	//danmuReg  := regexp.MustCompile("type@=chatmsg/.*rid@=(\\d*?)/.*uid@=(\\d*).*nn@=(.*?)/txt@=(.*?)/(.*)/")
	for {
		buf := make([]byte, 512)
		if _, err := io.ReadFull(client.conn, buf[:12]); err != nil {
			break
		}
		pl := binary.LittleEndian.Uint32(buf[:4])
		cl := pl - 8
		if cl > 512 {
			buf = make([]byte, cl)
		}

		if _, err := io.ReadFull(client.conn, buf[:cl]); err != nil {
			break
		}
		callback(common.Bytes2str(buf[:cl-1]))
	}

}
func (client *DouyuClient) heartBeat() {
	for {
		_, err := client.conn.Write(buildDouyuPkg("type@=mrkl/"))
		common.CheckErr(err)
		time.Sleep(time.Second * 40)
	}
}

func (client *DouyuClient) Close() {
	client.Close()
}


func buildDouyuPkg(str string) []byte {
	var err error
	data := new(bytes.Buffer)
	rawLen := len([]byte(str)) + 9
	err = binary.Write(data, binary.LittleEndian, int32(rawLen))
	err = binary.Write(data, binary.LittleEndian, int32(rawLen))
	err = binary.Write(data, binary.LittleEndian, int16(689))
	err = binary.Write(data, binary.LittleEndian, byte(0))
	err = binary.Write(data, binary.LittleEndian, byte(0))
	err = binary.Write(data, binary.LittleEndian, []byte(str))
	err = binary.Write(data, binary.LittleEndian, byte(0))
	common.CheckErr(err)
	return data.Bytes()
}