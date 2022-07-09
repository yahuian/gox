## errors

```go
func main() {
	fmt.Println(f())
}

func f() error {
	if err := foo(); err != nil {
		return err
	}
	if err := bar(); err != nil {
		return err
	}
	return nil
}

func foo() error {
	return errors.New("something error")
}

func bar() error {
	return errors.New("something error")
}
```

如此情况根本无法区分 ``something error`` 是哪个函数抛出的。


``fmt.Errorf("context: %w", err)`` 虽然可以增加上下文，形成错误链，但如果增加的上下文也相同了，依旧无法区分错误，这种本质还是手动行为。

## https://github.com/pkg/errors

由于 [Wrap() duplicates call stack](https://github.com/pkg/errors/issues/242) 堆栈重复打印的问题，必须使用 ``errors.WithMessage`` 方法本质还是手动行为。

## errorx

自动记录每一个 err 产生时的函数名，文件路径以及行号，即使不用 ``WrapMsg`` 也能定位到整个错误链，具体示例可以看测试文件。