package main

import "fmt"

func batcher(size int, shard_size int, jobs []string) {
	batch := make(map[int][]string)

	j := 0
	for i := 0; i < size; i++ {
		if i != 0 && i%shard_size == 0 {
			j++
		}

		batch[j] = append(batch[j], jobs[i])
	}
	fmt.Println(batch)

}

func main() {
	jobs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	a := len(jobs)
	batcher(a, 4, jobs)
}
