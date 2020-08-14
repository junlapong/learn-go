package main

/*

https://gist.github.com/moraes/2141121#gistcomment-2933277

*/

import (
	"container/list"
	"fmt"
)

func main() {
	// Stack: LIFO - last in first out, Stack use list
	stack := list.New()
	for i := 0; i < 10; i++ {
		stack.PushBack(i)
	}
	//
	for stack.Back() != nil {
		fmt.Println(stack.Back().Value)
		stack.Remove(stack.Back())
	}

	// Queue: FIFO - Fist In Fist Out,queue use list
	queue := list.New()
	for i := 0; i < 10; i++ {
		queue.PushBack(i)
	}

	for queue.Front() != nil {
		fmt.Println(queue.Front().Value)
		queue.Remove(queue.Front())
	}
}
