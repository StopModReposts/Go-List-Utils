# SMR List Utils

This is a Go CLI tool that helps with managing the StopModReposts blacklist.

## Downloading

Use GitHub Releases and download binary. ELF is for Linux, EXE is for windows  
**WARNING:** I have not tested this on anything other than Linux

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
