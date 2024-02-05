# Crawlies

A go concurrent programming exercise.

It reads URLS from an input file (one per line) and downloads them placing them in a directory structure inferred from url paths. It runs on the specified number of threads concurrently. It outputs using a multi progress bar.

It outputs any errors at the end after the downloads.

## Command

    -input string
          input file name

    -threadCnt int
          number of downloader threads (default 8)

## Example

    $ crawlies -input input.txt
    [========================================================================]                      base-3-2-any.pkg.tar.zst.sig
    [========================================================================]                 bash-5.2.026-2-x86_64.pkg.tar.zst
    [========================================================================]             bash-5.2.026-2-x86_64.pkg.tar.zst.sig
    [========================================================================]                binutils-2.42-2-x86_64.pkg.tar.zst
    [========================================================================]            binutils-2.42-2-x86_64.pkg.tar.zst.sig
    [============================================>---------------------------]                  bison-3.8.2-6-x86_64.pkg.tar.zst
    [========================================================================]              bison-3.8.2-6-x86_64.pkg.tar.zst.sig
    [========================================================================]             brotli-1.1.0-1-x86_64.pkg.tar.zst.sig
    [========================================================================]                 brotli-1.1.0-1-x86_64.pkg.tar.zst
    [===============>--------------------------------------------------------]        brotli-testdata-1.1.0-1-x86_64.pkg.tar.zst
    [========================================================================]    brotli-testdata-1.1.0-1-x86_64.pkg.tar.zst.sig
    [========================================================================]              btrfs-progs-6.7-1-x86_64.pkg.tar.zst
    [========================================================================]          btrfs-progs-6.7-1-x86_64.pkg.tar.zst.sig
    [========================================================================]                  bzip2-1.0.8-5-x86_64.pkg.tar.zst
    [========================================================================]              bzip2-1.0.8-5-x86_64.pkg.tar.zst.sig
    [========================================================================]        ca-certificates-20220905-1-any.pkg.tar.zst
    [========================================================================]    ca-certificates-20220905-1-any.pkg.tar.zst.sig
    [===>--------------------------------------------------------------------] ca-certificates-mozilla-3.97-1-x86_64.pkg.tar.zst
    [========================================================================]  ca-certificates-utils-20220905-1-any.pkg.tar.zst
    [==================>-----------------------------------------------------]                                           core.db     
    Network error Get "yyy": unsupported protocol scheme ""
    Network error Get "xxxx": unsupported protocol scheme ""
    Non 2xx response 404 for https://raw.githubusercontent.com/paulsonkoly/calc/main/log2.txt
    Non 2xx response 404 for https://raw.githubusercontent.com/paulsonkoly/calc/main/x.calc
    Non 2xx response 404 for https://raw.githubusercontent.com/paulsonkoly/calc/main/calc
    Non 2xx response 404 for https://raw.githubusercontent.com/paulsonkoly/calc/main/log.txt
