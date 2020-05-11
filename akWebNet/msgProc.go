package akWebNet

/*
	初级消息处理
*/

import (
	"github.com/gorilla/websocket"
	"fmt"
	"encoding/binary"
)

type receiveMsgProc func(sess *WebSession, msg *wsMessage)

var (
	_procMsgs = map[int]receiveMsgProc{
		websocket.TextMessage: 		TextMessageFunc,
		websocket.BinaryMessage: 	BinaryMessageFunc,
		websocket.CloseMessage:		CloseMessageFunc,
		websocket.PingMessage:		PingMessageFunc,
		websocket.PongMessage:		PongMessageFunc,
	}
)

func GetMessageHandler(id int)receiveMsgProc{
	return _procMsgs[id]
}

func TextMessageFunc(sess *WebSession, msg *wsMessage){
	fmt.Println("read TextMessage data: ", string(msg.data))
	sess.Write(websocket.TextMessage, []byte("hello,too!"))
}

func BinaryMessageFunc(sess *WebSession, msg *wsMessage){
	var pos int
	msgid := binary.LittleEndian.Uint32(msg.data[pos:])	//消息id
	pos+=4

	datalen := binary.LittleEndian.Uint32(msg.data[pos:]) //消息长度
	pos+=4

	var (
		params = []uint32{}	
	)
	for i := uint32(0); i < datalen; i++ {
		param := binary.LittleEndian.Uint32(msg.data[pos:])
		pos+=4

		params = append(params, param)
	}

	if len(params) <= 0 {
		fmt.Println("invalid params: ", params)
		return
	}
	
	fmt.Println("receive msgid: ", msgid)
	if proc := GetGameLogicProcMsg(int(msgid)); proc != nil {
		err, _ := proc(sess, params)
		if err != nil {
			fmt.Println("proc msg err: ", err)
		}
	}else{
		fmt.Println("invalid msg id: ", msgid)
	}
}

func CloseMessageFunc(sess *WebSession, msg *wsMessage){
	fmt.Println("read CloseMessage.")
	sess.offch <-sess
}

func PingMessageFunc(sess *WebSession, msg *wsMessage){
	fmt.Println("read PingMessage.")
}

func PongMessageFunc(sess *WebSession, msg *wsMessage){
	fmt.Println("read PongMessage.")
}