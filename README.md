> :warning: **THIS REPOSITORY IS NO LONGER MAINTAINED**

# Slack Announce Action

Sends a slack message announcing the completion of a workflow run with build details and links.

![Screenshot](/msg-screenshot.png)

## Inputs

| Input Name  | Required | Default | Description |
| ----------- | ----------- | ---------- | ---------- |
| `webhook_url` | YES       | N/A | Slack webhook url to use for sending the announcement |
| `channel` | YES | N/A | Slack channel to send the announcement to |
| `indicators` | no | "[]" | JSON array of `{ "name": "my name", "status": "my status" }` objects to report on. If any indicators have a status of `failed` or `cancelled`, the overall announcement status will be `failed` or `cancelled` respectively. Statuses must be one of `success`, `failure`, `skipped`, or `cancelled`. At least 1 indicator is required unless a `status_override` is given. |
| `status_override` | no | N/A | Can be used to override the status derived from the individual indicator statuses for the overall announcement. Statuses must be one of `success`, `failure`, `skipped`, or `cancelled`. |
| `username` | no | "GitHub Actions" | Username to display as the sender of the announcement |
| `title_entity` | no | "workflow" | The entity being reported on to be used in the announcement title, ex. "my-repo <title_entity> completed successfully!" (supports markdown) |

## Usage

In the example below, the announcement will report the build as failed due to the failure of the `say-goodbye` job (see screenshot above).

```yaml
jobs:

  say-hello:
    runs-on: ubuntu-latest
    steps:
      - name: Say Hello
        run: echo "Hello, ${NAME}!"

  say-goodbye:
    needs: [say-hello]
    runs-on: ubuntu-latest
    steps:
      - name: Say Goodbye
        run: |
          echo "Goodbye, ${NAME}!"
          exit 1

  slack-announce:
    needs:
      - say-hello
      - say-goodbye
    if: ${{ always() }}
    runs-on: ubuntu-latest
    steps:
      - uses: DrizlyInc/slack-announce-action@v0.1.0
        with:
          webhook_url: ${{ secrets.SLACK_WEBHOOK_URL }}
          channel: webhook-playground
          title_entity: build
          indicators: |
            [
              { "name": "`say-hello`", "status": "${{ needs.say-hello.result }}" },
              { "name": "`say-goodbye`", "status": "${{ needs.say-goodbye.result }}" }
            ]
```

This action can also be used to report the status of a single job using `status_override`:
```yaml
jobs:

  say-hello:
    runs-on: ubuntu-latest
    steps:

      - name: Say Hello
        run: echo "Hello, ${NAME}!"

      - uses: DrizlyInc/slack-announce-action@v0.1.0
        with:
          webhook_url: ${{ secrets.SLACK_WEBHOOK_URL }}
          channel: webhook-playground
          status_override: ${{ job.status }}
          title_entity: "`say-hello`"
```

# Releasing

To generate a new release of this action, simply update the version tag on the image designation at the end of the [action metadata file](./action.yml). The github workflow will automatically publish a new image and create a release upon merging to main.

# License

The contents of this repository are released under the [GNU General Public License v3.0](LICENSE)
