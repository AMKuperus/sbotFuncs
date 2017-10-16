//Package sbotFuncs contains functions for slack bots
package sbotFuncs
//stat with function name
// GetTimeOffDay returns time-of-day string for welcoming user
func GetTimeOffDay(h int) string {
  var tod string
  switch {
  case h < 12 && h >= 6:
      tod = "Goodmorning "
    case h < 18 && h >= 12:
      tod = "Goodafternoon "
    case h < 24 && h >= 18:
      tod = "Goodevening "
    case h <= 6:
      tod = "Goodnight "
  }
  return tod
}
