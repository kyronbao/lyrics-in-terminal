# 命令行终端显示播放歌词

## 功能
可以在编辑器的命令行播放歌词

## 用法
编译
```
go build -o lyric main.go
```
把可执行程序移到环境变量所在目录  
例如linux环境
```
mv lyric /usr/bin/
```
播放显示示例歌词
```
lyric demo(secret_base).lrc
```
![使用演示](./show.gif)

ps.  
单曲循环需要保证歌词的最后一行的时间和歌曲时间大小一致

## 歌词下载地址推荐
 https://www.22lrc.com/