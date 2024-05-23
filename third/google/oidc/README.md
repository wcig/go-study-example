# Google OIDC 示例

1. 浏览器打开 URL: http://localhost:28082/web/google/integration
2. 确认使用 google 账号登录
3. 登录成功，服务端接口获取到用户 idToken 信息，然后重定向至成功页面
4. 成功页面展示 idToken 的 jwt 解析后的 header 和 payload
5. 服务端日志打印 idToken 解析后的结果

参考:

* [Google login | Google for Developers](https://developers.google.com/identity/gsi/web/guides/overview?hl=zh-cn)
* [网站侧接入google一键登录 - 掘金 (juejin.cn)](https://juejin.cn/post/7281437377511637050)
