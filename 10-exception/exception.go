package exception

import "fmt"

func test()  {
	defer func(){
		if err := recover();err != nil {
			fmt.Println(err)
		}
	}()
	panic("error")
}


func test1()  {
		panic("error")
}
