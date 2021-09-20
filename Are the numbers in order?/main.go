package kata

var isSorted = true
var last int

func InAscOrder(numbers []int) bool {

  isSorted = true
  last = numbers[0]
  
  for i,v := range numbers {
    if i == 0 {
      last = v
    } else if last > v {
        isSorted = false
      }else{
        last = v
      }
    }
  
  
  return isSorted
}
