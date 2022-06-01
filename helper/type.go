package helper

type ESResultUser struct {
	Hits ESHitsUser `json:"hits"`
}

type ESHitsUser struct {
	Total int `json:"total"`
	Hits  []struct {
		Source ESSourceUser `json:"_source"`
	} `json:"hits"`
}

type ESSourceUser struct {
	FacebookIDFp             string `json:"facebook_id_fp"`
	FanpageID                string `json:"fanpage_id"`
	Project                  string `json:"project"`
	TConversationID          string `json:"t_conversation_id"`
	ConversationID           string `json:"conversation_id"`
	SenderID                 string `json:"sender_id"`
	Gender                   string `json:"gender"`
	Locale                   string `json:"locale"`
	Timezone                 int    `json:"timezone"`
	Avatar                   string `json:"avatar"`
	AvatarOrigin             string `json:"avatar_origin"`
	UpdatedAt                string `json:"updated_at"`
	Name                     string `json:"name"`
	SourceType               string `json:"source_type"`
	CreatedAt                string `json:"created_at"`
	MessageCount             int    `json:"message_count"`
	IDCrm                    string `json:"id_crm"`
	LastFbMessageTime        string `json:"last_fb_message_time"`
	ConversationURL          string `json:"conversation_url"`
	FirstFbMessageTime       string `json:"first_fb_message_time"`
	PushUserElasticCreatedAt string `json:"push_user_elastic_created_at"`
}
