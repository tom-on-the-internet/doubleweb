# Double Web - A few things I like about Golang

Double Web is a repo used for a lightning talk at Humi. It's a sample code repo.

## Running

```bash
go install github.com/tom-on-the-internet/doubleweb@latest && doubleweb
```

or

```bash
docker run --rm -it -p 3000:3000 golang bash -c 'go install github.com/tom-on-the-internet/doubleweb@latest && doubleweb'
```

## A few things I like about Go

Note: I am a novice Gopher.
Note: I like Go, but no language is perfect.

1. Can compile to a single statically linked binary.

1. Fast.

1. Has an incredible standard library. So, I have fewer dependencies.

1. Comes with a production-ready web server.

1. Allows embedding of files into binary.

1. Has its own templating language.

1. Strongly typed. Statically typed. I can almost always jump to definition.

1. Native support for concurrency.

1. Testing is built into the language. And it's parallel. And my tests can be generated for me.
