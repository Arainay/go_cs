package strategy

type EvictionAlgo interface {
	evict(c *Cache)
}

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func InitCache(algo EvictionAlgo) *Cache {
	storage := map[string]string{}
	return &Cache{
		storage:      storage,
		evictionAlgo: algo,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) SetEvictionAlgo(algo EvictionAlgo) {
	c.evictionAlgo = algo
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func (c *Cache) Add(key string, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}

	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	value, isExists := c.storage[key]

	if isExists {
		delete(c.storage, key)
		c.capacity--
	}

	return value, isExists
}
