/*
 * @Author: shengchao
 * @Date: 2023-07-24 16:36:49
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 13:46:36
 * @Description: --- 定义错误码 ---
 * @FilePath: /algo_proj/utils/errors.go
 */
package utils

import "errors"

var (
	ErrOom            = errors.New("out of memory")
	ErrListEmpty      = errors.New("list is empty")
	ErrQueueEmpty     = errors.New("queue is empty")
	ErrInputParameter = errors.New("input paramrters wrong")
)
