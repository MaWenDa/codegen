# codegen

#### 介绍

集成gin、gorm、mysql、text/template实现生成go web代码，提高开发效率。

#### 软件架构

软件架构说明

#### 目录说明

- codegen
    - logs 日志
    - db 数据库初始化SQL
    - pkg 公共工具包
    - api API公共参数
    - config 配置文件
    - commands 启动HTTP服务
    - router 路由
    - model 数据模型(结构体)
    - dao 数据库操作
    - service 业务服务
    - server 服务层
    - main.go 启动程序
    - go.mod go依赖
    - codegen-api.json 接口文档

#### 安装教程

1. go mod download

#### 使用说明

1. 将 db 目录下的sql文件导入到mysql中
2. 修改config/app.yml配置文件
3. 运行main.go启动服务
4. 使用数据源配置管理添加数据源
5. 根据数据源ID获取数据表列表
6. 输入表名生成代码

#### 参与贡献

1. Fork 本仓库
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request

#### 特技

1. 使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2. Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3. 你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4. [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5. Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6. Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
