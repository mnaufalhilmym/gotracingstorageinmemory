package gotracingstorageinmemory_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mnaufalhilmym/gotracing"
	"github.com/mnaufalhilmym/gotracingstorageinmemory"
)

func recursiveFn(i int) {
	i--
	if i > 0 {
		recursiveFn(i)
	} else {
		gotracing.Debug("Recursive Function", "Test")
	}
}

func TestStorage(t *testing.T) {
	levelFilter := gotracing.LevelFilterTrace

	gotracing.SetMinConsolePrintLevel(levelFilter)
	gotracing.SetMinStoreLevel(levelFilter)
	gotracing.SetMaxProgramCounters(1)

	storage := gotracingstorageinmemory.New(3)
	gotracing.SetStorage(&storage)

	fmt.Println(">>> Test Trace 1 <<<")
	recursiveFn(5)
	fmt.Println(">>> Test Trace 2 <<<")
	recursiveFn(5)
	fmt.Println(">>> Test Trace 3 <<<")
	recursiveFn(5)
	fmt.Println(">>> Test Trace 4 <<<")
	recursiveFn(5)

	time.Sleep(1 * time.Millisecond)

	debugTraces := storage.GetAll(gotracing.LevelDebug)

	fmt.Println(">>> Test GetAll <<<")
	fmt.Println("Length:", len(debugTraces))
	fmt.Println(debugTraces)
}
