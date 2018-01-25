package a1

import "testing"


func TestCountStrings(t *testing.T) { 
    
    
    tables := []struct {
        x string
        y int
    }{
			{"one", 1}, // table of inputs x and expected outputs y
			{"mind",0},
			{"as",2},
			{"but",2},
			{"the",4},
			{"of",9},
			{"in", 2},
			{"(test)",5},
			{"Each",1},
			{"each", 1},
			{"534",0},
	}
	counts := countStrings("filetext.txt")
	for _, table := range tables {
			//total_primes := countStrings(table.x)
		if  counts[table.x] != table.y {
				t.Errorf(" %d is the wrong amount of times the word \"%s\" is containd in this text", counts[table.x], table.x)
		}
	}

}