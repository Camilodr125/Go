package main
import ("reflect" 
		"fmt")
type contact struct {
	userID       string
	sendingLimit int32
	age          int32

}

type perms struct {
	canSend         bool
	canReceive      bool
	canManage       bool
	permissionLevel int
	
}

func main()  {
	typ := reflect.TypeOf(contact{})
fmt.Printf("Struct is %d bytes\n", typ.Size())
}

