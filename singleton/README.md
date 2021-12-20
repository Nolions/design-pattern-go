# singleton 單例

在kotlin一樣提供一個keyword `object` 可以讓我們很快速實作singleton，而在Golnag中標準函式庫Sync中的Once也可以讓我們很快速在Golang實作singleton

```Go
once.Do(func() {
    // TODO new instance
})
```
