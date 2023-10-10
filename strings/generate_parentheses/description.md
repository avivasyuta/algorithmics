# Генерация валидной скобочной последовательности (один вид скобок)

[leetcode](https://leetcode.com/problems/generate-parentheses/)

Сложность: Средняя

## Оглавление

- [Описание](#description)
- [Брутфорс решение](#solution_1)
- [Оптимизированный брутфорс](#solution_2)

---

## <a name="description"></a>Описание

Напишите функцию, которая будет генерировать все возможные валидные скобочные последовательности длиной `2n`.

Входные данные: int — количество пар скобок.

Выходные данные: массив строк — все возможные валидные скобочные последовательности длиной `2n`.

### Ограничения

- 1 <= n <= 8

### Примеры

#### Пример 1

```
Вход: 3
```

```
Ответ: ["((()))","(()())","(())()","()(())","()()()"]
```

#### Пример 2

```
Вход: 1
```

```
Ответ: ["()"]
```

## <a name="solution_1"></a> Брутфорс решение

Давайте разобьем задачу на 2 подзадачи:
- Генерация всех возможных строк длиной 2n, состоящих из символов «(» и «)»
- Проверка строки на то, что она является валидной скобочной последовательностью


### Генерация всех возможных строк длиной 2n, состоящих из символов «(» и «)»

Достаточно тривиальная задача:
- Создаем исходный массив. В нашем случае из одного элемента, так как валидная скобочная последовательность может начинаться только с открывающей скобки.
- Запускаем цикл до `2 * n`, где `n` — количество пар скобок.
- В каждой итерации цикла проходимся по всем элементам предыдущего результата и добавляем 2 новых варианта с открывающей и закрывающей скобками.

```go
func generateAllPossibleParenthesis(n int) []string {
    combinations := make([]string, 0, 100)
    combinations = append(combinations, "(")
    
    for i := 1; i < n*2; i++ {
        newArray := make([]string, 0, len(combinations))
        for _, item := range combinations {
            newArray = append(newArray, item+"(")
            newArray = append(newArray, item+")")
        }
        
        combinations = newArray
    }
    
    return combinations
}
```
### Проверка строки на валидную скобочную последовательность

Эту задачу мы уже разбирали [разбирали](https://github.com/avivasyuta/algorithmics/blob/main/stack/valid_parentheses/description.md#solution).
Чтобы не повторяться, возьмем готовое решение:

```go
func isValid(brackets string) bool {
    stack := make([]rune, 0, len(brackets))
    
    closeBracket := map[rune]rune{
    ')': '(',
    ']': '[',
    '}': '{',
    }
    
    for _, char := range brackets {
        if openBracket, isClose := closeBracket[char]; isClose {
            if len(stack) == 0 {
                return false
            }
            
            lastOpenBracket := stack[len(stack)-1]
            
            if lastOpenBracket != openBracket {
                return false
            }
            
            stack = stack[:len(stack)-1]
        } else {
        stack    = append(stack, char)
        }
    }
    
    return len(stack) == 0
}
```

### Оценка сложности

**По времени**

Оптимальная сложность проверки строки на правильность — `O(m)`, где `m` — длина строки (в нашем случае все строки имеют длину `2n`).
Проверка на валидность вызывается `k` раз, где k — количество сгенеренных строк.

Давайте оценим кол-во сгенерированных строк.
Каждый элемент строки кроме первого — произвольный символ из множества ["(", ")"], а количество символов в итоговой строке `2 * n`.
Таким образом, кол-во строк равно <code>2<sup>2n-1</sup></code>.

Итоговую сложность мы можем оценить как <code>n\*(2<sup>2n-1</sup>)</code>.

**По памяти**

Самый большой расход памяти в нашем случае — хранение всех возможных строк.
Для этого нам понадобится массив строк, размером <code>2<sup>2n-1</sup></code>, каждый элемент которого равен `2 * n`.
То есть мы можем говорить о <code>2n\*2<sup>2n-1</sup></code> дополнительной памяти.

### Комментарий

Чаще всего, в задачах с комбинаторикой решение с полным перебором всех возможных вариантов является сильно не оптимальным, так как скорость роста сложности решения относительно размера входных данных очень высока.
Так в этой задаче при n = 8 кол-во всех возможных вариантов строк — **32768** (<code>2<sup>15</sup></code>).


## <a name="solution_2"></a> Оптимизированный брутфорс

Основная проблема брутфорс решения — генерация большого количества заведомо невалидных строк.
Это можно исправить если мы перед добавлением очередного символа добавим проверки, откидывающие заведомо невалидные символы.

Задача сильно упрощается тем, что в условии сказано, что скобки могут быть только одного типа — круглые.
В этом случае мы можем использовать для валидации последовательности обычный каунтер, а не стек. Давайте опишем правила, когда мы можем добавлять разные скобки:
- Открывающую скобку мы можем добавить, если суммарное количество открытых скобок в строке не превышает числа `n`.
- Закрывающую скобку мы можем добавить, если перед ней есть незакрытая открывающая скобка (будем отслеживать с помощью счетчика открытых скобок).

Для того чтобы отслеживать эти кейсы, введем вспомогательную структуру.

```go
type parenthesisWithCounter struct {
    str              string // генерированная строка
    generalOpenCount int // суммарное кол-во открытых скобок в строке (не должно превышать n)
    notClosedCount   int // кол-во незакрытых скобок в строке
}
```

Теперь мы можем сделать функцию проверки, которая будет отбрасывать невалидные строки в момент попытки добавления.

```go
func possibleToAddParenthesis(parenthesis rune, prev parenthesisWithCounter, maxOpenCount int) bool {
	if parenthesis == '(' {
		if prev.generalOpenCount >= maxOpenCount {
			return false
		}

		return true
	}

	if parenthesis == ')' {
		if prev.notClosedCount <= 0 {
			return false
		}

		return true
	}

	return false
}
```

Теперь мы можем изменить функцию по генерации строк из брутфорс решения так, чтобы она сразу возвращала только валидные последовательности.

```go
for i := 1; i < n*2; i++ {
    newArray := make([]parenthesisWithCounter, 0, len(parenthesisWithCounters)*2)
    for _, item := range parenthesisWithCounters {
        if possibleToAddParenthesis('(', item, n) {
            newArray = append(newArray, parenthesisWithCounter{
                str:              item.str + "(",
                generalOpenCount: item.generalOpenCount + 1,
                notClosedCount:   item.notClosedCount + 1,
            })
        }
        
        if possibleToAddParenthesis(')', item, n) {
            newArray = append(newArray, parenthesisWithCounter{
                str:              item.str + ")",
                generalOpenCount: item.generalOpenCount,
                notClosedCount:   item.notClosedCount - 1,
            })
        }
    }
    
    parenthesisWithCounters = newArray
}
```

Более того, в этом решении не нужно постфактум проверять получившиеся скобки на валидность, так что мы сможем отбросить один из множителей в нашей оценке сложности.

### Оценка сложности

**По времени**

В данном решении не нужна валидация последовательности. Ее мы заменили на проверку допустимости добавления открывающей/закрывающей скобки в момент генерации строки. 
Более того, за счет сохранения счетчика суммарного количества открытых и счетчика незакрытых скобок мы можем проводить эту проверку за константное время.

Количество сгенерированных строк равно <code>2<sup>2n-1</sup></code>.

> :❕: На самом деле количество строк сильно меньше, так как мы заранее отбрасываем все невалидные скобочные последовательности.
> Но, оценить их количество сильно сложнее. Валидно будет сказать, что в худшем случае сложность <code>2<sup>2n-1</sup></code>.

**По памяти**

Самый большой расход памяти в нашем случае — хранение всех возможных строк.
Для этого нам понадобится массив строк размером <code>2<sup>2n-1</sup></code>, каждый элемент которого равен `2 * n`.
То есть мы можем говорить о <code>2n\*2<sup>2n-1</sup></code> дополнительной памяти.

> :❕: На самом деле количество строк сильно меньше, так как мы заранее отбрасываем все невалидные скобочные последовательности.
> Но, оценить их количество сильно сложнее. Валидно будет сказать, что в худшем случае сложность <code>2<sup>2n-1</sup></code>.