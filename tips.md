# tips.md

## fmt

### Format

```go
	// %+v 構造体の場合にフィールド名を出力する
	//
	// 1
	// {Transport:<nil> CheckRedirect:<nil> Jar:<nil> Timeout:0s}
	fmt.Printf("%+v\n", 1)
	fmt.Printf("%+v\n", http.Client{})

	// %#v 値のGoの文法での表現を出力する
	//
	// 1
	// http.Client{Transport:http.RoundTripper(nil), CheckRedirect:(func(*http.Request, []*http.Request) error)(nil), Jar:http.CookieJar(nil), Timeout:0}
	fmt.Printf("%#v\n", 1)
	fmt.Printf("%#v\n", http.Client{})

	// %T 値の型のGoの文法での表現を出力する
	//
	// int
	// http.Client
	fmt.Printf("%T\n", 1)
	fmt.Printf("%T\n", http.Client{})
```

* [fmt.Printfなんかこわくない - Qiita](https://qiita.com/rock619/items/14eb2b32f189514b5c3c)

## Numerical

### math/big

#### Over int64

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

#### Big binary digits

大きな2進数の文字列(a,b)を扱う。big.Intを基数を指定して文字列にする場合はTextを使う。String()は内部的にはText(10)を呼んでいる。

```
	i1, _ := new(big.Int).SetString(a, 2)
	i2, _ := new(big.Int).SetString(b, 2)
	i3 := new(big.Int).Add(i1, i2)
	return i3.Text(2)
```

* https://pkg.go.dev/math/big#Int.Text

## String

### Small things

文字列を1文字ずつのスライスにする

```go
	cs := strings.Split(s, "")
```

### String <> Int

```go
	i, _ := strconv.ParseInt("111", 2, 64)
	fmt.Println(i) // 7
```

```go
	s := strconv.FormatInt(7, 2)
	fmt.Println(s) // 111
```

### Reverse

文字列を反転させる

```go
package main

import "fmt"

func main() {
	// 界世 ,olleH
	fmt.Println(Reverse("Hello, 世界"))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
```
* https://seven-901.hatenablog.com/entry/2021/06/14/234000

## Array / Slice / Map

### Small Things

Mapの定義

```go
	sm := map[string]string{}
```

```go
	sm := map[string]int{}
```

### Queue

雑でもよければこんな感じ。

```go
	q := make([]int, 0)

	// push
	q = append(q, 1)
	q = append(q, 2)

	// pop
	v1 := q[0]
	q = q[1:]
	fmt.Println(v1, q)

	// pop
	v2 := q[0]
	q = q[1:]
	fmt.Println(v2, q)
```

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