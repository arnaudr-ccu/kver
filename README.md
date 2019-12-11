kver (aka kernel version) reads a file and looks for a string that looks like a
version, something like `\d\d*\.\d\d*\.\d\d*\S*`. The first match is returned.

Use-case:

    $ ./kver -kernel /boot/vmlinuz-linux
    5.3.0-2-amd64
    $ ls /usr/lib/modules
    5.3.0-2-amd64

    # !!! amazing, it matches !!!
    # now scripts know which modules
    # and which kernels go together

See: https://stackoverflow.com/q/3180029/776208
