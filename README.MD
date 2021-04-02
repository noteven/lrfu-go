# lrfu-go

Go based implementation of LRFU (least recently frequently used) cache as set forth in [Donghee Lee et al., "LRFU: a spectrum of policies that subsumes the least recently used and least frequently used policies,"](http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.55.1478).

### Example

```go
```

### Usages

## Reasons to use this implementation

 - It's free.

## Reasons to not use this implementation

 - No concurrent operations. No work was put into making the cache safe for concurrent or parallel access.
 - Simple cache implementation. No attention was given to optimizing data structures space or operations time complexity.
 - Lack of tests and benchmarks. At the moment benchmarking and test coverage are lacking.

### Alternatives
Below are some possible alternatives to this library.

 * [lfuda-go](https://github.com/bparli/lfuda-go)
 * [lfu-go](https://github.com/dgrijalva/lfu-go)
 * [lru-go](https://github.com/hashicorp/golang-lru)

## License

This implementation is licensed under the MIT License [LICENSE](http://opensource.org/licenses/MIT).

## Contributing

All constributions are licensed under MIT license, unless otherwise stated.
