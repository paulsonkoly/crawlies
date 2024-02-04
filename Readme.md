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

     crawlies -input input.txt
     [==============================================================================] tokentype_string.go
     [==============================================================================] evaluator.go
     [==============================================================================] go.mod
     [==============================================================================] token.go
     [==============================================================================] calc_test.go
     [==============================================================================] frame.go
     [==============================================================================] value.go
     [==============================================================================] node.go
     [==============================================================================] pretty_printer.go
     [==============================================================================] transaction.go
     [==============================================================================] LICENSE
     [==============================================================================] transaction_test.go
     [==============================================================================] lexer.go
     [==============================================================================] lexer_test.go
     [==============================================================================] states.go
     [==============================================================================] go.sum
     [==============================================================================] transformer.go
     [==============================================================================] Readme.md
     [==============================================================================] token_wrapper.go
     [==============================================================================] parser.go
     [==============================================================================] calc.go
    Network error Get "yyy": unsupported protocol scheme ""
    Network error Get "xxxx": unsupported protocol scheme ""
    Non 2xx response 404 for https://raw.githubusercontent.com/paulsonkoly/calc/main/log2.txt
    Non 2xx response 404 for https://raw.githubusercontent.com/paulsonkoly/calc/main/x.calc
    Non 2xx response 404 for https://raw.githubusercontent.com/paulsonkoly/calc/main/calc
    Non 2xx response 404 for https://raw.githubusercontent.com/paulsonkoly/calc/main/log.txt
