import { Trie } from "./solution";

describe('Trie_insert', () => {
    const suits: {
        name: string
        in: string[]
        expected: Trie
    }[] = [
        {
            name: "empty_tree",
            expected: new Trie(),
            in: [],
        },
        {
            name: "one_word",
            in: ["app"],
            expected: new Trie(
                new Map([
                    ['a', new Trie(
                        new Map([
                            ['p', new Trie(
                                new Map([
                                    ['p', new Trie(
                                        new Map(),
                                        true,
                                    )],
                                ]),
                            )],
                        ]),
                    )],
                ]),
            ),
        },
        {
            name: "several_worlds",
            in: ["apple", "app", "bio"],
            expected: new Trie(
                new Map([
                    ['a', new Trie(
                        new Map([
                            ['p', new Trie(
                                new Map([
                                    ['p', new Trie(
                                        new Map(
                                            new Map([
                                                ['l', new Trie(
                                                    new Map([
                                                        ['e', new Trie(
                                                            new Map(),
                                                            true,
                                                        )],
                                                    ]),
                                                )],
                                            ]),
                                        ),
                                        true,
                                    )],
                                ]),
                            )],
                        ]),
                    )],
                    ['b', new Trie(
                        new Map([
                            ['i', new Trie(
                                new Map([
                                    ['o', new Trie(
                                        new Map(),
                                        true,
                                    )],
                                ]),
                            )],
                        ]),
                    )],
                ]),
            )
        },
    ]

    suits.forEach(suit => {
        test(suit.name, () => {
            const tree = new Trie()

            suit.in.forEach((word) => {
                tree.insert(word)
            })

            expect(tree).toEqual(suit.expected);
        })
    })
})

describe('Trie_startsWith', () => {
    const suits: {
        name: string
        wordsToInsert: string[]
        expected: Map<string, boolean>
    }[] = [
        {
            name: "empty_tree",
            wordsToInsert: [],
            expected: new Map([
                ["ap", false],
                ["app", false],
                ["appl", false],
                ["apple", false],
                ["b", false],
                ["kio", false],
            ]),
        },
        {
            name: "one_word",
            wordsToInsert: ["app"],
            expected: new Map([
                ["ap", true],
                ["app", true],
                ["appl", false],
                ["apple", false],
                ["b", false],
                ["kio", false],
            ]),
        },
        {
            name: "several_word",
            wordsToInsert: ["app", "apple", "bio"],
            expected: new Map([
                ["ap", true],
                ["app", true],
                ["appl", true],
                ["apple", true],
                ["b", true],
                ["kio", false],
            ]),
        },
    ]

    suits.forEach(suit => {
        test(suit.name, () => {
            const tree = new Trie()

            suit.wordsToInsert.forEach((word) => {
                tree.insert(word)
            })

            suit.expected.forEach((expected, prefix) => {
                const isExists = tree.startsWith(prefix)
                expect(isExists).toEqual(expected);
            })
        })
    })
})

describe('Trie_search', () => {
    const suits: {
        name: string
        wordsToInsert: string[]
        expected: Map<string, boolean>
    }[] = [
        {
            name: "empty tree",
            wordsToInsert: [],
            expected: new Map([
                ["ap", false],
                ["app", false],
                ["appl", false],
                ["apple", false],
                ["b", false],
                ["kio", false],
            ]),
        },
        {
            name: "one word",
            wordsToInsert: ["app"],
            expected: new Map([
                ["ap", false],
                ["app", true],
                ["appl", false],
                ["apple", false],
                ["b", false],
                ["kio", false],
            ]),
        },
        {
            name: "several words",
            wordsToInsert: ["app", "apple", "bio"],
            expected: new Map([
                ["ap", false],
                ["app", true],
                ["appl", false],
                ["apple", true],
                ["b", false],
                ["bio", true],
                ["kio", false],
            ]),
        },
    ]

    suits.forEach(suit => {
        test(suit.name, () => {
            const tree = new Trie()

            suit.wordsToInsert.forEach((word) => {
                tree.insert(word)
            })

            suit.expected.forEach((expected, prefix) => {
                const isExists = tree.search(prefix)
                expect(isExists).toEqual(expected);
            })
        })
    })
})