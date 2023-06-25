FROM node:alpine
WORKDIR /app
COPY package.json .
COPY index.js .
RUN npm i
EXPOSE 65526
CMD [ "node", "index.js" ]
