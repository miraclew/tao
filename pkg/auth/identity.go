package auth

type Identity struct {
	UserID   int64
	Roles    []string
	DeviceID string
	App      string
	Internal string // internal source
}
