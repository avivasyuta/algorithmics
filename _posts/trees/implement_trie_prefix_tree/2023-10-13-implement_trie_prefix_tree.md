---
layout: post
title:  "Имплементировать префиксное дерево (Trie)"
complexity: medium
original_url: https://leetcode.com/problems/implement-trie-prefix-tree/description/
tag: medium
---

[wiki](https://ru.wikipedia.org/wiki/%D0%9F%D1%80%D0%B5%D1%84%D0%B8%D0%BA%D1%81%D0%BD%D0%BE%D0%B5_%D0%B4%D0%B5%D1%80%D0%B5%D0%B2%D0%BE)

## Оглавление

- [Описание задачи](#описание-задачи)
- [Ограничения](#ограничения)
- [Пример](#пример)
- [Решение](#решение)

---

## Описание задачи

Необходимо имплементировать интерфейс префиксного дерева:
- Функция Insert(word string) - вставка слова в дерево
- Search(word string) bool - проверка наличия слова в дереве 
- StartsWith(prefix string) bool - проверка наличия префикса в дереве

---

## Ограничения

- Длина слов/префиксов в диапазоне от 1 до 2000 символов
- Слова и префиксы состоят только из латинских букв в нижнем регистре
- Не более 30000 вызовов функций в тесте

---

## Пример

```go
tree := Constructor()

tree.Insert("apple")
tree.Search("apple")   // return True
tree.Search("app")     // return False
tree.StartsWith("app") // return True
tree.Insert("app")
tree.Search("app")     // return True
```

---

## Решение

Для решения данной задачи нужно знать что такое префиксное дерево и как работать с рекурсией (для реализации вставки значений и обхода по дереву).

Нода дерева в данном случае будет выглядеть следующим образом:

```go
type Trie struct {
	children   map[rune]*Trie
	isFullWord bool
}
```

Каждая нода префиксного дерева (кроме рутовой) представляет собой букву латинского алфавита. Само значение буквы хранить в ноде не нужно, так как значение будет ключом в хеш-мапе связей.  

Свойство isFullWord - флаг, который показывает, является ли текущая нода конечной для слова.

children — хеш-мапа со связями между текущей и дочерними нодами дерева. В качестве ключа используется буква латинского алфавита (значение дочерней ноды).

Тогда, если мы разложим слова acid, ai, aid, aide, aim и air в нашу структуру, в итоге мы получим следующее.

```go
Trie{
    children: map[rune]*Trie{
        'a': {
            children: map[rune]*Trie{
                'c': {
                    children: map[rune]*Trie{
                        'i': {
                            children: map[rune]*Trie{
                                'd': {
                                    isFullWord: true,
                                },
                            },
                        },
                    },
                },
                'i': {
                    children: map[rune]*Trie{
                        'd': {
                            children: map[rune]*Trie{
                                'e': {
                                    isFullWord: true,
                                },
                            },
                            isFullWord: true,
                        },
                        'm': {
                            isFullWord: true,
                        },
                        'r': {
                            isFullWord: true,
                        },
                    },
                    isFullWord: true,
                },
            },
        },
    },
}
```

Для вставки слова в префиксное дерево реализуем вспомогательный рекурсивный метод.

```go
func (t *Trie) insert(word []rune) {
	if len(word) == 0 {
		t.isFullWord = true
		return
	}

	child, exist := t.children[word[0]]
	if !exist {
		child = &Trie{
			children: make(map[rune]*Trie),
		}

		t.children[word[0]] = child
	}

	child.insert(word[1:])
}
```

Функции Search и StartsWith отличаются между собой только тем, что в Search у последней ноды префикса флаг isFullWord должен быть `true`.
Поэтому, для реализации обеих функций мы можем воспользоваться простой вспомогательной рекурсивной функцией обхода дерева.

Итоговое решение.

{% include code-example.md go_path="go/solution.go" ts_path="ts/solution.ts" %}

#### Оценка сложности

Так как задача - имплементировать структуру данных, мы будем оценивать сложность методов Insert, Search и StartsWith.

**По времени**

Функции Insert, Search и StartsWith используют рекурсивные вспомогательные функции `insert` и `traverse`. У обеих функций глубина рекурсии равна длине префикса.

Таким образом все 3 функции имеют сложность `O(n)`, где `n` - длина префикса.

**По памяти**

Для работы функций не используются промежуточные структуры, зависящие от длины префикса. Сложность по памяти `O(1)`.