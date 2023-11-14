package util

const (
	DEFAULT_SESSION_EXPIRE  = 3600
	SESSION_KEY_LOGIN_ADMIN = "login_admin"
)

var sessionExpire = map[string]int{
	SESSION_KEY_LOGIN_ADMIN: 3600,
}

func GetSessionExpire(sessionKey string) int {
	expire, ok := sessionExpire[sessionKey]
	if ok {
		return expire
	}
	return DEFAULT_SESSION_EXPIRE
}
