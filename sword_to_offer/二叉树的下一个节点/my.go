package next_node

type Trie struct {
	val    string
	parent *Trie
	left   *Trie
	right  *Trie
}

func (t *Trie) InOrderList() (list []string) {
	if t.left != nil {
		list = append(list, t.left.InOrderList()...)
	}
	list = append(list, t.val)
	if t.right != nil {
		list = append(list, t.right.InOrderList()...)
	}
	return list
}

func (t *Trie) First() string {
	if t.left != nil {
		return t.left.First()
	}
	return t.val
}

func (t *Trie) LeftNode() string {
	if t.parent.parent == nil {
		return ""
	}

	if t.parent.parent.left == t.parent {
		return t.parent.parent.val
	}
	return t.parent.LeftNode()
}

func my(in *Trie) string {
	//如果节点的右树不为空，则寻找右树的第一节点
	if in.right != nil {
		return in.right.First()
	}
	//如果节点右树为空，且是一个左节点，那么下一步就是打印它的父节点
	if in.parent != nil && in == in.parent.left {
		return in.parent.val
	}
	//如果节点右树为空，且是一个右节点，寻找它父节点的左树先辈
	if in.parent != nil && in == in.parent.right {
		return in.LeftNode()
	}
	return ""
}
