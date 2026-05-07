# A Github CLI that fetches from external URL 

This is a github cli that was built to fetch data from an external user

## What does this CLI do?

1. Authenticate users.
2. Refresh user token in case of expiry.
3. Creates a user profile.
4. Get users profile and get profiles based on specified id.
5. Export files from database to current directory.

## Running locally 

Requirements:
1. Go must be installed on PC.
2.

## Architecture

### Login Flow
User runs the command: 
`bash 
insighta login
`
Insighta creates a config file in user's config directory which is read for future access to backend.
When command runs, github opens directly in the browser.
When success is shown, the page can be closed.

## Create profile
To create user, the requesting user must be an admin
`bash
insighta create --name {" Name to be created "}

## Get user
To get user, the requesting user must be an admin
`bash
insighta get [id]
`

## List users
To list user, the requesting user must be an admin
`bash
insighta list --params value
`

## Export users
To list user, the requesting user must be an admin
`bash
insighta export --format value --params value
`

☕ Buy us a coffee
