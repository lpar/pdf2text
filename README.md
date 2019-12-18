# pdf2text

Experimental PDF to plain text which works as a pipe by using temporary files.

This is just a quick hack and probably shouldn't be used for anything.

`pdf2text < some-pdf-file.pdf > the-text.txt`

It doesn't try to do anything clever about collapsing repeated whitespace or wrapping lines sensibly, so you will likely have to postprocess the output fairly heavily.
