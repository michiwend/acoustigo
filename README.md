acoustigo
=========

a Go (Golang) AcousticBrainz client library


## Example usage
just a simple example, more to come...

```Go
client, _ := acoustigo.NewABClient("http://acousticbrainz.org")

result, err := client.HighLevel("96685213-a25c-4678-9a13-abd9ec81cf35")
if err != nil {
	log.Fatal(err)
}

for k, v := range result.HighLevel {
	fmt.Printf("Found high level data: %s\n", k)
	fmt.Printf("Value: %s with probability of %f\n\n", v.Value, v.Probability)
}
```
