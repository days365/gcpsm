## GCPSM

GCPSM parses Secret Manager value (only json format).

## How to Use

```
type Sec struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

c := secret.Config{
	Version:   "1",
	ProjectID: "your-project",
	Name:      "your-secret-manager-name",
}

v := Sec{}
if err := secret.NewSecret(c, &v); err != nil {
	// error handling
}
```
