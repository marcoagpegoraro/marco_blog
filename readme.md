
# Marco's Blog

My personal blog, it was coded by me using Go with Fiber, Gorm, AWS S3, Django Templates, FlatifyCSS and deploy at Railway.

Fell free to clone the project and see the code for yourself, maybe you can make some modifications and use this blog for yourself too 😉


## Article explaining how this blog was made

https://www.marcoagpegoraro.com.br/posts/15


## Authors

- [@marcoagpegoraro](https://www.github.com/marcoagpegoraro)


## Demo

https://www.marcoagpegoraro.com.br

## Deployment

To run this project with hot reload, run:

```bash
  CompileDaemon --command="./marco_blog" -include=Makefile -include="*.django" -include="*.css" -include="*.js"
```
To deploy this project run

```bash
  go build -ldflags '-s -w' -o app  
```

