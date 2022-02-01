官方文档：https://developer.fyne.io/ <br>
参考api: http://note.tyss.site/1633252 

#### 解决中文乱码
```go
func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simhei.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}
```