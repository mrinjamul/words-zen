# Words Zen

Words as a service written in Go.

---

> A simple API to generate unique randomized words & phrases.

## Usage

### Word Types

words-zen supports 8 word types that can be used to form a phrase:

- `adjective`
- `adverb`
- `animal`
- `bodyPart`
- `gerund`
- `noun`
- `pluralNoun`
- `verb` (imperative mood)

### Additional Features

When placed at the beginning of a phrase, the word `a` will be transformed into `an` if the second word in the phrase begins with a vowel (i.e. _a/an squid/octopus_). You can also use any additional words to form a phrase like _the_, _it_, _is_, etc.

### Forming a phrase

To form a phrase, connect the type(s) of words you want with a comma. For example, to form a phrase consisting of a verb, `the`, and a plural noun, you would use the following:
`the,pluralNoun,is,gerund` for the slug API or `the $pluralNoun is $gerund` for the query API.

## API

### Slug API

Make a `GET` request to [http://localhost:3000/api/PHRASE](http://localhost:3000/api), where PHRASE is the type of phrase desired.

```shell script
curl --request GET \
  --url 'http://localhost:3000/api/a $adjective $noun'
```

### Query API

or, `POST` to [http://localhost:3000/api](http://localhost:3000/api) with a query of the type of phrase desired.

```shell script
curl --request POST \
    --url http://localhost:3000/api \
    --header 'content-type: application/json' \
    --data '{
  	  "query": "the $pluralNoun is $gerund"
    }'
```

## Installation

Clone it from GitHub,

```bash
git clone https://github.com/mrinjamul/words-zen.git
```

And Run by,

```bash
    cd words-zen
    go mod download
    go build
    ./words-zen
```

## Technology Stack

Technologies used:

- [Golang](https://golang.org/): [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- HTML,CSS and JavaScript

## License

open-sourced under MIT license [MIT](LICENSE)
