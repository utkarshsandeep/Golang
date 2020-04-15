package organization

import (
	"errors"
	"fmt"
	"strings"
)

//TwitterHandler ...
//This will replace all Strings to TwitterHandler
type TwitterHandler string

/*Handler ...
type TwitterHandlerString = string

RedirectURL Since TwitterHandlerString is already a alias of string so it can't create a function of its type so to remove this remove =

func (th TwitterHandlerString) RedirectURL() string{
	return ""
}
*/
//Handler ...
type Handler struct {
	handle string
	name   string
}

//RedirectURL ...
func (th TwitterHandler) RedirectURL() string {
	CleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", CleanHandler)
}

/*
******************************************************************************************
 */
//Identifiable .
type Identifiable interface {
	ID() string
}

/*
********************************************************************************
 */
//Name ...
type Name struct {
	first string
	last  string
}

//EmbeddFullName ...
func (n Name) EmbeddFullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

//Employee ...
type Employee struct {
	firstName string
	lastName  string
}

type structEmbeddEmployee struct {
	name Name
}

//Person ..
type Person struct {
	firstName      string
	lastName       string
	twitterHandler TwitterHandler
}

type structEmbeddPerson struct {
	name           Name
	twitterHandler TwitterHandler
}

//NewFunction ...
func NewFunction(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func structEmbeddNewFunction(firstName, lastName string) Person {
	return structEmbeddPerson{
		name: Name{
			first: firstName,
			last:  lastName,
		},
	}
}

//FullName ...
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

//ID ...
func (p Person) ID() string {
	return "12345678"
}

//SetTwitterHandler ...
func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("Twitter Handle must start with an @ sumbol")
	}
	p.twitterHandler = handler
	return nil
}

//TwitterHandler ...
func (p Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
