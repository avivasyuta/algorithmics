package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie_Insert(t *testing.T) {
	testCases := []struct {
		name     string
		in       []string
		expected Trie
	}{
		{
			name: "empty_tree",
			expected: Trie{
				children: make(map[rune]*Trie),
			},
		},
		{
			name: "one_world",
			in:   []string{"app"},
			expected: Trie{
				children: map[rune]*Trie{
					'a': {
						children: map[rune]*Trie{
							'p': {
								children: map[rune]*Trie{
									'p': {
										children:   make(map[rune]*Trie),
										isFullWord: true,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "several_worlds",
			in:   []string{"apple", "app", "bio"},
			expected: Trie{
				children: map[rune]*Trie{
					'a': {
						children: map[rune]*Trie{
							'p': {
								children: map[rune]*Trie{
									'p': {
										children: map[rune]*Trie{
											'l': {
												children: map[rune]*Trie{
													'e': {
														children:   make(map[rune]*Trie),
														isFullWord: true,
													},
												},
											},
										},
										isFullWord: true,
									},
								},
							},
						},
					},
					'b': {
						children: map[rune]*Trie{
							'i': {
								children: map[rune]*Trie{
									'o': {
										children:   make(map[rune]*Trie),
										isFullWord: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tree := Constructor()
			for _, word := range testCase.in {
				tree.Insert(word)
			}

			assert.Equal(t, testCase.expected, tree)
		})
	}
}

func TestTrie_StartsWith(t *testing.T) {
	testCases := []struct {
		name          string
		wordsToInsert []string
		expected      map[string]bool
	}{
		{
			name: "empty_tree",
			expected: map[string]bool{
				"ap":    false,
				"app":   false,
				"appl":  false,
				"apple": false,
				"b":     false,
				"kio":   false,
			},
		},
		{
			name:          "one_word",
			wordsToInsert: []string{"app"},
			expected: map[string]bool{
				"ap":    true,
				"app":   true,
				"appl":  false,
				"apple": false,
				"b":     false,
				"kio":   false,
			},
		},
		{
			name:          "several_word",
			wordsToInsert: []string{"app", "apple", "bio"},
			expected: map[string]bool{
				"ap":    true,
				"app":   true,
				"appl":  true,
				"apple": true,
				"b":     true,
				"kio":   false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tree := Constructor()
			for _, word := range testCase.wordsToInsert {
				tree.Insert(word)
			}

			for prefix, expected := range testCase.expected {
				exist := tree.StartsWith(prefix)
				assert.Equal(t, expected, exist)
			}
		})
	}
}

func TestTrie_Search(t *testing.T) {
	testCases := []struct {
		name          string
		wordsToInsert []string
		expected      map[string]bool
	}{
		{
			name: "empty_tree",
			expected: map[string]bool{
				"ap":    false,
				"app":   false,
				"appl":  false,
				"apple": false,
				"b":     false,
				"kio":   false,
			},
		},
		{
			name:          "one_word",
			wordsToInsert: []string{"app"},
			expected: map[string]bool{
				"ap":    false,
				"app":   true,
				"appl":  false,
				"apple": false,
				"b":     false,
				"kio":   false,
			},
		},
		{
			name:          "several_word",
			wordsToInsert: []string{"app", "apple", "bio"},
			expected: map[string]bool{
				"ap":    false,
				"app":   true,
				"appl":  false,
				"apple": true,
				"b":     false,
				"bio":   true,
				"kio":   false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tree := Constructor()
			for _, word := range testCase.wordsToInsert {
				tree.Insert(word)
			}

			for prefix, expected := range testCase.expected {
				exist := tree.Search(prefix)
				assert.Equal(t, expected, exist)
			}
		})
	}
}
