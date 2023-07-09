# chatgpt-arkose-token-api

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
      - GO_CHATGPT_API_ARKOSE_TOKEN_URL=http://chatgpt-arkose-token-api:8081
    depends_on:
      - chatgpt-proxy-server-warp
      - chatgpt-arkose-token-api
    restart: unless-stopped

  chatgpt-proxy-server-warp:
    container_name: chatgpt-proxy-server-warp
    image: linweiyuan/chatgpt-proxy-server-warp
    environment:
      - LOG_LEVEL=OFF
    restart: unless-stopped

  chatgpt-arkose-token-api:
    container_name: chatgpt-arkose-token-api
    image: linweiyuan/chatgpt-arkose-token-api
    environment:
      - TZ=Asia/Shanghai
      - CHATGPT_ARKOSE_SERVER=http://chatgpt-proxy-server-arkose
    depends_on:
      - chatgpt-proxy-server-arkose
    restart: unless-stopped

  chatgpt-proxy-server-arkose:
    container_name: chatgpt-proxy-server-arkose
    image: linweiyuan/chatgpt-proxy-server-arkose
    environment:
      - LOG_LEVEL=OFF
    restart: unless-stopped
```
