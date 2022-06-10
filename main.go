package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"shrading/routes"
)

func main() {
	app := fiber.New()

	routes.RegisterAPI(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	errA := app.Listen(":3000")
	if errA != nil {
		log.Println("SERVICE START ERROR: " + errA.Error())
	} else {
		log.Println("SERVICE RUNNING ON PORT 3000")
	}
}

//var (
//	mutex   sync.Mutex
//	balance int
//)

//func deposit(value int, wg *sync.WaitGroup) {
//	mutex.Lock()
//	fmt.Printf("Depositing %d to account with balance %d", value, balance)
//	balance += value
//	mutex.Unlock()
//	wg.Done()
//}
//
//func withdraw(value int, wg *sync.WaitGroup) {
//	mutex.Lock()
//	fmt.Printf("Withdraw %d to account with balance %d", value)
//	balance -= value
//	mutex.Unlock()
//	wg.Done()
//}
//
//func main() {
//	balance = 1000
//
//	var wg sync.WaitGroup
//
//	wg.Add(2)
//	go withdraw(700, &wg)
//
//	go deposit(500, &wg)
//	wg.Wait()
//
//	fmt.Printf("New Balance %d\n", balance)
//}
