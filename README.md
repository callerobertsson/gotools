# gotools

_A small collection of Golang tools._

Some recreational Golang coding projects I implemented for fun or gain.

Moved them here from miscellanous locations to keep the in one place.

For me they are useful but do expect to encounter some bugs if you try
them out.

I have tested them on Mac OS and Arch Linux but there might be problems
running them on Windows.

## chkpath

Iterates the PATH variable and checks if each path is readable, a
directory, and contains executable files.

## ghich

Searches the path to find first (default) or all (-a) occurences of an
executable file like the `which` linux command. Reports inconcistencies
in the path as a bonus.

## gohuman

Command line tool for converting long numbers (int64) to human readable
form.

The exported functions 
`Bytes(int64, int, bool)` and `Kilos(int64, int, bool)` 
might be useful.

Nice to have in the shell when you want to read big numbers from some
script.

## gopastebin

A tool to send text from stdin to pastebin.com.

Useful to quickly share code with your friends. Don't use it that often
but sometimes is good to have.

## gotoolchain

Command used to test and check Golang packages.

Put it in `$PATH` somewhere and just run `gotoolchain` i a Golang
package directory.

This is the tool I use most as it is part of my development process
together with https://github.com/callerobertsson/eye

## mail2html

A tool to convert a mail message to HTML with fallback to text.

I use this in mutt to preprocess mail messages opened in elinks.

## mentalnote

Post messages to a Slack channel from command line. Can also list
messages from the channel.

I use this to quickly send me a mental reminder.

/Calle
