package gomatrix

// RespError is the standard JSON error response from Homeservers. It also implements the Golang "error" interface.
// See http://matrix.org/docs/spec/client_server/r0.2.0.html#api-standards
type RespError struct {
	ErrCode string `json:"errcode"`
	Err     string `json:"error"`
}

// Error returns the errcode and error message.
func (e RespError) Error() string {
	return e.ErrCode + ": " + e.Err
}

// RespCreateFilter is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-user-userid-filter
type RespCreateFilter struct {
	FilterID string `json:"filter_id"`
}

// RespVersions is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-versions
type RespVersions struct {
	Versions []string `json:"versions"`
}

// RespJoinRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-join
type RespJoinRoom struct {
	RoomID string `json:"room_id"`
}

// RespLeaveRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-leave
type RespLeaveRoom struct{}

// RespForgetRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-forget
type RespForgetRoom struct{}

// RespInviteUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-invite
type RespInviteUser struct{}

// RespKickUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-kick
type RespKickUser struct{}

// RespBanUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-ban
type RespBanUser struct{}

// RespUnbanUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-unban
type RespUnbanUser struct{}

// RespTyping is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#put-matrix-client-r0-rooms-roomid-typing-userid
type RespTyping struct{}

// RespJoinedRooms is the JSON response for TODO-SPEC https://github.com/matrix-org/synapse/pull/1680
type RespJoinedRooms struct {
	JoinedRooms []string `json:"joined_rooms"`
}

// RespJoinedMembers is the JSON response for TODO-SPEC https://github.com/matrix-org/synapse/pull/1680
type RespJoinedMembers struct {
	Joined map[string]struct {
		DisplayName *string `json:"display_name"`
		AvatarURL   *string `json:"avatar_url"`
	} `json:"joined"`
}

// RespMessages is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-r0-rooms-roomid-messages
type RespMessages struct {
	Start string  `json:"start"`
	Chunk []Event `json:"chunk"`
	End   string  `json:"end"`
}

// RespSendEvent is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#put-matrix-client-r0-rooms-roomid-send-eventtype-txnid
type RespSendEvent struct {
	EventID string `json:"event_id"`
}

// RespMediaUpload is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-media-r0-upload
type RespMediaUpload struct {
	ContentURI string `json:"content_uri"`
}

// RespUserInteractive is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#user-interactive-authentication-api
type RespUserInteractive struct {
	Flows []struct {
		Stages []string `json:"stages"`
	} `json:"flows"`
	Params    map[string]interface{} `json:"params"`
	Session   string                 `json:"string"`
	Completed []string               `json:"completed"`
	ErrCode   string                 `json:"errcode"`
	Error     string                 `json:"error"`
}

// HasSingleStageFlow returns true if there exists at least 1 Flow with a single stage of stageName.
func (r RespUserInteractive) HasSingleStageFlow(stageName string) bool {
	for _, f := range r.Flows {
		if len(f.Stages) == 1 && f.Stages[0] == stageName {
			return true
		}
	}
	return false
}

// RespRegister is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-register
type RespRegister struct {
	AccessToken  string `json:"access_token"`
	DeviceID     string `json:"device_id"`
	HomeServer   string `json:"home_server"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}

// RespLogin is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-login
type RespLogin struct {
	AccessToken string `json:"access_token"`
	DeviceID    string `json:"device_id"`
	HomeServer  string `json:"home_server"`
	UserID      string `json:"user_id"`
}

// RespLogout is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-logout
type RespLogout struct{}

// RespCreateRoom is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-createroom
type RespCreateRoom struct {
	RoomID string `json:"room_id"`
}

// RespSync is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-r0-sync
type RespSync struct {
	NextBatch   string `json:"next_batch"`
	AccountData struct {
		Events []Event `json:"events"`
	} `json:"account_data"`
	Presence struct {
		Events []Event `json:"events"`
	} `json:"presence"`
	Rooms struct {
		Leave map[string]struct {
			State struct {
				Events []Event `json:"events"`
			} `json:"state"`
			Timeline struct {
				Events    []Event `json:"events"`
				Limited   bool    `json:"limited"`
				PrevBatch string  `json:"prev_batch"`
			} `json:"timeline"`
		} `json:"leave"`
		Join map[string]struct {
			State struct {
				Events []Event `json:"events"`
			} `json:"state"`
			Timeline struct {
				Events    []Event `json:"events"`
				Limited   bool    `json:"limited"`
				PrevBatch string  `json:"prev_batch"`
			} `json:"timeline"`
		} `json:"join"`
		Invite map[string]struct {
			State struct {
				Events []Event
			} `json:"invite_state"`
		} `json:"invite"`
	} `json:"rooms"`
}

// RespKeysUpload is the JSON response for https://matrix.org/speculator/spec/drafts%2Fe2e/client_server/unstable.html#post-matrix-client-unstable-keys-upload
type RespKeysUpload struct {
}

// RespKeysQuery is the JSON response for https://matrix.org/speculator/spec/drafts%2Fe2e/client_server/unstable.html#post-matrix-client-unstable-keys-query
type RespKeysQuery struct {
	Failures   map[string]interface{} `json:"failures"`
	DeviceKeys map[string]map[string]struct {
		UserID     string                       `json:"user_id"`
		DeviceID   string                       `json:"device_id"`
		Algorithms []string                     `json:"algorithms"`
		Keys       map[string]string            `json:"keys"`
		Signatures map[string]map[string]string `json:"signatures"`
		Unsigned   struct {
			DeviceDisplayName string `json:"device_display_name"`
		} `json:"unsigned"`
	} `json:"device_keys"`
}

// RespKeysClaim is the JSON response for https://matrix.org/speculator/spec/drafts%2Fe2e/client_server/unstable.html#post-matrix-client-unstable-keys-claim
type RespKeysClaim struct {
	Failures    map[string]interface{}                       `json:"failures"`
	OneTimeKeys map[string]map[string]map[string]interface{} `json:"one_time_keys"`
}

// RespKeysChanges is the JSON response for https://matrix.org/speculator/spec/drafts%2Fe2e/client_server/unstable.html#get-matrix-client-unstable-keys-changes
type RespKeysChanges struct {
}
