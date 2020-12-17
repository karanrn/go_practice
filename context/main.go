package main

import (
	"context"
	"fmt"
	"time"
)

type key string

const (
	keyLvl1  key = "lvl1"
	keyLvl2a key = "lvl2a"
	keyLvl2b key = "lvl2b"
	notExist key = "non exist"
)

func childTask(ctx context.Context, origin string, d time.Duration) {
	const childKey key = "childTask"
	ctx2 := context.WithValue(ctx, childKey, "test-value")

	go grandChildTask(ctx2, origin, 6*time.Second)
	select {
	case <-ctx2.Done():
		fmt.Println("child task cancelled :( from " + origin + ", cause: " + ctx.Err().Error())
	case <-time.After(d):
		fmt.Println("child task done, from", origin)
	}
}

func grandChildTask(ctx context.Context, origin string, d time.Duration) {
	const grandChildKey key = "grandChildTask"
	ctx2 := context.WithValue(ctx, grandChildKey, "test-value")

	select {
	case <-ctx2.Done():
		fmt.Println("child task cancelled :( from " + origin + ", cause: " + ctx.Err().Error())
	case <-time.After(d):
		fmt.Println("child task done, from", origin)
	}
}

func Task1(ctx context.Context) {
	const (
		ctxKey key = "task1"
	)
	ctx2 := context.WithValue(ctx, ctxKey, "test-value")
	go childTask(ctx2, "Task1", 300*time.Millisecond)

	select {
	case <-ctx2.Done():
		fmt.Println("task 1 cancelled :(, err: ", ctx.Err())
	case <-time.After(3 * time.Second):
		fmt.Println("task 1 done")
	}
}

func Task2(ctx context.Context) {
	ctx2, cancel := context.WithCancel(ctx)
	go childTask(ctx2, "Task2", 9*time.Second)

	go func() {
		// simulating bad thing happens after 5 seconds
		// we cancel the context
		time.Sleep(5 * time.Second)
		fmt.Println("Oh no..., something bad happen... cancelling the context !!!")
		cancel()
	}()

	select {
	case <-ctx2.Done():
		fmt.Println("task 2 cancelled :(, cause: ", ctx2.Err())
	case <-time.After(7 * time.Second):
		fmt.Println("task 2 done")
	}
}

func demoWithContextCancellation() {
	// Create a root context
	ctx := context.Background()
	firstCtx, _ := context.WithCancel(ctx)

	go Task1(firstCtx)
	go Task2(firstCtx)

	time.Sleep(10 * time.Second)
}

func main() {
	// Create a root context
	ctx := context.Background()

	ctxLvl1 := context.WithValue(ctx, keyLvl1, "value-1")
	/*
		ctxLvl1, cancel := context.WithTimeout(ctx, time.Second*2)
		time.Sleep(2)
		defer cancel()
	*/
	value := ctxLvl1.Value(keyLvl1)
	fmt.Println(value)

	ctxLvl2a := context.WithValue(ctxLvl1, keyLvl2a, "value-2a")
	ctxLvl2b := context.WithValue(ctxLvl1, keyLvl2b, "value-2b")
	fmt.Println(ctxLvl2a.Value(keyLvl2a))
	fmt.Println(ctxLvl2b.Value(keyLvl2b), ctxLvl2b.Value(keyLvl1))

}
