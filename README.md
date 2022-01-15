<div align="center">

![Words as a Service](https://raw.githubusercontent.com/mrinjamul/words-zen/main/title.svg)

[![Hosted on Vercel](https://badgen.net/badge/%E2%96%B2%20Hosted%20on/Vercel/black)](https://vercel.com)
[![License: ISC](https://img.shields.io/badge/License-ISC-blue.svg)](https://opensource.org/licenses/ISC)
[![dependencies Status](https://david-dm.org/mrinjamul/words-zen/status.svg)](https://david-dm.org/mrinjamul/words-zen)
[![devDependencies Status](https://david-dm.org/mrinjamul/words-zen/dev-status.svg)](https://david-dm.org/mrinjamul/words-zen?type=dev)

[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/mrinjamul/words-zen/issues)
![mrinjamul/words-zen](https://badgen.net/github/last-commit/mrinjamul/words-zen/main)
[![Maintainability](https://api.codeclimate.com/v1/badges/913d463015f91a452b70/maintainability)](https://codeclimate.com/github/mrinjamul/words-zen/maintainability)

</div>

---

> A simple API to generate unique randomized words & phrases.

## Usage

### Word Types

words-zen supports 8 word types that can be used to form a phrase:

- `adjective`
- `adverb`
- `animal`
- `bodyParts`
- `gerunds`
- `pluralNouns`
- `verbs` (imperative mood)

### Additional Features

When placed at the beginning of a phrase, the word `a` will be transformed into `an` if the second word in the phrase begins with a vowel (i.e. _a/an squid/octopus_). You can also use any additional words to form a phrase like _the_, _it_, _is_, etc.

### Forming a phrase

To form a phrase, connect the type(s) of words you want with a comma. For example, to form a phrase consisting of a verb, `the`, and a plural noun, you would use the following:
`the,pluralNoun,is,gerund` for the slug API or `the $pluralNoun is $gerund` for the query API.

## API

### Slug API

Make a `GET` request to [https://words-zen.now.sh/api/PHRASE](https://words-zen.now.sh/api/), where PHRASE is the type of phrase desired.

```shell script
curl --request GET \
  --url 'https://words-zen.now.sh/api/the,$pluralNoun,is,gerund'
```

### Query API

or, `POST` to [https://words-zen.now.sh/api/](https://words-zen.now.sh/api/) with a query of the type of phrase desired.

```shell script
curl --request POST \
    --url https://words-zen.now.sh/api/ \
    --header 'content-type: application/json' \
    --data '{
  	  "query": "the $pluralNoun is $gerund"
    }'
```
