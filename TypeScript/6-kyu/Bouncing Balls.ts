export function bouncingBall(h: number, bounce: number, window: number): number {
  if (h <= 0 || bounce <= 0 || bounce >= 1 || window >= h) {
    return -1;
  }
  
  let count = 0;
  
  count++;
  
  while (h * bounce > window) {
    count += 2;
    h *= bounce;
  }
  
  return count;
}
