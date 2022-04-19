package main

import "fmt"

type Node struct {
    Val   int
    Left  *Node
    Right *Node
}

// 非递归版
func Insert1(root **Node, val int) bool {
    if *root == nil {
        node := &Node{
            Val:   val,
            Left:  nil,
            Right: nil,
        }
        *root = node
        return true
    }
    x := *root
    var y *Node
    for x != nil {
        y = x
        if val == x.Val {
            return false
        } else if val < x.Val {
            x = x.Left
        } else {
            x = x.Right
        }
    }
    node := &Node{
        Val:   val,
        Left:  nil,
        Right: nil,
    }
    if val < y.Val {
        y.Left = node
    } else {
        y.Right = node
    }
    return true
}

// 递归版
func Insert2(root **Node, val int) bool {
    if *root == nil {
        node := &Node{
            Val:   val,
            Left:  nil,
            Right: nil,
        }
        *root = node
        return true
    }
    if val < (*root).Val {
        return Insert2(&(*root).Left, val)
    }
    return Insert2(&(*root).Right, val)
}

func CreateBinarySearchTree(src []int) *Node {
    if len(src) == 0 {
        return nil
    }
    var root *Node
    for _, v := range src {
        //Insert1(&root, v)
        Insert2(&root, v)
    }
    return root
}

func Inorder(root *Node) {
    if root == nil {
        return
    }
    Inorder(root.Left)
    fmt.Println(root.Val)
    Inorder(root.Right)
}

// Search1 查找 非递归
func Search1(root *Node, val int) bool {
    if root == nil {
        return false
    }
    for root != nil {
        if root.Val == val {
            return true
        } else if val < root.Val {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return false
}

// Search2 查找 递归版
func Search2(root *Node, val int) bool {
    if root == nil {
        return false
    }
    if root.Val == val {
        return true
    } else if val < root.Val {
        return Search2(root.Left, val)
    } else {
        return Search2(root.Right, val)
    }
}

func Delete(root *Node, val int) *Node {
    if root == nil {
        return nil
    }
    if val == root.Val {
        // 1. 叶子节点删除
        if root.Right == nil && root.Left == nil {
            root = nil
            return nil
        }
        // 2. 只有左子树或者只有右子树
        if root.Right == nil {
            root = root.Left
            return root
        }
        if root.Left == nil {
            root = root.Right
            return root
        }
        // 3. 左右子树都存在
        // 找到当前节点root的左节点的最右节点
        q := root.Left
        for q.Right != nil {
            q = q.Right
        }
        root.Val = q.Val
        root.Left = Delete(root.Left, root.Val)
        return root
    } else if val < root.Val {
        // 如果小于，则在左子树中删除
        root.Left = Delete(root.Left, val)
        return root
    } else {
        // 如果大于，则在右子树中删除
        root.Right = Delete(root.Right, val)
        return root
    }
}

func remove(node *Node, key int) *Node {
    if node == nil {
        return nil
    }
    if key < node.Val {
        node.Left = remove(node.Left, key)
        return node
    }
    if key > node.Val {
        node.Right = remove(node.Right, key)
        return node
    }
    // key == node.key
    if node.Left == nil && node.Right == nil {
        node = nil
        return nil
    }
    if node.Left == nil {
        node = node.Right
        return node
    }
    if node.Right == nil {
        node = node.Left
        return node
    }
    /*minRightNode := node.Right
      for  minRightNode.Left != nil {
          //find smallest value on the right side
          minRightNode = minRightNode.Left
      }
      node.Val = minRightNode.Val
      node.Right = remove(node.Right, node.Val)*/
    maxLeftNode := node.Left
    for maxLeftNode.Right != nil {
        maxLeftNode = maxLeftNode.Right
    }
    node.Val = maxLeftNode.Val
    node.Left = remove(node.Left, node.Val)

    return node
}
