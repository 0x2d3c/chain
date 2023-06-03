package sender

import (
	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/bandprotocol/chain/v2/x/tss/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Detail represents the necessary details of a message for logging.
type Detail struct {
	GroupID   tss.GroupID
	SigningID tss.SigningID
	Type      string
}

// GetDetail extracts the details from a slice of SDK messages.
func GetDetail(msgs []sdk.Msg) []Detail {
	details := make([]Detail, 0)
	for _, msg := range msgs {
		var detail Detail
		switch t := msg.(type) {
		case *types.MsgSubmitDKGRound1:
			detail = Detail{
				GroupID: t.GroupID,
				Type:    t.Type(),
			}
		case *types.MsgSubmitDKGRound2:
			detail = Detail{
				GroupID: t.GroupID,
				Type:    t.Type(),
			}
		case *types.MsgConfirm:
			detail = Detail{
				GroupID: t.GroupID,
				Type:    t.Type(),
			}
		case *types.MsgComplain:
			detail = Detail{
				GroupID: t.GroupID,
				Type:    t.Type(),
			}
		case *types.MsgSubmitDEPairs:
			detail = Detail{
				Type: t.Type(),
			}
		case *types.MsgSign:
			detail = Detail{
				SigningID: t.SigningID,
				Type:      t.Type(),
			}
		default:
			detail = Detail{
				Type: "Unknown",
			}
		}

		details = append(details, detail)
	}

	return details
}
