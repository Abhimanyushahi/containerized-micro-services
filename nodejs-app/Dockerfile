FROM node:14.15.4 as builder
ENV NODE_ENV=production
WORKDIR /app
COPY [ "package*.json*","./"]
RUN npm ci --production
USER node
COPY app.js .

FROM node:14.15.4-slim
USER node
WORKDIR /app
COPY --from=builder /app .
CMD ["node","app.js"]
