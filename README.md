# URL Encode

This package provides a `urlencode` and `urldecode` utilities.

Input is given as an argument or via stdin.

## Usage

### Args

```
$ urlencode hello[there]=1
hello%5Bthere%5D%3D1
```

### Stdin

```
$ echo hello[there]=1 | urlencode
hello%5Bthere%5D%3D1
```

# License

The MIT License (MIT)

Copyright (c) 2021 Scott Barr

See [LICENSE](LICENSE)
