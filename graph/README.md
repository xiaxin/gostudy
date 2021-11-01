# Graph 图

- 加权图
- 有向图
- 无向图

**算法**
- 狄克斯特拉算法
- 贝尔曼-福特
- A*

**元素**
- Graph 图
- Edge 边 
- Node 节点


**图 Graph**
```
type Graph struct {
    nodes map[int]*Node  // 节点ID 关联 节点
    edges map[int]*Edge  // 节点ID 关联 边
}

func {
    AddNode
        添加 节点
    DelNode
        删除 节点
    SetEdge
        设置 边
    DelEdge
        删除 边
}
```

**节点 Node**
```
type Node struct {
    id  int
    val interface{}
}
```

**边 Edge**
```
type Edge Struct {
    f *Node
    t *Node
    s float64 // 权重分
}

func {
    From()  return e.F
    To()    return e.T
    Score() return e.S
}
```