# JRes

JRes is a JSON response helper to ensure usage of a consistent API structure.

## Examples

```go
// Default standard response with 200
jres.OK(w, map[string]string{"test":"data"})

// Add the location it was created at
jres.Created(w, "/api/v1/thing/1", nil)
```
