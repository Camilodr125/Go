package main
import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	users:= make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	for i := 0; i < len(names); i++ {
		users[names[i]]= user{names[i], phoneNumbers[i]}
	}

	return users, nil
	
}
type user struct {
	name        string
	phoneNumber int
}