/**
* @Author: fushisanlang
* @Date: 2020/5/30 2:23 下午
 */

package until

//错误输出
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
