package api

func (state AccountState) json() string {
	switch state {
	case AccountActive:
		return "active"
	case AccountInactive:
		return "inactive"
	case AccountClosed:
		return "closed"
	}
}

func (user *User) json() string {
	jsonObject := map[string]string{
		"id":           user.id,
		"state":        user.state.json(),
		"created-when": user.createdWhen.json(),
	}
}
