package html

import "io"

type Node struct {
	Type                  NodeType
	Data                  string
	Attr                  []Attribute
	FirstChild, NextChild *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	key, val string
}

func Parse(r io.Reader) (*Node, error)
