package main

import (
	"os"
	"path/filepath"

	"github.com/vadim-dmitriev/music-fs-web-interface/common"
)

// Node является представлением вершины дерева директории
type Node struct {
	FullPath string
	Name     string
	Children []*Node
	Parent   *Node
}

func main() {
	cfg, err := common.NewConfig()
	if err != nil {
		panic(err)
	}

	_, err = newTree(cfg.MusicDir)
	if err != nil {
		panic(err)
	}

}

func newTree(rootDir string) (result *Node, err error) {
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
			result = node
		} else {
			node.Parent = parent
			parent.Children = append(parent.Children, node)
		}
	}
	return
}
