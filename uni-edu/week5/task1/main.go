// Task 1: Notifications (+ Participation Task: io.Writer table)
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// 1. Interface
type Notifier interface {
	Notify(msg string) error
}

// 2. Three implementations
type Email struct{ Address string }
type SMS struct{ Phone string }
type Telegram struct {
	Chat string
	sent int // how many messages were sent
}

func (e Email) Notify(msg string) error {
	fmt.Printf("[email -> %s] %s\n", e.Address, msg)
	return nil
}

func (s SMS) Notify(msg string) error {
	if len(msg) > 20 {
		return fmt.Errorf("sms to %s: message too long (%d > 20)", s.Phone, len(msg))
	}
	fmt.Printf("[sms -> %s] %s\n", s.Phone, msg)
	return nil
}

// pointer receiver: we MUST change t.sent, a value receiver would change a copy
func (t *Telegram) Notify(msg string) error {
	t.sent++
	fmt.Printf("[telegram -> %s] %s\n", t.Chat, msg)
	return nil
}

// 3. Broadcast
func Broadcast(ns []Notifier, msg string) []error {
	var errs []error
	for _, n := range ns {
		if err := n.Notify(msg); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// 4. Compile-time checks
var _ Notifier = Email{}
var _ Notifier = SMS{}
var _ Notifier = (*Telegram)(nil)

// ---- Participation Task: accepts io.Writer ----
type Grade struct {
	Name  string
	Score int
}

func PrintGrades(w io.Writer, grades []Grade) {
	fmt.Fprintf(w, "%-12s | %5s\n", "Student", "Score")
	fmt.Fprintln(w, strings.Repeat("-", 20))
	for _, g := range grades {
		fmt.Fprintf(w, "%-12s | %5d\n", g.Name, g.Score)
	}
}

func main() {
	tg := &Telegram{Chat: "@group"}
	ns := []Notifier{Email{"aru@mail.kz"}, SMS{"+77001234567"}, tg}

	fmt.Println("--- short message ---")
	errs := Broadcast(ns, "hello")
	fmt.Println("errors:", errs)

	fmt.Println("--- long message ---")
	errs = Broadcast(ns, "this message is definitely longer than 20 chars")
	fmt.Println("errors:", len(errs))
	for _, err := range errs {
		fmt.Println("  ", err)
	}
	fmt.Println("telegram sent:", tg.sent) // 2
	_ = errors.New

	grades := []Grade{{"Aruzhan", 95}, {"Dias", 80}, {"Aibek", 72}}
	fmt.Println("\n--- grades to os.Stdout ---")
	PrintGrades(os.Stdout, grades)

	var sb strings.Builder
	PrintGrades(&sb, grades)
	fmt.Println("\n--- grades captured in strings.Builder ---")
	fmt.Print(sb.String())
}
