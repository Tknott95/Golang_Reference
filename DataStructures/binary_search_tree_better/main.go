package main

type node struct {
	left   *node
	right  *node
	val    int
	isTree bool
}

func insertBST(head *node, newVal int) {
	newNode := &node{
		val: newVal,
	}

	if head.isTree == false {
		head.val = newVal
		head.isTree = true
		return
	}

	for {
		if newNode.val > head.val {
			if head.right == nil {
				head.right = newNode
				return
			}
			head = head.right
		} else if newNode.val <= head.val {
			if head.left == nil {
				head.left = newNode
				return
			}
			head = head.left
		}
	}
}

func main() {
	head := new(node)

	insertBST(head, 5)
	insertBST(head, 1)
	insertBST(head, 5)
	insertBST(head, 3)
	insertBST(head, 2)
	insertBST(head, 8)
	insertBST(head, 90)
	insertBST(head, 4)

	println(head.right.right.val)
	println(head.left.right.val)

}
