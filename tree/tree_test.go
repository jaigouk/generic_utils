package tree_test

import (
	"fmt"
	"testing"
	"reflect"
	tree_util "github.com/jaigouk/generic_utils/tree"
)

func TestTree(t *testing.T) {
	t.Run("Insert", func(t *testing.T) {
		values := []string{"d", "b", "g", "g", "c", "e", "a", "h", "f", "i", "j", "l", "k"}

		data := []string{"delta", "bravo", "golang", "golf", "charlie", "echo", "alpha", "hotel", "foxtrot", "india", "juliett", "lima", "kilo"}

		tree := &tree_util.Tree[string, string]{}
		for i := 0; i < len(values); i++ {
			tree.Insert(values[i], data[i])
		}

		result := []string{}

		tree.Traverse(tree.Root, func(n *tree_util.Node[string, string]) {
			result = append(result, n.Data)
		})

		expected := []string{"alpha", "bravo", "charlie", "delta", "echo", "foxtrot", "golf", "hotel", "india", "juliett", "kilo", "lima"}

		if reflect.DeepEqual(result,expected) {
			fmt.Println("OK")
		} else {
			t.Error("Expected", expected, "got", result)
		}
		//fmt.Printf("res = %v\n", result)
		//fmt.Print("\n*** Tree with string search values and string data ***\n\n")
		//fmt.Print("Sorted values: | ")
		//tree.Traverse(tree.Root, func(n *tree_util.Node[string, string]) { fmt.Print(n.Value, ": ", n.Data, " | ") })
		//fmt.Println()

		//fmt.Println("Pretty print (turned 90° anti-clockwise):")
		//tree.PrettyPrint()
		////fmt.Printf("\nDump: %v\n", tree.Dump())
		//fmt.Println("================")

	})

	t.Run("Insert", func(t *testing.T) {
		// int tree

		keys := []int{4, 2, 7, 7, 3, 5, 1, 8, 6, 9, 10, 12, 11}

		data := []string{"delta", "bravo", "golang", "golf", "charlie", "echo", "alpha", "hotel", "foxtrot", "india", "juliett", "lima", "kilo"}
		intTree := &tree_util.Tree[int, string]{}
		for i := 0; i < len(keys); i++ {
			intTree.Insert(keys[i], data[i])
		}

		result := []string{}

		intTree.Traverse(intTree.Root, func(n *tree_util.Node[int, string]) {
			result = append(result, n.Data)
		})

		expected := []string{"alpha", "bravo", "charlie", "delta", "echo", "foxtrot", "golf", "hotel", "india", "juliett", "kilo", "lima"}

		if reflect.DeepEqual(result,expected) {
			fmt.Println("OK")
		} else {
			t.Error("Expected", expected, "got", result)
		}
		//fmt.Print("\n*** Tree with int search values and string data ***\n\n")
		//fmt.Print("Sorted values: | ")
		//intTree.Traverse(intTree.Root, func(n *tree_util.Node[int, string]) { fmt.Print(n.Value, ": ", n.Data, " | ") })
		//fmt.Println()

		//fmt.Println("Pretty print")
		//intTree.PrettyPrint()
	})
}
