FROM node:17-alpine

WORKDIR /app
COPY ./public ./public
COPY ./src/ ./src
COPY ./scripts/ ./scripts
COPY package.json ./
COPY rollup.config.js ./
RUN npm install
RUN npm run build
EXPOSE 8080
ENV HOST=0.0.0.0
CMD [ "npm", "start" ]