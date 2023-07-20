package Type

type Id struct {
	MainId   string
	MainType string
	User     string
	Group    string
	Name     string
}

type Data struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
