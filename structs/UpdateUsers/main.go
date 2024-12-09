package main

type Membership struct {
	Type string
	MessageCharLimit int

}
type User struct {
	Membership
	Name string
	
}


func newUser(name string, membershipType string) User {
	if membershipType == "premium" {
		userNew:= User{
			Name: name,
			Membership: Membership{
				Type: membershipType,
				MessageCharLimit: 1000,
			},
		}
		return userNew
	} else {
		userNew:= User{
			Name: name,
			Membership: Membership{
				Type: membershipType,
				MessageCharLimit: 100,
			},
		}
	
	return userNew

	
}

}
