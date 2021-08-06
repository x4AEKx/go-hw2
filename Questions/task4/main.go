package main

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch { // чтение из открытого и пустого канала приведет к блокировке goрутины,
		println(n) // до записи в канал или до закрытия канала
	}
}
