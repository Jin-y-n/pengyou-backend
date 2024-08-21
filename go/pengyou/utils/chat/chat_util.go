package chat

var establishRequestNode = make(map[struct {
	From string
	To   string
}]bool)

// AddEstablishRequestNode Adds an entry to the establishRequestNode map
func AddEstablishRequestNode(to, from string) {
	establishRequestNode[struct {
		From string
		To   string
	}{from, to}] = false
}

func GetEstablishRequestNode(to, from string) bool {
	return establishRequestNode[struct {
		From string
		To   string
	}{from, to}]
}

// SetEstablishRequestNode Sets an entry in the establishRequestNode map
func SetEstablishRequestNode(to, from string, value bool) {
	establishRequestNode[struct {
		From string
		To   string
	}{from, to}] = value
}

// RemoveEstablishRequestNode Removes an entry from the establishRequestNode map
func RemoveEstablishRequestNode(to, from string) {
	delete(establishRequestNode, struct {
		From string
		To   string
	}{from, to})
}
