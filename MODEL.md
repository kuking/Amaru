# Search engine

## Model

### Token

`tid` is 32 bits unsigned (4b uniques)
A token can be of different classes: text, tag.

- _Text token_: Normalised, plurals removed, etc.
- _Tag token_: i.e. pic has a dog, image is BW.

### Tokens Catalogue

The collection of tokens.

### Document

Fields: `did` a 32 bits unsigned (4b uniques), `ranking` float32,  `URL` string.

### Documents Catalogue

The collection of documents.

### Dossier

A collection of documents sharing a token. `did` are sorted in the dossier, so dossiers intersections can be optimised.

### Anthology

A list of dossiers, one for each token that has at least one document.

## Implementation

### Token Catalogue

Fixed size record i.e. 1 byte token type, then the token string `NULL` terminated with record size of `Amaru.MaxTokenLen` (25 bytes).
`tid` is implicit by order. .e. 450K tokens * (25 + 1) = 11.4MiB file.

_Filename:_ `/tokens`

### Document Catalogue

A text file: a float as string with six decimal points `\t` URL `\n`. `did` is implicit by position.

_Filename_: `/documents`

### Anthology

There can be 500K dossiers per anthology therefore they have to be packed into one file. A `MMAP`, holding all the dossiers.

- `n` records of:
    - `tid`: uint32
    - `capacity`: uint32
    - `count`: uint32
    - an ordered list of `did`
    - filename: `/anthology`

A MMAP, with an array of `uint64`, holding the offset to each `tid` in the anthology file, `tid` value is implicit.

_Filename_: `/anthology.idx`.

### Files

- `/tokens`
- `/documents`
- `/anthology`
- `/anthology.idx`

## Algorithm

### Basic without search intent

- a search query is converted to a list of tokens tid
- some tag tokens might also be added i.e. intent or doc type
- an anthology is selected
- intersect all documents present in all relevant dossiers (as in tid) for positive tokens
- for each negative token, do a filter fast (needs thinking)
- for all resulting documents, sort them by ranking.
    - ranking can be an array indexed by did.
- Truncate to a maximum of X (i.e. 10k)
- cache the result (query -> array of did)
- Open the documents, and lookup for exact sequences, that would rank more .. (how?)

### With search intent

Search intent can be handled with tags in documents.
This needs more thinking.

## Further, pending

- geo locations
- complex logic in search

### See Also

- http://blevesearch.com/ - super slow
- https://github.com/blugelabs/bluge
