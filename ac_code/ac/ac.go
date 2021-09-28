package ac

type AC struct {
    Trie        *Trie
    FailurePath map[uint32]*Node
}

func NewAC(trie *Trie) *AC {
    return &AC{
        Trie:        trie,
        FailurePath: make(map[uint32]*Node),
    }
}

// BuildFailurePath 创建失败路径
func (ac *AC) BuildFailurePath() {
    for node := range ac.Trie.bfs() {
        if node.IsRootNode() {
            ac.FailurePath[node.ID] = nil
        }
        parent := node.Parent
        if parent.IsRootNode() {
            ac.FailurePath[node.ID] = ac.Trie.Root
        }
        var temp = ac.getFailurePath(parent)
        for ac.FailurePath[node.ID] == nil{
            if temp.IsRootNode() {
                ac.FailurePath[node.ID] = ac.Trie.Root
            }
            if v,ok := temp.Children[node.Character];ok{
                ac.FailurePath[node.ID] = v
            } else {
                temp = ac.getFailurePath(temp)
            }
        }
    }
}

func (ac *AC) getFailurePath(n *Node) *Node {
    if n == nil {
        return nil
    }
    if v, ok := ac.FailurePath[n.ID]; ok {
        return v
    }
    return nil
}

func (ac *AC) next(node *Node, c rune) *Node {
    next, ok := node.Children[c]
    if ok {
        return next
    }
    return nil
}

func (ac *AC) fail(node *Node, c rune) *Node {
    var next *Node
    for {
        next = ac.next(ac.getFailurePath(node), c)
        if next == nil {
            if node.IsRootNode() {
                return node
            }
            node = ac.getFailurePath(node)
            continue
        }
        return next
    }
}

func (ac *AC) output(node *Node, runes []rune, position int, results []string) []string {
    if node.IsRootNode() {
        return results
    }

    if node.IsPathEnd() {
        results = append(results, string(runes[position+1-node.depth:position+1]))
    }

    return ac.output(ac.getFailurePath(node), runes, position, results)
}

func (ac *AC) Find(text string) []string {
    var (
        node  = ac.Trie.Root
        next  *Node
        runes = []rune(text)
    )
    res := make([]string, 0)
    for position := 0; position < len(runes); position++ {
        next = ac.next(node, runes[position])
        if next == nil {
            next = ac.fail(node, runes[position])
        }

        node = next
        res = ac.output(node, runes, position, res)
    }
    return res
}
