package maze

import (
	"io"
	"encoding/gob"
	"errors"
)

// The record holding info about individual maze agent performance at the end of simulation
type AgentRecord struct {
	// The ID of agent
	AgentID    int
	// The agent position at the end of simulation
	X, Y       float64
	// The agent fitness
	Fitness    float64
	// The flag to indicate whether agent reached maze exit
	GotExit    bool
	// The population generation when agent data was collected
	Generation int
	// The novelty value associated
	Novelty    float64

	// The ID of species to whom individual belongs
	SpeciesID  int
	// The age of species to whom individual belongs at time of recording
	SpeciesAge int
}

// The maze agent records storage
type RecordStore struct {
	// The array of agent records
	Records          []AgentRecord
	// The array of the solver agent path points
	SolverPathPoints []Point
}

// Writes record store to the provided writer
func (r *RecordStore) Write(w io.Writer) error {
	if len(r.Records) == 0 {
		return errors.New("No records to store")
	}
	enc := gob.NewEncoder(w)
	err := enc.Encode(r)
	return err
}

// Reads record store data from provided reader
func (rs *RecordStore) Read(r io.Reader) error {
	dec := gob.NewDecoder(r)
	err := dec.Decode(&rs)
	return err
}