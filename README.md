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
| `--fe-framework`            | Define framework. Values: ["reactjs", "nextjs"]        | `None` | `reactjs`       |
| `--styles`                  | Define the library for styles. Values: ["TailwindCss", "StyledComponents"] | `-s` | `tailwindcss`|
| `--package-manager`         | Define framework. Values: ["npm", "yarn", "pnpm"]| `-pm` | `pnpm`|
| `--use-linter`              | Define framework.                                | `-lt` | `-` |
| `--use-typescript`              | Use typescript instead of javascript.            | `-ts` | `-` |
| `--be-framework`            | Define back-end framework. Values: ["Express", "Fastify", "NestJS"] | `None` | `None` |
| `--version`                 | See the current version of cli app | `-v` | `None` |
| `--help`                    | Get help from the cli app | `-h` | `None` |

### TO-DO

- [ ] Do a complete front-end boilerplate com ReactJS, TailwindCSS, Eslint
