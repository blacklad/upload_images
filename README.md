# typora上传图片
a go client for typora, you can upload your images to alioss

## mac使用方法
1. 下载代码
2. 将代码中的 conf/conf_example.yaml 复制到用户目录（/home/user/）下并改名为 ./upload.conf
3. 找到代码中的upload_images, 复制本地文件路径
4. 将路径配置到typora中, 并测试


## conf文件内容
```yaml
// oss相关配置
oss:
  endpoint: oss-cn-beijing.aliyuncs.com
  key: aa
  secret: bb

// 图片上传到oss的目录
ossBasePath: test/

// 生成链接的域名
ossDomain: https://image.blacklad.com
```

## 其他平台
自行编译代码测试使用~