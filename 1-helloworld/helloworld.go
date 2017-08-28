package main
import (
		"errors"
		"fmt"
)

func Division(a, b float64) (float64, error) {
		if b == 0 {
				return 0, errors.New("除数不能为0")
		}

		return a / b, nil
}

func main()  {
		i,err:= Division(9,3)
		if err == nil {
				fmt.Println(i)
		} else{
				fmt.Println(err)
		}
}