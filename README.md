# `twitfol.go` and `folgat.sh`

`twitfol.go` is a small program for getting you the list of your Twitter followers on stdout.

It prints the list in format `<screenname><tab><userid>`.
This behavior can be changed in the code, there are many other properties that can be chosen.

It prints the list of the followers of the user of the consumer key.
This behavior can be changed in the code. But I'm pretty sure you need to provide *some* consumer key.

```
 diff twitfol_scrnam-tab-userid_2019-03-29T175454+0100.dat twitfol_scrnam-tab-userid_2019-03-29T185527+0100.dat
34d33
< _CancerCoaches	1099867811390746624
                 /^\
                  |
 tab separated____|
```
