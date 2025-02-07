package generator

// Generator is an interface that defines the signature of a data generation function.
// It returns a value of type `interface{}`, allowing flexibility in the type of data generated.
type Generator func() interface{}
