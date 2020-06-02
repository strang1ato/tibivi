<div align="center">
  <img src="/docs/logo-530x200.png">
  <p>Terminal based, inspired by vi/vim time blocking tool</p>
  <br>
  <a href="https://goreportcard.com/report/github.com/oltarzewskik/tibivi">
    <img src="https://goreportcard.com/badge/github.com/oltarzewskik/tibivi">
  </a>
  <a href="http://makeapullrequest.com/">
    <img src="https://img.shields.io/badge/PRs-welcome-brightgreen">
  </a>
  <br>
  <br>
  <br>
</div>

<div align="center">
  <img src="/docs/demo.gif">
</div>

## Advantages of tibivi over other time blocking tools

- ### Easy to use

  If you know basic vi/vim features you already know how to use tibivi

- ### Speed

  Tibivi works almost instantly because of being terminal based and being written in fast Go programming language

- ### All data is stored human-readable text files instead of database

  All data is stored in text files in `$HOME/.tibivi/` directory, which means that you can modify your schedule in any text editor

  Example content of datafile:
  ```
  14:00-17:00 Contribute to FOSS 💻
  17:00-18:00 Play chess ♟️
  18:00-19:00 Play some computer game 🎮
  ```

- ### Use emoji instead of images

  In tibivi you use emoji instead of images in order to arrange any time block

## Installation

### Binary

Download latest binary from [releases](https://github.com/oltarzewskik/tibivi/releases)

then set `0755` permission to file by for example:
```bash
   chmod a+x <path-to-tibivi>
```

and move tibivi binary to directory in `$PATH`


### From source

Make sure you have Go 1.14 installed and `$GOPATH/bin` added to `$PATH`, then execute:
```bash
  go get -u github.com/oltarzewskik/tibivi
```

## Keybindings

### Schedule

| Keybinding                | Action                                |
| --------------------------|---------------------------------------|
| <kbd>h</kbd> <kbd>l</kbd> | Navigate in selection of day schedule |

### Bar

| Keybinding         | Action       |
| -------------------|--------------|
| <kbd>:</kbd>       | Focus on bar |
| <kbd>Enter</kbd>   | Run command  |
| <kbd>Esc</kbd>     | Exit bar     |

### Menu

| Keybinding                | Action                            |
| --------------------------|-----------------------------------|
| <kbd>m</kbd>              | Open add/modify/remove block menu |
| <kbd>j</kbd> <kbd>k</kbd> | Navigate in menu                  |
| <kbd>Enter</kbd>          | Run selected menu option          |
| <kbd>Esc</kbd>            | Exit menu                         |

### Add/modify block form

| Keybinding                                          | Action                                                       |
| ----------------------------------------------------|--------------------------------------------------------------|
| <kbd>Esc</kbd>                                      | Switch to normal mode or exit form if in normal mode already |
| <kbd>i</kbd>                                        | Switch to insert mode                                        |
| <kbd>h</kbd> <kbd>j</kbd> <kbd>k</kbd> <kbd>l</kbd> | Navigate in form                                             |
| <kbd>Enter</kbd>                                    | Submit form or go to next field                              |

### Block selection

| Keybinding                | Action                       |
| --------------------------|------------------------------|
| <kbd>j</kbd> <kbd>k</kbd> | Navigate in block selection  |
| <kbd>Enter</kbd>          | Run specified in menu action |
| <kbd>Esc</kbd>            | Exit block selection         |

## Commands

| Command | Action                         |
| --------|--------------------------------|
| :w      | Write changes to datafiles     |
| :q      | Exit tibivi                    |
| :q!     | Exit tibivi and ignore changes |
| :wq     | Write and exit tibivi          |
