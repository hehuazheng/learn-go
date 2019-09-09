### 泛型
1. Go不支持泛型，只能用 interface{} 模拟返回的类型安全需要用户自己保证，
2. .(type) 的类型必须匹配
3. interface{} 是运行时泛型，性能没有编译
</pre>

```
func tuple(condition bool, a , b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// 调用方法, 参照StructDemo2
var s string = tuple(true, "nil", "not nil").(string)
```

### 匿名struct <=> 隐性引入匿名struct中的所有属性

### json 解析

