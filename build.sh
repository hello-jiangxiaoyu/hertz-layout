
#!/bin/bash

if [ -d "output/" ]; then
    rm -rf output
fi
mkdir -p output/pkg/conf


# 下载依赖
echo "Start download package"
go mod download
if [ $? -ne 0 ]; then
    echo "Failed to download Go modules."
    exit 1
fi

# 构建
echo "Start build server"
go build -o server .
if [ $? -ne 0 ]; then
    echo "Failed to build server."
    exit 2
fi


# 打包
echo "Start make zip package"
mv server output/server
cp ./pkg/conf/conf.online.yaml ./output/pkg/conf/conf.online.yaml

tar -czvf hertz-layout.tar.gz -C output .
