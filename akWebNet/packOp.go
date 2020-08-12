package akWebNet

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

var (
	msgOps = map[PACK_TYPE]*ProtoBufMsgOp{
		PACK_PROTO: &ProtoBufMsgOp{},
	}
)

func PackMsgOp(mainid, subid uint16, msg proto.Message, pt PACK_TYPE) (out []byte, err error) {
	op, exist := msgOps[pt]
	if !exist {
		err = fmt.Errorf("can not find msg op, src info[mainid: %v, subid: %v, pt: %v].", mainid, subid, pt)
		return
	}
	out, err = op.Pack(mainid, subid, msg)
	return
}

func UnPackMsgOp(data []byte, pt PACK_TYPE) (msg proto.Message, cb reflect.Value, err error) {
	op, exist := msgOps[pt]
	if !exist {
		err = fmt.Errorf("can not find msg op, src info[pt: %v].", pt)
		return
	}
	msg, cb, err = op.UnPack(data)
	return
}
