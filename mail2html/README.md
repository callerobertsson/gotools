# mail2html

_Simple raw mail to HTML converter with fallback to text_

Takes a raw mail message on stdin and extracts the text/html part. If no
text/html is present it will output the text/plain part. Otherwise it
will exit with status 1.

I use it in a mutt macro that pipes a mail to elinks without involving
mailcap.

/Calle

