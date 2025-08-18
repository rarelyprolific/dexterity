package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// TODO: Add login via signon.
	// TODO: Add database interactions.

	createIncidentSummary := flag.String("create-incident", "", "incident summary Description")
	incidentResolutionSummary := flag.String("resolve-incident", "", "incident resolution description")
	createTaskSummary := flag.String("create-task", "", "task summary description")
	exampleJsonRequest := flag.Bool("example-json-request", false, "make an example request for JSON content")

	flag.Usage = showHelpText

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("ERROR! No arguments given!")
		showHelpText()
		os.Exit(1)
	}

	if *createIncidentSummary != "" {
		createIncident(*createIncidentSummary)
	}

	if *incidentResolutionSummary != "" {
		resolveIncident(*incidentResolutionSummary)
	}

	if *createTaskSummary != "" {
		createTask(*createTaskSummary)
	}

	if *exampleJsonRequest {
		getJsonContent("http://localhost:8080/realms/master/.well-known/openid-configuration")
	}
}

// createIncident creates a new incident with the given summary.
func createIncident(incidentSummary string) {
	if strings.HasPrefix(incidentSummary, "-") {
		printError(incidentSummary, "incident summary")
		return
	}

	fmt.Println("An incident has been created:", incidentSummary)
}

// resolveIncident resolves an existing incident with the given incident resolution summary.
func resolveIncident(incidentResolutionSummary string) {
	if strings.HasPrefix(incidentResolutionSummary, "-") {
		printError(incidentResolutionSummary, "incident resolution summary")
		return
	}

	fmt.Println("An incident has been resolved:", incidentResolutionSummary)
}

// createTask creates a new task with the given summary.
func createTask(taskSummary string) {
	if strings.HasPrefix(taskSummary, "-") {
		printError(taskSummary, "task summary")
		return
	}

	fmt.Println("A task has been created:", taskSummary)
}

// printError prints an error relating to an invalid command line argument.
func printError(invalidValue string, description string) {
	errorMsg := fmt.Sprintf("ERROR! The value [%s] is not a valid %s!", invalidValue, description)
	fmt.Println(errorMsg)
}

// showHelpText shows the help text for the application.
func showHelpText() {
	var text strings.Builder

	text.WriteString("\nDexterity Shell\n\n")
	text.WriteString("This is the command line tool for interacting with the Dexterity system!\n\n")

	text.WriteString("Usage:\n")
	text.WriteString("  -create-incident 'The website is broken!'\n")
	text.WriteString("  -resolve-incident 'The database is working again!'\n")
	text.WriteString("  -create-task 'Add user login'\n")

	text.WriteString("\nDebug:\n")
	text.WriteString("  -example-json-request  -  Makes an example HTTP call to get Keycloak sign on discovery JSON\n")

	fmt.Print(text.String())
}
