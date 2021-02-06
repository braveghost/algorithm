package _01_200

import "math"

//字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列：
// 序列中第一个单词是 beginWord 。
// 序列中最后一个单词是 endWord 。
// 每次转换只能改变一个字母。
// 转换过程中的中间单词必须是字典 wordList 中的单词。
// 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，找到从 beginWord 到 endWord 的 最短转换序列 中
//的 单词数目 。如果不存在这样的转换序列，返回 0。
//
//
// 示例 1：
//输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","lo
//g","cog"]
//输出：5
//解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
//
//
// 示例 2：
//输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
//输出：0
//解释：endWord "cog" 不在字典中，所以无法进行转换。
//
//
//
// 提示：
// 1 <= beginWord.length <= 10
// endWord.length == beginWord.length
// 1 <= wordList.length <= 5000
// wordList[i].length == beginWord.length
// beginWord、endWord 和 wordList[i] 由小写英文字母组成
// beginWord != endWord
// wordList 中的所有字符串 互不相同
//
// Related Topics 广度优先搜索
// 👍 689 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func LadderLength(beginWord string, endWord string, wordList []string) int {
		l := len(wordList)
		if l == 0 {
			return 0
		}
		wordListMap := map[string]bool{}
		for _, str := range wordList {
			wordListMap[str] = true
		}
		type node struct {
			word string
			res  int
		}
		chars := "abcdefghijklmnopqrstuvwxyz"
		queue := []*node{{word: beginWord, res: 1}}
		for len(queue) != 0 {
			n := queue[0]
			queue = queue[1:]
			if n.word == endWord {
				return n.res
			}
			for i, c := range n.word {
				var word string
				for _, tmp := range chars {
					if tmp == c {
						continue
					}
					word = n.word[0:i] + string(tmp) + n.word[i+1:]
					if wordListMap[word] {
						queue = append(queue, &node{word: word, res: n.res + 1})
	                    delete(wordListMap,word)
					}
				}
			}
		}
		return 0
}

// todo 虚拟节点、图
func ladderLength(beginWord string, endWord string, wordList []string) int {
	 wordIds := map[string]int{}
	var graph [][]int

	addWord := func(word string) int {
		id, ok := wordIds[word]
		if !ok {
			id = len(wordIds)
			wordIds[word] = id
			graph = append(graph, []int{})
		}
		return id
	}

	buildGraph := func(word string) int{
		id := addWord(word)
		for i := range word{
			tmp := word[0:i] + "*"+  word[i+1:]
			id2 := addWord(tmp)
			graph[id] = append(graph[id], id2)
			graph[id2] = append(graph[id2], id)
		}
		return id
	}

	for _, l := range wordList{
		buildGraph(l)
	}
	firstId := buildGraph(beginWord)
	endId, ok :=wordIds[endWord]
	if !ok{
		return 0
	}

	dict := make([]int, len(wordIds))

	for i:= range dict{
		dict[i] = math.MaxInt64
	}
	dict[firstId] = 0
	queue := []int{firstId}
	for len(queue) != 0{
		val := queue[0]
		if val == endId{
			return dict[val]/2+1
		}
		queue = queue[1:]

		for _, w := range graph[val]{
			if dict[w] == math.MaxInt64{
				dict[w] =  dict[val] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0

}

