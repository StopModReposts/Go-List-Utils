# SMR List Utils

This is a Go CLI tool that helps with managing the StopModReposts blacklist.

## Install

Use GitHub Releases and download binary.

### Linux Quickinstall

```bash
wget https://github.com/StopModReposts/Go-List-Utils/releases/latest/download/smr-list-utils.elf -O ~/.local/bin/smr-list-utils
chmod +x ~/.local/bin/smr-list-utils
```

### Windows Install

Download binary from releases, place in any folder. Run via cmd/powershell

## Usage

```
smr-list-utils <command> <list_file>
```

### Commands

`issue` - adds all the sites that have the `accepted` label on their issue to the end of the list.  
`sort` - sorts sites alphabetically.

### Examples

`smr-list-utils issue ../Illegal-Mod-Sites/lists/minecraft.yaml`  
`smr-list-utils sort ../Illegal-Mod-Sites/lists/minecraft.yaml`

### Workflow

First, add `approved` label to all issues that are good  
Then run the `issue` command  
In your IDE, change the domains that weren't parsed successfully  
Then run `sort`
