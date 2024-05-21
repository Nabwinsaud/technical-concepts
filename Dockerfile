FROM node:20.0-alpine



WORKDIR /usr/src/app



COPY package*.json ./

RUN npm install -g pnpm

RUN pnpm i



RUN pnpm build
# # Bundle app source
COPY . .


CMD [ "node","dist/index.js" ]