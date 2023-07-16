# chatgpt-arkose-token-api

---

依赖 `Firefox`
无法单独运行，需配合 [linweiyuan/chatgpt-proxy-server-arkose](https://github.com/linweiyuan/docker-services/tree/main/linweiyuan/chatgpt-proxy-server-arkose)
一起使用，资源占用较大

---

本地调试环境变量

`PLAYWRIGHT_NODEJS_PATH=/usr/bin/node`

---

```yaml
services:
  go-chatgpt-api:
    container_name: go-chatgpt-api
    image: linweiyuan/go-chatgpt-api
    ports:
      - 8080:8080
    environment:
      - TZ=Asia/Shanghai
      - GO_CHATGPT_API_PROXY=socks5://chatgpt-proxy-server-warp:65535
      - GO_CHATGPT_API_ARKOSE_TOKEN_URL=http://chatgpt-proxy-server-arkose:8081/token
    depends_on:
      - chatgpt-proxy-server-warp
      - chatgpt-proxy-server-arkose
    restart: unless-stopped

  chatgpt-proxy-server-warp:
    container_name: chatgpt-proxy-server-warp
    image: linweiyuan/chatgpt-proxy-server-warp
    restart: unless-stopped

  chatgpt-proxy-server-arkose:
    container_name: chatgpt-proxy-server-arkose
    image: linweiyuan/chatgpt-proxy-server-arkose
    restart: unless-stopped
```
