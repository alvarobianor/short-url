# 🔗 ShortURL - URL Shortener Service

A simple and efficient URL shortener service built with Go. This service allows you to create shortened URLs and redirect to the original URLs.

## ✨ Features

- 🚀 Fast and lightweight
- 🔒 Simple and secure
- 📝 RESTful API
- 🔄 URL redirection
- 🎯 Unique code generation
- 📊 JSON responses

## 🛠️ Installation

1. Clone the repository:

```bash
mkdir short-url
cd short-url

git clone https://github.com/alvarobianor/short-url.git
```

2. Install dependencies:

```bash
go mod download
```

3. Run the server:

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## 📡 API Endpoints

### Create Short URL

```http
POST /v1/create
Content-Type: application/json

{
    "url": "https://example.com"
}
```

### Get Original URL

```http
GET /v1/{code}
```

## 📝 License

MIT License

Copyright (c) 2024 Álvaro Bianor

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
