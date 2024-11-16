package orders

import "context"

func main() {
	store := NewStore()
	svc := NewService(store)

	svc.CreateOrder(context.Background())
}
