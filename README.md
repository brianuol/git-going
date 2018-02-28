# git-going

Some fun with golang

## get it!

Clone or download this repo

## build it!

run `go build`

## run it

Get some usage help: `git-going -h`
```
Usage of git-going:
  -authToken string
        Optional.  The auth token to provide for faster search execution (anonymous runs 3x slower)
  -orgName string
        Required. The github org to analyze
  -topN int
        Optional (default 25).  The number of results (top-n) to display for each analysis set (default 25)
```

In addition to the above flags, the analysis types of `starred, forked, pulled, contributed` can be used to run any or all of the outputs supported by `git-going`.

Example: `git-going -orgName Netflix -topN 29 starred forked` would produce the following:
```
c:\Users\bob\go\src\github.com\brianuol\git-going>git-going -orgName Netflix -topN 29 starred forked
Starting...


### Top 29 Repositories, by Total Stars ###

Rank            Repo                            Stars
====            ===========                     ===========
1               Hystrix                         12687
2               falcor                          8564
3               SimianArmy                      6241
4               eureka                          4855
5               zuul                            3883
6               chaosmonkey                     3763
7               vector                          2639
8               dynomite                        2361
9               Scumblr                         2320
10              asgard                          2181
11              fast_jsonapi                    2147
12              vizceral                        2105
13              security_monkey                 2075
14              ribbon                          1885
15              curator                         1802
16              atlas                           1677
17              stethoscope                     1580
18              archaius                        1565
19              bless                           1375
20              servo                           1191
21              conductor                       981
22              astyanax                        969
23              vectorflow                      902
24              genie                           859
25              sleepy-puppy                    844
26              sketchy                         827
27              Priam                           824
28              rend                            823
29              lemur                           818


### Top 29 Repositories, by Total Forks ###

Rank            Repo                            Forks
====            ===========                     ===========
1               Hystrix                         2534
2               eureka                          1136
3               SimianArmy                      873
4               zuul                            848
5               Cloud-Prize                     501
6               ribbon                          462
7               asgard                          435
8               security_monkey                 426
9               curator                         393
10              falcor                          376
11              astyanax                        367
12              archaius                        339
13              dynomite                        294
14              Scumblr                         281
15              servo                           250
16              netflix.github.com              250
17              chaosmonkey                     248
18              Priam                           238
19              conductor                       227
20              genie                           211
21              vizceral                        206
22              Turbine                         194
23              vector                          184
24              karyon                          161
25              suro                            160
26              atlas                           155
27              aminator                        153
28              lemur                           144
29              governator                      142

c:\Users\bob\go\src\github.com\brianuol\git-going>

```

To gain faster execution time, provide an access token at the command line:

`git-going -orgName octocat -n55 -authToken [an oauth token]`

# share it!

`git-going` can be compiled on (or for) any modern OS

