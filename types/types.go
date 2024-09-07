package types

const (
	Django int = iota
	Express
	Nextjs
	Go
	Flask
	Springboot
)

type MessageStruct struct {
	Role    string   `json:"role"`
	Content []string `json:"content"`
}
