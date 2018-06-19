## go-options

Credits: Better Design of Options for Go Types (https://medium.com/@betakuang/better-design-of-options-for-go-types-d5a19d5309c)

As always, to best understand something is to dig into it. After reading the above article,
I implemented the same thing using a puppet ordering service that lets you set the puppet's
features (skin, hair, height etc), to amend your specification and change your mind and go back to the previous one.
There are several implementations of the bulk update and the related restoration of the previous values.

## Implementation

Puppet features are initially set up when creating a new one with New function:

	pup := puppet.New(puppet.Features{
		Skin:       "burnt",
		Eyes:       "blue",
		HairColour: "yellow",
		HairLength: "short",
		HairStyle:  "afro",
		Height:     160,
		Weight:     57,
	})

  Puppet's initial specification is set through the Features instance. There is no direct access to Puppet's
  features/fields. Individual features are retrieved through getters. Features can be changed individually or in bulk:

  - func (p *Puppet) Option(opt option) option
  - func (p *Puppet) Options(opts ...option) option
  - func (p *Puppet) Options2(opts ...option) option

  Each method outputs a restore function, that can be supplied to the method to restore the previous state.

  The bulk update methods, Options and Options2, uses merge to ensure all required restores are performed. There are two implementations of the merge functionality for the Options method. One is written as a method with an option receiver, and one as a function taking two option function arguments. The Options2 method uses mergeAll function, that does a bulk merge.

  merge method/function:
  - runs the two feature updates (both option methods)
  - creates a new restore function, which is a closure returned by merge, and returns it, allowing to reset back

  mergeAll function:
  - runs all feature updates (all option methods) and gathers all restore functions
  - creates a new restore function, which is a closure returned by mergeAll, and returns it, allowing to reset back

  ### Performace

  Benchmarks for the bulk update methods, Options and Options2, show that the Options one, updating and merging restore functions successively is on average 3 times faster than the Options2 one, that collects option functions and calls the mergeAll method once.

  ### Restoring
  
  - current settings: eyes: brown, skin: dar
  - change the eyes and skin features
	    restoreAll = pup.Options(puppet.SetEyes("green"), puppet.SetSkin("fair"))
  - reset the eyes and skin features back to: eyes: brown, skin: dark
	    restoreAll1 = pup.Options(restoreAll)
  - reset the eyes and skin features back to: eyes: green, skin: fair
	    restoreAll2 = pup.Options(restoreAll1)
  - reset the eyes and skin features back to: eyes: brown, skin: dark
	    restoreAll3 = pup.Options(restoreAll2)
   - reset the eyes and skin features back to: eyes: brown, skin: dark
	   restoreAll4 = pup.Options2(restoreAll3)
