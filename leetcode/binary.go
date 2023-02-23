package leetcode

//给你两个整数 n 和 start。你的任务是返回任意 (0,1,2,,...,2^n-1) 的排列 p，并且满足：
//p[0] = start
//p[i] 和 p[i+1] 的二进制表示形式只有一位不同
//p[0] 和 p[2^n -1] 的二进制表示形式也只有一位不同

//输入：n = 2, start = 3
//输出：[3,2,0,1]
//解释：这个排列的二进制表示是 (11,10,00,01)
//所有的相邻元素都有一位是不同的，另一个有效的排列是 [3,1,0,2]

func CircularPermutation(n int, start int) []int {
	ans := make([]int, 1, 1<<n)
	ans[0] = start
	for i := 1; i <= n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ((ans[j]^start)|(1<<(i-1)))^start)
		}
	}
	return ans

}

// 给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。
//
//示例 1：
//
//输入: num = 1775(11011101111)
//输出: 8

func ReverseBits(num int) int {
	return num
}
