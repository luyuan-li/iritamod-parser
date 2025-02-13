package node

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritamod-parser/modules"
)

type NodeClient struct {
}

func NewClient() NodeClient {
	return NodeClient{}
}

func (nft NodeClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgNodeCreate:
		docMsg := DocMsgCreateValidator{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgNodeUpdate:
		docMsg := DocMsgUpdateValidator{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgNodeRemove:
		docMsg := DocMsgRevokeNode{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgNodeGrant:
		docMsg := DocMsgGrantNode{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgNodeRevoke:
		docMsg := DocMsgRevokeNode{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
