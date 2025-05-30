String.prototype.toJadenCase = function () {
  return this.split(' ')
    .map(word => word.charAt(0).toUpperCase() + word.slice(1))
    .join(' '); 
};

interface String {
  toJadenCase(): string;
}
