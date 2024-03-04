package gee

import "strings"

/**
  前缀数路由解析
*/

type node struct {
	//待匹配路由， 例如 /p/:lang  只有末尾节点才有值
	pattern string
	//路由中的一部分 例如:lang
	part string
	//子节点  例如 [doc, tutorial, intro]
	children []*node
	//是否精确匹配 part 含有 : 或 * 时为true
	isWild bool
}

// 第一个匹配成功的节点， 用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	//匹配失败
	return nil
}

// 所有匹配度节点 用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 将路由规则插入到子节点中
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		//结尾节点
		n.pattern = pattern
		return
	}
	part := parts[height]
	//获取判断子节点
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
	}
	child.insert(pattern, parts, height+1)
}

// 进行搜索查找
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.pattern, "*") {
		//如果不是最后一个节点 返回 nil
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	//查找
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
