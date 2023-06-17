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

| Argument                    | Description                                      | Default     |
| :-------------------------- | :----------------------------------------------- | :---------- |
| `--path`                    | Path where project will be build                 | `.`       |
| `--framework`               | Define framework. Values: React and Next         | `React`       |
| `--styles`                  | Define the library for styles. Values: ["TailwindCss", "StyledComponents"]   | `TailwindCss`|
| `--package-manager`         | Define framework. Values: ["npm", "yarn", "pnpm"]| `pnpm`|
| `--use-linter`              | Define framework.                                | | `-` |
| `--typescript`              | Use typescript instead of javascript.            | `-`|
| `--use-backend`             | Define that project has a Back-end               | `-`|
 

