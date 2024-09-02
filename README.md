# obs_task
Integration google Tasks on Obsidian.
Using [Google Tasks API](https://pkg.go.dev/google.golang.org/api/tasks/v1).

# Prerequisites
- Latest version of [Go](https://golang.org/).
- Latest version of [Git](https://git-scm.com/).
- [A Google Cloud project](https://developers.google.com/workspace/guides/create-project).
- A Google account with Google Tasks enabled.

## Set up environment
This is a copy of the [Go quickstart](https://developers.google.com/tasks/quickstart/go)
### 1. [Enable the Google Tasks API](https://console.cloud.google.com/flows/enableapi?apiid=tasks.googleapis.com).
### 2. Configure the OAuth consent screen
If you're using a new Google Cloud project to complete this quickstart, configure the OAuth consent screen and add yourself as a test user. If you've already completed this step for your Cloud project, skip to the next section.

1. In the Google Cloud console, go to Menu menu > APIs & Services > OAuth consent screen.
[Go to OAuth consent screen](https://console.cloud.google.com/apis/credentials/consent)

2. For User type select Internal, then click Create.
3. Complete the app registration form, then click Save and Continue.
4. For now, you can skip adding scopes and click Save and Continue. In the future, when you create an app for use outside of your Google Workspace organization, you must change the User type to External, and then, add the authorization scopes that your app requires.
5. Review your app registration summary. To make changes, click Edit. If the app registration looks OK, click Back to Dashboard.

### 3. Authorize credentials for a desktop application
To authenticate end users and access user data in your app, you need to create one or more OAuth 2.0 Client IDs. A client ID is used to identify a single app to Google's OAuth servers. If your app runs on multiple platforms, you must create a separate client ID for each platform.
1. In the Google Cloud console, go to Menu menu > APIs & Services > Credentials.
  [Go to Credentials](https://console.cloud.google.com/apis/credentials)
2. Click Create Credentials > OAuth client ID.
3. Click Application type > Desktop app.
4. In the Name field, type a name for the credential. This name is only shown in the Google Cloud console.
5. Click Create. The OAuth client created screen appears, showing your new Client ID and Client secret.
6. Click OK. The newly created credential appears under OAuth 2.0 Client IDs.
7. Save the downloaded JSON file as credentials.json, and move the file to your working directory.

### 4. Prepare the workspace
1. Create a working directory:
```bash
mkdir quickstart
```
2. Change to the working directory:
```bash
cd quickstart
```
4. Initialize the new module:
```bash
go mod init quickstart
```
5. Get the Google Tasks API Go client library and OAuth2.0 package:
```bash
go get google.golang.org/api/tasks/v1
go get golang.org/x/oauth2/google
```

## Set up the program
At first lunch...
