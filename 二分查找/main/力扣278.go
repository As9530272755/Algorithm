/*
力扣 278

你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。

假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。

你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

示例:

给定 n = 5，并且 version = 4 是第一个错误的版本。

调用 isBadVersion(3) -> false
调用 isBadVersion(5) -> true
调用 isBadVersion(4) -> true

所以，4 是第一个错误的版本。 

题解：
1.已知发行版本是从 1-n 版，由此定义最低和最高值
2.low <= higt 定义查找范围，因为不排除最后一个是错误版本所以一直查找到 higt 最后一个版本
3.就检查中间的值
4.调用 isBadVersion() 函数，如果为中就表示错误版本在 low 这边，由此 higt 递减
5.假如中间这个值为 false 也就是错误版本在 higt 这边，由此 low 递增
6.最后返回 low 的值，因为 low 向 higt 方向移动，遵循 左闭右开 的原则，所以 low 就是我们要找的那个值

*/



func firstBadVersion(n int) int {
    low := 1
    higt := n
    
    for low <= higt {
        mid := (low + higt) / 2 
        if isBadVersion(mid)  {
            higt = mid - 1
        } else {
            low = mid + 1
        }
    }
    return low
}
