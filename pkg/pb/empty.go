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

type LimitOffset struct {
	Offset int64
	Limit  int32
}

type QueryRequest struct {
	Offset int64
	Limit  int32
	Filter map[string]interface{}
	Sort   string
}

type GetRequest struct {
	Id     int64
	Filter map[string]interface{}
}

type UpdateRequest struct {
	Id     int64
	Values map[string]interface{}
}
