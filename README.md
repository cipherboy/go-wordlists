# go-wordlist
    u: Alexander Scheel
    e: <alexander.m.scheel@gmail.com>
    l: BSD: 2-clause


Basic word list generator taking the original input and adding numbers of
lengths 1-4 (0-9, 00-99, 000-999, and 0000-9999) and writing to STDOUT using
buffio (performance reasons).

Note, output is not sorted in any order. 

    hello
    hello0
    ...
    hello9
    hello00
    ...
    hello99
    hello000
    ...
    hello999
    hello0000
    ...
    hello9999
