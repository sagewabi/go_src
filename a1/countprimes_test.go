package a1

import "testing"


func TestCountPtrings(t *testing.T) { 
    n := 40
    if n<=1 && countPrimes(n) != 0{
        t.Errorf("error number of primes should equal 0")
    }else if n==2 && countPrimes(n) != 1{
     t.Errorf("error there should be 1 prime")
     }
 
    tables := []struct {
        x int
        y int
    }{
        {-40, 0}, // table of inputs x and expected outputs y
        {1,0},
        {2,1},
        {3,2},
        {30,10},
        {40,12},
        {346346, 29683},
        {4563,618},
        {97883,9410},
        {454574564, 24082464},
        {100000000,5761455},
    }

    for _, table := range tables {
        total_primes := countPrimes(table.x)
        if  total_primes != table.y {
            t.Errorf(" %d is the wrong amount of primes containd in %d", total_primes, table.x)
        }
    }
   
  
     

    
}