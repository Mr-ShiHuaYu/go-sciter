# FFmpeg官网
[https://ffmpeg.org/](https://ffmpeg.org/)

------

# 图片
## 图片格式转换
```
指令：
ffmpeg -i input.jpg output.png
说明：
input.jpg: 原图
output.png：转换后的格式
```

------

# 视频
## 给视频添加封面
```
指令：
ffmpeg -i demo.mp4 -i logo.jpg -map 0 -map 1 -c copy  -disposition:v:1 attached_pic output.mp4
说明：
ffmpeg -i input.mp4 -i cover.jpg \
-map 0 -map 1 \          # 合并视频和图片流
-c copy \                # 复制原始流不重新编码
-disposition:v:1 attached_pic \  # 标记图片为封面
output.mp4
```

