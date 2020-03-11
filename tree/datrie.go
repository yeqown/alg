// trie 前缀树
// http://www.cnblogs.com/huangxincheng/archive/2012/11/25/2788268.html

/*
前缀树有如下特点：
	1. root节点不包含字符
	2. 从根节点到某一节点，路径上经过的字符连接起来，就是该节点对应的字符串。
	3. 每个单词的公共前缀作为一个字符节点保存。
使用场景：
	1. 词频统计（节省内存）
	2. 前缀匹配（高效）
*/

// ！！！！！！！！ 实现动态前缀树！！！！！！！！只针对英文

package tree

// "fmt"

// DATrie ...
// TODO:
type DATrie struct {
	Base  []int // 表示该位置的状态
	Check []int // Check[i]表示该状态的前一状态，用于检查状态转移的正确
}

// NewDATrie ...
func NewDATrie(words ...string) *DATrie {
	root := &DATrie{}
	root.Base[0] = 1
	return root
}

// CharCode ...
func CharCode(char byte) int {
	return int(char) - 97
}

// Add ...
func (d *DATrie) Add(word string) {
	code := CharCode(word[0])
	base := d.Base[0] + code

	for i, c := range word {
		// 处理叶子节点
		if i == len(word)-1 {

		}

		offset := CharCode(byte(c)) + d.Base[base]
		pos := d.Base[1] + offset

		// 如果该位置为空
		if d.Base[pos] == 0 {

		}
		// 出现冲突
	}
}

// func (d *DATrie) AllocateBase() int {
// 	dbBase := make([]int, 2*len(d.Base))
// 	// dbBase := append()
// }

// Delete ...
func (d *DATrie) Delete(word string) {

}

// Search ...
func (d *DATrie) Search(prefix string) {

}
