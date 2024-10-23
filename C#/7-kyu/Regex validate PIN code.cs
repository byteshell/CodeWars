using System;
using System.Text.RegularExpressions;

public class Kata
{
  public static bool ValidatePin(string pin)
  {
    return Regex.IsMatch(pin, @"^\d{4}$|^\d{6}$") && pin.Trim().Length == pin.Length;
  }
}
