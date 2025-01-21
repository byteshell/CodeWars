package kata

func MinMax(arr []int) [2]int {
    if len(arr) == 0 {
        return [2]int{}
    }

    min, max := arr[0], arr[0]
    for _, num := range arr {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }

    return [2]int{min, max}
}
