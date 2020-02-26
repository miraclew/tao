package mapper

type Locator struct {
	Module    string
	Resources []Resource
}

type Resource struct {
	Name     string
	HasEvent bool
}
