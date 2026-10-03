# Week 3: Strings & Maps (Go): подготовка к защите

## Содержание
1. Строки, байты, руны
2. Trim-функции
3. Медленная конкатенация и `strings.Builder`
4. Конвертации `[]byte` ↔ `string` (задача с бенчмарком)
5. Maps: теория
6. Maps: задачи с кодом
7. Вопросы для защиты
8. Шпаргалка

---

## 1. Строки, байты, руны

### Главная идея
- **string**: неизменяемая (read-only) последовательность байтов. Внутри заголовок из указателя и длины.
- Строки обычно хранят текст в **UTF-8**: один символ занимает от 1 до 4 байт.
- **byte** = alias для `uint8`.
- **rune** = alias для `int32`, один Unicode code point.
- Отдельного типа `char` в Go нет.

### Задача: что выведет `len("世界")`?
**6.** `len` считает байты, а не символы. Каждый иероглиф в UTF-8 это 3 байта, 2 × 3 = 6.
Число символов: `utf8.RuneCountInString(s)` → **2**.

```go
s := "世界"
fmt.Println(len(s))                    // 6
fmt.Println(utf8.RuneCountInString(s)) // 2
```

### Неправильная итерация (`"hêllo"`)
```go
s := "hêllo"
for i := range s {
    fmt.Printf("position %d: %c\n", i, s[i])
}
```
Вывод:
```
position 0: h
position 1: Ã
position 3: l
position 4: l
position 5: o
```
Почему так:
1. `ê` это U+00EA, в UTF-8 два байта `c3 aa`.
2. `range s` с одной переменной даёт **байтовый индекс начала каждой руны**: 0, 1, 3, 4, 5. Индекс 2 пропущен: он внутри `ê`.
3. `s[i]` это **байт**, а не руна. Байт `0xc3` через `%c` печатается как `Ã` (U+00C3).

Таблица из слайда:

| | h | ê | | l | l | o |
|---|---|---|---|---|---|---|
| `[]byte(s)` | 68 | c3 | aa | 6c | 6c | 6f |
| индекс `i` | 0 | 1 | 2 | 3 | 4 | 5 |

### Правильно
```go
for i, r := range s {
    fmt.Printf("position %d: %c\n", i, r)
}
```
`range` по строке сам декодирует UTF-8. `r` это `rune`, `i` это байтовая позиция начала руны.

### Вопрос 1: как Go хранит Unicode и почему `s[i]` даёт байт?
- Строка хранится как байты в кодировке UTF-8. Исходный код Go тоже всегда UTF-8.
- Руна (`int32`) нужна, чтобы работать с code point отдельно от кодировки.
- `s[i]` возвращает байт, потому что UTF-8 имеет переменную ширину. Чтобы найти i-й **символ**, пришлось бы идти с начала строки, то есть за O(n). Индексация по байтам работает за O(1), без скрытых затрат.
- Если нужны символы, это делается явно: `range`, `[]rune(s)` или `unicode/utf8`.

### Вопрос 2: как получить вторую руну в читаемом виде?
```go
s := "hêllo"
r := []rune(s)             // декодируем UTF-8 в []int32
fmt.Println(string(r[1]))  // "ê"
fmt.Printf("%c\n", r[1])   // ê
```
Как это работает: `[]rune(s)` декодирует строку и **создаёт новый срез**, по 4 байта на руну (O(n) по времени и памяти). Затем `r[1]` это второй символ, а `string(r[1])` превращает руну обратно в строку.

Вариант без аллокации:
```go
_, size := utf8.DecodeRuneInString(s)          // размер первой руны в байтах
second, _ := utf8.DecodeRuneInString(s[size:]) // вторая руна
fmt.Printf("%c\n", second)
```

---

## 2. Trim-функции

| Функция | Что делает |
|---|---|
| `TrimLeft(s, cutset)` | убирает слева **любые** символы из набора |
| `TrimRight(s, cutset)` | убирает справа **любые** символы из набора |
| `Trim(s, cutset)` | то же с обеих сторон |
| `TrimPrefix(s, prefix)` | убирает **ровно этот префикс**, один раз |
| `TrimSuffix(s, suffix)` | убирает **ровно этот суффикс**, один раз |

Демонстрация (на защите просят показать вживую):
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.TrimRight("123oxo", "xo"))   // "123"  (набор символов x и o)
    fmt.Println(strings.TrimSuffix("123oxo", "xo"))  // "123o" (именно суффикс "xo")

    fmt.Println(strings.TrimLeft("xoxo123", "xo"))   // "123"
    fmt.Println(strings.TrimPrefix("xoxo123", "xo")) // "xo123"
}
```
**Ловушка:** второй аргумент `TrimLeft`/`TrimRight` это **набор символов**, а не подстрока. Поэтому `TrimRight("123oxo", "xo")` съедает `o`, `x`, `o`, пока не дойдёт до `3`.

---

## 3. Медленная конкатенация

Плохо:
```go
func concat(values []string) string {
    s := ""
    for _, v := range values {
        s += v
    }
    return s
}
```
Строки неизменяемы, поэтому каждый `+=` **выделяет новую строку и копирует всё накопленное**. Итого O(n²) по копированию и много мусора для GC.

Хорошо:
```go
func concat(values []string) string {
    var sb strings.Builder
    for _, v := range values {
        sb.WriteString(v)
    }
    return sb.String()
}
```
`strings.Builder` держит внутри `[]byte` и растит его с запасом (амортизированно O(n)). `String()` возвращает строку без копирования.

### Задача: показать `Grow`
```go
func concatGrow(values []string) string {
    total := 0
    for _, v := range values {
        total += len(v)
    }
    var sb strings.Builder
    sb.Grow(total) // одна аллокация нужного размера
    for _, v := range values {
        sb.WriteString(v)
    }
    return sb.String()
}
```
`Grow(n)` заранее резервирует место, внутренний буфер не перевыделяется: меньше аллокаций и копирований.

**Нюанс:** `strings.Builder` нельзя копировать после первого использования (panic), передавай его по указателю.

---

## 4. Конвертации `[]byte` ↔ `string`

### Что происходит в памяти
- `string(b)` и `[]byte(s)` **копируют данные**.
- Причина: строка неизменяема, срез байтов изменяем. Если бы они делили память, изменение среза «меняло» бы строку.
- Копия обычно идёт в куче, значит нагрузка на GC. Компилятор иногда это оптимизирует (небольшие данные, которые не «убегают», `m[string(b)]`, сравнения, `range`), но рассчитывать на это нельзя.
- Большинство I/O API (`io.Reader`, `os.ReadFile`, `net/http` и т.д.) работают с `[]byte`. Лишние конвертации туда-обратно это лишние аллокации.
- Пакет `bytes` содержит аналоги функций `strings`: `Split`, `Contains`, `Index`, `TrimSpace` и др.

```go
// плохо
s := string(bytes.TrimSpace([]byte(data)))
// хорошо
b := bytes.TrimSpace(data)
```

### Рефакторинг `processPayload`
Оригинал:
```go
func processPayload(data []byte) string {
    trimmed := string(bytes.TrimSpace(data))       // копия #1
    clean := strings.ReplaceAll(trimmed, "\r", "") // копия #2
    return clean
}
```
Оптимизированный:
```go
func processPayloadOptimized(data []byte) []byte {
    trimmed := bytes.TrimSpace(data) // возвращает подсрез, без копирования
    return bytes.ReplaceAll(trimmed, []byte("\r"), nil)
}
```
`bytes.TrimSpace` не копирует, а возвращает подсрез исходных данных. Остаётся одна аллокация внутри `bytes.ReplaceAll`.

Бонус: версия без аллокаций, но она **портит входной срез** (осознанный компромисс):
```go
func processPayloadInPlace(data []byte) []byte {
    trimmed := bytes.TrimSpace(data)
    out := trimmed[:0]
    for _, b := range trimmed {
        if b != '\r' {
            out = append(out, b)
        }
    }
    return out
}
```

### Бенчмарк `BenchmarkProcessPayload`
Файл `payload_test.go`:
```go
package main

import (
    "bytes"
    "strings"
    "testing"
)

func processPayload(data []byte) string {
    trimmed := string(bytes.TrimSpace(data))
    return strings.ReplaceAll(trimmed, "\r", "")
}

func processPayloadOptimized(data []byte) []byte {
    trimmed := bytes.TrimSpace(data)
    return bytes.ReplaceAll(trimmed, []byte("\r"), nil)
}

// Данные должны содержать \r, иначе strings.ReplaceAll вернёт строку без копирования
// и сравнение получится нечестным.
var payload = []byte("  " + strings.Repeat("some line of text\r\n", 200) + "  ")

var (
    sinkString string
    sinkBytes  []byte
)

func BenchmarkProcessPayload(b *testing.B) {
    b.Run("original", func(b *testing.B) {
        b.ReportAllocs()
        for i := 0; i < b.N; i++ {
            sinkString = processPayload(payload)
        }
    })
    b.Run("optimized", func(b *testing.B) {
        b.ReportAllocs()
        for i := 0; i < b.N; i++ {
            sinkBytes = processPayloadOptimized(payload)
        }
    })
}
```
Запуск:
```
go test -bench=ProcessPayload -benchmem
```
Смотри колонки `allocs/op` и `B/op`. Ожидаемо у оригинала 2 аллокации и примерно вдвое больше байт на операцию, у оптимизированной 1 аллокация. Точные числа возьми из своего запуска и называй их на защите. `sink`-переменные нужны, чтобы компилятор не выкинул вызов как неиспользуемый.

---

## 5. Maps: теория

### Базовые операции
```go
m := map[string]int{"one": 1, "two": 2}
fmt.Println(m["one"]) // чтение
m["two"] = 20         // обновление / вставка
delete(m, "one")      // удаление (если ключа нет, ничего не произойдёт)
```

### Как проверить наличие ключа
```go
v, ok := m["key"]
if ok { /* ключ есть */ }
```
Нужно, потому что для отсутствующего ключа возвращается **нулевое значение типа** (0, "", nil), и оно неотличимо от реально сохранённого нуля.

### Как устроена map внутри
- Для read/update/delete Go считает **хеш ключа**, по нему находит корзину (bucket / группу) и ищет слот внутри через **control words** (метаданные о слотах).
- В среднем операции O(1).
- В слайдах упомянуты **Swiss Tables**. Это новая реализация, начиная с Go 1.24. В старых версиях была классическая схема: корзины по 8 элементов, load factor ~6.5, overflow-корзины. На слайде оба описания смешаны, поэтому на защите уточни: «в старых версиях так, в новых Swiss Tables».
- При росте элементы **перераспределяются** по новым корзинам. Отдельная вставка в худшем случае может стоить до O(n), поэтому важна предварительная ёмкость.

### Оптимизация инициализации
```go
m := make(map[string]int, 1_000_000)
```
Второй аргумент это подсказка размера. Память выделяется сразу, поэтому нет многократного роста и перехеширования. Полезно, когда число элементов известно заранее.

---

## 6. Maps: задачи с кодом

### Задача 1: `scores` (Bob vs Charlie)
```go
package main

import "fmt"

func main() {
    scores := map[string]int{
        "Alice": 90,
        "Bob":   0, // сдал экзамен на 0
    }

    if _, ok := scores["Bob"]; ok {
        fmt.Println("Bob is in the map")
    } else {
        fmt.Println("Bob is missing")
    }

    if _, ok := scores["Charlie"]; ok {
        fmt.Println("Charlie is in the map")
    } else {
        fmt.Println("Charlie is missing")
    }
}
```
Вывод:
```
Bob is in the map
Charlie is missing
```
Смысл: `scores["Bob"]` и `scores["Charlie"]` оба дают 0, но только `ok` их различает.

### Задача 2: инвентарь
```go
package main

import "fmt"

func main() {
    inventory := map[string]int{
        "apples":  10,
        "bananas": 5,
    }

    fmt.Println("apples:", inventory["apples"]) // Read: 10

    inventory["bananas"] = 12   // Update
    inventory["oranges"] = 8    // Insert
    delete(inventory, "apples") // Delete

    fmt.Println(inventory) // map[bananas:12 oranges:8]
}
```
`fmt` печатает map с ключами по возрастанию, поэтому вывод детерминирован. Но **порядок обхода через `range` случайный**.

### Задача 3: Tour of Go, Exercise: Maps (WordCount)
```go
package main

import (
    "strings"

    "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    for _, w := range strings.Fields(s) {
        m[w]++
    }
    return m
}

func main() {
    wc.Test(WordCount)
}
```
`strings.Fields` режет строку по пробельным символам. `m[w]++` читает нулевое значение для нового слова (0) и увеличивает до 1, отдельная проверка наличия не нужна.

---

## 7. Вопросы для защиты

1. **Что такое строка в Go?** Неизменяемая последовательность байтов, обычно UTF-8.
2. **Чем rune отличается от byte?** `rune` = `int32` = code point, `byte` = `uint8` = один байт. Символ занимает 1–4 байта.
3. **Почему `len` не равен числу символов?** Он считает байты. Для символов: `utf8.RuneCountInString` или `len([]rune(s))`.
4. **`for i := range s` vs `for i := 0; i < len(s); i++`?** `range` декодирует руны и даёт индекс начала каждой, обычный цикл идёт по байтам.
5. **`TrimRight` vs `TrimSuffix`?** Набор символов против точной подстроки.
6. **Почему `+=` в цикле плохо?** Новая строка и копирование на каждой итерации, O(n²).
7. **Зачем `Builder.Grow`?** Предвыделяет буфер, убирает перевыделения.
8. **Что происходит при `string(b)` / `[]byte(s)`?** Копирование данных, потенциальная аллокация в куче.
9. **Что вернёт чтение отсутствующего ключа?** Нулевое значение типа, без паники.
10. **Можно ли писать в `nil` map?** Нет, будет panic. Читать можно.
11. **Потокобезопасна ли map?** Нет. Параллельная запись приводит к fatal error. Нужны `sync.Mutex`/`RWMutex` или `sync.Map`.
12. **Какие типы могут быть ключами?** Только comparable (числа, строки, bool, указатели, структуры и массивы из comparable). Срезы, map и функции нельзя.
13. **Порядок обхода?** Не гарантирован, намеренно рандомизирован.
14. **Зачем `make(map[K]V, n)`?** Избежать роста и перехеширования при известном числе элементов.
15. **Сложность операций?** O(1) в среднем.
16. **Что за «copy substrings to prevent memory leaks»?** Подстрока `s[a:b]` разделяет память с исходной строкой и удерживает её целиком. Чтобы этого избежать, используют `strings.Clone`. Тему разберут в следующих неделях.

---

## 8. Шпаргалка (key takeaways)

- Строка это последовательность байтов, руна это Unicode-символ.
- Для корректной итерации по символам используй `range`.
- `TrimRight` убирает набор символов, `TrimSuffix` точный суффикс.
- Для конкатенации используй `strings.Builder` (+ `Grow`, если размер известен).
- Избегай лишних конвертаций `[]byte` ↔ `string`: каждая это копия.
- `v, ok := m[key]` отличает «нет ключа» от «значение 0».
- `make(map[K]V, n)`, если размер известен заранее.
- Порядок обхода map случайный, map не потокобезопасна, запись в nil map паникует.
