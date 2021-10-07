# wispeeer博客生成器

+ 目标：自定义博客生成流程

## Build
```
make
```

## Usage
```bash
./bin/wispeeer -h
```

## Theme 
[wisper](https://github.com/Wispeeer/wisper)

## 默认渲染规则
+ 文章详情  Article     *.html ---> /webiste/*.html
+ 文章列表  []Article   index.html ---> /index.html, [2,3].html ---> /articles/*

默认规则：
```
2 {{ .PublicDir }}/{{ .PaginationDir }}/1.html {{ .PublicDir }}/index.html
```
第一列序号：渲染模式
第二列路径：愿渲染文件路径
第三列路径：修改后的目标路径
如果想要修改默认渲染规则，可以在主题文件夹下创建rule.txt，来自定义渲染规则。

## mode
+ 1. 文章详情页
+ 2. 页面详情页：如about
+ 3. 文章列表页

## tree public
```
public
├── index.html
├── about
│   └── index.html
├── articles
│   ├── 2.html
│   └── 3.html
├── links
│   └── index.html
└── website
    ├── Markdown_1.0.1.html
    ├── Wispeeer_User_Guide
    │   └── screenshot.png
    └── Wispeeer_User_Guide.html

5 directories, 9 files
```
