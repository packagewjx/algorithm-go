package contest

const qumo = 1000000007

type node struct {
	parentNode *node
	coin       int
	xiashuCoin int
	child      []*node
	numXiashu  int
}

func addCoinToAll(root *node, coin int) {
	if root == nil {
		return
	}
	root.coin = (root.coin + coin) % qumo
	root.xiashuCoin = ((root.numXiashu*coin)%qumo + root.xiashuCoin) % qumo
	for i := 0; i < len(root.child); i++ {
		addCoinToAll(root.child[i], coin)
	}
}

func addXiashuCoinToParent(n *node, coin int) {
	par := n.parentNode
	for par != nil {
		par.xiashuCoin = (par.xiashuCoin + coin) % qumo
		par = par.parentNode
	}
}

func addCoinToMe(root *node, coin int) {
	root.coin = (root.coin + coin) % qumo
	addXiashuCoinToParent(root, coin)
}

func countNodes(root *node) int {
	if root == nil {
		return 0
	}
	xiashu := 0
	for i := 0; i < len(root.child); i++ {
		xiashu += countNodes(root.child[i])
	}
	root.numXiashu = xiashu
	return xiashu + 1
}

func bonus(n int, leadership [][]int, operations [][]int) []int {
	nodes := make([]*node, n+1)
	isXiashu := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		nodes[i] = &node{
			coin:  0,
			child: make([]*node, 0, 4),
		}
	}
	for i := 0; i < len(leadership); i++ {
		shangsi := nodes[leadership[i][0]]
		xiashu := nodes[leadership[i][1]]
		isXiashu[leadership[i][1]] = true
		shangsi.child = append(shangsi.child, xiashu)
		xiashu.parentNode = shangsi
	}
	// 寻找根节点
	rootId := 0
	for i := 1; i < len(isXiashu); i++ {
		if !isXiashu[i] {
			rootId = i
			break
		}
	}
	// 计算下属数量
	countNodes(nodes[rootId])

	result := make([]int, 0, 4)
	for i := 0; i < len(operations); i++ {
		switch operations[i][0] {
		case 1:
			addCoinToMe(nodes[operations[i][1]], operations[i][2])
		case 2:
			i3 := nodes[operations[i][1]]
			addCoinToAll(i3, operations[i][2])
			addXiashuCoinToParent(i3, i3.coin+i3.xiashuCoin)
		case 3:
			i2 := nodes[operations[i][1]]
			result = append(result, (i2.coin+i2.xiashuCoin)%qumo)
		}
	}

	return result
}
