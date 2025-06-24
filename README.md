# Learning repository

Just a repo to record learning go lang for fun.

## How to add a new project to the mono repository

```bash
mkdir xx-name-of-project
cd xx-name-of-project
go mod init github.com/dlebee/learning-go/xx-name-of-project
go work use .
```