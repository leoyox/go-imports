/*
 * @Author: leoking
 * @Date: 2022-07-28 15:28:56
 * @LastEditTime: 2022-07-28 16:49:19
 * @LastEditors: your name
 * @Description:
 */
package test

import (
	"testing"

	sb "github.com/leoyox/go-imports/submodule_b"

	sc "github.com/leoyox/go-imports/submodule_c"
)

func TestSubmoduleC(t *testing.T) {
	sc.Print()
}

/**
 * @description:
	事实证明有模块嵌套时，自模块也需要以绝对路径声明才行
 * @param {*testing.T} t
 * @return {*}
*/
func TestSubmoduleB(t *testing.T) {
	sb.Print()
}
