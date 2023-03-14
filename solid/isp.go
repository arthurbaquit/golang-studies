package solid

// The Interface Segregation Principle
//
// The Interface Segregation Principle (ISP) states that no client should be
// forced to depend on methods it does not use. In other words, make fine
// grained interfaces that are client specific.

type Document struct {
}

type Machine interface {
	Print(d *Document)
	Fax(d *Document)
	Scan(d *Document)
}

type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(d *Document) {
	// implement something
}

func (m *MultiFunctionPrinter) Fax(d *Document) {
	// implement something
}

func (m *MultiFunctionPrinter) Scan(d *Document) {
	// implement something
}

type OldFashionedPrinter struct {
}

func (o *OldFashionedPrinter) Print(d *Document) {
	// implement something
}

func (o *OldFashionedPrinter) Fax(d *Document) {
	panic("operation not supported")
}

func (o *OldFashionedPrinter) Scan(d *Document) {
	panic("operation not supported")
}

// As it can be seen, the OldFashionedPrinter does not support the Fax and Scan,
// but the MultiFunctionPrinter does. So, the MultiFunctionPrinter is a subtype of
// the OldFashionedPrinter. In this sense, the Interface Segregation Principle
// is violated, because the OldFashionedPrinter does not need to implement the
// Fax and Scan methods. We can fix this by creating a new interface that
// contains only the Print method, and then the OldFashionedPrinter can
// implement this interface, and the MultiFunctionPrinter can implement the
// Machine interface, which extends the Printer.

type Printer interface {
	Print(d *Document)
}
type Scanner interface {
	Scan(d *Document)
}

type MyPrinter struct {
}

func (m *MyPrinter) Print(d *Document) {
	// implement something
}

type Photocopier struct {
	printer Printer
	scanner Scanner
}

// In this way, the Photocopier can implement both the Printer and Scanner
// interfaces, and the OldFashionedPrinter can implement only the Printer
// interface.
