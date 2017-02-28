関数には果たしてオブジェクトを渡した方がいいか、値を渡したほうがいいか
=====================================

1. 値を渡す
2. ポインタ構造体
3. 構造体

```
$ go test -bench . -benchmem
BenchmarkVariable-8    	 5000000	       227 ns/op	      64 B/op	       3 allocs/op
BenchmarkPtrStruct-8   	 5000000	       222 ns/op	      64 B/op	       3 allocs/op
BenchmarkStructr-8     	 5000000	       219 ns/op	      64 B/op	       3 allocs/op
```

おもったより差はなかったし、実行毎に結果がブレる。このレベルは大きな差はうまれないのか
