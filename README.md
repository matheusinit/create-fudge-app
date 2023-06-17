## create-fudge-app

create-fudge-app is a tool builder to boilerplate common used stacks to not waste so much time in configuring project and configuring ESLint, Typescript, etc.

### How to use

 + See version

 ```bash
  create-fudge-app version 
 ```

 + Create project boilerplate (interactive)

 ```bash
  create-fudge-app init
 ```

 + Create project boilerplate (non-interactive)

 ```
 create-fudge-app --framework reactjs --styles tailwindcss
 ```

### Arguments

| Argument                    | Description                                      | Shorthand | Default     |
| :-------------------------- | :----------------------------------------------- |:--:|:----------:|
| `--path`                    | Path where project will be build                 | `-p` | `.`       |
| `--fe-framework`            | Define framework. Values: React and Next         | `None` | `React`       |
| `--styles`                  | Define the library for styles. Values: ["TailwindCss", "StyledComponents"] | `-s` | `TailwindCss`|
| `--package-manager`         | Define framework. Values: ["npm", "yarn", "pnpm"]| `-pm` | `pnpm`|
| `--use-linter`              | Define framework.                                | `-ul` | `-` |
| `--typescript`              | Use typescript instead of javascript.            | `-t` | `-` |
| `--use-backend`             | Define that project has a Back-end               | `-b` | `-`|
| `--use-frontend`            | Define that project has a Front-end | `-f` | `-`|
| `--be-framework`            | Define back-end framework. Values: ["Express", "Fastify", "NestJS"] | `None` | `None` |
| `--orm`                     | Define ORM to use. Values: ["Prisma", "Drizzle", "TypeORM"] | `-o` | `None` |
| `--be-folder-name`          | Define name for back-end folder | `None` | `server-app` |
| `--be-folder-name`          | Define name for front-end folder | `None` | `web` |

