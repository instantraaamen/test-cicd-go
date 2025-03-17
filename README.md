# 🚀 test-cicd-go

_Environment_

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?logo=go&logoColor=white)
![Debian](https://img.shields.io/badge/debian-%23A81D33.svg?logo=debian&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?logo=docker&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-%232088FF.svg?logo=github-actions&logoColor=white)

## 🛠️ Tool

| *Tool* | *Explanation* |
|-------------|------|
| ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?logo=go&logoColor=white) | v1.22 |
| ![Debian](https://img.shields.io/badge/debian-%23A81D33.svg?logo=debian&logoColor=white) | bookworm-slim |
| ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?logo=docker&logoColor=white) | multistage build |
| ![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-%232088FF.svg?logo=github-actions&logoColor=white) | CI/CD |


## 📦 Third-Party Library

- 🧪 [testify](https://github.com/stretchr/testify)  
    Test

- 📚 [samber/io](https://github.com/samber/io)  
    File I/O


## 🚦 Quick Start

```bash
git clone https://github.com/instantraaamen/test-cicd-go.git
cd test-cicd-go

# docker start
docker build -t test-cicd-go .
docker run --rm test-cicd-go

# local start
go run ./main.go
'''
