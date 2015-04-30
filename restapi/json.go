package api

func jsonAccountState(jsonString string) AccountState {
	var state string
	json.Unmarshal(jsonString, &state)

	switch state {
	case "active":
		return AccountActive
	case "inactive":
		return AccountInactive
	case "closed":
		return AccountClosed
	}
}

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

func jsonUser() User {
}

func (user *User) json() string {
	jsonObject := map[string]string{
		"id":           user.id,
		"state":        user.state.json(),
		"created-when": user.createdWhen.json(),
	}
}
