package transformers

import (
	"log"

	"github.com/yuin/goldmark/ast"
	gast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// TaskTransformer ...
type TaskTransformer struct {
	Task      int
	taskItems int
}

// Transform ...
func (t *TaskTransformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	source := reader.Source()

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		// Each node will be visited twice, once when it is first encountered (entering), and again
		// after all the node's children have been visited (if any). Skip the latter.
		if !entering {
			return ast.WalkContinue, nil
		}

		if node.Kind() == gast.KindTaskCheckBox {
			t.taskItems++

			if t.Task == t.taskItems {
				checkboxNode := node.(*gast.TaskCheckBox)
				t.CheckTaskCheckBox(checkboxNode, source)
			}
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		log.Fatal("Error encountered while transforming AST:", err)
	}
}

// CheckTaskCheckBox ...
func (t *TaskTransformer) CheckTaskCheckBox(node *gast.TaskCheckBox, source []byte) {
	node.IsChecked = true
}
