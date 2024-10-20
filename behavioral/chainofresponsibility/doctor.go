package chainofresponsibility

type Doctor struct {
	Next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckupDone {
		d.Next.Execute(p)
	}
	p.DoctorCheckupDone = true
	d.Next.Execute(p)
}
