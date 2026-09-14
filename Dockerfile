FROM ghcr.io/pnpm/pnpm:12

RUN pnpm runtime set node 24 -g

WORKDIR /app
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./

RUN pnpm install --frozen-lockfile

COPY . .
ENV NODE_ENV=production

EXPOSE 3000
ENV HOSTNAME="0.0.0.0"

CMD ["pnpm", "start"]
