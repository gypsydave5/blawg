---
layout: post
title: "(Even More) Memoization in JavaScript"
date: 2016-03-28 00:01:36
tags:
    - JavaScript
    - Learnings
    - Functional Programming
published: true
---

A while ago I wrote a piece on [basic lazy evaluation and memoization in
JavaScript][firstBlog]. I'd like to continue some of the thoughts there on
memoization, and to look at how some popular JS libraries handle it.

### Memoization refresher

According to [Wikipedia][memWik], memoization is:

> an optimization technique used primarily to speed up computer programs by
> storing the results of expensive function calls and returning the cached
> result when the same inputs occur again.

We want the program to remember the results of its operation, so it doesn't
have to expend any extra time and effort in doing them all over again. It all
seems a little obvious really.

So the first time we saw memoization it was in the context of [lazy
evaluation][firstBlog], where the memoized function had already been prepared
to be lazy. The lazily evaluated function was executed within our memoizer, and
the result was retained within the scope of the memoized function, to be
returned each subsequent time the function was called. As a refresher:

```javascript
function lazyEvalMemo (fn) {
  const args = arguments;
  let result;
  const lazyEval = fn.bind.apply(fn, args);
  return function () {
    if (result) {
      return result
    }
    result = lazyEval()
    return result;
  }
}
```

Which is great if want the function to be memoized with a specific set of
arguments. But that isn't going to work all of the time - probably not even most
of the time. It's rare that there's just one set of parameters we want to be
eternally bound to our function. A bit of freedom would be nice.

### Keeping track of your arguments

If we want to be able to memoize a function which can take a variety of
arguments, we need a way to keep track of them. Let's take one of the simplest
cases:

```javascript
function addOne (x) {
  return x + 1;
};
```

That shouldn't take too much explanation. Now, in order to have a memoized
version of it, we need to keep track of all the different values of `x` passed
to it, and to associate them with the return value of `addOne`.

A JavaScript object should suffice, taking `x` as a key, and the result as the
value. Let's give it a go:

```javascript

const memo = {};

function memoAddOne (x) {
  if (memo[x]) {
    return memo[x];
  }
  return memo[x] = x + 1;
};
```

This takes advantage of the fact that the value of an assignment to an object is
the value assigned (in this case `x + 1`), saving a line or so. The only issue
here is that the value of `memo` is floating around in public, just waiting for
other functions to overwrite and mutate. We need a way to hide it.

Well, we could place `memo` on the function itself - it is an object after all
- and just put an underscore in front of its name to try and let the world know
that it's private (even though it isn't really private):

```javascript
function memoAddOne (x) {
  memoAddOne._memo = memoAddOne._memo || {};

  if (memoAddOne._memo[x]) {
    return memoAddOne._memo[x];
  }
  return memoAddOne._memo[x] = x + 1;
};
```

This isn't beautiful, but it works. `_memo` gets defined as an empty object on
initialization and gets filled up with results on each application of the
function - throw some `console.log()`s in there to prove it to yourself. That's
exciting - although we're still a little exposed with the `_memo` property being
available on the function.

That said, we've got what we came for - a memoized version of our function that
works for many different arguments.

### Fun with strings

Problem is, we're reliant on `x` being used as the property for our `_memo`
object. As all good school children are taught, JavaScript, like the universe,
is just a load of strings held together by poorly-understood forces. When `x`
is used in `_memo[x]`, JavaScript handily casts it to a string to be used as the
property name.

I say handily - but try this...

```javascript
memoAddOne([55]) // => '551'
memoAddOne(55) // => '551'
memoAddOne(66) // => 67
memoAddOne([66]) // => 67
```

Because...

```javascript
55.toString() // => '55'
[55].toString() // => '55'
```

Ah, JavaScript: thou givest with one hand... `toString()`, which is what
JavaScript is using under the hood to cast the non-string property identifier to
a string, does not uniquely identify that argument. So our function behaves
inconsistently depending on whether it was given the array or the number that
`toString()` converts to the same string.

We need a more predictable way of parsing our argument. Happily, we can borrow
one of the built-in functions of JavaScript to do this, `JSON.stringify()`.

```javascript
JSON.stringify(55) // => '55'
JSON.stringify('55') // => '"55"'
JSON.stringify([55]) // => '[55]'
JSON.stringify(['55']) // => '["55"]'
```

Pretty good - let's give it a whirl:

```javascript
function memoAddOne (x) {
  memoAddOne._memo = memoAddOne._memo || {};
  const jsonX = JSON.stringify(x);

  if (memoAddOne._memo[jsonX]) {
    return memoAddOne._memo[jsonX];
  }
  return memoAddOne._memo[jsonX] = x + 1;
};

memoAddOne([55]) // => '551'
memoAddOne(55) // => 56
```

Sorted.

### The General Case

Now let's put together a function that can memoize _any_ function - and as
a bonus, we can also hide that nasty `_memo` property behind a closure:

```javascript
function memoize (fn) {
  const memo = {};

  return function () {
    const args = Array.prototype.slice.call(arguments);
    const jsonArgs = JSON.stringify(args);

    if (memo[jsonArgs]) {
      return memo[jsonArgs];
    }
    return memo[jsonArgs] = fn.apply(null, args);
  };
};
```

Much of this should now be familiar, but let's dig in. We take a single
argument, hopefully a function, and bind it to the variable `fn`. We now get to
declare `memo` inside our function - and hooray it's now unavailable to anyone
but the function we're returning. Now that's what I call private - thank you
closures!

The function we give back, well we don't know how many arguments it's going to
be given so why bind them to any parameters? We'll just leave its parameters
empty. Any arguments we do get we'll instantly turn into an array by using the
funky `Array.prototype.slice.call(arguments)` dance. And that array we can
`stringify()` on the next line and call something useful like `jsonArgs`.

Then we do much the same as above, only instead of giving `x + 1` as the
value of `_memo[jsonX]`, we set the value of `memo[jsonArgs]` as the result of
applying the `args` array to the original function we were given to memoize. Job
done.

Again, throw some `console.log`s in there to see what's really going on.

### Here's One I <del>Made Earlier</del> Installed With `npm`

The above is so incredibly useful that you'll not be surprised to learn that
it's implemented, with slight modifications, in functional JavaScript libraries
like [Underscore], [Lodash] and (personal niche favourite) [Ramda].[^1]

Let's take a look at [the Lodash implementation]:

```
function memoize(func, resolver) {
  if (typeof func != 'function' || (resolver && typeof resolver != 'function')) {
    throw new TypeError(FUNC_ERROR_TEXT);
  }
  var memoized = function() {
    var args = arguments,
        key = resolver ? resolver.apply(this, args) : args[0],
        cache = memoized.cache;

    if (cache.has(key)) {
      return cache.get(key);
    }
    var result = func.apply(this, args);
    memoized.cache = cache.set(key, result);
    return result;
  };
  memoized.cache = new memoize.Cache;
  return memoized;
}
```

Now, although this is a little more <del>long-winded</del> complicated than the
code above, it should be similar enough for us to see that they're doing the
same thing. The difference being that in the above we are to supply an external function to
hash the arguments (the `resolver` function), and that Lodash offers a custom
caching object with a `get()` and `set()` interface, which we can overwrite if
we like.

(Bonus Question: why does this implementation of `memoize` suck if we don't supply
a `resolver` argument?)

The above library code will save us all the hassle of writing a memoization
function ourselves - but now we can see how they work under the hood, we can
take a more informed decision about whether we need to create a dependency on
external library, or whether we just put together the (relatively) simple piece
of code ourselves.

[^1]: Seriously, this is the one to go for. It's amazing, it's beautiful - it's _functional_.
[the Lodash implementation]: https://github.com/lodash/lodash/blob/4.6.1-npm/memoize.js#L48
[firstBlog]: {% post_url 2015-03-21-lazy-eval-and-memo %}
[Ramda]: http://ramdajs.com/0.20.0/index.html
[Underscore]: http://underscorejs.org/
[Lodash]: https://lodash.com/
[memWik]: https://en.wikipedia.org/wiki/Memoization
