# Evolutionary Computation in Golang
This is really a re-implementation of a set of libraries I've developed
over the years to perform evolutionary computation.  I've had versions
in C++, Java, C#, LISP, and even Ada.  I'm re-implementing it in Go for a couple
of reasons.  The first is I just like the subject matter.  It's fun to 
play "guess what converges faster."  

The second is because I've always solved this problem with generics.
Except for LISP (where I passed in specific algorithms as functions), 
the easy way to tackle a generic evolutionary computing
problem was to parameterize a generic EC framework with specific 
implementations of behaviors.  The behaviors involve how populations 
are constructed, how reproduction occurs, and how offspring and parents
are selected.

While the behavior of specific subtype of evolutionary algorithms might
involve sexual or asexual reproduction, mutation, encoding as bit-strings,
etc., they all had roughly the same general outline.  

1. Generate a candidate population
2. Produce some number of offspring
3. Have the offspring and parents compete of space in the new generation.

Being able to parameterize an EA so as to mimic a (highly explorative)
Genetic Algorithm with sexual reproduction and fitness-proportional 
selection or more of a <span>&mu; + &lambda;</span> (and highly exploitive) Evolutionary
Strategy is possible in a framework that views them as points along a 
continuum.  That allows you to "shift the slider" left or right as you
explore a problem space.

So what's this like without generics (which may be coming in Go 2.0).
Without opening up a can of worms, generics are actually a hard language
feature to implement correctly.  (cough cough Java).  Even then, not all
parameterized constructs are equally generic.  For example, a generic 
parameter may be required to support a specific set of operations or 
inherit from a specifc parent.  

So it's also just interesting to take a problem that I've solved multiple
times with generics and solve it with nothing more than interfaces and first-class
functions.  
