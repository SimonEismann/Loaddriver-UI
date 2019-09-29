FROM node:8

WORKDIR /usr/src/app

COPY package*.json ./
RUN ["npm", "install"]

COPY . .
RUN ["npm", "run", "build"]

ENV PORT 80
ENV API_ROOT "http://localhost"
ENV CONSOLE_URI "ws://localhost/jobs/current/output"
EXPOSE ${PORT}

CMD ["npm", "run", "serve"]
