<p></p>
<p align="center">
  <img src="find-suid.png" />
</p>

## Examples
```bash
  # Find Setuid binaries that are in $PATH. This is 100x faster than using find to scan
  # an entire file system. More often than not there is some attack surface in this area
  ./find-suid --path

  # Supply a set of paths to search. More often than not we also have an idea where a Setuid
  # binary may be. If we cant find it at this point, then we can resort to scanning the entire
  # file system
  ./find-suid --find "/sbin" "/home/sweetbbak/bin" "/root/workdir"
```

100 results, truncated:
> time ./find-suid -p 0.01s user 0.02s system 96% cpu 0.028 total

```bash
Found Setuid Binary: /home/sweet/bin/doit
Found Setuid Binary: /home/sweet/bin/suwu
Found Setuid Binary: /bin/sg
Found Setuid Binary: /bin/umount
Found Setuid Binary: /bin/unix_chkpwd
Found Setuid Binary: /usr/bin/chage
Found Setuid Binary: /usr/bin/chsh
Found Setuid Binary: /usr/bin/expiry
Found Setuid Binary: /usr/bin/fusermount
...
Found Setuid Binary: /usr/bin/passwd
Found Setuid Binary: /usr/bin/pkexec
Found Setuid Binary: /usr/bin/readcd
Found Setuid Binary: /usr/bin/rscsi
Found Setuid Binary: /usr/bin/sg
Found Setuid Binary: /usr/bin/su
Found Setuid Binary: /usr/sbin/sg
Found Setuid Binary: /usr/sbin/su
Found Setuid Binary: /usr/sbin/sudo
Found Setuid Binary: /usr/sbin/umount
Found Setuid Binary: /usr/sbin/unix_chkpwd
```

## Installation
```bash
  go build *.go
```
TODO: Add a releases page, depending on need for it.

## Features
idk it finds Setuid binaries.

## TODO
- add max and min depth
- check for more file modes
- more robust error handling

send PR's if you want to. They are definitely welcomed.
