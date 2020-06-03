package pb

type Empty struct {
}

type Response struct {
	Success bool
	Message string
	Code    int32
}

type UserAvatar struct {
	Id       int64
	Nickname string
	Avatar   string
}
