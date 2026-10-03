# Результаты выполнения (Go 1.22.2, linux/amd64)

| Папка | Задание | Запуск |
|---|---|---|
| t1_strings | len("世界"), итерация по "hêllo", вторая руна | `go run ./t1_strings` |
| t2_trim | TrimRight / TrimLeft / TrimPrefix / TrimSuffix | `go run ./t2_trim` |
| t3_builder | `+=` vs Builder vs Builder.Grow | `go run ./t3_builder` |
| t4_payload | Рефакторинг processPayload + BenchmarkProcessPayload | `go test ./t4_payload -bench=. -benchmem` |
| t5_scores | Bob / Charlie, двузначное присваивание | `go run ./t5_scores` |
| t6_inventory | CRUD на map inventory | `go run ./t6_inventory` |
| t7_wordcount | Tour of Go: Exercise Maps (WordCount) | `go run ./t7_wordcount` |

## Ключевые числа

len("世界") = 6, RuneCountInString = 2.

Конкатенация 1000 строк:

| Способ | ns/op | B/op | allocs/op |
|---|---|---|---|
| `+=` | 661670 | 4717224 | 999 |
| Builder | 10687 | 34296 | 15 |
| Builder + Grow | 5390 | 9472 | 1 |

BenchmarkProcessPayload:

| Версия | ns/op | B/op | allocs/op |
|---|---|---|---|
| original | 5679 | 8192 | 2 |
| optimized | 4153 | 4096 | 1 |

Числа ns/op зависят от машины, на защите запусти бенчмарк у себя.
