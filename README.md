# diff

See commit diffs between Staging and Production.

## Assumptions

* tagged releases represent current production state.
* master branch represent current staging state.

## Running

The makefile is your best friend!

### Setup

```bash
# setup dependencies
make dep

# build binary
make build
```

Copy and modify [.env.sample](.env.sample) to your own settings under a `.env` file.

| environment variable | remarks |
| --- | --- |
| `GITHUB_OWNER` | string value of your GitHub user or organization handler |
| `GITHUB_REPO` | comma-separated string value of repo names, e.g., `repo1,repo2` |
| `GITHUB_ACCESS_TOKEN` | string value of your GitHub access token, preferably of a bot account |


### run on local

```bash
make run
# or alternatively, in a verbose manner

GITHUB_ACCESS_TOKEN=yourToken GITHUB_OWNER=handler GITHUB_REPO=repo1,repo2 ./diff

open http://localhost:9999
```

or point your browser to http://localhost:9999 to see the JSON response.

### Example response

See [example_output.json](example_output.json)
