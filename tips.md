# tips.md

## Numerical

### Over int64

int64を超えた数値を扱いたい場合は math/big を使うが NewInt の引数がint64となっている。

```go
func NewInt(x int64) *Int
```

そのため SetString を利用する

```go
func (z *Int) SetString(s string, base int) (*Int, bool)
```

こんな感じになる。大きな数値を文字列からbig.Intで定義して、それに1を加え、文字列に戻している。

```go
	num, _ := new(big.Int).SetString(s, 10)
	num.Add(num, big.NewInt(1))
	s = num.String()
```

* [Golangのmath/bigでIntを直接入力したい - 逆さまにした](https://cipepser.hatenablog.com/entry/2017/04/15/100914)
* https://pkg.go.dev/math/big#NewInt
* https://pkg.go.dev/math/big#Int.SetString


## String

文字列を1文字ずつのスライスにする

```go
	cs := strings.Split(s, "")
```

## Slice

### Binary Search

```go
func binarySearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var index int
	for {
		if end < start {
			return -1
		}
		index = (start + end) / 2

		if nums[index] == target {
			return index
		}
		if nums[index] < target {
			start = index + 1
		} else {
			end = index - 1
		}
	}
}
```