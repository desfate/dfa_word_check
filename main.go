/*
 * @Author: your name
 * @Date: 2022-04-04 14:42:58
 * @LastEditTime: 2022-04-05 19:38:19
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /dfa/main.go
 */
package main

import (
	"Github/data"
	"Github/dfa"
)

func main() {
	dfaword := dfa.DfaForWord{}
	dfa.NewConstruct(&dfaword).Construct(data.GetBadWord())

	// print(dfaword.CheckWordAndShield("shit"))
	print(dfaword.CheckWord("shitfalun"))
	// print(dfaword.CheckWordAndShield("shit324342shit"))
	// print(dfaword.CheckWordAndShield("sdfshite23423"))
	print("hello world")
}
