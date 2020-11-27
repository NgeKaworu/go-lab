package graph

// Graph 图结构
type Graph struct {
	start *Vertex
}

// Vertex 顶点
type Vertex struct {
	key int             // 键
	in  map[*Vertex]int // 入度
	out map[*Vertex]int // 出度
}

/*
Connect 连接两个顶点
degree 度 —— 0: 出度; 1: 入度; 2:双向
weight 权重
*/
func (v *Vertex) Connect(n *Vertex, degree, weight int) *Vertex {
	switch degree {
	case 0:
		v.out[n] = weight
		break
	case 1:
		v.in[n] = weight
		break
	case 2:
		v.out[n] = weight
		v.in[n] = weight
		break

	}
	return v
}

/*
Disconnect 取消连接两个顶点
degree 度 —— 0: 出度; 1: 入度; 2:双向
*/
func (v *Vertex) Disconnect(n *Vertex, degree int) *Vertex {
	switch degree {
	case 0:
		delete(v.out, n)
		break
	case 1:
		delete(v.in, n)
		break
	case 2:
		delete(v.out, n)
		delete(v.in, n)
		break

	}
	return v
}
