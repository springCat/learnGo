package exception

import "fmt"

func test()  {
	defer func(){
		if err := recover();err != nil {
			fmt.Println(err.(string))
		}
	}()
	panic("error")
}
