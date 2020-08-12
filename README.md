# decompress-run-length-encoded-list

## 題目解讀：

### 題目來源:
[decompress-run-length-encoded-list](https://leetcode.com/problems/decompress-run-length-encoded-list/)

### 原文:
We are given a list nums of integers representing a list compressed with run-length encoding.

Consider each adjacent pair of elements [freq, val] = [nums[2*i], nums[2*i+1]] (with i >= 0).  For each such pair, there are freq elements with value val concatenated in a sublist. Concatenate all the sublists from left to right to generate the decompressed list.

Return the decompressed list.

### 解讀：
給定一個正整數陣列nums 長度是2的倍數

其中 對於每2個 元素個別代表 出現的頻率 跟 數值

[freq val] = [nums[2 * i] nums[2 * i+1]] 其中 i > 0

舉例來說：

[1 2 3 4] => 1 [1 2] => [2]  值為2 出現次數1
             2 [3 4] => [4 4 4] 值為4 出現次數3
所以最後結果為 [2 4 4 4]
求給定的nums 的 還原陣列 

## 初步解法:
### 初步觀察:
對於每個nums
其中需要每次讀取兩個元素來做處理

然後每次需要產生新的array

最後再join成新的陣列
### 初步設計:
given an integer array nums

step 0: let an integer array result = []

step 1: let an integer length = length of nums /2, a integer idx = 0

step 2: if idx > length go to step 6

step 3: let value = nums[2 * idx + 1], create a integer array extract with length of nums[2 * idx]

step 4: loop array extract for index = 0 to index < length of extract set extract[index] = value

step 5: result = append(result, extract...) 

step 6: return result


## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package decompress_list

func decompressRLElist(nums []int) []int {
	result := []int{}
	length := len(nums) / 2
	for idx := 0; idx < length; idx++ {
		val := nums[2*idx+1]
		extract := make([]int, nums[2*idx])
		for i := range extract {
			extract[i] = val
		}
		result = append(result, extract...)
	}
	return result
}

```
## 測資的撰寫
```golang
package decompress_list

import (
	"reflect"
	"testing"
)

func Test_decompressRLElist(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{2, 4, 4, 4},
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 1, 2, 3},
			},
			want: []int{1, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decompressRLElist(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decompressRLElist() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang ironman 30 day 11th day](https://hackmd.io/o5F2CiZeSIycxnJHYN_NXw?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)