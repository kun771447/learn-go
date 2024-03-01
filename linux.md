

- 获取端口和 pid
```
lsof -i -P -n | grep LISTEN
```

- 关闭端口

```
kill pid
```