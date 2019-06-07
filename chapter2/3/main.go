package main

import "fmt"

type Element struct {
	data int
	next *Element
}

func deleteAmongElements(e *Element) {
	next := e.next

	e.data = next.data
	e.next = next.next
}

func main() {
	var head = &Element{data: -1}

	var cur *Element = head
	var cur3 *Element
	for i := 0; i < 10; i++ {
		b := &Element{data: i}
		cur.next = b
		if i == 3 {
			cur3 = cur
		}
		cur = b
	}

	{
		deleteAmongElements(cur3)

		c := head
		for {
			if c.next == nil {
				break
			}
			fmt.Println(c)
			c = c.next
		}
	}
}
