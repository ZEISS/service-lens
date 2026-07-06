FROM node:lts-alpine

WORKDIR /app  

COPY package.json package-lock.json ./

RUN npm install --only=production
 
COPY . .
ENV NODE_ENV=production
 
EXPOSE 3000
ENV HOSTNAME="0.0.0.0"
  
CMD ["npm", "start"]  