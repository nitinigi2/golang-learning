1) Are arrays can be compared with == ?

    Ans - Yes, if both of them hold same data type with same length


        func main() {
   
	      arr1 := [4]int{1, 2, 3, 4}
	      arr2 := [4]int{4, 5, 6, 7}
          arr3 := [5]int{5, 6, 7, 8, 9}

	      fmt.Println(arr1 == arr2)
          fmt.Println(arr2 == arr3) // this will give error
       }

2) Are slices can be compared with == ?

    Ans - No. dynamic length types can't be compared with ==

        func main() {
	      arr1 := []int{1, 2, 3, 4}
	      arr2 := []int{4, 5, 6, 7}

	      fmt.Println(arr1 == arr2) // this will give error 
        }

3) Are maps can be compared with == ?

    Ans - No

        func main() {
	          dict1 := map[string]int{
              "abc": 1,
              "def": 2,
            }

            dict2 := map[string]int{
              "abc": 1,
              "def": 2,
            }

            fmt.Println(dict1 == dict2) // this will give error
        }

4) How to sort slices with custom types?

        func sortMe() {
              // custom type
            type Person struct {
              Name string
              Age  int
            }
              // slice of custom data
            people := []Person{
              {"Bob", 31},
              {"John", 42},
              {"Aichael", 31},
              {"Jenny", 26},
            }
              // sort based on age first
              // if age is same then sort based on name
              sort.Slice(people, func(i, j int) bool {
                if people[i].Age == people[j].Age {
                  return people[i].Name < people[j].Name
                }
                return people[i].Age < people[j].Age
              })

            fmt.Println(people)
        }

5) How to copy a slice?

   Ans-  Use copy(dst, src []Type) int 

        func main() {
            arr := []int{1, 2, 3, 4}
            brr := make([]int, 4)

            copy(brr, arr)

            fmt.Println(brr)
        }

6) How to copy a map?

    Ans-  Iterate over map and add key, value pair

        func main() {
          dict1 := map[int]int{1: 1, 2: 2, 3: 3}
          destination := make(map[int]int)

          for k, v := range dict1 {
            destination[k] = v
          }

          fmt.Println(destination)
        }

7) How to copy and array?

   Ans - assign new array with old

        func main() {
          arr := [4]int{1, 2, 3, 4}

          brr := arr

          fmt.Println(brr, arr)
        }

8) Which types can be used as a key in a map?

    Ans - All type which are comparable like (int, string, etc)

9) Is it safe to share a single slice & array across several Goroutines?

    Ans - If there's only read operation then it's fine. If read/write then not allowed (unexpected error/panic)

        func main() {
          arr := []int{1, 2, 3, 4}

          go func() {
            for _, v := range arr {
              fmt.Println(v)
            }
          }()

          go func() {
            for _, v := range arr {
              fmt.Println(v)
            }
          }()
        }

Above above works fine as we are doing read operation only. But if any one or multiple try to write to the same slice then above problems will occur.

```Run using this -> go run -race filename```

```Above command will notify is there's any race condition arises```

10) What happens if a few goroutines write to a single map instance?

    Ans- If multiple go routines do only read operations then it's fine. If any of the go routine perform write operation then above mentioned problem will occur.

11) Is it safe to insert values into a map that isn’t initialized with the make function?

    Ans - No. Map not initialized with make is like pointing to nil in memory. So if someone try to do so, this results in panic

12) Is it safe to append to a slice that isn't initialized with the make function?

    Ans- Yes.

    ```The only reason why append function works with nil, is that append function can do reallocation for the given slice. For example, if you trying to to append 6th element to slice of 5 elements with current capacity 5, it will create the new array with new capacity, copy all the info from old one, and swap the data array pointers in the given slice.```
    ```But, if you try to directly assign a value to nil slice this will result in painc (same as map). ```

13) Is it required to always care about the return value from the append function?

    Ans- Yes, we should always store the retuned result. Appned function internally allocates new array(if cap is full) and return the resulted slice.

14) In which order performed traversal over a map(for range loop), How to traverse the map in order sorted by keys?

    Ans - map stores data based on hash function (i.e no order maintained).

      ```To traverse map in sorted by keys, follow below steps:
      1.1) store keys in slice
      1.2) sort this slice
      1.3) iterate over this slice and find corresponding value from map.```

        func main() {
          dict := map[string]int{
            "abc": 2,
            "ref": 5,
            "adf": 1,
          }

          slice := make([]string, 0)

          for key := range dict {
            slice = append(slice, key)
          }

          sort.Strings(slice)

          for _, v := range slice {
            fmt.Println(v, dict[v])
          }
        }
