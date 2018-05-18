# diff

See commit diffs between Staging and Production.

## Assumptions

* tagged releases represent current production state.
* master branch represent current staging state.

## Running

The makefile is your best friend!

### Setup

```shell
# setup dependencies
make dep

# build Docker image
make docker_build
```

Copy and modify [.env.sample](.env.sample) to your own settings under a `.env` file.

| environment variable | remarks |
| --- | --- |
| `GITHUB_OWNER` | string value of your GitHub user or organization handler |
| `GITHUB_REPO` | comma-separated string value of repo names, e.g., `repo1,repo2` |
| `GITHUB_ACCESS_TOKEN` | string value of your GitHub access token, preferably of a bot account |


# run image
make docker_run
```
