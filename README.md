# jotto

Wordle/jotto for your terminal! Uses the official wordle word list, answer list,
and defaults to the daily wordle for your local time. You can play older or
newer wordles by passing the number as an argument. Your results are
automatically copied into your clipboard using the same format as wordle.

## HOW TO PLAY

Guess a random daily 5 letter word in 6 tries.

Each guess must be a valid 5 letter word. Hit the enter button to submit.

After each guess, the color of the letters show how close your guess was to the
answer. A green letter indicates the correct letter in the correct spot. A
yellow letter indicates a correct letter. Finally, a list of letter you haven't
used yet is printed at the end of each guess.

[Recording of solving Wordle 232](https://asciinema.org/a/05xmNs8QcKFSAtKPUNWFUoYIu)

## INSTALL

Ensure you have `golang`, `scdoc` and `make` installed. Then run `sudo make
install` to compile and install `jotto`. You can run `sudo make uninstall` to
uninstall. If you're missing `scdoc` make will print an error, but will still
install `jotto` just without the man page.

Alternatively, you can run `go build` and copy the `jotto` binary anywhere you'd
like.

## CONTRIBUTING

This project is licensed under GPL3 or later. You can
[email patches](https://git-send-email.io) or questions to
[~kota/public-inbox@lists.sr.ht](https://lists.sr.ht/~kota/public-inbox) or
browse the list [here](https://lists.sr.ht/~kota/public-inbox).
