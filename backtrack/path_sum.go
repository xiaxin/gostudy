package backtrack

// 113. 路径总和 II（中等）
func PathSum113(root *TreeNode, targetSum int) [][]int {
	var (
		result [][]int
		track  []int

		fn func(root *TreeNode, sum int)
	)

	fn = func(root *TreeNode, sum int) {
		if nil == root {
			return
		}
		if root.Left == nil && root.Right == nil && root.Val == sum {
			temp := append([]int{}, track...)
			result = append(result, append(temp, root.Val))
			return
		}

		track = append(track, root.Val)
		if root.Left != nil {
			fn(root.Left, sum-root.Val)
		}

		if root.Right != nil {
			fn(root.Right, sum-root.Val)
		}
		track = track[0 : len(track)-1]
	}

	fn(root, targetSum)
	return result
}
