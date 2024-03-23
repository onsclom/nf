# nf

`nf` makes it easier to create new files. `touch` lets you create files, but it doesn't create directories. `nf` does! and it gives you a nice tree view of the directories it creates as output.

## Example

```bash
$ pwd
/home/user
$ nf foo/bar/baz.txt
/home/user/
└─foo/ (existed)
  └─bar/ (created)
    └─baz.txt (created)
```
