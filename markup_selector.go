package ewa

import "github.com/apex/log"

//Waves makes waves structure
func (m *Markup) Waves() Waves {
	return Waves{Impulses: m.Impulses, Corrections: m.Corrections}
}

//FromTo finds waves that start and end at specified price
func (in Waves) FromTo(from, to float64) (out Waves) {
	for _, one := range in.Impulses {
		if one.Starts() == from && one.Tops() == to {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Starts() == from && one.Tops() == to {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//From finds waves that start there
func (in Waves) From(from float64) (out Waves) {

	for _, one := range in.Impulses {
		if one.Starts() == from {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Starts() == from {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//To finds waves that end there
func (in Waves) To(to float64) (out Waves) {

	for _, one := range in.Impulses {
		if one.Tops() == to {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Tops() == to {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//Sub finds waves that have subdivision
func (in Waves) Sub(has bool) (out Waves) {

	for _, one := range in.Impulses {
		if one.Sub() == has {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Sub() == has {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

func (in Waves) dir(dir bool) (out Waves) {

	for _, one := range in.Impulses {
		if one.Up() == dir {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Up() == dir {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//Up finds waves that go up
func (in Waves) Up() (out Waves) {
	return in.dir(true)
}

//Down finds waves that go down
func (in Waves) Down() (out Waves) {
	return in.dir(false)
}

//Degree finds waves that end there
func (in Waves) Degree(degree DegreeType) (out Waves) {

	for _, one := range in.Impulses {
		if one.Degree == degree {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Degree == degree {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//DegreeGreaterOr finds waves that degree GreaterOr
func (in Waves) DegreeGreaterOr(degree DegreeType) (out Waves) {

	for _, one := range in.Impulses {
		if one.Degree >= degree {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Degree >= degree {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//DegreeGreater finds waves that degree GreaterOr
func (in Waves) DegreeGreater(degree DegreeType) (out Waves) {

	for _, one := range in.Impulses {
		if one.Degree > degree {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Degree > degree {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//DegreeLessOr finds waves that degree GreaterOr
func (in Waves) DegreeLessOr(degree DegreeType) (out Waves) {

	for _, one := range in.Impulses {
		if one.Degree <= degree {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Degree <= degree {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//DegreeLess finds waves that degree GreaterOr
func (in Waves) DegreeLess(degree DegreeType) (out Waves) {

	for _, one := range in.Impulses {
		if one.Degree < degree {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Degree < degree {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//Print selector
func (in Waves) Print() {
	for _, one := range in.Impulses {
		log.WithFields(log.Fields{
			"M": one.Move,
			"D": one.Degree,
		}).Info("Impulse")
	}

	for _, one := range in.Corrections {
		log.WithFields(log.Fields{
			"M": one.Move,
			"D": one.Degree,
			"T": one.Type(),
		}).Info("Correction")
	}

	return
}

//Imp gets only impulses
func (in Waves) Imp() Impulses {
	return in.Impulses
}

//Extended gets only extended impulses
func (in Impulses) Extended(extended bool) (out Impulses) {

	for _, one := range in {
		if one.Extended() == extended {
			out = append(out, one)
		}
	}

	return
}

//Diagonal gets only diagonal impulses
func (in Impulses) Diagonal(diagonal bool) (out Impulses) {

	for _, one := range in {
		if one.Diagonal() == diagonal {
			out = append(out, one)
		}
	}

	return
}

//Type gets corrections by type
func (in Corrections) Type(ct CorrectionType) (out Corrections) {

	for _, one := range in {
		if one.Type() == ct {
			out = append(out, one)
		}
	}

	return
}

//Zigzag corrections only
func (in Corrections) Zigzag() (out Corrections) {
	return in.Type(CTZigzag)
}

//Flat corrections only
func (in Corrections) Flat() (out Corrections) {
	return in.Type(CTFlat)
}

//Triangle corrections only
func (in Corrections) Triangle() (out Corrections) {
	return in.Type(CTTriangle)
}

//Combo corrections only
func (in Corrections) Combo() (out Corrections) {
	return in.Type(CTCombo)
}

//Triple corrections only
func (in Corrections) Triple() (out Corrections) {
	return in.Type(CTTriple)
}

//Corr gets only impulses
func (in Waves) Corr() Corrections {
	return in.Corrections
}

//First gets first impulse
func (in Impulses) First() (*Impulse, bool) {
	if len(in) > 0 {
		return in[0], true
	}

	return nil, false
}

//Last gets first impulse
func (in Impulses) Last() (*Impulse, bool) {
	if len(in) > 1 {
		return in[len(in)-1], true
	}

	if len(in) == 1 {
		return in[0], true
	}

	return nil, false
}

//First gets first correction
func (in Corrections) First() (*Correction, bool) {
	if len(in) > 0 {
		return in[0], true
	}

	return nil, false
}

//Last gets first correction
func (in Corrections) Last() (*Correction, bool) {
	if len(in) > 1 {
		return in[len(in)-1], true
	}

	if len(in) == 1 {
		return in[0], true
	}

	return nil, false
}
