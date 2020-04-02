package pb

type Empty struct {
}

type Response struct {
	Success bool
	Message string
	Code    int32
}
