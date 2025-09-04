package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rarelyprolific/dexterity/shell/incident"
	"github.com/rarelyprolific/dexterity/shell/task"
)

func main() {
	// TODO: Add login via signon.

	listIncidentsCommand := flag.Bool("list-incidents", false, "list incidents")
	showIncidentIdentifier := flag.String("show-incident", "", "show incidents")
	createIncidentSummary := flag.String("create-incident", "", "incident summary description")
	incidentResolutionSummary := flag.String("resolve-incident", "", "incident resolution description")

	listTasksCommand := flag.Bool("list-tasks", false, "list tasks")
	createTaskSummary := flag.String("create-task", "", "task summary description")
	exampleJsonRequest := flag.Bool("example-json-request", false, "make an example request for JSON content")

	flag.Usage = showHelpText

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("ERROR! No arguments given!")
		showHelpText()
		os.Exit(1)
	}

	if *listIncidentsCommand {
		listIncidents()
	}

	if *showIncidentIdentifier != "" {
		showIncident(*showIncidentIdentifier)
	}

	if *createIncidentSummary != "" {
		createIncident(*createIncidentSummary)
	}

	if *incidentResolutionSummary != "" {
		resolveIncident(*incidentResolutionSummary)
	}

	if *listTasksCommand {
		listTasks()
	}

	if *createTaskSummary != "" {
		createTask(*createTaskSummary)
	}

	if *exampleJsonRequest {
		fmt.Println(getJsonContent("http://localhost:8080/realms/master/.well-known/openid-configuration", true))
	}
}

// listIncidents lists incidents.
func listIncidents() {
	incidentApiUri := os.Getenv("INCIDENT_API_URI")

	if incidentApiUri == "" {
		incidentApiUri = "http://localhost:8910"
	}

	incidents := getJsonContent(incidentApiUri+"/incidents", false)

	var incidentSummaryList []incident.Summary

	err := json.Unmarshal([]byte(incidents), &incidentSummaryList)

	if err != nil {
		panic(err)
	}

	fmt.Println("Listing all incidents (ID  |  Status  |  Created On  |  Created By  |  Summary):")

	for _, incident := range incidentSummaryList {
		fmt.Printf("%s  |  %s  |  %s  |  %s  |  %s\n", incident.ID, incident.Status, incident.CreatedOn.Format("2006-01-02 15:04:05"), incident.CreatedBy, incident.Summary)
	}

	fmt.Println(len(incidentSummaryList), "incidents found!")
}

// showIncident shows a specific incident.
func showIncident(incidentIdentifier string) {
	if strings.HasPrefix(incidentIdentifier, "-") {
		fmt.Println(formatError(incidentIdentifier, "incident identifier"))
		return
	}

	fmt.Println("Showing incident:", incidentIdentifier)
}

// createIncident creates a new incident with the given summary.
func createIncident(incidentSummary string) {
	if strings.HasPrefix(incidentSummary, "-") {
		fmt.Println(formatError(incidentSummary, "incident summary"))
		return
	}

	fmt.Println("An incident has been created:", incidentSummary)
}

// resolveIncident resolves an existing incident with the given incident resolution summary.
func resolveIncident(incidentResolutionSummary string) {
	if strings.HasPrefix(incidentResolutionSummary, "-") {
		fmt.Println(formatError(incidentResolutionSummary, "incident resolution summary"))
		return
	}

	fmt.Println("An incident has been resolved:", incidentResolutionSummary)
}

// listIncidents lists tasks.
func listTasks() {
	taskApiUri := os.Getenv("TASK_API_URI")

	if taskApiUri == "" {
		taskApiUri = "http://localhost:8920"
	}

	tasks := getJsonContent(taskApiUri+"/tasks", false)

	var taskSummaryList []task.Summary

	err := json.Unmarshal([]byte(tasks), &taskSummaryList)

	if err != nil {
		panic(err)
	}

	fmt.Println("Listing all tasks (ID  |  Status  |  Created On  |  Created By  |  Summary):")

	for _, task := range taskSummaryList {
		fmt.Printf("%s  |  %s  |  %s  |  %s  |  %s\n", task.ID, task.Status, task.CreatedOn.Format("2006-01-02 15:04:05"), task.CreatedBy, task.Summary)
	}

	fmt.Println(len(taskSummaryList), "tasks found!")
}

// createTask creates a new task with the given summary.
func createTask(taskSummary string) {
	if strings.HasPrefix(taskSummary, "-") {
		fmt.Println(formatError(taskSummary, "task summary"))
		return
	}

	fmt.Println("A task has been created:", taskSummary)
}

// formatError builds an error string relating to an invalid command line argument.
func formatError(invalidValue string, description string) string {
	return fmt.Sprintf("ERROR! The value [%s] is not a valid %s!", invalidValue, description)
}

// showHelpText shows the help text for the application.
func showHelpText() {
	var text strings.Builder

	text.WriteString("\nDexterity Shell\n\n")
	text.WriteString("This is the command line tool for interacting with the Dexterity system!\n\n")

	text.WriteString("Usage:\n")
	text.WriteString("  -list-incidents\n")
	text.WriteString("  -show-incident 'ICD001'\n")
	text.WriteString("  -create-incident 'The website is broken!'\n")
	text.WriteString("  -resolve-incident 'The database is working again!'\n")

	text.WriteString("\n")

	text.WriteString("  -list-tasks\n")
	text.WriteString("  -create-task 'Add user login'\n")

	text.WriteString("\nDebug:\n")
	text.WriteString("  -example-json-request  -  Makes an example HTTP call to get Keycloak sign on discovery JSON\n")

	fmt.Print(text.String())
}
