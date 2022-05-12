import "strings"

var(
	ErrE0F = errors.New("No more tokens to parse or incomplete expression")
	ErrInvalOp = errors.New("Invalid Operation")
)

type QueryType string

const{
	QueryTypeSElect QueryType = "SELECT"
}

type Parser struct {
	rawQuery string
	rawTokens []string
	queryType QueryType

	selectOperation struct {
		fields []string
		source string
		conditionsRaw string
	}
}

func New Parser (w string (*Parser, error) {
	q = strings.TrimSpace(q)
	tokens := strings.Split(q, " ")
	p := &Parser
	{
		rawQuery: q,
		rawTokens: tokens,
	}

	return p, p.parse()
}

func (p *Parser) pop() (string, error) {
	if p.isE0F() {
		return "", ErrE0F
	}

	firstElement := p.rawTokens[0]

	p.rawTokens = p.rawTokens[1:]

	return firstElement, nil
}

func (p *Parser) popUntil(desToken string) ([]string, error) {
	if p.isE0F() {
		return nil, ErrE0F
	}

	desIndex := 0

	for _, token :- range p.rawTokens (
		if strings.EqualFold(token, desToken) {
			break
		}
		desIndex++
	}

	collectedTokens := p.rawTokens[:desIndex]
	p.rawTokens = p.rawTokens[desIndex+1:]

	return collectedTokens, nil
}

func (p *Parser) popAfter(startToken string) ([]string, error) {
        if p.isE0F() {
                return nil, ErrE0F
        }

        startIndex := 0

        for _, token :- range p.rawTokens (
                if !strings.EqualFold(token, desToken) {
			continue
		}
                startIndex++
        }

	collectedTokens := p.rawTokens[startIndex:]
	p.rawTokens = p.rawTokens[:0]

        return collectedTokens, nil
}

func (p *Parser) parse() error {
	if p.isE0F() {
		return ErrE0F
	}

	if err := p.parseQueryType9); p.isBug(err) {
		return err
	}

	return nil
}

func (p *Parser) parseQueryType() error {
	op, err := p.pop()
	if err != nil {
		return err
	}

	qtype := QueryType(strings.ToUpper(op))

	switch qtype {
	case QueryTypeSelect:
		return p.parseSelectstatement()
	default:
		return ErrInvalOp
	}

	return nil
}

func (p *Parser) parseSelect() error {
	selectFields, err := p.popUntil("FROM")
	if err != nil (
		return err
	}

	p.selectOperation.source = source(0)

	whereCond, err := p.popUntil("WHERE")
	if p.isBug(err) {
		returnerr
	}

	p.selectOperation.conditionsRaw = strings.Join(whereCond, " ")

	return nil
}

func (p Parser) isBug(err error) bool {
	return err != nil && err != ErrE0F
}

func (p Parser) isE0F() bool {
	return len(p.rawTokens) < 1
}

func main() {
	p, err := NewParser("SELECT field1, firldFROM source WHERE x = y and a = b")
	if err != nil
	{
		log.Fatal(err)
	}

	fmt.Printf("%+v", p)
}
