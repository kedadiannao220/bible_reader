## 圣经阅读命令行工具

### 后台接口

```
FindBookList
FindTextByBookAndChapter
FindTextByBookAndChapterAndVerse
FindFuzzyText
FindFuzzyBookText
FindFuzzyBookAndChapterText
ConvertBookID
```

### 命令列表

- book
  - -l list bible book(bool)

- text

  - -b book (string)

  - -c chapter (int)

  - -v verse (int)

  - -q like query (string)

### 样式输出
样式输出引入了tabwrite的库

### TODO
- 字体大小设置
- copy当前经节的圣经
- 制作经文卡片
- 加载一些图片库



