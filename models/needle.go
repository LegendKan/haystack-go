package models

import(
	"fmt"
)

type Needle struct {
	Header 		int32	//Magic number
	Cookie 		int32	//Random number to mitigate brute force lookups
	Key	   		int64	//photo id
	AlternateKey 	int32
	Flags  		int8
	Size   		int64	//Data size
	Data   		[]byte
	Footer 		int32
	CheckSum 	string
	Padding 	[]byte
}

func init() {
	fmt.Println("Needle Init()")
}
