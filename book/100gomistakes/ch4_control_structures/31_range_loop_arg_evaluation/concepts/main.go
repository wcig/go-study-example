package main

func main() {
	{
		// 无限循环
		s1 := []int{0, 1, 2}
		for range s1 {
			s1 = append(s1, 10)
		}
	}

	{
		// 无限循环
		s2 := []int{0, 1, 2}
		for i := 0; i < len(s2); i++ {
			s2 = append(s2, 10)
		}
	}
}
