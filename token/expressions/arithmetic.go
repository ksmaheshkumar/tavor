package expressions

import (
	"strconv"

	"github.com/zimmski/tavor/rand"
	"github.com/zimmski/tavor/token"
	"github.com/zimmski/tavor/token/lists"
)

type AddArithmetic struct {
	a token.Token
	b token.Token
}

func NewAddArithmetic(a, b token.Token) *AddArithmetic {
	return &AddArithmetic{
		a: a,
		b: b,
	}
}

// Clone returns a copy of the token and all its children
func (e *AddArithmetic) Clone() token.Token {
	return &AddArithmetic{
		a: e.a.Clone(),
		b: e.b.Clone(),
	}
}

func (e *AddArithmetic) Fuzz(r rand.Rand) {
	// do nothing
}

func (e *AddArithmetic) FuzzAll(r rand.Rand) {
	e.Fuzz(r)

	e.a.FuzzAll(r)
	e.b.FuzzAll(r)
}

func (e *AddArithmetic) Parse(pars *token.InternalParser, cur int) (int, []error) {
	panic("TODO implement")
}

func (e *AddArithmetic) Permutation(i uint) error {
	permutations := e.Permutations()

	if i < 1 || i > permutations {
		return &token.PermutationError{
			Type: token.PermutationErrorIndexOutOfBound,
		}
	}
	// do nothing

	return nil
}

func (e *AddArithmetic) Permutations() uint {
	return 1
}

func (e *AddArithmetic) PermutationsAll() uint {
	return e.a.PermutationsAll() * e.b.PermutationsAll()
}

func (e *AddArithmetic) String() string {
	a, err := strconv.Atoi(e.a.String())
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(e.b.String())
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(a + b)
}

// List interface methods

func (e *AddArithmetic) Get(i int) (token.Token, error) {
	switch i {
	case 0:
		return e.a, nil
	case 1:
		return e.b, nil
	default:
		return nil, &lists.ListError{
			Type: lists.ListErrorOutOfBound,
		}
	}
}

func (e *AddArithmetic) Len() int {
	return 2
}

func (e *AddArithmetic) InternalGet(i int) (token.Token, error) {
	return e.Get(i)
}

func (e *AddArithmetic) InternalLen() int {
	return e.Len()
}

// InternalLogicalRemove removes the referenced internal token and returns the replacement for the current token or nil if the current token should be removed.
func (e *AddArithmetic) InternalLogicalRemove(tok token.Token) token.Token {
	if tok == e.a || tok == e.b {
		return nil
	}

	return e
}

// InternalReplace replaces an old with a new internal token if it is referenced by this token
func (e *AddArithmetic) InternalReplace(oldToken, newToken token.Token) {
	if oldToken == e.a {
		e.a = newToken
	}
	if oldToken == e.b {
		e.b = newToken
	}
}

type SubArithmetic struct {
	a token.Token
	b token.Token
}

func NewSubArithmetic(a, b token.Token) *SubArithmetic {
	return &SubArithmetic{
		a: a,
		b: b,
	}
}

// Clone returns a copy of the token and all its children
func (e *SubArithmetic) Clone() token.Token {
	return &SubArithmetic{
		a: e.a.Clone(),
		b: e.b.Clone(),
	}
}

func (e *SubArithmetic) Fuzz(r rand.Rand) {
	// do nothing
}

func (e *SubArithmetic) FuzzAll(r rand.Rand) {
	e.Fuzz(r)

	e.a.FuzzAll(r)
	e.b.FuzzAll(r)
}

func (e *SubArithmetic) Parse(pars *token.InternalParser, cur int) (int, []error) {
	panic("TODO implement")
}

func (e *SubArithmetic) Permutation(i uint) error {
	permutations := e.Permutations()

	if i < 1 || i > permutations {
		return &token.PermutationError{
			Type: token.PermutationErrorIndexOutOfBound,
		}
	}
	// do nothing

	return nil
}

func (e *SubArithmetic) Permutations() uint {
	return 1
}

func (e *SubArithmetic) PermutationsAll() uint {
	return e.a.PermutationsAll() * e.b.PermutationsAll()
}

func (e *SubArithmetic) String() string {
	a, err := strconv.Atoi(e.a.String())
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(e.b.String())
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(a - b)
}

// List interface methods

func (e *SubArithmetic) Get(i int) (token.Token, error) {
	switch i {
	case 0:
		return e.a, nil
	case 1:
		return e.b, nil
	default:
		return nil, &lists.ListError{
			Type: lists.ListErrorOutOfBound,
		}
	}
}

func (e *SubArithmetic) Len() int {
	return 2
}

func (e *SubArithmetic) InternalGet(i int) (token.Token, error) {
	return e.Get(i)
}

func (e *SubArithmetic) InternalLen() int {
	return e.Len()
}

// InternalLogicalRemove removes the referenced internal token and returns the replacement for the current token or nil if the current token should be removed.
func (e *SubArithmetic) InternalLogicalRemove(tok token.Token) token.Token {
	if tok == e.a || tok == e.b {
		return nil
	}

	return e
}

// InternalReplace replaces an old with a new internal token if it is referenced by this token
func (e *SubArithmetic) InternalReplace(oldToken, newToken token.Token) {
	if oldToken == e.a {
		e.a = newToken
	}
	if oldToken == e.b {
		e.b = newToken
	}
}

type MulArithmetic struct {
	a token.Token
	b token.Token
}

func NewMulArithmetic(a, b token.Token) *MulArithmetic {
	return &MulArithmetic{
		a: a,
		b: b,
	}
}

// Clone returns a copy of the token and all its children
func (e *MulArithmetic) Clone() token.Token {
	return &MulArithmetic{
		a: e.a.Clone(),
		b: e.b.Clone(),
	}
}

func (e *MulArithmetic) Fuzz(r rand.Rand) {
	// do nothing
}

func (e *MulArithmetic) FuzzAll(r rand.Rand) {
	e.Fuzz(r)

	e.a.FuzzAll(r)
	e.b.FuzzAll(r)
}

func (e *MulArithmetic) Parse(pars *token.InternalParser, cur int) (int, []error) {
	panic("TODO implement")
}

func (e *MulArithmetic) Permutation(i uint) error {
	permutations := e.Permutations()

	if i < 1 || i > permutations {
		return &token.PermutationError{
			Type: token.PermutationErrorIndexOutOfBound,
		}
	}
	// do nothing

	return nil
}

func (e *MulArithmetic) Permutations() uint {
	return 1
}

func (e *MulArithmetic) PermutationsAll() uint {
	return e.a.PermutationsAll() * e.b.PermutationsAll()
}

func (e *MulArithmetic) String() string {
	a, err := strconv.Atoi(e.a.String())
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(e.b.String())
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(a * b)
}

// List interface methods

func (e *MulArithmetic) Get(i int) (token.Token, error) {
	switch i {
	case 0:
		return e.a, nil
	case 1:
		return e.b, nil
	default:
		return nil, &lists.ListError{
			Type: lists.ListErrorOutOfBound,
		}
	}
}

func (e *MulArithmetic) Len() int {
	return 2
}

func (e *MulArithmetic) InternalGet(i int) (token.Token, error) {
	return e.Get(i)
}

func (e *MulArithmetic) InternalLen() int {
	return e.Len()
}

// InternalLogicalRemove removes the referenced internal token and returns the replacement for the current token or nil if the current token should be removed.
func (e *MulArithmetic) InternalLogicalRemove(tok token.Token) token.Token {
	if tok == e.a || tok == e.b {
		return nil
	}

	return e
}

// InternalReplace replaces an old with a new internal token if it is referenced by this token
func (e *MulArithmetic) InternalReplace(oldToken, newToken token.Token) {
	if oldToken == e.a {
		e.a = newToken
	}
	if oldToken == e.b {
		e.b = newToken
	}
}

type DivArithmetic struct {
	a token.Token
	b token.Token
}

func NewDivArithmetic(a, b token.Token) *DivArithmetic {
	return &DivArithmetic{
		a: a,
		b: b,
	}
}

// Clone returns a copy of the token and all its children
func (e *DivArithmetic) Clone() token.Token {
	return &DivArithmetic{
		a: e.a.Clone(),
		b: e.b.Clone(),
	}
}

func (e *DivArithmetic) Fuzz(r rand.Rand) {
	// do nothing
}

func (e *DivArithmetic) FuzzAll(r rand.Rand) {
	e.Fuzz(r)

	e.a.FuzzAll(r)
	e.b.FuzzAll(r)
}

func (e *DivArithmetic) Parse(pars *token.InternalParser, cur int) (int, []error) {
	panic("TODO implement")
}

func (e *DivArithmetic) Permutation(i uint) error {
	permutations := e.Permutations()

	if i < 1 || i > permutations {
		return &token.PermutationError{
			Type: token.PermutationErrorIndexOutOfBound,
		}
	}
	// do nothing

	return nil
}

func (e *DivArithmetic) Permutations() uint {
	return 1
}

func (e *DivArithmetic) PermutationsAll() uint {
	return e.a.PermutationsAll() * e.b.PermutationsAll()
}

func (e *DivArithmetic) String() string {
	a, err := strconv.Atoi(e.a.String())
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(e.b.String())
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(a / b)
}

// List interface methods

func (e *DivArithmetic) Get(i int) (token.Token, error) {
	switch i {
	case 0:
		return e.a, nil
	case 1:
		return e.b, nil
	default:
		return nil, &lists.ListError{
			Type: lists.ListErrorOutOfBound,
		}
	}
}

func (e *DivArithmetic) Len() int {
	return 2
}

func (e *DivArithmetic) InternalGet(i int) (token.Token, error) {
	return e.Get(i)
}

func (e *DivArithmetic) InternalLen() int {
	return e.Len()
}

// InternalLogicalRemove removes the referenced internal token and returns the replacement for the current token or nil if the current token should be removed.
func (e *DivArithmetic) InternalLogicalRemove(tok token.Token) token.Token {
	if tok == e.a || tok == e.b {
		return nil
	}

	return e
}

// InternalReplace replaces an old with a new internal token if it is referenced by this token
func (e *DivArithmetic) InternalReplace(oldToken, newToken token.Token) {
	if oldToken == e.a {
		e.a = newToken
	}
	if oldToken == e.b {
		e.b = newToken
	}
}
