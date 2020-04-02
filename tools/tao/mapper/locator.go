package mapper

type Locator struct {
	Module    string // module of the locator
	Resources []Resource
}

type Resource struct {
	Module   string
	Pkg      string // pkg name
	Name     string // Title camel case, e.g. ContentReport
	HasEvent bool
}
