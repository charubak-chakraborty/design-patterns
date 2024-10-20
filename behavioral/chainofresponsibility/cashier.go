package chainofresponsibility

type Cashier struct {
	Next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		c.Next.Execute(p)
	}
	p.PaymentDone = true
	c.Next.Execute(p)
}
