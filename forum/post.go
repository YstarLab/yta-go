package forum

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewPost is an action representing a simple message to be posted
// through the chain network.
func NewPost(poster yta.AccountName, postUUID, content string, replyToPoster yta.AccountName, replyToPostUUID string, certify bool, jsonMetadata string) *yta.Action {
	a := &yta.Action{
		Account: ForumAN,
		Name:    ActN("post"),
		Authorization: []yta.PermissionLevel{
			{Actor: poster, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Post{
			Poster:          poster,
			PostUUID:        postUUID,
			Content:         content,
			ReplyToPoster:   replyToPoster,
			ReplyToPostUUID: replyToPostUUID,
			Certify:         certify,
			JSONMetadata:    jsonMetadata,
		}),
	}
	return a
}

// Post represents the `eosio.forum::post` action.
type Post struct {
	Poster          yta.AccountName `json:"poster"`
	PostUUID        string          `json:"post_uuid"`
	Content         string          `json:"content"`
	ReplyToPoster   yta.AccountName `json:"reply_to_poster"`
	ReplyToPostUUID string          `json:"reply_to_post_uuid"`
	Certify         bool            `json:"certify"`
	JSONMetadata    string          `json:"json_metadata"`
}
