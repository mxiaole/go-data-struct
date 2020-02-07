package graph

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 邻接表表示的图
type AdjList struct {
	v    int           // 节点的个数
	e    int           // 边的个数
	list map[int][]int // k表示节点的编号，v表示与该节点相连的节点的下标
}

// 通过文件创建图
func (g *AdjList) Init(fileName string) {
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	br := bufio.NewReader(f)
	line, _, err := br.ReadLine()
	s := strings.Split(string(line), " ")
	g.v, _ = strconv.Atoi(s[0])
	g.e, _ = strconv.Atoi(s[1])
	g.list = make(map[int][]int)

	for i := 0; i < g.v; i++ {
		var l []int
		g.list[i] = l
	}

	for {
		a, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}

		l := strings.Split(string(a), " ")
		e1, _ := strconv.Atoi(l[0])
		e2, _ := strconv.Atoi(l[1])

		g.list[e1] = append(g.list[e1], e2)
		g.list[e2] = append(g.list[e2], e1)
	}
}

// 获取图中节点的个数
func (g *AdjList) GetVertex() int {
	return g.v
}

// 获取图中边的个数
func (g *AdjList) GetEdge() int {
	return g.e
}

// 判断两个顶点是否有边
func (g *AdjList) HasEdge(v, w int) bool {
	s := g.list[v]
	for _, v := range s {
		if v == w {
			return true
		}
	}

	return false
}

// 获取一个顶点相邻的所有顶点
func (g *AdjList) Ajd(v int) []int {
	return g.list[v]
}

// 获取一个顶点的度
func (g *AdjList) Degree(v int) int {
	return len(g.list[v])
}

// 打印邻接表
func (g *AdjList) Print() {
	for k, v := range g.list {
		fmt.Print(k, ":", v)
		fmt.Println()
	}
}
