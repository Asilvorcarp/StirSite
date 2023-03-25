# Stir Site

A script to make my website up to date by stirring it.

(like my blog with Notion as CMS)

## Usage

You can change global variables like `WEBSITE` in `.env` file:

```
WEBSITE=https://www.example.com
```

Then run the script:

```bash
$ go install
$ go run main.go
```

## Todo

- [ ] github action version
- [ ] docker version
- [ ] list of sites to stir
- [x] read .env file