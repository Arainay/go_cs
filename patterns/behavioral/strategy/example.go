package strategy

func Example() {
	lfu := &Lfu{}
	lru := &Lru{}
	fifo := &Fifo{}

	cache := InitCache(lfu)

	cache.Add("a", "first")
	cache.Add("b", "second")
	cache.Add("c", "third")

	cache.SetEvictionAlgo(lru)

	cache.Add("d", "fourth")

	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "fifth")
}
