package chainofresponsibility

type Medical struct {
	Next Department
}

func (m *Medical) Execute(p *Patient) {
	if p.MedicineDone {
		m.Next.Execute(p)
	}
	p.MedicineDone = true
	m.Next.Execute(p)
}
