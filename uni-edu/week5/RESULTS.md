# Week 5: results and notes for the defence

All code was run with Go 1.22.2 (linux/amd64, Xeon 2.1 GHz VM). Re-run on your own machine and use YOUR numbers.

## How to run
    go run ./task1      go run ./task2      go run ./task3      go run ./task4
    go run ./task5      go run ./task7      go run ./bonus
    go test -bench=. -benchmem ./task6
    go vet ./task4/recursive      # shows the recursion warning
    go build ./task5/notilde      # shows the missing ~ error

## Task 2: expected vs actual (all matched)
nil -> nil | 7 -> integer 7 | int64(7) -> integer 7 | "привет" -> string of 12 bytes
[]int{1,2,3} -> slice with 3 ints | errors.New("disk full") -> error: disk full
Celsius(21.5) -> Stringer: 21.5°C | 3.14 -> unknown type float64 | struct{}{} -> unknown type struct {}
"привет" is 12 bytes: UTF-8, each Cyrillic letter = 2 bytes, 6 letters.
Timeout: kind(Timeout{5}) = "Stringer" (first matching case wins; swap the first two cases -> "error").
fmt.Println(e) for an error prints Error(): "timeout after 3s".

## Task 3
Buggy validate(20) prints:  error: <nil>
err == nil -> false, %T -> *main.ValidationError.
Reason: interface holds (*ValidationError, nil). fmt calls Error() on a nil receiver, e.Field panics, fmt recovers and prints <nil>.

## Task 4: go vet
    task4/recursive/main.go:10:9: fmt.Sprint arg e causes recursive call to (...ErrNegativeSqrt).Error method
Sprint(e) sees that e has Error(), calls it, which calls Sprint(e) again... stack overflow. Fix: float64(e).

## Task 5, step 5: compiler error without ~
    Score does not satisfy Number (possibly missing ~ for int in Number)

## Task 6: benchmark (go test -bench=. -benchmem, -count=3, typical values)
| Benchmark                 | ns/op       | B/op  | allocs/op |
|---------------------------|-------------|-------|-----------|
| BenchmarkSumInts          | ~410-455    | 0     | 0         |
| BenchmarkSumGeneric       | ~560-620    | 0     | 0         |
| BenchmarkSumAny           | ~730-800    | 0     | 0         |
| BenchmarkSumAnyWithBoxing | ~17000-21500| 22336 | 745       |

Explanations
- Generic vs concrete: neither allocates and both are plain loops over []int/[]T, no boxing, no type checks.
  HONEST NOTE: in my runs generic was ~30% slower than concrete (Go 1.22). Go does not fully stamp out a copy per type:
  it compiles per "GC shape" and may pass a hidden dictionary, so a small overhead is possible. The point for the
  defence: same order of magnitude, 0 allocs, and far cheaper than any-based code. Report what YOUR run shows.
- SumAny is slower even with 0 allocs: a type assertion n.(int) on every element (compare dynamic type), the value
  is read through a pointer, and each element is 16 bytes (type+pointer) instead of 8, so worse cache use.
- Allocations in the boxing benchmark: 745 = 1 (the []any slice) + 744 boxed ints. Values 0..255 are not allocated:
  Go points the interface at a static table for small values. 1000 - 256 = 744. So not exactly 1001.

## Bonus
a == b -> true (both int 1) | c == d -> false (int vs int64) | e == f -> PANIC: comparing uncomparable type []int
It compiles because == is always allowed on interface values; the check happens at run time.
