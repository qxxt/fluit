# fluit
fluit is a text formatter that adds margin to a text and limits a text length to specified breakpoints. It is based on the overflow property from CSS. The adjustable breakpoint makes it possible to print text that is responsive to the user's console width. It also has a CLI usage builder. Which is the main point for me making this.

### Basic Usage
``` go
    fluit.UserBreakpoint = 80

    fmt.Println(fluit.Wrap(4, "mQINBGG7QgUBEACsgsf8oWpmx56TahIt71x8aaFRUMZ8MRG+LJ7+zreJZOmaxjshfQN85aIwArUmH1DqDQLf6Jq0dnvVNrhR1c7iZE0r6K569dFUwpeCy3n5toWwj2JyjZSdxED/ODpQZ7y8u0moAGVM9/c7DeMZm17aqBXqr5TglilI2FX/Sq5C44FBfKN40qBBOw6b1lSV4B5fM6cypARgDuSGkOuA0OL+oThBjiykyZg+FbhFOxzS5vcfULvq\n"))

    fmt.Println(fluit.Wrap(4, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut diam quam nulla porttitor massa id. Porta lorem mollis aliquam ut. Mollis nunc sed id semper risus. Viverra vitae congue eu consequat. Dis parturient montes nascetur ridiculus mus. Aliquet porttitor lacus luctus accumsan tortor posuere. Amet consectetur adipiscing elit ut aliquam purus sit. Rhoncus urna neque viverra justo. Sagittis vitae et leo duis ut diam quam. Donec et odio pellentesque diam volutpat. Consectetur a erat nam at lectus urna duis convallis. Integer quis auctor elit sed vulputate. At tellus at urna condimentum. Vulputate eu scelerisque felis imperdiet proin fermentum leo vel. Dolor sit amet consectetur adipiscing elit pellentesque habitant morbi tristique."))
```

The code above will create the following output
```
    mQINBGG7QgUBEACsgsf8oWpmx56TahIt71x8aaFRUMZ8MRG+LJ7+zreJZOmaxjshfQN85aIwArUm
    H1DqDQLf6Jq0dnvVNrhR1c7iZE0r6K569dFUwpeCy3n5toWwj2JyjZSdxED/ODpQZ7y8u0moAGVM
    9/c7DeMZm17aqBXqr5TglilI2FX/Sq5C44FBfKN40qBBOw6b1lSV4B5fM6cypARgDuSGkOuA0OL+
    oThBjiykyZg+FbhFOxzS5vcfULvq

    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua. Ut diam quam nulla
    porttitor massa id. Porta lorem mollis aliquam ut. Mollis nunc sed id semper
    risus. Viverra vitae congue eu consequat. Dis parturient montes nascetur
    ridiculus mus. Aliquet porttitor lacus luctus accumsan tortor posuere. Amet
    consectetur adipiscing elit ut aliquam purus sit. Rhoncus urna neque viverra
    justo. Sagittis vitae et leo duis ut diam quam. Donec et odio pellentesque
    diam volutpat. Consectetur a erat nam at lectus urna duis convallis. Integer
    quis auctor elit sed vulputate. At tellus at urna condimentum. Vulputate eu
    scelerisque felis imperdiet proin fermentum leo vel. Dolor sit amet
    consectetur adipiscing elit pellentesque habitant morbi tristique.
```

### Breakpoint
You can specify breakpoints at which the text will be wrapped to a new line. This can be used to print text that is responsive to the user's console width. To get user's console width, you can use [GetSize](https://pkg.go.dev/golang.org/x/term#GetSize) function from [golang.org/x/term](https://pkg.go.dev/golang.org/x/term) module.

```go
    if w, _, err := term.GetSize(1); err != nil {
        fluit.UserBreakpoint = w
    }
```

This will make the output responsive to user's console width.
![Peek 2023-06-08 18-27](https://github.com/qxxt/fluit/assets/57898942/01a0634c-5b10-41f3-89ce-66859d713ae0)

However not all terminal is supported. If you're trying to use it on emacs minibuffer it create an ignorable errors.

### Usage Builder

```go
    fluit.UserBreakpoint = 80

    fluit.PrintlnWrap(4, "Run Emacs, the extensible, customizable, self-documenting real-time display editor.  The recommended way to start Emacs for normal editing is with no options at all.\n")

    fmt.Println("Initialization options:")
    u := fluit.Usages{}
    u.ArgumentRowLength = 30
    u.AddOption("--batch", "do not do interactive display; implies -q")
    u.AddOption("--chdir DIR", "change to directory DIR")
    u.AddOption("--daemon, --bg-daemon[=NAME]", "start a (named) server in the background")
    u.AddOption("--fg-daemon[=NAME]", "start a (named) server in the foreground")
    u.AddOption("--debug-init", "enable Emacs Lisp debugger for init file")
    u.PrintUsages()
    u.UsageGroup = nil

    fmt.Println("\nDisplay options:")
    u.ArgumentRowLength = 25
    u.AddOption("--background-color COLOR", "window background color")
    u.AddOption("--basic-display", "disable many display features;\nused for debugging Emacs")
    u.AddOption("--border-color, -bd COLOR", "main border color")
    u.AddOption("--border-width, -bw WIDTH", "width of main border")
    u.AddOption("--color, --color=MODE", "override color mode for character terminals;\nMODE defaults to `auto', and\ncan also be `never', `always',\nor a mode name like `ansi8'")
    u.PrintUsages()

    fmt.Print("\n")

    fluit.PrintlnWrap(0, "You can generally also specify long option names with a single -; for example, -batch as well as --batch. You can use any unambiguous abbreviation for a --option.\n")
    fluit.PrintlnWrap(0, "Various environment variables and window system resources also affect the operation of Emacs. See the main documentation.\n")
    fluit.PrintlnWrap(0, "Report bugs to bug-gnu-emacs@gnu.org. First, please see the Bugs section of the Emacs manual or the file BUGS.")
```

The output will be:
```
    Run Emacs, the extensible, customizable, self-documenting real-time display
    editor.  The recommended way to start Emacs for normal editing is with no
    options at all.

Initialization options:
  --batch                         do not do interactive display; implies -q
  --chdir DIR                     change to directory DIR
  --daemon, --bg-daemon[=NAME]    start a (named) server in the background
  --fg-daemon[=NAME]              start a (named) server in the foreground
  --debug-init                    enable Emacs Lisp debugger for init file

Display options:
  --background-color COLOR   window background color
  --basic-display            disable many display features;
                             used for debugging Emacs
  --border-color, -bd COLOR  main border color
  --border-width, -bw WIDTH  width of main border
  --color, --color=MODE      override color mode for character terminals;
                             MODE defaults to `auto', and
                             can also be `never', `always',
                             or a mode name like `ansi8'

You can generally also specify long option names with a single -; for example,
-batch as well as --batch. You can use any unambiguous abbreviation for a
--option.

Various environment variables and window system resources also affect the
operation of Emacs. See the main documentation.

Report bugs to bug-gnu-emacs@gnu.org. First, please see the Bugs section of the
Emacs manual or the file BUGS.
```

For more about documentation, please go to [This package's documentary page](https://pkg.go.dev/github.com/qxxt/fluit)
