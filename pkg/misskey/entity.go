package misskey

import (
	"fmt"
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

func (mr *MisskeyRequest) BuildTweetText(domain string) string {
	return fmt.Sprintf("%s %s/notes/%s", mr.Body.Note.Text, domain, mr.Body.Note.ID)
}

func (mr *MisskeyRequest) GetFileUrls() []string {
	ret := []string{}
	for _, v := range mr.Body.Note.Files {
		ret = append(ret, v.URL)
	}
	return ret
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
	FileIds  []any         `json:"fileIds"`
	Files    []MisskeyFile `json:"files"`
	ReplyID  any           `json:"replyId"`
	RenoteID any           `json:"renoteId"`
}

type MisskeyFile struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Md5         string    `json:"md5"`
	Size        int       `json:"size"`
	IsSensitive bool      `json:"isSensitive"`
	Blurhash    string    `json:"blurhash"`
	Properties  struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"properties"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
	Comment      any    `json:"comment"`
	FolderID     any    `json:"folderId"`
	Folder       any    `json:"folder"`
	UserID       any    `json:"userId"`
	User         any    `json:"user"`
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
