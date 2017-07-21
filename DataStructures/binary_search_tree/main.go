package main

type node struct {
	left  *node
	right *node
	val   int
}

func insert(head *node, val int) {
	newNode := new(node)
	newNode.val = val
	temp := head

	for {
		if val < temp.val {
			// if val is less than temp-val check if left node is nil. If it is make this left node
			if temp.left == nil {
				temp.left = newNode
				return
			}
			temp = temp.left
		} else if val >= temp.val {
			// if val is greater than temp val check if right node is nil, if it is insert into node
			if temp.right == nil {
				temp.right = newNode
				return
			}
			temp = temp.right
		}

	}
}

func main() {
	var head *node
	head = new(node)
	head.val = 1

	insert(head, 1)
	insert(head, 2)

	for i := 0; i < 100; i++ {
		insert(head, i)
	}

	println(head.val)
	println(head.right.val)
	println(head.right.right.val)
	println(head.right.right.left.val)
	println(head.right.right.right.right.right.val)
	println(head.right.right.right.right.right.right.right.right.right.right.right.val)

}
