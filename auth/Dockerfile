FROM node:12.14.1

WORKDIR /usr/src/app
COPY package*.json ./
RUN yarn install
COPY . .

CMD ["node", "index.js"]