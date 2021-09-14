# Slack Announce Action

Sends a slack message announcing the completion of a workflow run with build details and links.

![Screenshot](/msg-screenshot.png)

# Usage

```yaml
- uses: DrizlyInc/slack-announce-action@v0.1.0
  with:

    # Slack webhook url to use for sending the notification
    # required
    webhook_url: ${{ secrets.GHA_SLACK_WEBHOOK_URL }}

    # Status to be included in the message, from ${{ job.status }} ("success", "failure", or "cancelled")
    # required
    status: ${{ job.status }}

    # JSON array of { "title": "my title", "status": "my status" } objects to provide statuses of individual steps or jobs
    # optional, default "[]"
    steps: |
      [
        { "title": "`terraform-plan` on dev-general", "status": "${{ needs.terraform-plan-dev-general.result }}" },
        { "title": "`terraform-apply` on dev-general", "status": "${{ needs.terraform-apply-dev-general.result }}" },
        { "title": "`terraform-plan` on old-prod", "status": "${{ needs.terraform-plan-old-prod.result }}" },
      ]

    # Slack channel to send the announcement to
    # required
    channel: dev-releases

    # Username to display as the sender of the notification
    # optional, default "GitHub Actions"
    username: GitHub Actions
```

# License

The contents of this repository are released under the [GNU General Public License v3.0](LICENSE)