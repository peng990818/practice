package tree

import (
    "fmt"
)

const (
    RED   = 0
    BLACK = 1
)

type RBT struct {
    nValue  int
    nColor  int
    pLeft   *RBT
    pRight  *RBT
    pFather *RBT
}

type RBTree struct {
    Root *RBT
}

// RightRotate 右旋
func (r *RBTree) RightRotate(pTree *RBT) {
    // 1. 校验
    if pTree == nil || pTree.pLeft == nil {
        return
    }

    pNode := pTree
    pMark := pNode.pLeft

    // 2. 三个孩子
    pNode.pLeft = pMark.pRight // A 的左为E
    pMark.pRight = pNode
    if pNode.pFather != nil {
        if pNode == pNode.pFather.pLeft {
            pNode.pFather.pLeft = pNode
        } else {
            pNode.pFather.pRight = pNode
        }
    } else {
        // 根
        r.Root = pMark
    }

    // 3. 三个父亲
    if pNode.pLeft != nil {
        pNode.pLeft.pFather = pNode
    }
    pMark.pFather = pNode.pFather
    pNode.pFather = pMark
}

// LeftRotate 左旋
func (r *RBTree) LeftRotate(pTree *RBT) {
    // 1. 校验
    if pTree == nil || pTree.pRight == nil {
        return
    }

    pNode := pTree
    pMark := pNode.pRight

    // 2. 三个孩子
    pNode.pRight = pMark.pLeft
    pMark.pLeft = pNode
    if pNode.pFather != nil {
        if pNode == pNode.pFather.pLeft {
            pNode.pFather.pLeft = pMark
        } else {
            pNode.pFather.pRight = pMark
        }
    } else {
        // 根
        r.Root = pMark
    }

    // 3. 三个父亲
    if pNode.pRight != nil {
        pNode.pRight.pFather = pNode
    }
    pMark.pFather = pNode.pFather
    pNode.pFather = pMark
}

// Search 查找插入节点的位置
func (r *RBTree) Search(pTree *RBT, nNum int) *RBT {
    if pTree == nil {
        return nil
    }
    for pTree != nil {
        if pTree.nValue > nNum {
            // 左侧
            if pTree.pLeft == nil {
                return pTree
            }
            pTree = pTree.pLeft
        } else if pTree.nValue < nNum {
            // 右侧
            if pTree.pRight == nil {
                return pTree
            }
            pTree = pTree.pRight
        } else {
            fmt.Println("data error")
            return nil
        }
    }
    return nil
}

// GetUncle 获取叔叔节点
func (r *RBTree) GetUncle(pNode *RBT) *RBT {
    if pNode == pNode.pFather.pLeft {
        return pNode.pFather.pRight
    } else {
        return pNode.pFather.pLeft
    }
}

// AddNode 插入节点
func (r *RBTree) AddNode(pTree *RBT, nNum int) {
    // 查找
    pNode := r.Search(pTree, nNum)

    pTemp := new(RBT)
    pTemp.nValue = nNum
    pTemp.nColor = RED
    pTemp.pLeft = nil
    pTemp.pRight = nil
    pTemp.pFather = pNode

    // 1. 根 直接插入
    if pNode == nil {
        r.Root = pTemp
        r.Root.nColor = BLACK
        return
    }

    // 连接
    if pNode.nValue > nNum {
        pNode.pLeft = pTemp
    } else {
        pNode.pRight = pTemp
    }

    // 2. 父亲是黑色的
    // 直接插入
    if pNode.nColor == BLACK {
        return
    }

    // 3. 父亲是红色的
    var pGrandFather *RBT
    var pUncle *RBT

    for pNode.nColor == RED {
        pGrandFather = pNode.pFather
        pUncle = r.GetUncle(pNode)

        // 3.1 叔叔是红的
        // 祖父节点变成红色
        // 父亲节点变成黑色
        // 叔叔节点变成黑色
        // 将祖父节点作为新插入节点进行回溯
        if pUncle != nil && pUncle.nColor == RED {
            pUncle.nColor = BLACK
            pNode.nColor = BLACK
            pGrandFather.nColor = RED

            pTemp = pGrandFather
            pNode = pTemp.pFather

            // 根
            if pNode == nil {
                r.Root.nColor = BLACK
                break
            }
            continue
        }

        // 3.2 叔叔是黑色的
        if pUncle != nil || pUncle.nColor == BLACK {
            // 3.2.1 父亲是爷爷的左
            if pNode == pGrandFather.pLeft {
                // 3.2.1.1 插入节点为右
                // 父节点为中心左旋
                // 祖父节点为中心右旋
                // 插入节点涂黑，祖父节点变红
                if pTemp == pNode.pRight {
                    pTemp = pNode
                    r.LeftRotate(pTemp)
                    pNode = pTemp.pFather
                }

                // 3.2.1.2 插入节点为左
                // 以祖父节点为中心右旋
                // 父节点变成黑色，祖父节点变成红色
                if pTemp == pNode.pLeft {
                    pNode.nColor = BLACK
                    pGrandFather.nColor = RED
                    r.RightRotate(pGrandFather)
                    break
                }
            }
            // 3.2.2 父亲是爷爷的右
            if pNode == pGrandFather.pRight {
                // 3.2.2.1 当前节点是父亲的左
                // 父节点为中心，右旋
                // 祖父节点为中心，左旋
                // 当前节点变成黑色
                // 祖父节点变成红色
                if pTemp == pNode.pLeft {
                    pTemp = pNode
                    r.RightRotate(pTree)
                    pNode = pTemp.pFather
                    continue
                }

                // 3.2.2.2 当前节点是父亲的右
                // 祖父节点左旋
                // 父节点变成黑色
                // 祖父节点变成红色
                if pTemp == pNode.pRight {
                    pNode.nColor = BLACK
                    pGrandFather.nColor = RED
                    r.LeftRotate(pGrandFather)
                    break
                }
            }
        }
    }
}

// FindNode 查找节点
func (r *RBTree) FindNode(pTree *RBT, nNum int) *RBT {
    for pTree != nil {
        if pTree.nValue == nNum {
            return pTree
        } else if pTree.nValue > nNum {
            pTree = pTree.pLeft
        } else {
            pTree = pTree.pRight
        }
    }
    return nil
}

func (r *RBTree) DeleteNode(pTree *RBT, nNum int) {
    if pTree == nil {
        return
    }

    // 查找
    pTemp := r.FindNode(pTree, nNum)
    if pTemp == nil {
        return
    }

    // 孩子情况（2个）
    var pMark *RBT
    if pTemp.pLeft != nil && pTemp.pRight != nil {
        pMark = pTemp

        pTemp = pTemp.pLeft
        for pTemp.pRight != nil {
            pTemp = pTemp.pRight
        }
        pMark.nValue = pTemp.nValue
    }

    pNode := pTemp.pFather

    // 1. 根
    if pNode == nil {
        // 没有孩子
        if pTemp.pLeft == nil && pTemp.pRight == nil {
            pTemp = nil
            r.Root = nil
            return
        }

        // 有一个红孩子
        if pTemp.pLeft != nil || pTemp.pRight != nil {
            if pTemp.pLeft != nil {
                r.Root = pTemp.pLeft
            } else {
                r.Root = pTemp.pRight
            }
            r.Root.nColor = BLACK
            r.Root.pFather = nil
            pTemp = nil
            return
        }
    }

    // 2. 非根且红色
    if pTemp.nColor == RED {
        if pTemp == pNode.pLeft {
            pNode.pLeft = nil
        } else {
            pNode.pRight = nil
        }
        pTemp = nil
        return
    }

    // 3. 非根 黑色 且有一个红孩子
    if pTemp.nColor == BLACK && (pTemp.pLeft != nil || pTemp.pRight != nil) {
        if pTemp == pNode.pLeft {
            if pTemp.pLeft != nil {
                pNode.pLeft = pTemp.pLeft
            } else {
                pNode.pLeft = pTemp.pRight
            }
            pNode.pLeft.nColor = BLACK
            pNode.pLeft.pFather = pNode
        } else {
            if pTemp.pLeft != nil {
                pNode.pRight = pTemp.pLeft
            } else {
                pNode.pRight = pTemp.pRight
            }
            pNode.pRight.nColor = BLACK
            pNode.pRight.pFather = pNode
        }
        pTemp = nil
        return
    }

    // 4. 非根 黑 且无孩子
    pBrother := r.GetUncle(pTemp)
    // 假删除
    if pTemp == pNode.pLeft {
        pNode.pLeft = nil
    } else {
        pNode.pRight = nil
    }

    pMark = pTemp

    for {
        // 4.1 兄弟是红色
        if pBrother.nColor == RED {
            pBrother.nColor = BLACK
            pNode.nColor = RED
            // 4.1.1 兄弟是父亲的右
            if pBrother == pNode.pRight {
                r.LeftRotate(pNode)
                pBrother = pNode.pRight
                continue
            }

            // 4.1.2 兄弟是父亲的左
            if pBrother == pNode.pLeft {
                r.RightRotate(pNode)
                pBrother = pNode.pLeft
                continue
            }
        }

        // 4.2 兄弟是黑色的
        if pBrother.nColor == BLACK {
            // 4.2.1 两个侄子全黑
            if (pBrother.pLeft == nil && pBrother.pRight == nil) ||
                ((pBrother.pLeft != nil && pBrother.pLeft.nColor == BLACK) && (pBrother.pRight != nil && pBrother.pRight.nColor == BLACK)) {
                // 4.2.1.1 父亲是红色的
                if pNode.nColor == RED {
                    pNode.nColor = BLACK
                    pBrother.nColor = RED
                    break
                }

                // 4.2.1.2 父亲是黑色的
                if pNode.nColor == BLACK {
                    pBrother.nColor = RED
                    pTemp = pNode
                    pNode = pTemp.pFather

                    // 根
                    if pNode == nil {
                        break
                    }

                    pBrother = r.GetUncle(pTemp)
                    continue
                }
            }

            // 4.2.2 左侄子红 右侄子黑
            if (pBrother.pLeft != nil && pBrother.pLeft.nColor == RED) &&
                (pBrother.pRight == nil || pBrother.pRight.nColor == BLACK) {
                // 4.2.2.1 兄弟是父亲的右
                if pBrother == pNode.pRight {
                    pBrother.nColor = RED
                    pBrother.pLeft.nColor = BLACK

                    r.RightRotate(pBrother)

                    pBrother = pNode.pRight
                    continue
                }
                // 4.2.2.2 兄弟是父亲的左
                if pBrother == pNode.pLeft {
                    pBrother.nColor = pNode.nColor
                    pNode.nColor = BLACK
                    pBrother.pLeft.nColor = BLACK

                    r.RightRotate(pNode)
                    break
                }
            }

            // 4.2.3 右侄子是红色的
            if pBrother.pRight != nil && pBrother.pRight.nColor == RED {
                // 4.2.3.1 兄弟是父亲的左
                if pBrother == pNode.pLeft {
                    pBrother.nColor = RED
                    pBrother.pRight.nColor = BLACK

                    r.LeftRotate(pBrother)
                    pBrother = pNode.pLeft
                    continue
                }
                // 4.2.3.2 兄弟是父亲的右
                if pBrother == pNode.pRight {
                    pBrother.nColor = pNode.nColor
                    pNode.nColor = BLACK
                    pBrother.pRight.nColor = BLACK

                    r.LeftRotate(pNode)
                    break
                }
            }
        }
    }

    pMark = nil
}

func CreateRBT(src []int, nLength int) *RBTree {
    rbt := &RBTree{Root: nil}
    for i := 0; i < nLength; i++ {
        rbt.AddNode(rbt.Root, src[i])
    }
    return rbt
}

func Traversal(pTree *RBT) {
    if pTree == nil {
        return
    }

    fmt.Println(fmt.Sprintf("val == %d\tcol == %d\n", pTree.nValue, pTree.nColor))
    Traversal(pTree.pLeft)
    Traversal(pTree.pRight)
}
