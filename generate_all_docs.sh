#!/bin/bash

# 设置项目路径和输出目录
PROJECT_PATH="/mnt/c/Users/Administrator/GolandProjects/GoAllInOne"
OUTPUT_DIR="./docs"

# 创建输出目录
mkdir -p $OUTPUT_DIR

# 启动 godoc 服务
godoc -http=:6060 &
GODOC_PID=$!

# 等待 godoc 服务启动
sleep 2

# 抓取项目根目录的文档并保存为 HTML 文件
ROOT_URL="http://localhost:6060/pkg/GoAllInOne/"
ROOT_OUTPUT_FILE="$OUTPUT_DIR/index.html"
curl -s $ROOT_URL > $ROOT_OUTPUT_FILE

if [ $? -eq 0 ]; then
  echo "Documentation for project root has been saved to $ROOT_OUTPUT_FILE"
else
  echo "Failed to generate documentation for project root"
  rm -f $ROOT_OUTPUT_FILE
fi

# 遍历 pkg 目录下的所有子目录（包）
for PACKAGE in $(find $PROJECT_PATH/pkg -mindepth 1 -type d); do
  # 获取相对路径
  RELATIVE_PACKAGE=${PACKAGE#$PROJECT_PATH/}

  # 构建 URL 和输出文件名
  PACKAGE_URL="http://localhost:6060/pkg/GoAllInOne/pkg/$RELATIVE_PACKAGE/"
  OUTPUT_FILE="$OUTPUT_DIR/${RELATIVE_PACKAGE//\//_}.html"

  # 抓取文档并保存为 HTML 文件
  curl -s $PACKAGE_URL > $OUTPUT_FILE

  if [ $? -eq 0 ]; then
    echo "Documentation for $RELATIVE_PACKAGE has been saved to $OUTPUT_FILE"
  else
    echo "Failed to generate documentation for $RELATIVE_PACKAGE"
    rm -f $OUTPUT_FILE
  fi
done

# 停止 godoc 服务
kill $GODOC_PID

echo "All documentation has been saved to $OUTPUT_DIR"
