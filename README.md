# bbi2github
Super simple way to migrate your BitBucket issues to GitHub account

## Usage

1. Export your Bitbucket issues via Settings -> Export function.
2. Extract your ZIP file and copy db-jira-cloud.json to your go project folder
3. Update CONST values with:
  - new GitHub account credentials (you can generate token by following these instructions [https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token]
  - repo owner
  - repo name
  
4. go run main.go
