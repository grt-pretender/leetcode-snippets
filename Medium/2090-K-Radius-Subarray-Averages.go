// You are given a 0-indexed array nums of n integers, and an integer k.
// The k-radius average for a subarray of nums centered at some index i with the radius k 
// is the average of all elements in nums between the indices i - k and i + k (inclusive). 

// If there are less than k elements before or after the index i, then the k-radius average is -1.
// Build and return an array avgs of length n where avgs[i] 
// is the k-radius average for the subarray centered at index i.

// The average of x elements is the sum of the x elements divided by x, using integer division. 
// The integer division truncates toward zero, which means losing its fractional part.


func getAverages(nums []int, k int) []int {
  
  // using sliding window
  _default := -1
  avg_result := make([]int, len(nums))
  for i := range avg_result {
    avg_result[i] = _default
        }

  left_ptr, window_sum := 0, 0
  window_size := 2 * k + 1
  
  for right_ptr := range nums {
    window_sum += nums[right_ptr] 

    if right_ptr - left_ptr + 1 == window_size {
      center_index := right_ptr - k
      avg_result[center_index] = window_sum / window_size
      window_sum -= nums[left_ptr]
      left_ptr += 1
    }
  }
  return avg_result 
 
}
