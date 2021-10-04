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
