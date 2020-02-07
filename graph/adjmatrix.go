package graph

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// 邻接矩阵表示的图

type AdjMatrix struct {
	v      int     // 顶点的个数
	e      int     // 边的个数
	matrix [][]int // 数组的下标表示节点的编号，值表示顶点间的关系
}

// 1. 从文件初始化一个图
func (g *AdjMatrix) Init(file string) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	br := bufio.NewReader(f)
	line, _, err := br.ReadLine()
	s := strings.Split(string(line), " ")
	g.v, _ = strconv.Atoi(s[0])
	g.e, _ = strconv.Atoi(s[1])
	g.matrix = make([][]int, g.v)

	for i := 0; i < g.v; i++ {
		var a = make([]int, g.v)
		g.matrix[i] = a
	}

	for {
		a, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}

		l := strings.Split(string(a), " ")
		e1, _ := strconv.Atoi(l[0])
		e2, _ := strconv.Atoi(l[1])
		g.matrix[e1][e2] = 1
		g.matrix[e2][e1] = 1
	}

}

// 获取图中顶点的个数
func (g *AdjMatrix) GetVertex() int {
	return g.v
}

// 获取图中边的个数
func (g *AdjMatrix) GetEdge() int {
	return g.e
}

// 判断给定的两个顶点是否有边
func (g *AdjMatrix) HasEdge(v, w int) bool {
	return g.matrix[v][w] == 1
}

// 获取一个顶点相邻的所有顶点
func (g *AdjMatrix) Adj(v int) (adj []int) {
	a1 := g.matrix[v]
	for i, v := range a1 {
		if v == 1 {
			adj = append(adj, i)
		}
	}

	return
}

// 获取一个顶点的度
func (g *AdjMatrix) Degree(v int) int {
	return len(g.Adj(v))
}

// 打印邻接矩阵
func (g *AdjMatrix) Print() {
	for i := 0; i < g.v; i++ {
		for j := 0; j < g.v; j++ {
			fmt.Print(g.matrix[i][j], "\t")
		}
		fmt.Println()
	}
}
