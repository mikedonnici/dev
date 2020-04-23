# Performance

## Magic Commands

Magic commands are enhancements in [iPython](https://ipython.org/) on top of normal Python syntax.

Note: `ipython`, not the standard interactive python shell. 

- Prefixed by the `%` character
- list all with `%lsmagic`
- Ref: https://ipython.readthedocs.io/en/stable/interactive/magics.html


### Examining runtime with %timeit

Timing code at runtime is a good way to compare different pieces of code for efficiency.

- _line magic_ mode for a single line of code: `%timeit`
- _cell magic_ mode for multiple liens of code: `%%timeit` 
- provides a mean and std of multiple runs
- set number of runs (`-r`) and loops (`-n`):
- output can be saved to a variable with `-o`, then can access attributes `.best`, `.worst`

```python
%timeit rand_nums = np.random.rand(1000)
%timeit -r2 -n10 rand_nums = np.random.rand(1000)

# compare times for creating a dict using formal vs literal syntax
f_time = %timeit -o d1 = dict()
# 149 ns ± 1.8 ns per loop (mean ± std. dev. of 7 runs, 10000000 loops each)
l_time = %time_it -o d2 = {}
# 41.3 ns ± 0.138 ns per loop (mean ± std. dev. of 7 runs, 10000000 loops each)
print("f_time worst = {}, l_time worst = {}".format(f_time.worst, l_time.worst))
# f_time worst = 1.5262773020000394e-07, l_time worst = 4.152891470000668e-08
```

## Code Profiling

### Run time

- Done in iPython shell
- Code profilers gather stats of frequency and duration of function calls
- Do line-by-line analysis
- Better for analysing performance of larger pieces of code

Can use package `line_profiler`:

```python
%load_ext line_profiler
%lprun -f func_name func_name(arg1, arg2)
```

### Memory usage

- quick and dirty approach, can use `sys.getsizeof()` to see size of vars 
- better to use a code profiler
- `memory-profiler` package
- requires the function being profiled is imported, so the func itself must be defined in a physical file and then 
imported into iPython shell 
- Small funcs may show up as 0 MiB

```python
from mod_name import func_name
%load_ext memory_profiler
%mprun -f func_name func_name(arg1, arg2)
```

### Use alternative libraries for combining and iterating

There packages that provide more efficient ways to perform operations with collections and iterators. 

For example, `collections` and `itertools`.

Some built-ins are also very efficient, such as `zip` and `set`.

### Eliminating loops

> "Flat is better than nested"

Looping is often inefficient so, where possible, should be avoided.

Using built-in list comprehensions and `map` will result in less code, and faster run times.

Using packages such as `NumPy` provides many functions which are significantly more efficient than using loops.

When loops have to be used, there are better ways to approach them.

- analyse what is being done in the loop
- anything that can be done _once_ should be moved out of the loop
- operations on a list built in a loop can be moved _below_ and done with `map`

```python
# 'holistic' operation foo() inside loop
results = []
for i in lst:
    i = foo(i)  
    results.append(i)
  
# more efficient
results = []
for i in lst:
    results.append(i)
results = map(foo, results)
```

hp_avg = hps.mean()
hp_std = hps.std()
z_scores = (hps - hp_avg)/hp_std
poke_zscores2 = [*zip(names, hps, z_scores)]
highest_hp_pokemon2 = [(name, hps, z_score) for name,hps,z_score in poke_zscores2 if z_score > 2]







 



















