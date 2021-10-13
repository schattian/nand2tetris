package parser

import (
	"reflect"
	"testing"

	"github.com/schattian/nand2tetris/compiler/parse"
)

var (
	fooFieldSubset0 = &fieldSchema{subset: 0}
	fooFieldSubset1 = &fieldSchema{subset: 1}
	fooFieldSubset2 = &fieldSchema{subset: 2}
)

func Test_node_fieldsBySubset(t *testing.T) {
	type fields struct {
		children        []parse.Node
		fields          []*field
		token           *parse.Token
		nodeType        parse.NodeType
		closed          bool
		lastField       int
		lastFieldSubset int
	}
	tests := []struct {
		name        string
		node        *node
		wantSubsets [][]*field
	}{
		{
			name: "3 subsets",
			node: &node{
				schema: &nodeSchema{
					fieldsSchema: []*fieldSchema{fooFieldSubset0, fooFieldSubset1, fooFieldSubset2},
				},
			},
			wantSubsets: [][]*field{
				{&field{schema: fooFieldSubset0}},
				{&field{schema: fooFieldSubset1}},
				{&field{schema: fooFieldSubset2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSubsets := tt.node.getFieldsBySubset(); !reflect.DeepEqual(gotSubsets, tt.wantSubsets) {
				t.Errorf("node.fieldsBySubset() = %v, want %v", gotSubsets, tt.wantSubsets)
			}
		})
	}
}
