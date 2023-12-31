=== Advent of Code 2019: Retrospective ===

[[home]](/)

# Advent of Code 2019: Retrospective

_December 27, 2019_

**Note: Per-problem commentary is [here](/advent-of-code-2019-commentary.html)**

This year I participated in [Advent of
Code](https://adventofcode.com/2019) for the first time, and yesterday, on Christmas Day, I finished all fifty
stars!

After finishing part 1 of [Day
25](/advent-of-code-2019-commentary.html#day-25), upon
discovering that part 2 was just "have you gotten all the other
stars," I went back and finished part 2 of [Day
22](/advent-of-code-2019-commentary.html#day-22) -- the only
thing I had remaining. After doing that I "activated the warp
drive" and got the fiftieth star!

<img src="img/aoc2019stars.png">

Except for Day 18 and part 2 of Day 22, I [finished every
problem](img/aoc2019leaderboard.png) on the day it came out and
blogged about it, which was my initial goal. Although sometimes
the blogging took a couple of days :) All the posts are in a
single enormous [blog
post](/advent-of-code-2019-commentary.html).

## Topics

Overall I can say that this is my proudest programming-related
acheivement and the most fun I've ever had programming. I had to
relearn a bunch of math and computer science techniques, like 2d
and 3d geometry, modular arithmetic, bit twiddling, basic data
structures and algorithms, BFS/DFS and pathfinding, dynamic
programming, matrix algebra, assembly programming, and cellular
automata. I also learned a bit about computer organization in
practice, something we only briefly touched in university, by
building a virtual machine and
[disassembling](/advent-of-code-2019-commentary.html#intcode-reverse-engineering)
an Intcode program to reverse engineer a problem. The breadth and
depth of the topics was truly remarkable. I feel like I used most
of my formal education in one single month. This was so
refreshing after years of iterative large-system building that is
mostly devoid of theoretically-interesting problems. (Not to say
it's not just as important :)

I've been skeptical of my formal education for the last few
years, but it certainly did help this month: I was able to
at least identify the underlying theoretical problem or relevant
approach each day within a minute or two, even if I had to read
Wikipedia a bit to remember the details. For shortest paths I
immediately went to BFS, for the card shuffling I immediately
thought "group theory," and so on.

## Problem Solving

I learned a lot about problem solving through countless false
starts, wrong answers, and hours of debugging and refactoring.
Here's some of the things that came up repeatedly:

### Solve the specific problem and not accidentally a harder one

This bit me on Day 18 (among other things): keeping track of the
key visit path at each step was a harder problem than just
tracking total distance traveled, since relying on distances
alone makes it possible to solve with dynamic programming, which
is what made it tractable! I was tracking the paths for
debugging, but after dropping this extra state, the actual distance
problem was easier to solve.

### Don't generalize early

Generalizing early means introducing early complexity when you're
supposed to be solving a specific problem. Early complexity can
mean bugs unrelated to your problem solving! Solve the problem
first and refactor later (also often a useful lesson in real
engineering).

### Use brute force unless you can't

As far as I can recall, the naive solution worked for every
single part 1, and complicated algorithms like A\* were never
strictly necessary. Using the most straightforward solution and
only doing something clever once it's clear that the program
isn't going to terminate, even with shortcuts and minor tuning,
means that there's less unnecessary complexity (and fewer bugs)
in the way of the problem solving. In other words, don't optimize
prematurely!

### Optimize/Generalize _after_ the code is correct

Related to avoiding premature generalization and optimization:
ensure the algorithm is correctly implemented beforehand!
Several times I wrote a bunch of code and then started optimizing
it before I'd done sufficient testing to know it was correct. I
then spent a bunch of time debugging the optimization before
realizing the algorithm was wrong.

### Don’t special case in recursion

Recursion is fundamentally about reducing a problem to itself,
but slightly smaller. This self-similarity is the heart of it,
and trying to add some special casing (beyond base cases) for
certain types or nodes or elements or means writing code that's
harder to understand and involves duplicated logic.

### Reduce state in recursion to the minimum required

Recursion requires careful reasoning, and the more state you
introduce, the harder it is to understand why your code isn't
working. Figure out the minimum amount of data you need to solve
the problem and then compute additional metadata/debugging info
as necessary -- but not in the main algorithm implementation!
I'm convinced that one of the reasons that Day 18 was so hard for
was that I was carrying around so much state in each recursive
step.

### Your input data might change the problem complexity

Maybe your specific input doesn't require solving the general
problem outlined in the day's description! An example of this Day
17, when you could just do the droid path compression by hand
instead of trying to write a custom compression algorithm for
this use case. Another example is Day 16, when you don't actually
need the full N-repeated input and M matrix multiplications,
since there's a trick based on the specific matrix involved that
simplifies the problem drastically. Take time to understand the
specific _instance_ of the problem instead of the general
theoretical problem .

### Write simple concurrency that fits in your head

Most of the concurrency I wrote was very simple: a handful of
goroutines that wrap a library function and a handful of
coordinating channels. The only time I had trouble was on Day 23,
when I used a ton of goroutines with function literals and
multiple channels per machine for each step in the network.
When I had trouble with deadlock and contention it was hard for
me to understand what was going on, because there was simply too
much code and the concurrency was too complex to fit in my head.
I eventually got an answer, but the code is in serious need of
refactoring to abstract out some functions and move more stuff
into the main thread with basic for loops and so on.

## Go

Having written [~4800 lines of code in
Go](https://github.com/dhconnelly/advent-of-code-2019), I now
feel as fluent in Go as I do after seven years of
near-uninterrupted C++. The language itself is a refreshing
change of pace: it can fit in one person's head, and the
[spec](https://golang.org/ref/spec) is helpful and clear! I
really felt my problem solving, and not the language, was the
barrier each day -- the code just flowed, and I never sat and
wondered how to express something.

Apart from the first few days, when it seemed like the input
parsing would vary less from day-to-day and so generics might
have been nice, I never seriously missed it. The lack of
expressiveness that comes with e.g. no generics or sum types or
something meant that I just fluently wrote out more explicit code
without thinking hard about it. I understand there's a trade-off
here, and I'm not saying this is necessarily the right side of
it, but it is nice to generally have one way to do something that
is explicit and comprehensible. I'll be interested to see how
they grow the language for Go 2 in a way that keeps it this
small. I really think this is a large part of what makes it
feel refreshing to me after all this C++ and several periods
experimenting with Haskell and Rust, not to mention the
occasional Java and Objective-C I also use at work from time to
time.

Outside of the language, the standard library is easy to use and
comprehensive: the only external dependency I pulled in was
[`tcell`](https://github.com/gdamore/tcell), for [Day
13](/advent-of-code-2019-commentary.html#day-13).

I also developed simple patterns for sets, stacks, queues, trees
and graphs using just slices and maps, which makes it easy to
avoid dependencies entirely for data representation. I wrote that
up a bit in a
[Gist](https://gist.github.com/dhconnelly/7c99902b00874900ef8dff374a04f477).

Channels also made several problems easy, particularly the
Intcode problems, although I need to spend some more time
developing architectural patterns; [Day
23](/advent-of-code-2019-commentary.html#day-23), for example,
needs a bunch of refactoring to make it simpler, more
comprehensible, and eliminate thrashing and contention. But
overall I'm very happy with the way the language approaches
concurrency, which mostly avoids the worst bits of the buggy and
incomprehensible multithreaded code I see all the time at work.

Anyway, all in all I think Go is a great general purpose language
and it was a joy to use it for Advent of Code.

## Conclusion

I'm very glad Advent of Code 2019 is over. I got behind at work,
I was stuck on the computer while visiting family, and towards
the end I was getting a bit burnt out. Still, this is absolutely
the most fun I've ever had programming, and I'm so happy to have
practiced all this material and refreshed my skills, and proud to
have finished it all by the last day! Looking forward to
participating next year!

[[home]](/)
