package primitives

import (
	"fmt"
	"math"
	"strconv"

	"github.com/zimmski/tavor/rand"
	"github.com/zimmski/tavor/token"
)

// ConstantInt implements an integer token which holds a constant integer
type ConstantInt struct {
	value int
}

// NewConstantInt returns a new instance of a ConstantInt token
func NewConstantInt(value int) *ConstantInt {
	return &ConstantInt{
		value: value,
	}
}

// SetValue sets the value of the token
func (p *ConstantInt) SetValue(v int) {
	p.value = v
}

// Value returns the value of the token
func (p *ConstantInt) Value() int {
	return p.value
}

// Token interface methods

// Clone returns a copy of the token and all its children
func (p *ConstantInt) Clone() token.Token {
	return &ConstantInt{
		value: p.value,
	}
}

// Fuzz fuzzes this token using the random generator by choosing one of the possible permutations for this token
func (p *ConstantInt) Fuzz(r rand.Rand) {
	// do nothing
}

// FuzzAll calls Fuzz for this token and then FuzzAll for all children of this token
func (p *ConstantInt) FuzzAll(r rand.Rand) {
	p.Fuzz(r)
}

// Parse tries to parse the token beginning from the current position in the parser data.
// If the parsing is successful the error argument is nil and the next current position after the token is returned.
func (p *ConstantInt) Parse(pars *token.InternalParser, cur int) (int, []error) {
	v := strconv.Itoa(p.value)
	vLen := len(v)

	nextIndex := vLen + cur

	if nextIndex > pars.DataLen {
		return cur, []error{&token.ParserError{
			Message: fmt.Sprintf("expected %q but got early EOF", v),
			Type:    token.ParseErrorUnexpectedEOF,
		}}
	}

	if got := pars.Data[cur:nextIndex]; v != got {
		return cur, []error{&token.ParserError{
			Message: fmt.Sprintf("expected %q but got %q", v, got),
			Type:    token.ParseErrorUnexpectedData,
		}}
	}

	return nextIndex, nil
}

// Permutation sets a specific permutation for this token
func (p *ConstantInt) Permutation(i uint) error {
	permutations := p.Permutations()

	if i < 1 || i > permutations {
		return &token.PermutationError{
			Type: token.PermutationErrorIndexOutOfBound,
		}
	}

	// do nothing

	return nil
}

// Permutations returns the number of permutations for this token
func (p *ConstantInt) Permutations() uint {
	return 1
}

// PermutationsAll returns the number of all possible permutations for this token including its children
func (p *ConstantInt) PermutationsAll() uint {
	return p.Permutations()
}

func (p *ConstantInt) String() string {
	return strconv.Itoa(p.value)
}

// RandomInt implements an integer token which holds a random integer which gets newly generated on every permutation
type RandomInt struct {
	value int
}

// NewRandomInt returns a new instance of a RandomInt token
func NewRandomInt() *RandomInt {
	return &RandomInt{
		value: 0,
	}
}

// Clone returns a copy of the token and all its children
func (p *RandomInt) Clone() token.Token {
	return &RandomInt{
		value: p.value,
	}
}

// Fuzz fuzzes this token using the random generator by choosing one of the possible permutations for this token
func (p *RandomInt) Fuzz(r rand.Rand) {
	p.value = r.Int()
}

// FuzzAll calls Fuzz for this token and then FuzzAll for all children of this token
func (p *RandomInt) FuzzAll(r rand.Rand) {
	p.Fuzz(r)
}

// Parse tries to parse the token beginning from the current position in the parser data.
// If the parsing is successful the error argument is nil and the next current position after the token is returned.
func (p *RandomInt) Parse(pars *token.InternalParser, cur int) (int, []error) {
	panic("TODO implement")
}

// Permutation sets a specific permutation for this token
func (p *RandomInt) Permutation(i uint) error {
	permutations := p.Permutations()

	if i < 1 || i > permutations {
		return &token.PermutationError{
			Type: token.PermutationErrorIndexOutOfBound,
		}
	}

	// TODO this could be done MUCH better
	p.value = 0

	return nil
}

// Permutations returns the number of permutations for this token
func (p *RandomInt) Permutations() uint {
	return 1 // TODO maybe this should be like RangeInt
}

// PermutationsAll returns the number of all possible permutations for this token including its children
func (p *RandomInt) PermutationsAll() uint {
	return p.Permutations()
}

func (p *RandomInt) String() string {
	return strconv.Itoa(p.value)
}

// RangeInt implements an integer token holding a range of integers
// Every permutation generates a new value within the defined range and step. For example the range 1 to 10 with step 2 can hold the integers 1, 3, 5, 7 and 9.
type RangeInt struct {
	from int
	to   int
	step int

	value int
}

// NewRangeInt returns a new instance of a RangeInt token with the given range and step value of 1
func NewRangeInt(from, to int) *RangeInt {
	if from > to {
		panic("TODO implement that From can be bigger than To")
	}

	return &RangeInt{
		from: from,
		to:   to,
		step: 1,

		value: from,
	}
}

// NewRangeIntWithStep returns a new instance of a RangeInt token with the given range and step value
func NewRangeIntWithStep(from, to, step int) *RangeInt {
	if from > to {
		panic("TODO implement that From can be bigger than To")
	}
	if step < 1 {
		panic("TODO implement 0 and negative step")
	}

	return &RangeInt{
		from: from,
		to:   to,
		step: step,

		value: from,
	}
}

// From returns the from value of the range
func (p *RangeInt) From() int {
	return p.from
}

// To returns the to value of the range
func (p *RangeInt) To() int {
	return p.to
}

// Step returns the step value
func (p *RangeInt) Step() int {
	return p.step
}

// Token interface methods

// Clone returns a copy of the token and all its children
func (p *RangeInt) Clone() token.Token {
	return &RangeInt{
		from: p.from,
		to:   p.to,
		step: p.step,

		value: p.value,
	}
}

// Fuzz fuzzes this token using the random generator by choosing one of the possible permutations for this token
func (p *RangeInt) Fuzz(r rand.Rand) {
	i := r.Int63n(int64(p.Permutations()))

	p.permutation(uint(i))
}

// FuzzAll calls Fuzz for this token and then FuzzAll for all children of this token
func (p *RangeInt) FuzzAll(r rand.Rand) {
	p.Fuzz(r)
}

// Parse tries to parse the token beginning from the current position in the parser data.
// If the parsing is successful the error argument is nil and the next current position after the token is returned.
func (p *RangeInt) Parse(pars *token.InternalParser, cur int) (int, []error) {
	if cur == pars.DataLen {
		return cur, []error{&token.ParserError{
			Message: fmt.Sprintf("expected integer in range %d-%d with step %d but got early EOF", p.from, p.to, p.step),
			Type:    token.ParseErrorUnexpectedEOF,
		}}
	}

	i := cur
	v := ""

	for {
		c := pars.Data[i]

		if c < '0' || c > '9' {
			break
		}

		v += string(c)

		if ci, _ := strconv.Atoi(v); ci > p.to {
			v = v[:len(v)-1] // remove last digit

			break
		}

		i++

		if i == pars.DataLen {
			break
		}
	}

	i--

	ci, _ := strconv.Atoi(v)

	if v == "" || (ci < p.from || ci > p.to) || ci%p.step != 0 {
		// is the first character already invalid
		if i < cur {
			i = cur
		}

		return cur, []error{&token.ParserError{
			Message: fmt.Sprintf("expected integer in range %d-%d with step %d but got %q", p.from, p.to, p.step, pars.Data[cur:i]),
			Type:    token.ParseErrorUnexpectedData,
		}}
	}

	p.value = ci

	return i + 1, nil
}

func (p *RangeInt) permutation(i uint) {
	p.value = p.from + (int(i) * p.step)
}

// Permutation sets a specific permutation for this token
func (p *RangeInt) Permutation(i uint) error {
	permutations := p.Permutations()

	if i < 1 || i > permutations {
		return &token.PermutationError{
			Type: token.PermutationErrorIndexOutOfBound,
		}
	}

	p.permutation(i - 1)

	return nil
}

// Permutations returns the number of permutations for this token
func (p *RangeInt) Permutations() uint {
	// TODO FIXME this
	perms := (p.to-p.from)/p.step + 1

	if perms < 0 {
		return math.MaxUint32
	}

	return uint(perms)
}

// PermutationsAll returns the number of all possible permutations for this token including its children
func (p *RangeInt) PermutationsAll() uint {
	return p.Permutations()
}

func (p *RangeInt) String() string {
	return strconv.Itoa(p.value)
}
