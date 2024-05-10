# dict-counter

Reads a word list from a file (or STDIN), counts character frequency, and outputs histograms.

Normalizes on:
- Case
- Diacritics (character accents)

Ignores:
- Possessives (any line ending in "'s"), which essentially duplicate many words and skew counts
- Non-alphabetic characters (punctuation/digits/whitespace/...)

Notably, it does not discard plural forms. This gets kinda tricky and language-specific.
