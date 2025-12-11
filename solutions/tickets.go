package solutions

import (
	"encoding/json"
	"fmt"
	"strings"
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

func (s *Solutions) GetPoolByName(name string) (Pool, error) {
	var pool Pool
	poolParams := map[string]string{
		"name":  name,
		"_size": "1",
	}
	poolResBytes, err := s.App.GetRecords("support.pools", poolParams)
	if err != nil {
		return pool, fmt.Errorf("failed to fetch pools: %w", err)
	}

	var pools ManyResponse[Pool]
	err = json.Unmarshal(poolResBytes, &pools)
	if err != nil {
		return pool, fmt.Errorf("failed to unmarshal pools: %w", err)
	}

	if len(pools.Items) == 0 {
		return pool, fmt.Errorf("pool '%s' not found", name)
	}

	return pools.Items[0], nil
}

func (s *Solutions) GetTicketsByPool(pool Pool) ([]Ticket, error) {
	var projectIDs []string
	for _, p := range pool.Projects {
		projectIDs = append(projectIDs, p.ID)
	}
	joinedIDs := strings.Join(projectIDs, ",")

	ticketParams := map[string]string{
		"project":    joinedIDs,
		"_size":      SIZE,
		"_sortField": "createdTimestamp",
		"_sortType":  "desc",
	}

	ticketResBytes, err := s.App.GetRecords("support.tickets", ticketParams)
	if err != nil {
		return []Ticket{}, fmt.Errorf("failed to fetch tickets: %w", err)
	}

	var ticketsResponse ManyResponse[Ticket]
	err = json.Unmarshal(ticketResBytes, &ticketsResponse)
	if err != nil {
		return []Ticket{}, fmt.Errorf("failed to unmarshal tickets: %w", err)
	}
	return ticketsResponse.Items, nil
}
