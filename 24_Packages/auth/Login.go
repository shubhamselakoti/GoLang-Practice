package auth

func Login(email string, pass string) string {
	if matchPass(pass) {
		return "Logged In"
	}
	return "Failed"
}

func matchPass(pass string) bool {
	if pass == "Choti@123" {
		return true
	}
	return false
}
