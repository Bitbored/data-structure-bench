A Go data structure benchmark
=============================

This is an experiment to determine the best way to store elements indexed by an integer in Go.

TL;DR, skip to [the results](#user-content-test-results)

The "best way" is of course relative to the scenario in which the storage will be used. My focus will mostly be on the following scenario:

- The storage system may be static
- Accessing elements is done a lot more frequently than adding elements
- The lookup should be able to happen by multiple parallel go routines, so it should not be blocking

This repository contains some general benchmarks that allows to compare the storage costs and lookup time of a few different data structures in relation to the amount of elements stored in the structure.

If you want to add other data structures to the test, or if you spot a mistake, feel free to send me a pull request.

Following data structures are currently tested:

## [The simple Binary Search Tree (BST)](http://en.wikipedia.org/wiki/Binary_search_tree)

A Binary search tree is a simple way to store data indexed by an integer, with a focus on fast retrieval (compared to e.g a linked list).

In theory, it has an average lookup time complexity of O(log n), and a worst case lookup time of O(n). This worst case occurs when elements are added in-order (based on their index). This makes it a pretty bad in a lot of common scenario's.

A good way to prevent a worst case scenario would be to sort the elements before inserting them in the tree. The downside of this is that then the tree is not entirely dynamic, meaning we need to rebuild the tree every time when we want to add new elements after the initial build. In some use-cases (like in my case), this is not a problem.

I did not implement this sorting mechanism, but I did add an algorithm to the benchmark that inserts elements in an ideal order. This allows to measure only the characteristics of the "tree part" of the data structure in a best case scenario.

## [The Left-Leaning Red-Black (LLRB) tree](http://en.wikipedia.org/wiki/Left-leaning_red%E2%80%93black_tree)

The LLRB tree is a type of tree similar to the binary search tree, but it is self-balancing. This means it has a O(log n) lookup time complexity regardless of the order the elements where added.

I used @petar's [GoLLRB](https://github.com/petar/GoLLRB) library and added some functions to make it compatible with the `DataStrucure` interface.

Note that I didn't write benchmarks to see if this implementation acts like an ideal LLRD implementation.

## [The built in map](https://golang.org/doc/effective_go.html#maps)

The Go programming languages comes with a built-in implementation of a hash table, called a map.

[The implementation of Go's map](https://code.google.com/p/go/source/browse/src/runtime/hashmap.go) is pretty complex, so I had a hard time predicting what it's behavior would be. However, in theory it should be able to offer an average lookup time complexity of 0(1), and a worst-case of 0(n).

One of the goals of this experiment is to find out if it is a good idea to use the map for a small / extremely large dataset.

The biggest downside of the map is that accessing it is not guaranteed to be an atomic operation. However, when writing correct concurrent code in Go, this should not be a problem.  

## [The built in slice](https://golang.org/doc/effective_go.html#slices)

This data structure is mainly here to provide a benchmark for the fastest possible way to store data that is indexed using an integer.

It is very fast to store data in a slice, which is essentially an array with some meta-data, so it should always offer 0(n) lookup times. The downsides are pretty straightforward:
- You can not store elements with a negative index
- Gaps in the indexes of elements can causes a lot of useless memory overhead
- You need to reserve enough memory in advance, meaning you should know the amount of elements you want to store if you want to have an interface like my `DataStructure`. It is of course possible to re-allocate data, like the way the `add` function works, but this is not compatible with my `DataStructure.Add` 

# Test results

## Insert times

The average time required to add a single element to the structure relative to the size of the structure.

### [Raw data](results/insert_times.csv)

elements | BST worst case | BST best case | LLRB    | Map    | Slice
---------|----------------|---------------|---------|--------|------
      10 |         161.20 |        647.10 |  565.70 | 112.50 | 80.60
     100 |         441.18 |        622.67 |  871.13 | 117.20 | 78.07
    1000 |        3030.04 |        507.17 | 1241.27 | 111.59 | 72.29
   10000 |       30624.00 |        824.87 | 1621.54 | 114.34 | 65.32
  100000 |      449889.02 |        712.54 | 2093.55 | 171.63 | 64.96
 1000000 |    60759419.24 |        658.21 | 2973.17 | 275.94 | 66.09

### Graph
![Lookup times](https://cdn.rawgit.com/Bitbored/data-structure-bench/master/results/insert_times.svg)

## Lookup times

The average time required to get a single element from the structure relative to the size of the structure.

### [Raw data](results/lookup_times.csv)

elements | BST worst case | BST best case | LLRB    | Map   | Slice
---------|----------------|---------------|---------|-------|------
      10 |          23.70 |         12.90 |  456.50 | 17.50 |  5.26
     100 |         265.19 |         23.59 |  747.51 | 19.18 |  5.27
    1000 |        2736.37 |         67.26 | 1098.09 | 30.83 |  4.96
   10000 |       28900.79 |         85.12 | 1484.10 | 41.30 |  5.00
  100000 |      382708.02 |        109.89 | 1799.66 | 64.97 |  5.08
 1000000 |     4213983.23 |        162.99 | 2186.23 | 91.01 |  5.07

### Graph
![Memory usage](https://cdn.rawgit.com/Bitbored/data-structure-bench/master/results/lookup_times.svg)


## Memory usage

The average memory usage required per element relative to the size of the structure.

### [Raw data](results/memory_usage.csv)

elements | BST worst case | BST best case | LLRB  | Map   | Slice
---------|----------------|---------------|-------|-------|------
      10 |          32.00 |         73.40 | 48.00 | 16.00 | 16.00
     100 |          32.00 |         70.72 | 48.00 | 16.00 | 16.00
    1000 |          32.06 |         57.41 | 48.05 | 16.01 | 16.00
   10000 |          38.40 |         91.94 | 48.48 | 16.10 | 16.01
  100000 |          64.00 |         75.00 | 57.60 | 16.81 | 16.08
 1000000 |          64.00 |         74.72 | 96.00 | 40.27 | 16.80

### Graph
![Memory usage](https://cdn.rawgit.com/Bitbored/data-structure-bench/master/results/memory_usage.svg)

# Replicating the experiment

To execute the benchmark for a specific type, use:

```
go test -benchmem -test.bench . -elements [amount of elements (default=100)]
```
*in the folder of the structure type you want to test.*

Note that the benchmarks for the BST will usually take a **long** time to execute for big data sets(> 100000 elements), so unless you have a crazy fast supercomputer, you will have to use the `-benchtime` [test flag](https://golang.org/cmd/go/) (something like 12h should suffice). 


You can than store the results in the CSV files under `/results`, and if you want you can generate nice some plots from the test data using:
```
gnuplot plot.gnu
```

# Conclusion

A first conclusion is that Go's implementation of a hash-table, the map, is surprisingly fast. The complex nature of the Go map would suggest that it should be pretty inefficient for small data sets, but this is not the case. 

In a best-case scenario the BST is faster than the map for extremely small data sets (<50 elements), but the difference is pretty neglectable.

A second strange result is the pretty bad performance of the LLRD implementation compared to the performance of my primitive BST implementation. An LLRD tree should be extremely good at handling big amounts of data, but for smaller datasets it seems to suffer from it's complexity. Also note that this LLRD implementation needs to allocate some data every time it does a lookup. This is a pretty big downside when doing a lot of lookups.

A final conclusion is that there is sadly no such thing as magic. The only data structure capable of true O(1) lookups is the slice.

However, when the limitations of a slice aren't feasible, a map is a fast and extremely flexible alternative.

I will probably use a map for my use case.

