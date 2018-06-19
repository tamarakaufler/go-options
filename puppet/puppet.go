package puppet

type hair struct {
	colour, length, style string
}

type Puppet struct {
	skin, eyes     string
	hair           hair
	height, weight float32
}

type Features struct {
	Skin, Eyes                        string
	HairColour, HairLength, HairStyle string
	Height, Weight                    float32
}

func New(f Features) *Puppet {
	if f.Skin == "" {
		f.Skin = "olive"
	}
	if f.Eyes == "" {
		f.Eyes = "blue"
	}
	if f.HairColour == "" {
		f.HairColour = "blond"
	}
	if f.HairLength == "" {
		f.HairLength = "long"
	}
	if f.HairStyle == "" {
		f.HairStyle = "pony tail"
	}
	if f.Height == 0 {
		f.Height = 165
	}
	if f.Weight == 0 {
		f.Weight = 60
	}

	return &Puppet{
		skin: f.Skin,
		eyes: f.Eyes,
		hair: hair{
			colour: f.HairColour,
			length: f.HairLength,
			style:  f.HairStyle,
		},
		height: f.Height,
		weight: f.Weight,
	}
}

type option func(f *Puppet) option

func (p *Puppet) Hair() hair {
	return p.hair
}

func (p *Puppet) Skin() string {
	return p.skin
}

func (p *Puppet) Eyes() string {
	return p.eyes
}

func (p *Puppet) Height() float32 {
	return p.height
}

func (p *Puppet) Weight() float32 {
	return p.weight
}

// Functions provided for feature update.
// These are applied through the Option, Options or
// Options2 methods. They do the feature update
// and return a closure (for access to the Puppet instance),
// that, when run, restores the previous value.
func SetHairColour(c string) option {
	return func(p *Puppet) option {
		prevHair := p.hair
		p.hair.colour = c

		return SetHairColour(prevHair.colour)
	}
}

func SetHairLength(l string) option {
	return func(p *Puppet) option {
		prevHair := p.hair
		p.hair.length = l

		return SetHairLength(prevHair.length)
	}
}

func SetHairStyle(s string) option {
	return func(p *Puppet) option {
		prevHair := p.hair
		p.hair.style = s

		return SetHairStyle(prevHair.style)
	}
}

func SetSkin(s string) option {
	return func(p *Puppet) option {
		prevSkin := p.skin
		p.skin = s

		return SetSkin(prevSkin)
	}
}

func SetEyes(e string) option {
	return func(p *Puppet) option {
		prevEyes := p.eyes
		p.eyes = e

		return SetEyes(prevEyes)
	}
}

func SetHeight(h float32) option {
	return func(p *Puppet) option {
		prevHeight := p.height
		p.height = h

		return SetHeight(prevHeight)
	}
}

func SetWeight(w float32) option {
	return func(p *Puppet) option {
		prevWeight := p.weight
		p.weight = w

		return SetWeight(prevWeight)
	}
}

// Option accepts a predeclared option function and outputs
// the overriden one
func (p *Puppet) Option(opt option) option {
	return opt(p)
}

// Options provides a possibility of a bulk update.
// It returns a function that restores back all features,
// that were updated, to values before the update.
func (p *Puppet) Options(opts ...option) (restore option) {
	for i, opt := range opts {
		if i == 0 {
			restore = opt(p)
		} else {
			// --> method implementation of merge
			//restore = restore.merge(opt(p))
			// --> function implementation of merge
			restore = merge(restore, opt(p))
		}
	}
	return restore
}

// Options provides a possibility of a bulk update.
// It returns a function that restores back all features,
// that were updated, to values before the update.
func (p *Puppet) Options2(opts ...option) (restore option) {
	restores := []option{}

	for _, opt := range opts {
		restore = opt(p)
		restores = append(restores, restore)
	}
	restore = mergeAll(restores...)

	return restore
}

// merge runs both the receiver and the argument functions
// which sets the respective features/fields.
// The method returns a closure, that does a reset to previous values, when run.
func (opt option) merge(opt2 option) option {
	return func(p *Puppet) option {
		// running setting of the option opt
		r := opt(p)
		// running setting of the option opt2
		r2 := opt2(p)

		// a new restore function is returned
		//     rAll is the closure returned by merge method
		//		 within which the r and r2 functions are applied
		rAll := r.merge(r2)
		return rAll
	}
}

// merge runs both argument functions
// which sets the respective features/fields.
// The function returns a closure, that does a reset to previous values, when run.
func merge(opt option, opt2 option) option {
	return func(p *Puppet) option {
		// running setting of the option opt
		r := opt(p)
		// running setting of the option opt2
		r2 := opt2(p)

		// a new restore function is returned
		//     rAll is the closure returned by merge method
		//		 within which the r and r2 functions are applied
		rAll := merge(r, r2)
		return rAll
	}
}

// mergeAll runs all argument functions
// which sets the respective features/fields. The returned
// restore functions are gathered and used when returning  a closure,
// that does a reset to previous values, when run.
func mergeAll(opts ...option) option {
	return func(p *Puppet) option {
		restores := []option{}
		for _, opt := range opts {
			restores = append(restores, opt(p))
		}
		// a new restore function is returned
		//     rAll is the closure returned by mergeAll function
		//		 within which all restore functions are applied
		rAll := mergeAll(restores...)
		return rAll
	}
}
