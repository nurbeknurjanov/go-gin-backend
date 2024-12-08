package manuals

/*func RunTest() {
	bctx := context.Background()

	ctxValue := context.WithValue(bctx, "key", "testValue")

	ctxCancel, cancel := context.WithCancel(bctx)
	_ = cancel

	deadlineCtx, cancelDead := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancelDead() // Освобождаем ресурсы

	go func() {
		fmt.Println("Value", ctxValue.Value("key"))
		//time.Sleep(1 * time.Second)
		select {
		case <-ctxCancel.Done(): //подслушивает
			fmt.Println("Operation cancelled")
		case <-deadlineCtx.Done():
			fmt.Println("Deadline exceeded:", deadlineCtx.Err())
		case <-time.After(3 * time.Second):
			fmt.Println("3 second passed")
		}

	}()

	//cancel()
	time.Sleep(5 * time.Second)
}
*/
