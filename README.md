# bcrypt-hasher

## Usage with docker:

Hash the password `asdf`:

```bash
$ docker run --rm ghcr.io/bjarkt/bcrypt-hasher -pass asdf
<hash>
```

Check that the generated hash matches, or not:

```bash
$ docker run --rm ghcr.io/bjarkt/bcrypt-hasher -pass password -hash <hash>
Match
```

Output the hash base 64 encoded:

```bash
$ docker run --rm ghcr.io/bjarkt/bcrypt-hasher -pass asdf -b64
<base64 encoded hash>
```