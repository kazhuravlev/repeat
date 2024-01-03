# repeat

This is a very simple tool that allow you to repeat some command several times.

## Installation

```shell
go install github.com/kazhuravlev/repeat/cmd/repeat@latest
```

## Usage

Repeat `curl https://example.com` 10 times sequentially.

```shell
repeat 10 curl https://example.com
```

Repeat `date` 10 times sequentially with 5 seconds delay between each repeat.

```shell
repeat 10 5s date
```

## Corner cases

Repeat non-exists/bad program will repeat it even when not exists.

```shell
repeat 10 not-exists-program
```
