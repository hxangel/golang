package main

import (
	"fmt";
	"encoding/json"
)

type Tweet struct {
	Text string;
}


func main() {
	s := `{"text" : "\u0422\u0443\u043b\u0430"}`;
	var tw Tweet;
	by := make([]byte,len(s))
	copy(by,s)
	json.Unmarshal(by, &tw);
	fmt.Printf("Escaped: %s\n", tw.Text);

}