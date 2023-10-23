# 代码生成命令

新建工程
```bash
hz new -module hertz-layout -I idl -idl idl
```

更新接口
```bash
hz update -I idl -idl tools/idl/user.proto
hz update -I idl -idl tools/idl/oauth.proto
```

生成文档
```bash
swag init -o internal/docs
```


