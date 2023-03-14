package generatetoken

type Payload struct {
	UserName string `json:"user_name"`
}

func (s *Payload) GetUserName() string {
	return s.UserName
}
