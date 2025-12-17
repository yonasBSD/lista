# 🤔 What is this?
This is a simple todo list with notes cli program, as the name suggests. It's simple to use but also pleasant to interact with, thanks to [bubbletea](https://github.com/charmbracelet/bubbletea).
It has two modes, "normal" mode, where we can just pass our strings and command line flags to the program to accept our inputs. But it also as "TUI" mode, where you can keep engaging with the CLI
program interactively.

# 🤔 Why Lista?
Well, I made it for myself, because I found that when I am programming sometimes I like to take small notes here and there so I don't forget something, sometimes it might be just what to work on next,
other times something pops up and I would like to remind myself. 

I know there are other tools but, I have been going deep in the rabbit hole of NVIM + TMUX + BASH, so I want to make a program that I
can call inside a TMUX session to act as a todo list, and Lista was born. nothing crazy; its a simple program that helps me, so why not share it with others who might be into or getting into terminals like I am.

 # 🔨 Installation
 you can install using homebrew, right now its pointing at my personal tap:
 ```bash
  brew tap kwame-Owusu/lista
 ```
 # 🛠️ Development
 To build the project locally, you must:
 ```bash
  git clone https://github.com/kwame-Owusu/lista.git
  cd lista
 ```
Build the binary and run it:
```bash
  go build -o lista
  ./
 ```

 # Tmux + Nvim + Lista demo
 add link to the bash script which we use in TMUX

