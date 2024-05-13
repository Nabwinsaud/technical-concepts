FROM node:20.0-alpine



WORKDIR /usr/src/app



COPY package*.json ./



RUN npm install -g pnpm

RUN npm install -g ts-node-dev


RUN pnpm i




# # Bundle app source
COPY . .


CMD [ "ts-node-dev","index.ts" ]