# Adapted from example at https://docs.astro.build/en/recipes/docker/
FROM ghcr.io/pnpm/pnpm:11 AS base
WORKDIR /app

# Install Node.js via pnpm's bundled runtime manager (no corepack needed/available here)
RUN pnpm runtime set node 22 -g

# By copying only the package.json and pnpm-lock.yaml here, we ensure that the following `-deps` steps are independent of the source code.
# Therefore, the `-deps` steps will be skipped if only the source code changes.
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./

FROM base AS prod-deps
RUN pnpm install --frozen-lockfile --prod

FROM base AS build-deps
RUN pnpm install --frozen-lockfile

FROM build-deps AS build
COPY . .
RUN pnpm run build

FROM base AS runtime
COPY --from=prod-deps /app/node_modules /app/node_modules
COPY --from=build /app/dist /app/dist

ENV HOST=0.0.0.0
ENV PORT=4321
EXPOSE 4321
CMD ["node", "./dist/server/entry.mjs"]
