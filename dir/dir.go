package dir

import (
	"os"
	"path/filepath"
)

// Node является представлением вершины дерева директории
type Node struct {
	FullPath string
	Name     string
	Children []*Node
	Parent   *Node
}

// NewTree создает древовидную стркутуру на основе директории с музыкой.
// Возвращает указатель на корень дерева
func NewTree(rootDir string) (root *Node, err error) {
	absRoot, err := filepath.Abs(rootDir)
	if err != nil {
		return
	}
	parents := make(map[string]*Node)

	err = filepath.Walk(absRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		parents[path] = &Node{
			Name:     info.Name(),
			FullPath: path,
			Children: make([]*Node, 0),
		}
		return nil
	})

	if err != nil {
		return
	}

	for path, node := range parents {
		parentPath := filepath.Dir(path)

		parent, exists := parents[parentPath]
		if !exists { // Если нет родителя, то это корень.
			root = node
		} else {
			node.Parent = parent
			parent.Children = append(parent.Children, node)
		}
	}

	return
}
