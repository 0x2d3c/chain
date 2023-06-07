package types

var MsgGrants = []string{
	"/tss.v1beta1.MsgCreateGroup",
	"/tss.v1beta1.MsgSubmitDKGRound1",
	"/tss.v1beta1.MsgSubmitDKGRound2",
	"/tss.v1beta1.MsgComplain",
	"/tss.v1beta1.MsgConfirm",
	"/tss.v1beta1.MsgSubmitDEs",
	"/tss.v1beta1.MsgSign",
}

const (
	AddrLen   = 20
	uint64Len = 8
)
