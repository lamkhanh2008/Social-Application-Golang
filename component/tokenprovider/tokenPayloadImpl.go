package tokenprovider

type TokenPayLoadImpl struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

func (t TokenPayLoadImpl) UserId() int {
	return t.UId
}

func (t TokenPayLoadImpl) Role() string {
	return t.URole
}
