package main

import "fmt"

func main() {
	channel := make(chan string)
	defer close(channel)

	go GoroutinePrimeira(channel)
	go GoroutineSegunda(channel)
	go func(channel chan string) {
		for {
			select {
			case message := <-channel:
				fmt.Println("Terceira: ", message)
			default:
				break
			}
		}
	}(channel)

	for i := 0; i < 1000; i++ {
		channel <- fmt.Sprint(i)
	}
}

func GoroutinePrimeira(channel chan string) {
	for {
		select {
		case message := <-channel:
			fmt.Println("Primeira: ", message)
		default:
			break
		}
	}
}

func GoroutineSegunda(channel chan string) {
	for {
		select {
		case message := <-channel:
			fmt.Println("Segunda: ", message)
		default:
			break
		}
	}
}
