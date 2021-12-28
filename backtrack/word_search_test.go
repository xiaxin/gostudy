package backtrack

import "testing"

func TestWordSearch(t *testing.T) {
	var items = []struct {
		name  string
		board [][]byte
		word  string
	}{
		{"79 word ABCCED",
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCCED",
		},
		{"79 word SEE",
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"SEE",
		},
		{"79 word ABCB",
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCB",
		},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			result := WordSearch79(item.board, item.word)
			t.Log(result)
		})
	}

}
