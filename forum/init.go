package forum

import yta "github.com/YstarLab/yta-go"

func init() {
	yta.RegisterAction(ForumAN, ActN("clnproposal"), CleanProposal{})
	yta.RegisterAction(ForumAN, ActN("expire"), Expire{})
	yta.RegisterAction(ForumAN, ActN("post"), Post{})
	yta.RegisterAction(ForumAN, ActN("propose"), Propose{})
	yta.RegisterAction(ForumAN, ActN("status"), Status{})
	yta.RegisterAction(ForumAN, ActN("unpost"), UnPost{})
	yta.RegisterAction(ForumAN, ActN("unvote"), UnVote{})
	yta.RegisterAction(ForumAN, ActN("vote"), Vote{})
}

var AN = yta.AN
var PN = yta.PN
var ActN = yta.ActN

var ForumAN = AN("eosio.forum")
