# Dexterity

A system to help you design, build and support software systems.

> [!WARNING]
> **This is not a real application!** It doesn't really work. It will likely
> never be in a finished state. It's a dummy project I'm hacking on
> and using to learn stuff.

## Guiding Principles

### Statements we believe to be facts

- Complex systems are designed incrementally and initial assumptions are likely to change.
- Building systems is not just about bugs and features. Design, documentation, deployment, infrastructure and support work are also required.
- The start of many changes begin with asking a question.
- All functional areas may not be apparent at the origins of a system. They may be discovered as the system evolves.

### We care about the following

- Anything that helps you figure out and define what you need to do.
- Designing, defining, deciding and documenting the system.
- Troubleshooting the deployed system and finding effective resolutions.
- Not repeating the same mistake more than once.
- Minimal admin. The system should help you and not hinder you unnecessarily.
- Prioritisation. Rather than only justifying why a task should be done.. What is the cost of **not** doing something? What gets more difficult if you defer a task?

### We don't care about the following (either not at all or not yet)

- Estimation. Hard to do and usually wrong. Can be very wrong unless you've figured out **what** you need to do.
- Actual time taken or any metrics. May be useful to look back retrospectively but not a current focus of the system.

## Development Technologies

- Develop in `Go`.
- Deploy in `Docker Compose` (localdev) and `Kubernetes` via `Helm` chart.
- Use `Keycloak` for auth *(later)*
- Find a lightweight database for persistence *(later - use file system with grep until we have enough of a database schema to pick a database technology)*
- Some sort of web and/or application GUI **(later)**

## Components

- **dexterity-api** - REST APIs *(could return hardcoded responses during initial design phase then read/write files to a volume/folder)*. May have
individual API services for:
  - **question-api** - Asking and answering questions.
  - **incident-api** - Creating and maintaining incidents.

- **dexterity-shell** - A command line tool which will call the REST API. This is the UI until we have a GUI.
*(This will conveniently be a shell that runs inside the container network initially: We don't want to deal with public/private network concerns until later)*

- **dexterity-db** - An actual database server **(later)**

## Domain Model

The entities of this model will be exposed as resources via **Open API specifications** which will be implemented by the **dexterity-api(s)**.

### Entities

- **incident** - Something has happened! We don't know the full details yet! It's an investigation.
- **incident-status** [Unstarted, InProgress, Paused, Abandoned, Resolved]
- **task** - Something specific needs to be done!
- **task-status** [Undefined, ReadyToStart, InProgress, Paused, Completed]
- **question** - We have a question that needs to be answered.
- **answer** - We have an answer that resolves a question.

### Incident

An **incident** is an event which is occurring and requires attention. We may not have all the details yet and we may need to ask **questions**.
The resolution may not be obvious until more it known about the incident. It could be a non-issue. It may need a bug fix. It may just need instructions
to help troubleshoot.

### Task

A **task** is any reasonably small, specific and fully defined item of work which can be started. A task which has been created but is still
**Undefined** may have one or more **questions** which need **answers** before it is **ReadyToStart**.

### Brainstorming some higher-level entities

*These are very subject to change!*

To make the system meaningful we need to group **tasks** into structures which describe why we are doing them. These will help inform
and organised how we are designing and documenting the system.

- **system** - Defines and explains what the system does and what the core responsibilities are.
- **functional-area** - An area of an application with a specific responsibility within a larger system.
- **feature** - A feature of an application within a functional area.

### Types of Task

- **application** - A task related to the business logic and general high-level behaviours of the system.
- **bugfix** - A task related to small change which is a pure bugfix. No features are added.
- **build** - A task related to building the system (CI) such as compiling and running automated tests.
- **deployment** - A task related to deploying the system (CD).
- **infrastructure** - A task related to adding low-level infrastructure or cross-cutting concerns (could be for application or deployment).
- **docs** - A task to add documentation for design, technical detail or troubleshooting instructions.

### Potential example work-in-progress **dexterity-shell** commands and shortcuts

```sh
dexterity incident create --summary="user login has broken"
dex i c "user login has broken"

dexterity incident resolve --summary="fixed config"
dex i r "fixed config"

dexterity task create --summary="add basic logging"
dex t c "add basic logging"

dexterity task delete --id=2
dex t d 2
```
