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

	createIncidentSummary := flag.String("create-incident", "", "Incident Summary Description")
	incidentResolutionSummary := flag.String("resolve-incident", "", "Incident Resolution Description")
	createTaskSummary := flag.String("create-task", "", "Task Summary Description")
	exampleHttpCall := flag.Bool("example-http-call", false, "Make an example HTTP call")

	flag.Usage = showHelpText

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("ERROR! No arguments given!")
		showHelpText()
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

	if *exampleHttpCall {
		makeExampleHttpCall()
	}
}

func createIncident(incidentSummary string) {
	if strings.HasPrefix(incidentSummary, "-") {
		printError(incidentSummary, "incident summary")
		return
	}

	fmt.Println("An incident has been created:", incidentSummary)
}

func resolveIncident(incidentResolutionSummary string) {
	if strings.HasPrefix(incidentResolutionSummary, "-") {
		printError(incidentResolutionSummary, "incident resolution summary")
		return
	}

	fmt.Println("An incident has been resolved:", incidentResolutionSummary)
}

func createTask(taskSummary string) {
	if strings.HasPrefix(taskSummary, "-") {
		printError(taskSummary, "task summary")
		return
	}

	fmt.Println("A task has been created:", taskSummary)
}

func printError(invalidValue string, description string) {
	errorMsg := fmt.Sprintf("ERROR! The value [%s] is not a valid %s!", invalidValue, description)
	fmt.Println(errorMsg)
}

func showHelpText() {
	var text strings.Builder

	text.WriteString("\nDexterity Shell\n\n")
	text.WriteString("This is the command line tool for interacting with the Dexterity system!\n\n")

	text.WriteString("Usage:\n")
	text.WriteString("  -create-incident 'The website is broken!'\n")
	text.WriteString("  -resolve-incident 'The database is working again!'\n")
	text.WriteString("  -create-task 'Add user login'\n")

	text.WriteString("\nDebug:\n")
	text.WriteString("  -example-http-call  -  Makes an example HTTP call to 'https://www.rarelyprolific.co.uk'\n")

	fmt.Print(text.String())
}
