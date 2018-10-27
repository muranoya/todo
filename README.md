# todo
`todo` is simple Todo manager written in Golang.

I make this application to study Golang, I study Golang package, slice, basic grammar, how to use basic Golang package and how to write unit-test in Golang through developing this application.

## Feature

- Add TODO with title, description and created time.
- Show all TODO, incompleted TODO.
- Change status of TODO, from incomplete to done, from done to incomplete.
- Great portability. Only one execute file and one saving json file.

# Installation

Install Golang through apt, brew, [official package](https://golang.org/) or other package manager.

And `git clone` from github.com.

```bash
git clone git@github.com:muranoya/todo
```

or

```bash
git clone https://github.com/muranoya/todo.git
```

Build with Makefile

```bash
cd todo && make
```

# Usage

Add new TODO.

```bash
td add -msg "Read books"
```

```bash
td add -msg "Read books" -detail "地球の歩き方"
```

Show incompleted TODOs.

```bash
td show
1 Read books:  (Created: 2018-10-27 21:52)
2 Read books: 地球の歩き方 (Created: 2018-10-27 21:52)
```

Show all TODOs.

```bash
td show -all
1 Read books:  (Created: 2018-10-27 21:52)
2 Read books: 地球の歩き方 (Created: 2018-10-27 21:52)
```

Make incomplete TODO to completed TODO.

```bash
td done -id 1
td show -all
1 Done	Read books:  (Created: 2018-10-27 21:52)
2 Read books: 地球の歩き方 (Created: 2018-10-27 21:52)
```

Make completed TODO to incompleted TODO.

```bash
td done -unset 1
td show -all
1 Read books:  (Created: 2018-10-27 21:52)
2 Read books: 地球の歩き方 (Created: 2018-10-27 21:52)
```

Delete completed TODOs.

```bash
td done -clean
```

Help

```bash
td add -h
td show -h
td done -h
```

