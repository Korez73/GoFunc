package lib

import "fmt"

func OuterLabel() {

	samples := []string{"hello", "apple_n!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}

}
