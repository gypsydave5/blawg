---
layout: post
title: "Lambda Calculus 3 - Logic with Church Booleans"
date: 2017-10-21 20:54:21
tags:
    - Mathematics
    - Logic
    - Functional Programming
    - Lambda Calculus
published: true
---

I found Church numbers pretty tough, and I'm still not sure I fully understand
them. But this post should be a little bit easier. I promised logic for this
post, and logic I will give you. But not right now. First, it's...

## Data structures with functions

As modern 21st century software developers, we're used to a strong divide
between _data_ and _process_. Even with object orientation we consider an object
to be made of things it knows (the data) and things it does (the methods).[^1]

I always picture my programs as big old conveyor belts, where my program-workers
each beat the hell out of a piece of data as it goes past, until it comes out as
a shiny ~~new BMW~~ piece of JSON.

Meanwhile, here in the lambda calculus... well, we've got numbers for sure. But
how can we get data structures like a piece of shiny JSON? Or even just a list?

Take heart - through the lambda, all things are possible![^2]

Let's try the simplest of all data structures - the pair. Simple, sure - but
powerful. Every good Lisper knows that the can build any data interface you can
conceive using enough pairs. If Lisp was written by Archimedes, he'd say "Give
me a place to stand and enough `cons` cells and I shall move the earth".

To get a pair as a data structure up and flying, we need three functions. One to
make a pair out of two arguments, one that returns the first item in the pair,
and another that returns the second item. Lispers will say "Ah! `cons`, `car`
and `cdr`!", but we will say $pair$, $first$ and $second$.

First $pair$:

$$
pair\quad \equiv \quad \lambda p.\lambda q.\lambda f.\ f\ p\ q
$$

We take the two things we're pairing - that much makes sense - but then we take
one more argument and apply it, first to the first argument, and then the result
of that to the second argument.

What's going on?

We want something like $ first \< aPair \> $ to give us back the first item that we
gave to the pair. So the final $f$ in pair is going to be offered both $p$ and
$q$ and left to decide which one it wants.

$$
first\quad \equiv \quad \lambda pair. pair\ (\lambda a.\lambda b.\ a)
$$

$first$ takes a pair, and then gives that pair a function that takes two
arguments and returns... well, the first one. Given that, it's a doddle to write
$second$:

$$
second\quad \equiv \quad \lambda pair. pair\ (\lambda a.\lambda b.\ b)
$$

Same again, but this time we evaluate to the second of the arguments.

Feel free to stick it into a programming language and play with it. Here it is
in Scheme:

```scheme
(define pair
    (lambda (p)
        (lambda (q)
            (lambda (f)
                ((f p) q)))))
```

```scheme
(define first
    (lambda (p) (p (lambda (a) (lambda (b) a)))))
```

```scheme
(define second
    (lambda (p) (p (lambda (a) (lambda (b) b)))))
```

## Booleans

> What is truth? said jesting Pilate, and would not stay for an answer.
> &mdash; Francis Bacon, _On Truth_

Well, I hope you do stay for an answer. And there _is_ an answer, at least for
the concerns of the lambda calculus, but it's not going to be particularly life
altering. In fact, it's probably going to remind you most of the solution to
what a number is.

## Numbers... AGAIN!

When we defined numbers a few posts ago, I was being fairly adamant that the
best way to encode a number using functions would be to count the number of
applications of a function. Remember - $zero$ was no applications, $one$ was
one, etc.

But this ain't necessarily so - we could implement numbers using the definition
of pairs above:

$$
zero \quad \equiv \quad \lambda x\ x

succ \quad \equiv \quad \lambda n.\ pair\ <FALSE>\ n
$$

Here we've defined $zero$ as the identity function, and $succ$, the successor,
as a pair of $\<FALSE\>$ to whatever the previous number was. Each number is now
'counting' using the number of times that $zero$ has been paired up with
$\<FALSE\>$.

We can now go on to define other functions around this implementation - and we
will - but the key thing I'd like to stress is that what makes a number a number
isn't really what it _is_, but rather how it _behaves_ - how it behaves when
being used with other functions like $add$ and $multiply$.

Let's use that insight to imagine what $true$ and $false$ might be.

## if ... then ... else

Every programming language I've ever worked in has some sort of `if` expression
or statement - a way of choosing one bit of code or another based on whether
something was true or false. You know:

```ruby
if (1 + 1 == 2)
    puts "One and one is two!"
else
    puts "Maths is broken!"
end
```

We could think of `if` as being a function in the lambda calculus:

$$
if \quad \equiv \quad \lambda bool.\lambda t.\lambda f.\ <SOMETHING>
$$

This is fine, but gets us nowhere. But what if there were two different
functions, one of which we used for true booleans, and one of which we used
for false ones. Yes, I know, that would make no sense - you'd have to know which
one to use. But humour me.

$$
if-true \quad \equiv \quad \lambda bool.\lambda t.\lambda f.\ t

if-false \quad \equiv \quad \lambda bool.\lambda t.\lambda f.\ f
$$

We're not even using the boolean any more, we're just saying that if the
boolean is true, we evaluate to first argument, and if it's false we evaluate to
the second argument.

If we're not using the boolean, we can get rid of it from the end.

$$
true \quad \equiv \quad \lambda a.\lambda b.\ a

false \quad \equiv \quad \lambda a.\lambda b.\ b
$$

And there we have it. We can just say that $true$ is the function that returns
the first, and $false$ is the function that returns the second.

Wait, can we? Well, why not? All we need really is a function that will signal
one of two things - true or false we can call them. This 'signal' we choose to
be the return of the first or the second of the arguments it is applied to. Who
cares _how_ truth works - this is a mechanism that does what it needs to do.

## Truth Tables

### And

So now we've got truth going, let's have some fun defining some boolean
operations. First, an easy one - $and$. And once again we're going to use the
behaviour of $and$ to give us a clue as to the implementation.

What' s the behaviour? This might be easier to do if we construct a truth
table. What's a truth table I hear you cry? Well, in logic we can draw up a
table showing the truth or falsity of a proposition (sentence that is either
true or false) based upon the truth or falsity of the propositions from which it
is composed.

The truth table is just the exhaustive table of true and false values that can
exist in the proposition, along with the resulting truthfulness of the overall
proposition. A proposition involving 'and' will be made of two
sub-propositions - the two being 'anded' together. Traditionally these are
written as '$P$' and '$Q$' - and who are we to disagree with tradition? The symbol
'$\land$' is often used for 'and', so we'll do the same here. Finally, true and
false will be '$T$' and '$F$'.

$$
\begin{array}{ c  c | c }
P & Q & P \land Q \\\
\hline
T & T & T \\\
T & F & F \\\
F & T & F \\\
F & F & F
\end{array}
$$

What can we learn from this? Well, two things:

- If $P$ is false, then the proposition is always false.
- If $P$ is true, then the proposition has the same value as $Q$

So we could say something like "if $P$ then $Q$ else $false$". Which can be
written quite as:

$$
and \quad \equiv \quad \lambda p.\lambda q.\ p\ q\ false
$$

Or, even more concisely:

$$
and \quad \equiv \quad \lambda p.\lambda q.\ p\ q\ p
$$

As if $p$ is false we can just evaluate to $p$

### Or
'Or' is represented by '$\lor$':

$$
\begin{array}{ c  c | c }
P & Q & P \lor Q \\\
\hline
T & T & T \\\
T & F & T \\\
F & T & T \\\
F & F & F
\end{array}
$$

The pattern here should be clearer after doing $and$

- If $P$ is true, then the proposition is always true.
- If $P$ is false, then the proposition has the same value as $Q$

Which can be written in lambdas as:

$$
or \quad \equiv \quad \lambda p.\lambda q.\ p\ p\ q
$$


### Not
'Not' is nice and short as a truth table. We will use $\lnot$ to represent it

$$
begin{array}{ c | c }
P & \lnot P \\\
\hline
T & F  \\\
F & T
\end{array}
$$

We just need to flip $P$ around - if it was true (returning the first argument),
we make it return false (return the second argument), and vice versa.

Like this:

$$
not \quad \equiv \quad \lambda p.\lambda a.\lambda b.\ p\ b\ a
$$

### if ... then

$$
\begin{array}{ c  c | c }
P & Q & P \limp Q \\\
\hline
T & T & T \\\
T & F & F \\\
F & T & T \\\
F & F & T
\end{array}
$$

$$
implies \quad \equiv \quad \lambda p.\lambda q. p\ q\ \p
$$

### if and only if

$$
\begin{array}{ c  c | c }
P & Q & P \liff Q \\\
\hline
T & T & T \\\
F & F & T \\\
F & T & F \\\
T & F & F
\end{array}
$$

$$
true \quad \equiv \quad \lambda a.\lambda b.\ a
$$

$$
false \quad \equiv \quad \lambda a.\lambda b.\ b
$$

$$
and \quad \equiv \quad \lambda p.\lambda q.\ p q p
$$

$$
or \quad \equiv \quad \lambda p.\lambda q.\ p p q
$$

$$
if \quad \equiv \quad \lambda p.\lambda a.\lambda b.\ p a b
$$

[^1]: I will be the first to admit that this is a terrible definition of OO, mainly driven by my ignorance.
[^2]: All things not necessarily possible with the lambda. Terms and conditions apply.
