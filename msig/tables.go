package msig

import yta "github.com/YstarLab/yta-go"

type ProposalRow struct {
	ProposalName       yta.Name              `json:"proposal_name"`
	RequestedApprovals []yta.PermissionLevel `json:"requested_approvals"`
	ProvidedApprovals  []yta.PermissionLevel `json:"provided_approvals"`
	PackedTransaction  yta.HexBytes          `json:"packed_transaction"`
}
