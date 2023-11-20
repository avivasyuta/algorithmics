export class Trie {
    children: Map<string, Trie>
    isFullWord: boolean

    constructor(children?: Map<string, Trie>, isFullWord?: boolean) {
        this.children = children ?? new Map()
        this.isFullWord = isFullWord ?? false
    }

    insert(word: string): void {
        debugger
        if (word.length === 0) {
            this.isFullWord = true
            return
        }

        let child = this.children.get(word[0])

        if (!child) {
            child = new Trie()
            this.children.set(word[0], child)
        }

        child.insert(word.slice(1))
    }

    search(word: string): boolean {
        const node = this.traverse(word)
        if (!node) {
            return false
        }

        return node.isFullWord
    }

    startsWith(prefix: string): boolean {
        const node = this.traverse(prefix)

        return node !== null
    }

    private traverse(prefix: string): Trie | null {
        if (prefix.length === 0) {
            return this
        }

        const child = this.children.get(prefix[0])

        if (!child) {
            return null
        }

        return child.traverse(prefix.slice(1))
    }
}