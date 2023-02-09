package main




import (
	"net/http"
	routes "newProject/Routes"
)

func main() {

    routes.Routes()
	http.ListenAndServe(":8000", nil)

}












/*
import (
	"fmt"
	"runtime"
	"sync"
)


func main() {

	     fmt.Println("CPUs: ", runtime.NumCPU())
		 fmt.Println("Goroutines:", runtime.NumGoroutine())

        contador := 0
		totaldegoroutines := 15


	 var wg sync.WaitGroup
           wg.Add(totaldegoroutines)

       var mu sync.Mutex

		   for i := 0 ; i < totaldegoroutines; i++ {
			go func(){
				mu.Lock()
				v := contador
                            runtime.Gosched()    //yield
				v++
				contador = v
                   mu.Unlock()
				wg.Done()
			}()
			fmt.Println("Goroutines:", runtime.NumGoroutine())
		   }
		   wg.Wait()
		   fmt.Println("Goroutines:", runtime.NumGoroutine())
		   fmt.Println("valor final:", contador)

}
-----------------------------------------------------------------------------------------------------------


import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// add () total de funções

	wg.Add(2)

	go func1()
	 go func2()

	wg.Wait()
}

func func1() {

	for i := 0; i < 10; i++ {

		fmt.Println("Func1:", i)
	}
	//deu
	wg.Done()
}

func func2() {

	for i := 0; i < 10; i++ {

		fmt.Println("Func2:", i)
	}
	wg.Done()
}
*/



