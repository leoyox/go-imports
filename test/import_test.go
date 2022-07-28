/*
 * @Author: leoking
 * @Date: 2022-07-28 15:28:56
 * @LastEditTime: 2022-07-28 17:09:25
 * @LastEditors: your name
 * @Description:
 */
package test

import (
	"testing"

	sc "github.com/leoyox/go-imports/submodule_c"

	sb "github.com/leoyox/go-imports/submodule_b"
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
