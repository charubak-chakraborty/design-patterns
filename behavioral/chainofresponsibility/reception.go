package chainofresponsibility

type Reception struct {
	Next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		r.Next.Execute(p)
	}
	p.RegistrationDone = true
	r.Next.Execute(p)
}
