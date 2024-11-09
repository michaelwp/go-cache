package go_cache

type cacheMapValue struct {
	Id    int32
	Value interface{}
}

type cacheMap map[string]cacheMapValue
type cacheMapIndex map[int32]string

type Cache interface {
	Add(key string, value interface{})
	Get(key string) (interface{}, bool)
	View() interface{}
}

type cache struct {
	Capacity      int32
	cacheMap      cacheMap
	cacheMapIndex cacheMapIndex
	NewestId      int32
	OldestId      int32
}

func NewCache(capacity int32) Cache {
	return &cache{
		Capacity:      capacity,
		cacheMap:      make(cacheMap),
		cacheMapIndex: make(cacheMapIndex),
		NewestId:      0,
		OldestId:      1,
	}
}

func (c *cache) Add(key string, value interface{}) {
	c.NewestId++
	c.cacheMap[key] = cacheMapValue{Id: c.NewestId, Value: value}
	c.cacheMapIndex[c.NewestId] = key

	if len(c.cacheMap) > int(c.Capacity) {
		// if cache is full, evict the oldest value
		delete(c.cacheMap, c.cacheMapIndex[c.OldestId])
		delete(c.cacheMapIndex, c.OldestId)

		c.updateOldestId()
	}
}

func (c *cache) Get(key string) (interface{}, bool) {
	if value, ok := c.cacheMap[key]; ok {
		// set this item as the newest
		if value.Id == c.OldestId {
			c.OldestId++
		}

		delete(c.cacheMapIndex, value.Id)
		c.Add(key, value.Value)

		return value, ok
	}

	return nil, false
}

func (c *cache) View() interface{} {
	return c.cacheMap
}

func (c *cache) updateOldestId() {
	// find the oldest id by incrementing oldestId by 1
	for {
		if _, ok := c.cacheMapIndex[c.OldestId]; ok {
			return
		}

		c.OldestId++
	}
}
