/*
 * @Author: your name
 * @Date: 2022-04-04 12:35:37
 * @LastEditTime: 2022-04-05 19:31:04
 * @LastEditors: Please set LastEditors
 * @Description: Deterministic Finite Automaton 确定有穷自动机
 * @FilePath: /dfa/dfa.go
 */
package dfa

import (
	"fmt"
	"sync"
)

// DFA工具 用于敏感词汇校验
type DfaToolsForWord interface {
	AddData(data []string)                         // 添加数据
	Init()                                         // 整理数据并初始化
	CheckWord(word string) bool                    // 校验词汇 返回是否命中
	CheckWordAndShield(word string) (bool, string) // 校验词汇并屏蔽对应词汇
}

// 定义一个构建器
type Director struct {
	dfaTool DfaToolsForWord
}

// 构造DFA工具
func (d *Director) Construct(data []string) {
	d.dfaTool.AddData(data)
	d.dfaTool.Init()
}

func NewConstruct(bui DfaToolsForWord) *Director {
	return &Director{
		dfaTool: bui,
	}
}

type DfaForWord struct {
	data []string   // 目标数据
	trie *Trie      // 字典树
	lock sync.Mutex // 读写锁 防止功能扩展后的一些冲突
}

/**
 * @description: 写入数据
 * @param {[]string} data
 * @return {*}
 */
func (dfaWord *DfaForWord) AddData(data []string) {
	dfaWord.data = data
}

/**
 * @description: 初始化工具并分析数据
 * @param {*}
 * @return {*}
 */
func (dfaWord *DfaForWord) Init() {
	// 初始化Trie
	dfaWord.trie = NewTrie()
	// 分析数据 将数据转换成Trie
	// 非递归添加
	for _, element := range dfaWord.data {
		dfaWord.trie.Add(element)
	}

	// // 递归添加
	// for _, element := range dfaWord.data {
	// 	dfaWord.trie.AddByRec(element)
	// }
}

/**
 * @description: 校验是否匹配 并返回校验结果
 * @param {string} word
 * @return {*}
 */
func (dfaWord *DfaForWord) CheckWord(word string) bool {
	target, found := dfaWord.trie.Check(word)
	fmt.Print("%r,%r", target, found)
	return true
}

/**
 * @description: 校验是否匹配 并返回校验结果 且进行敏感词屏蔽
 * @param {string} word
 * @return {*}
 */
func (dfaWord *DfaForWord) CheckWordAndShield(word string) (bool, string) {
	// ok, start, end := dfaWord.trie.IsPrefix((word))
	// return ok, fmt.Sprintf("%d $ %d", start, end)
	return true, "123"
}
