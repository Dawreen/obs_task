package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// Creating new tasklist
func createTasksList(title string, srv *tasks.Service) {
	fmt.Println("Creating new tasklist!")
	var testTaskList tasks.TaskList
	testTaskList.Title = title
	_, err := srv.Tasklists.Insert(&testTaskList).Do()
	if err != nil {
		log.Fatalf("Error insert new task list %v", err)
	}
}

func main() {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, tasks.TasksScope) // TasksReadonlyScope
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := tasks.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
	}

	r, err := srv.Tasklists.List().MaxResults(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve task lists. %v", err)
	}

	fmt.Println("Task Lists:")
	ids := map[string]string{}
	if len(r.Items) > 0 {
		for _, i := range r.Items {
			fmt.Printf("%s (%s)\n", i.Title, i.Id)
			ids[i.Title] = i.Id
			t, err := srv.Tasks.List(i.Id).Do()
			if err != nil {
				log.Fatalf("Unable to retrieve tasks from task lists. %v", err)
			}
			for _, j := range t.Items {
				// j.ForceSendFields = append(j.ForceSendFields, "Completed")
				fmt.Printf("Title: %s (%s)\n", j.Title, j.Id)

				// fmt.Printf("\tAssignmentInfo: %s", j.AssignmentInfo)
				// fmt.Printf("\tCompleted: %s", j.Completed)
				fmt.Printf("\tDeleted: %t\n", j.Deleted)
				fmt.Printf("\tDue: %s\n", j.Due)
				fmt.Printf("\tEtag: %s\n", j.Etag)
				fmt.Printf("\tHidden: %t\n", j.Hidden)
				fmt.Printf("\tId: %s\n", j.Id)
				fmt.Printf("\tKind: %s\n", j.Kind)
				fmt.Printf("\tLinks: %d\n", len(j.Links))
				fmt.Printf("\tNotes: %s\n", j.Notes)
				fmt.Printf("\tParent: %s\n", j.Parent)
				fmt.Printf("\tPosition: %s\n", j.Position)
				fmt.Printf("\tSelfLink: %s\n", j.SelfLink)
				fmt.Printf("\tStatus: %s\n", j.Status)
				fmt.Printf("\tTitle: %s\n", j.Title)
				fmt.Printf("\tUpdated: %s\n", j.Updated)
				fmt.Printf("\tWebViewLink: %s\n", j.WebViewLink)
				// fmt.Printf("\tForceSendFields: %s\n",j.ForceSendFields)
				// fmt.Printf("\tNullField: %ss\n"j.NullField)
			}
			fmt.Println("-- END task list --")
		}
	} else {
		fmt.Print("No task lists found.")
	}
	// Creating new tasklist
	createTasksList("Testing new fun", srv)

	// Creating a new Task
	fmt.Println("Creating new Task")
	var testTask tasks.Task
	testTask.Title = "Test API creation"
	testTask.Notes = "Created using the google Go tasks API"
	_, err = srv.Tasks.Insert(ids["Test creating TaskList"], &testTask).Do()
	if err != nil {
		log.Fatalf("Error creating new task %v", err)
	}
}
