package ast

func FindFirstParentOfType[T Node](node Node) T {
	var null T
	n := node
	for {
		if n == nil {
			return null
		}

		if v, ok := n.(T); ok {
			return v
		}

		n = n.GetParent()
	}
}
