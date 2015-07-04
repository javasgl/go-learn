package main 
import(
	"fmt"
	"encoding/json"
)
type Msg struct{
	code int
	data string
	time int64
}
func main(){
	var m Msg
	m.code = 1 
	m.data = "json data"
	m.time = 12919212212
	b,err := json.Marshal(m)
	if err!=nil {
		fmt.Println(b);	
	}else{
		fmt.Println(err);
	}
}