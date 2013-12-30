package main

import (
	"fmt"
	"time"
	"./structures/queue"
	"./structures/stack"
	"./structures/dictionary"
)

func main() {
	// This is used to recover from the error so we
	// can print a clean error message. It is not
	// really needed, using it more to just try out
	// the recover functionality of Go
	defer func() {
        if err := recover(); err != nil {
            out("Caught error: %s\n", err)
        }
    }()
	
	out("Testing queue\n")
	q := queue.InitQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Print()
	q.Dequeue()
	q.Print()
	q.Dequeue()
	q.Print()
	q.Dequeue()
	q.Print()
	
	out("\nTesting stack\n")
	s := stack.InitStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	// Should Print in reverse order (3 2 1) since it is a stack
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()

	out("\nTesting dictionary\n")
	d := dictionary.InitDictionary()
	d.Insert("a", 1)
	d.Insert("b", 2)
	d.PrintKeys()
	out("d.Search('a') = %d\n", d.Search("a"))
	out("d.Count = %d\n", d.Count)
	out("d.Remove('a') = %d\n", d.Remove("a"))
	out("d.Count = %d\n", d.Count)
	//d.Search("2345")

	out("\nTesting stackA\n")
	s1 := stack.InitStackA()
	s1.Push(1)
	s1.Push(2)
	s1.Push(3)
	// Should Print in reverse order (3 2 1) since it is a stack
	s1.Print()
	s1.Pop()
	s1.Print()
	s1.Pop()
	s1.Print()
	s1.Pop()
	s1.Print()


	n := 1000000
	testStack(n)
	testStackA(n)
	// On my laptop the tests were pretty conclusive.
	// The slice based stack (stackA) was about
	// 2x faster than the pointer based stack (stack)
	out("Done testing!\n")
}

func testStack(n int) {
	t0 := time.Now()
	defer displayFuncRunTime("testStack", t0)

	s := stack.InitStack()
	for i:=0;i<n;i++ {
		s.Push(i)
	}
	for i:=0;i<n;i++ {
		s.Pop()
	}
}

func testStackA(n int) {
	t0 := time.Now()
	defer displayFuncRunTime("testStackA", t0)

	s := stack.InitStackA()
	for i:=0;i<n;i++ {
		s.Push(i)
	}
	for i:=0;i<n;i++ {
		s.Pop()
	}
}

func displayFuncRunTime(funcName string, t0 time.Time) {
	t1 := time.Now()
	out("%s executed in %v\n", funcName, t1.Sub(t0))
}

// Shorthand function for fmt.Printf
func out(format string, v...interface{}) {
	fmt.Printf(format, v...)
}