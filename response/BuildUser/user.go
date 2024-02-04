package BuildUser

type BuildUser struct {
	Token  string `json:"token"`
	UserId uint   `json:"user_id"`
}

// ResponseBuildUser 给用户返回值
func ResponseBuildUser(token string, userId uint) BuildUser {
	return BuildUser{
		Token:  token,
		UserId: userId,
	}
}
