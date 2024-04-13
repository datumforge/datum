package webhook

import "regexp"

type Config struct {
	OutgoingWebhook         OutgoingWebhook         `json:"outgoing_webhook"`
	OutgoingWebhookResponse OutgoingWebhookResponse `json:"outgoing_webhook_response"`
	OutgoingWebhookPayload  OutgoingWebhookPayload  `json:"outgoing_webhook_payload"`
	SlackAttachment         SlackAttachment         `json:"slack_attachment"`
	SlackAttachmentField    SlackAttachmentField    `json:"slack_attachment_field"`
}

type OutgoingWebhook struct {
	Id          string `json:"id"`
	Token       string `json:"token"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
	DeleteAt    int64  `json:"delete_at"`
	CreatorId   string `json:"creator_id"`
	ChannelId   string `json:"channel_id"`
	TeamId      string `json:"team_id"`
	TriggerWhen int    `json:"trigger_when"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	ContentType string `json:"content_type"`
	Username    string `json:"username"`
	IconURL     string `json:"icon_url"`
}

type OutgoingWebhookPayload struct {
	Token       string `json:"token"`
	TeamId      string `json:"team_id"`
	TeamDomain  string `json:"team_domain"`
	ChannelId   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Timestamp   int64  `json:"timestamp"`
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	PostId      string `json:"post_id"`
	Text        string `json:"text"`
	TriggerWord string `json:"trigger_word"`
	FileIds     string `json:"file_ids"`
}

type OutgoingWebhookResponse struct {
	Text         *string         `json:"text"`
	Username     string          `json:"username"`
	IconURL      string          `json:"icon_url"`
	Attachments  SlackAttachment `json:"attachments"`
	Type         string          `json:"type"`
	ResponseType string          `json:"response_type"`
}

var linkWithTextRegex = regexp.MustCompile(`<([^<\|]+)\|([^>]+)>`)

type SlackAttachment struct {
	Id         int64                `json:"id"`
	Fallback   string               `json:"fallback"`
	Color      string               `json:"color"`
	Pretext    string               `json:"pretext"`
	AuthorName string               `json:"author_name"`
	AuthorLink string               `json:"author_link"`
	AuthorIcon string               `json:"author_icon"`
	Title      string               `json:"title"`
	TitleLink  string               `json:"title_link"`
	Text       string               `json:"text"`
	Fields     SlackAttachmentField `json:"fields"`
	ImageURL   string               `json:"image_url"`
	ThumbURL   string               `json:"thumb_url"`
	Footer     string               `json:"footer"`
	FooterIcon string               `json:"footer_icon"`
	Timestamp  any                  `json:"ts"` // This is either a string or an int64
}

type SlackAttachmentField struct {
	Title string `json:"title"`
	Value any    `json:"value"`
}
