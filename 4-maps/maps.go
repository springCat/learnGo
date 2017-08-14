package maps

import "fmt"

func MapDemo() {
	m := make(map[string]string, 1000)

	//map的插入：
	m["a"] = "a"

	//map的遍历
	for k, v := range m {
		fmt.Println(k, v)
	}

	if v, ok := m["a"]; ok {
		fmt.Println(v, ok)
	}

	//map的删除：
	delete(m, "a")

}
