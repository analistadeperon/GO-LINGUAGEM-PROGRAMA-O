package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}

package main
------------------------------------------------------
import (
	"fmt"
)

func main() {
  fmt.Println("Please enter your name.")
}

------------------------------------------------------------

package main

import (
		"fmt"
		"strings"
)

func main() {
	fmt.Println("Please enter your name.")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hi, %s! I'm Go!", name)
	name = strings.TrimSpace(name)
}
-----------------------------------------------------------

package main

import (
	"fmt"
	"time"

	)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// estrutura de controle específica para concorrencia
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		//	default:
		//		return "Sem resposta ainda"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.cade.com",
		"https://www.uol.com",
		"https://www.google.com",
	)

	fmt.Println(campeao)
}

-------------------------------------------------------------------------


package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(5, 4))
}

---------------------------------------------------

package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(5, 4))
}

---------------------------------------------------------

package main
import "time"
import "fmt"
func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "um"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "dois"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("recebido", msg1)
        case msg2 := <-c2:
            fmt.Println("recebido", msg2)
        }
    }
}

----------------------------------------------------------------------

package main

import "fmt"

func main() {

	a := 10
	b := 20

	c := Add(a, b)

	fmt.Println("O resultado dessa soma é: ", c)

}

// Add é um função que recebe dois inteiros e retorna a soma entre eles
func Add(a int, b int) int {
	return a + b
}

--------------------------------------------------------------------
