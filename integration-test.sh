#!/bin/bash

# 设置博客系统的基础URL
BASE_URL="http://localhost:8080"

# 测试首页是否可访问
curl_response=$(curl -s -o /dev/null -w "%{http_code}" "${BASE_URL}/")

if [ "$curl_response" != "200" ]; then
  echo "Error: Home page is not accessible. Expected HTTP status code 200, but got ${curl_response}."
  exit 1
else
  echo "Home page is accessible. HTTP status code: ${curl_response}."
fi

# 测试博客文章列表
curl_response=$(curl -s -o /dev/null -w "%{http_code}" "${BASE_URL}/posts")

if [ "$curl_response" != "200" ]; then
  echo "Error: Posts page is not accessible. Expected HTTP status code 200, but got ${curl_response}."
  exit 1
else
  echo "Posts page is accessible. HTTP status code: ${curl_response}."
fi

# 测试获取单篇博客文章
# 假设第一篇文章的ID是1
curl_response=$(curl -s -o /dev/null -w "%{http_code}" "${BASE_URL}/posts/1")

if [ "$curl_response" != "200" ]; then
  echo "Error: Cannot get post with ID 1. Expected HTTP status code 200, but got ${curl_response}."
  exit 1
else
  echo "Successfully retrieved post with ID 1. HTTP status code: ${curl_response}."
fi

# 测试创建新的博客文章
# 请根据你的博客系统API进行调整
# 这里只是一个示例，可能需要根据你的API定义和参数进行更详细的设置
new_post_data='{"title": "Test Post", "content": "This is a test post."}'
curl_response=$(curl -s -o /dev/null -w "%{http_code}" -X POST -H "Content-Type: application/json" -d "${new_post_data}" "${BASE_URL}/posts")

if [ "$curl_response" != "201" ]; then
  echo "Error: Cannot create new post. Expected HTTP status code 201, but got ${curl_response}."
  exit 1
else
  echo "Successfully created new post. HTTP status code: ${curl_response}."
fi

echo "All integration tests passed successfully."
exit 0