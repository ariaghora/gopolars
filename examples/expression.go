package main

import (
	"fmt"

	pl "github.com/ariaghora/gopolars"
)

func main() {
	df, _ := pl.ReadCSV("examples/data/iris.csv")

	// Show sepal length and width of Setosa samples,
	// take top 5 rows
	df, _ = df.Lazy().
		Filter(
			pl.Col("variety").Eq(pl.Str("Setosa")),
		).
		Select(
			pl.Col("sepal.length"),
			pl.Col("sepal.width"),
			pl.Col("variety").Alias("Class"),
		).
		Collect()

	df = df.Head(5)

	fmt.Println(df)
}
