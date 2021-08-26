## GIN-VUE-ADMIN插件开发规范

### 请勿提交任何具有危险性的代码

### 详细介绍

插件分类（pr时请标注）
- 前端功能类插件
- 前端美化类插件
- 后端工具性插件
- 后端中间件
- 业务性插件 (完整的后端业务链条，配备前端vue静态文件页面)
- 其他

### 开发规范

插件需按照gin-vue-admin基本目录结构书写 插件前端部分如果需要则在views、api、utils等需要到对应功能的目录下分别创建同插件名的文件例如  必须配备readme 使用说明

```
  web
    --src
      --api
        --plugs.js
      --view
        --plugsTable.vue
        --plugsForm.vue
      --utils
        --plugs.js
```

后端开发需要在plugin下需创建等同于gin-vue-admin后端目录的结构并安放你开发的插件代码  必须配备readme 使用说明

```
  server
    --plugin
      --你的插件名字
        --api
        --你需要的功能包 对应gva的目录结构和功能结构即可
```

### 加入开发者需求

向gin-vue-admin的plugins分支提交一款插件且通过pr被采纳即可被拉入成为本仓库开发者

将向gin-vue-admin提交的插件拆分出来 配备成套的readme在本仓库建立分支平且推送到自己的插件分支即可
