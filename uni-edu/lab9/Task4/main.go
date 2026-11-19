// Task 4: Equivalent Binary Trees (Tour of Go)
//
// Пакет golang.org/x/tour/tree не подключаем (модуль без зависимостей),
// поэтому tree.New(k) реализован здесь так же, как в оригинале:
// дерево из значений k, 2k, ..., 10k, вставленных в случайном порядке.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

// Walk отправляет все значения дерева в ch по возрастанию и закрывает ch.
// Walk рекурсивен, поэтому close нельзя ставить в рекурсию (закрыли бы
// канал после первого поддерева). Рекурсия вынесена в walk, а close
// делает внешняя функция — ровно один раз, когда обход закончен.
//
// done нужен, чтобы Same мог остановить обход досрочно: без него при
// разных деревьях Same выходит, а Walk навсегда висит на ch <- v (утечка).
func Walk(t *Tree, ch chan<- int, done <-chan struct{}) {
	defer close(ch)
	walk(t, ch, done)
}

func walk(t *Tree, ch chan<- int, done <-chan struct{}) bool {
	if t == nil {
		return true
	}
	if !walk(t.Left, ch, done) {
		return false
	}
	select {
	case ch <- t.Value:
	case <-done:
		return false
	}
	return walk(t.Right, ch, done)
}

// Same — с остановкой через done: горутины не утекают.
func Same(t1, t2 *Tree) bool {
	done := make(chan struct{})
	defer close(done) // при любом выходе останавливаем оба Walk
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1, done)
	go Walk(t2, c2, done)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			return true
		}
	}
}

// SameLeaky — «наивная» версия из тура: при первом расхождении выходит,
// а обе горутины Walk остаются заблокированными на отправке навсегда.
func SameLeaky(t1, t2 *Tree) bool {
	never := make(chan struct{}) // никогда не закрывается
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1, never)
	go Walk(t2, c2, never)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			return true
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(New(1), ch, nil) // nil done: case <-done выключен навсегда
	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println("Same(New(1), New(1)):", Same(New(1), New(1)))
	fmt.Println("Same(New(1), New(2)):", Same(New(1), New(2)))

	settle := func() int { time.Sleep(50 * time.Millisecond); return runtime.NumGoroutine() }

	fmt.Println("\ngoroutines at start:", settle())
	for i := 0; i < 100; i++ {
		Same(New(1), New(2))
	}
	fmt.Println("after 100 × Same(different):      ", settle())
	for i := 0; i < 100; i++ {
		SameLeaky(New(1), New(2))
	}
	fmt.Println("after 100 × SameLeaky(different): ", settle(), "(+200 утёкших Walk)")
}
