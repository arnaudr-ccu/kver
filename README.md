kver (aka kernel version) reads a file and looks for a string that looks like a
version, something like `\d\d*\.\d\d*\.\d\d*\S*`. The first match is returned.

    $ kver -kernel /boot/vmlinuz-linux
    5.3.0-2-amd64
    $ ls /usr/lib/modules
    5.3.0-2-amd64

    # !!! amazing, it matches !!!
    # knowing a particular kernel file,
    # scripts can find the corresponding modules

    $ for f in /boot/vmlinuz-*; do echo "$f -> $(kver -kernel $f)"; done
    /boot/vmlinuz-linux -> 5.4.2-arch1-1
    /boot/vmlinuz-linux-hardened -> 5.3.15.a-1-hardened
    /boot/vmlinuz-linux-lts -> 4.19.88-1-lts
    /boot/vmlinuz-linux-zen -> 5.4.2-zen1-1-zen

kver can also take a string, and then scan installed kernels for this string.
It returns the path of the first kernel where the string was found.

    $ ls /usr/lib/modules
    5.3.0-2-amd64
    $ kver -release 5.3.0-2-amd64
    /boot/vmlinuz-linux

    # !!! amazing !!!
    # knowing a particular kernel release,
    # scripts can find the corresponding kernel

    $ for f in /lib/modules/[0-9]*; do echo "$f -> $(kver -release $(basename $f))"; done
    /lib/modules/4.19.88-1-lts -> /boot/vmlinuz-linux-lts
    /lib/modules/5.3.15.a-1-hardened -> /boot/vmlinuz-linux-hardened
    /lib/modules/5.4.2-arch1-1 -> /boot/vmlinuz-linux
    /lib/modules/5.4.2-zen1-1-zen -> /boot/vmlinuz-linux-zen

See: https://stackoverflow.com/q/3180029/776208
