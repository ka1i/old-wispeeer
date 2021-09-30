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
注：对于页面模板，你也可以直接创建像这样：`about.html`模板，这样程序会转而使用首先此模板，而不再使用`page.html`

主题需包含以下模板
```
index.html # 文章列
post.html # 文章渲染模板
page.html # 页面：比如about，渲染模板
```

## tree public
public
├── index.html
├── about
│   └── index.html
├── articles
│   ├── index.html
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