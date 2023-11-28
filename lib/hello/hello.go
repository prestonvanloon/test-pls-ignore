package hello

// Hello returns a friendly message.
func Hello(name string) string {
  if len(name) == 0 {
    return "Hello!"
  }
  return "Hello " + name + "!"
} 
