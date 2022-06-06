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

//func main() {
//	fmt.Println("Start")
//	var wg sync.WaitGroup
//
//	wg.Add(2)
//
//	go func(wg *sync.WaitGroup) {
//		defer wg.Done()
//		fmt.Println(1)
//		//time.Sleep(100*time.Microsecond)
//	}(&wg)
//	//
//	//go func(wg *sync.WaitGroup) {
//	//	defer wg.Done()
//	//	fmt.Println(2)
//	//}(&wg)
//
//	go func(wg *sync.WaitGroup) {
//		defer wg.Done()
//		time.Sleep(1 * time.Microsecond)
//		log.Print(1)
//	}(&wg)
//
//	wg.Wait()
//
//	fmt.Println("Endline")
//
//}

//func main() {
//
//	worker1CH := make(chan int, 1)
//	worker2CH := make(chan int, 1)
//
//	// worker for even numbers
//	go func(in chan int) {
//		for i := range in {
//			log.Print(i)
//		}
//	}(worker1CH)
//
//	// worker for odd numbers
//	go func(in chan int) {
//		for i := range in {
//			log.Print(i)
//		}
//	}(worker2CH)
//
//	// sender which sends even numbers to worker1CH, and odd numbers to worker2CH
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func(wg *sync.WaitGroup, evenChan chan int, oddChan chan int) {
//		defer wg.Done()
//
//		data := rand.Perm(10)
//		for _, i := range data {
//			switch i%2 {
//			case 0:
//				evenChan <- i
//			default:
//				oddChan <- i
//			}
//		}
//	}(&wg, worker1CH, worker2CH)
//	wg.Wait()
//
//}
