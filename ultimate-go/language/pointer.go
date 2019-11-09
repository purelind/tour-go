package main

type user struct {
	name	string
	email	string
}

func main() {
	// 声明一个变量值为10，该值存放在stack中
	count := 10

	// 获取值的地址
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// 传入count的值
	increment1(count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	increment2(&count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	stayOnStack()
	escapeToHeap()
}

func increment1(inc int) {
	inc++
	println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

// the * here is not an operator, it is part of the type name
func increment2(inc *int) {
	// the * is an operator, it tells us the value of the pointer points to
	*inc++
	println("inc2:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}

func stayOnStack() user {
	u := user{
		name:	"Hoanh An",
		email:	"hoanhan@bennington.edu",
	}
	return u
}

func escapeToHeap() *user {
	u := user{
		name:	"Hoanh An",
		email:	"hoanhan@benington.edu",
	}

	return &u
}


