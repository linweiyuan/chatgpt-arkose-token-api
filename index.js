// noinspection SpellCheckingInspection

const http = require('http');
const fun = require('funcaptcha');

http.createServer(async (req, res) => {
  console.log('...')
  const token = await fun.getToken({
    pkey: "35536E1E-65B4-4D96-9D97-6ADB7EFF8147",
    surl: "https://tcr9i.chat.openai.com",
    headers: {
      "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
    },
    site: "https://chat.openai.com"
  });
  res.write(JSON.stringify(token));
  res.end();
}).listen(65526);
