package main

import (
	"fmt"
)

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	m map[string]*Customer
}

func main() {
	{
		s := Store{
			m: make(map[string]*Customer),
		}
		s.storeCustomers1([]Customer{
			{ID: "1", Balance: 10},
			{ID: "2", Balance: -10},
			{ID: "3", Balance: 0},
		})
		print(s.m)
	}

	{
		s := Store{
			m: make(map[string]*Customer),
		}
		s.storeCustomers2([]*Customer{
			{ID: "1", Balance: 10},
			{ID: "2", Balance: -10},
			{ID: "3", Balance: 0},
		})
		print(s.m)
	}

	{
		s := Store{
			m: make(map[string]*Customer),
		}
		s.storeCustomers3([]*Customer{
			{ID: "1", Balance: 10},
			{ID: "2", Balance: -10},
			{ID: "3", Balance: 0},
		})
		print(s.m)
	}

	// Output:
	// 0x14000114018
	// 0x14000114018
	// 0x14000114018
	// key=1, value=&main.Customer{ID:"3", Balance:0}
	// key=2, value=&main.Customer{ID:"3", Balance:0}
	// key=3, value=&main.Customer{ID:"3", Balance:0}
	// 0x14000114078
	// 0x14000114090
	// 0x140001140a8
	// key=2, value=&main.Customer{ID:"2", Balance:-10}
	// key=3, value=&main.Customer{ID:"3", Balance:0}
	// key=1, value=&main.Customer{ID:"1", Balance:10}
	// 0x14000114108
	// 0x14000114120
	// 0x14000114138
	// key=1, value=&main.Customer{ID:"1", Balance:10}
	// key=2, value=&main.Customer{ID:"2", Balance:-10}
	// key=3, value=&main.Customer{ID:"3", Balance:0}
}

func (s *Store) storeCustomers1(customers []Customer) {
	for _, customer := range customers {
		fmt.Printf("%p\n", &customer)
		s.m[customer.ID] = &customer
	}
}

func (s *Store) storeCustomers2(customers []*Customer) {
	for _, customer := range customers {
		fmt.Printf("%p\n", customer)
		s.m[customer.ID] = customer
	}
}

func (s *Store) storeCustomers3(customers []*Customer) {
	for i := range customers {
		customer := customers[i]
		fmt.Printf("%p\n", customer)
		s.m[customer.ID] = customer
	}
}

func print(m map[string]*Customer) {
	for k, v := range m {
		fmt.Printf("key=%s, value=%#v\n", k, v)
	}
}
