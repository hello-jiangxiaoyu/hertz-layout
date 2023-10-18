
#!/bin/bash
RUN_NAME=server
mkdir -p output/bin
cp tools/script/* output 2>/dev/null
cp internal/conf/conf.online.yaml  output/internal/conf/conf.online.yaml  2>/dev/null

chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}
