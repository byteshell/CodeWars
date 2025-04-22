package kata

func Grow(arr []int) int {
  result := 1
  
  for _, value := range arr {
      result *= value
  }
  
  return result
}
