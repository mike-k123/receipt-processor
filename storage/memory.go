package storage

type InMemoryStore struct {
	receipts map[string]int
}

// Create a new in-memory store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		receipts: make(map[string]int),
	}
}

// Save the points for a receipt ID
func (store *InMemoryStore) Save(id string, points int) {
	store.receipts[id] = points
}

// Get the points for a receipt ID
func (store *InMemoryStore) Get(id string) (int, bool) {
	points, exists := store.receipts[id]
	return points, exists
}
