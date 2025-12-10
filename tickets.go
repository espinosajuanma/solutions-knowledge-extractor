package notebook

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/espinosajuanma/solutions-knowledge-extractor/parser"
)

const (
	SIZE = "1000"
)

type Pool struct {
	Projects []Relationship `json:"projects"`
}

type Ticket struct {
	Number        int           `json:"number"`
	Title         string        `json:"title"`
	Project       Relationship  `json:"project"`
	Status        string        `json:"status"`
	Priority      string        `json:"priority"`
	Assignee      *Relationship `json:"assignee"` // Pointer to handle null/nil
	Customer      Relationship  `json:"customer"`
	Description   string        `json:"description"`
	Attachments   []Attachment  `json:"attachments"`
	Notes         []Note        `json:"notes"`
	InternalNotes []Note        `json:"internalNotes"`
}

type Attachment struct {
	File        File         `json:"file"`
	Description string       `json:"description"`
	Link        string       `json:"link"`
	ID          string       `json:"id"`
	Label       string       `json:"label"`
	AddedBy     Relationship `json:"addedBy"`
}

type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Note struct {
	Label string `json:"label"`
	Note  string `json:"note"`
}

func (s *Solutions) GetTicketsByPoolName(name string, outputMode string) (string, error) {
	if outputMode == "" {
		outputMode = "markdown"
	}
	outputMode = strings.ToLower(outputMode)
	if outputMode != "markdown" && outputMode != "html" {
		return "", fmt.Errorf("invalid output mode. Use 'markdown' or 'html'")
	}

	poolParams := map[string]string{
		"name":  name,
		"_size": "1",
	}
	poolResBytes, err := s.App.GetRecords("support.pools", poolParams)
	if err != nil {
		return "", fmt.Errorf("failed to fetch pools: %w", err)
	}

	var pools ManyResponse[Pool]
	err = json.Unmarshal(poolResBytes, &pools)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal pools: %w", err)
	}

	if len(pools.Items) == 0 {
		return "", fmt.Errorf("pool '%s' not found", name)
	}

	var projectIDs []string
	for _, p := range pools.Items[0].Projects {
		projectIDs = append(projectIDs, p.ID)
	}
	joinedIDs := strings.Join(projectIDs, ",")

	fmt.Println(joinedIDs)

	ticketParams := map[string]string{
		"project":    joinedIDs,
		"_size":      SIZE,
		"_sortField": "createdTimestamp",
		"_sortType":  "desc",
	}

	ticketResBytes, err := s.App.GetRecords("support.tickets", ticketParams)
	if err != nil {
		return "", fmt.Errorf("failed to fetch tickets: %w", err)
	}

	var ticketsResponse ManyResponse[Ticket]
	err = json.Unmarshal(ticketResBytes, &ticketsResponse)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal tickets: %w", err)
	}

	if outputMode == "html" {
		return parser.ToHTML("tickets", ticketsResponse.Items)
	}
	return parser.ToMarkdown("tickets", ticketsResponse.Items)
}
