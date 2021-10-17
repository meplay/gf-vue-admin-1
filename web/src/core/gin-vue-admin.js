/*
 * gin-vue-admin web框架组
 *
 * */
// 加载网站配置文件夹
import { register } from './global'

export default {
  install: (app) => {
    register(app)
    console.log(`
       欢迎使用 Gf-Vue-Admin
       当前版本:V2.4.5
       加群方式:微信：SliverHorn QQ群：1040044540
       默认自动化文档地址:http://127.0.0.1:${import.meta.env.VITE_SERVER_PORT}/swagger
       默认前端文件运行地址:http://127.0.0.1:${import.meta.env.VITE_CLI_PORT}
       如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gf-vue-admin.com/docs/coffee
    `)
  }
}