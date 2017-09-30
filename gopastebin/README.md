# gopastebin

_Simple command line tool for piping text to pastebin.com_

Note that you need a user/developer key to use this tool. It can easily
be obtain as a registered user on the pastebin.com/api page.

# Usage example

    $ gopastebin -k <YOUR KEY> < you-text-file.txt
    
or

    $ ls -alF | gopastebin -k <YOUR KEY> -x "1Y" -n "Files"
    
A good ide is to make an alias or script that adds the key. See the
provided `bash-script-example` file.

/Calle

