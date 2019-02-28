// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: /usr/share/vpp/api/stn.api.json

/*
 Package stn is a generated from VPP binary API module 'stn'.

 It contains following objects:
	  4 messages
	  2 services

*/
package stn

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// Services represents VPP binary API services:
//
//	"services": {
//	    "stn_add_del_rule": {
//	        "reply": "stn_add_del_rule_reply"
//	    },
//	    "stn_rules_dump": {
//	        "reply": "stn_rules_details",
//	        "stream": true
//	    }
//	},
//
type Services interface {
	DumpStnRules(*StnRulesDump) ([]*StnRulesDetails, error)
	StnAddDelRule(*StnAddDelRule) (*StnAddDelRuleReply, error)
}

/* Messages */

// StnAddDelRule represents VPP binary API message 'stn_add_del_rule':
//
//	"stn_add_del_rule",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u8",
//	    "is_ip4"
//	],
//	[
//	    "u8",
//	    "ip_address",
//	    16
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	[
//	    "u8",
//	    "is_add"
//	],
//	{
//	    "crc": "0x9f0bbe21"
//	}
//
type StnAddDelRule struct {
	IsIP4     uint8
	IPAddress []byte `struc:"[16]byte"`
	SwIfIndex uint32
	IsAdd     uint8
}

func (*StnAddDelRule) GetMessageName() string {
	return "stn_add_del_rule"
}
func (*StnAddDelRule) GetCrcString() string {
	return "9f0bbe21"
}
func (*StnAddDelRule) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// StnAddDelRuleReply represents VPP binary API message 'stn_add_del_rule_reply':
//
//	"stn_add_del_rule_reply",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "i32",
//	    "retval"
//	],
//	{
//	    "crc": "0xe8d4e804"
//	}
//
type StnAddDelRuleReply struct {
	Retval int32
}

func (*StnAddDelRuleReply) GetMessageName() string {
	return "stn_add_del_rule_reply"
}
func (*StnAddDelRuleReply) GetCrcString() string {
	return "e8d4e804"
}
func (*StnAddDelRuleReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// StnRulesDump represents VPP binary API message 'stn_rules_dump':
//
//	"stn_rules_dump",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "client_index"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	{
//	    "crc": "0x51077d14"
//	}
//
type StnRulesDump struct{}

func (*StnRulesDump) GetMessageName() string {
	return "stn_rules_dump"
}
func (*StnRulesDump) GetCrcString() string {
	return "51077d14"
}
func (*StnRulesDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// StnRulesDetails represents VPP binary API message 'stn_rules_details':
//
//	"stn_rules_details",
//	[
//	    "u16",
//	    "_vl_msg_id"
//	],
//	[
//	    "u32",
//	    "context"
//	],
//	[
//	    "u8",
//	    "is_ip4"
//	],
//	[
//	    "u8",
//	    "ip_address",
//	    16
//	],
//	[
//	    "u32",
//	    "sw_if_index"
//	],
//	{
//	    "crc": "0x5eafa31e"
//	}
//
type StnRulesDetails struct {
	IsIP4     uint8
	IPAddress []byte `struc:"[16]byte"`
	SwIfIndex uint32
}

func (*StnRulesDetails) GetMessageName() string {
	return "stn_rules_details"
}
func (*StnRulesDetails) GetCrcString() string {
	return "5eafa31e"
}
func (*StnRulesDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*StnAddDelRule)(nil), "stn.StnAddDelRule")
	api.RegisterMessage((*StnAddDelRuleReply)(nil), "stn.StnAddDelRuleReply")
	api.RegisterMessage((*StnRulesDump)(nil), "stn.StnRulesDump")
	api.RegisterMessage((*StnRulesDetails)(nil), "stn.StnRulesDetails")
}
