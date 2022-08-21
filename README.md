# Bookmarks

Bookmarks is a tool to easily jump between folders in the cli, using bookmarks.

## Usage

```bash
> bookmarks help

Bookmark manages bookmarks in the cli

Usage:
  bookmark [flags]
  bookmark [command]

Available Commands:
  add         Add a new bookmark to the bookmark list
  completion  Generate the autocompletion script for the specified shell
  get         retrieve path for a given bookmark
  help        Help about any command
  remove      Removes a bookmark from the bookmark list
  update      Updates a bookmark in the bookmark list

Flags:
  -h, --help   help for bookmark

Use "bookmark [command] --help" for more information about a command.
```

## Examples

```bash
bookmarks add go $HOME/go/src/github.com/mrjerz/ # add bookmark

bookmarks get go # return path of a bookmark.
/Users/jerzy/go/src/github.com/mrjerz

bookmarks update go $HOME/go/ # update a bookmark key

bookmarks remove go # delete a bookmark key
```

## Configuration

Bookmarks uses a configuration file `$HOME/.bookmarks`. It is JSON encoded, and contains the key name and the correspondig folders.

```json
{
  "bookmarks": [
    {
      "name": "go",
      "path": "/Users/jerzy/go/src/github.com/mrjerz"
    }
  ]
}
```

## Add Aliases to .zshrc

Make the bookmark tool easier to use with zsh aliases

```bash
g() { cd `bookmarks get $1` }
a() { cd `bookmarks add $1 $2` }
u() { cd `bookmarks update $1 $2` }
alias r ='bookmarks remove'
```
