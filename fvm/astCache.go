package fvm

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
	"github.com/onflow/cadence/runtime/ast"
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/sema"
)

// ASTCache is an interface to a cache for parsed program ASTs.
type ASTCache interface {

	// GetProgram returns a program AST from the cache
	GetProgram(common.Location) (*ast.Program, error)

	// SetProgram adds a program AST to the cache.
	SetProgram(common.Location, *ast.Program) error

	// GetProgram returns an elaboration from the cache.
	GetElaboration(common.Location) (*sema.Checker, error)

	// SetElaboration adds an elaboration to the cache.
	SetElaboration(common.Location, *sema.Checker) error

	// Clear removes all the entries from the cache.
	Clear()
}

// LRUASTCache implements a program AST cache with a LRU cache.
type LRUASTCache struct {
	programs     *lru.Cache
	elaborations *lru.Cache
}

// NewLRUASTCache creates a new LRU cache that implements the ASTCache interface.
func NewLRUASTCache(size int) (*LRUASTCache, error) {
	programs, err := lru.New(size)
	if err != nil {
		return nil, fmt.Errorf("failed to create program cache, %w", err)
	}

	elaborations, err := lru.New(size)
	if err != nil {
		return nil, fmt.Errorf("failed to create elaboration cache, %w", err)
	}

	return &LRUASTCache{
		programs:     programs,
		elaborations: elaborations,
	}, nil
}

// GetProgram retrieves a program AST from the LRU cache.
func (cache *LRUASTCache) GetProgram(location common.Location) (*ast.Program, error) {
	program, found := cache.programs.Get(location.ID())
	if found {
		cachedProgram, ok := program.(*ast.Program)
		if !ok {
			return nil, fmt.Errorf("could not convert cached value to ast.Program")
		}
		// Return a new program to clear importedPrograms.
		// This will avoid a concurrent map write when attempting to
		// resolveImports.
		return &ast.Program{Declarations: cachedProgram.Declarations}, nil
	}
	return nil, nil
}

// SetProgram adds a program AST to the LRU cache.
func (cache *LRUASTCache) SetProgram(location common.Location, program *ast.Program) error {
	_ = cache.programs.Add(location.ID(), program)
	return nil
}

// GetChecker retrieves an elaboration from the LRU cache.
func (cache *LRUASTCache) GetElaboration(location common.Location) (*sema.Checker, error) {
	program, found := cache.elaborations.Get(location.ID())
	if found {
		checker, ok := program.(*sema.Checker)
		if !ok {
			return nil, fmt.Errorf("could not convert cached value to sema.Checker")
		}

		// Return a new elaborations to avoid any concurrency issues
		// TODO: replace with elaboration
		tempChecker := &sema.Checker{
			Program:           checker.Program,
			Location:          checker.Location,
			PredeclaredValues: checker.PredeclaredValues,
			PredeclaredTypes:  checker.PredeclaredTypes,
			AllCheckers:       checker.AllCheckers,
			GlobalValues:      checker.GlobalValues,
			GlobalTypes:       checker.GlobalTypes,
			TransactionTypes:  checker.TransactionTypes,
			Occurrences:       checker.Occurrences,
			MemberAccesses:    checker.MemberAccesses,
			Elaboration:       checker.Elaboration,
		}

		return tempChecker, nil
	}
	return nil, nil
}

// SetChecker adds an elaboration to the LRU cache.
func (cache *LRUASTCache) SetElaboration(location common.Location, checker *sema.Checker) error {
	_ = cache.elaborations.Add(location.ID(), checker)
	return nil
}

func (cache *LRUASTCache) Clear() {
	cache.elaborations.Purge()
	cache.programs.Purge()
}
