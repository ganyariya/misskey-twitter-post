package misskey

import (
	"time"
)

type MisskeyRequest struct {
	HookId    string             `json:"hookId"`
	UserId    string             `json:"userId"`
	EventId   string             `json:"eventId"`
	CreatedAt int                `json:"createdAt"`
	Type      string             `json:"type"`
	Body      MisskeyRequestBody `json:"body"`
}

type MisskeyRequestBody struct {
	Note MisskeyNote `json:"note"`
}

type MisskeyNote struct {
	ID           string      `json:"id"`
	CreatedAt    time.Time   `json:"createdAt"`
	UserID       string      `json:"userId"`
	User         MisskeyUser `json:"user"`
	Text         string      `json:"text"`
	Cw           any         `json:"cw"`
	Visibility   string      `json:"visibility"`
	LocalOnly    bool        `json:"localOnly"`
	RenoteCount  int         `json:"renoteCount"`
	RepliesCount int         `json:"repliesCount"`
	Reactions    struct {
	} `json:"reactions"`
	ReactionEmojis struct {
	} `json:"reactionEmojis"`
	FileIds  []any `json:"fileIds"`
	Files    []any `json:"files"`
	ReplyID  any   `json:"replyId"`
	RenoteID any   `json:"renoteId"`
}

type MisskeyUser struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	Host           any    `json:"host"`
	AvatarURL      string `json:"avatarUrl"`
	AvatarBlurhash string `json:"avatarBlurhash"`
	IsBot          bool   `json:"isBot"`
	IsCat          bool   `json:"isCat"`
	Emojis         struct {
	} `json:"emojis"`
	OnlineStatus string `json:"onlineStatus"`
	BadgeRoles   []any  `json:"badgeRoles"`
}
