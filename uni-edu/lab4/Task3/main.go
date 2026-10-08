//Without running the code, mark every numbered line as ✅ compiles or ❌
//does not compile, and explain why. Then check yourself with go build
package main


type Counter struct{
	n int
}

func (c Counter) Get() int {
	return c.n
}

func (c Counter) Inc() {
	c.n++
}

func NewCounter() Counter {
	return Counter{}
}

func main(){
	var c Counter
	c.Inc()
	(&c).Inc()
	Counter{}.Get()
	Counter{}.Inc()
	NewCounter().Inc()

	m := map[string]Counter{"a": {}}
	m["a"].Get()
	m["a"].Inc()

	s := []Counter{{}}
	s[0].Inc()

	pm := map[string]*Counter{"a":{}}
	pm["a"].Inc()
}
