# Evolutionary Computation in Golang
This is really a re-implementation of a set of libraries I've developed
over the years to perform evolutionary computation.  I've had versions
in C++, Java, C#, LISP, and even Ada.  I'm re-implementing it in Go for a couple
of reasons.  The first is I just like the subject matter.  

The second is because I've always solved this problem with generics.
Except for LISP, the easy way to tackle a generic evolutionary computing
problem was to parameterize a generic EC framework with specific 
implementations of behaviors.  

While the behavior of specific subtype of evolutionary algorithms might
involve sexual or asexual reproduction, mutation, encoding as bit-strings,
etc., they all had roughly the same general outline.  

1. Generate a candidate population
2. Produce some number of offspring
3. Have the offspring and parents compete of space in the new generation.

Being able to parameterize an EA so as to mimic a (highly explorative)
Genetic Algorithm with sexual reproduction and fitness-proportional 
selection or more of a \mu \plus \lambda (and highly exploitive) Evolutionary
Strategy is possible in a framework that views them as points along a 
continuum.
