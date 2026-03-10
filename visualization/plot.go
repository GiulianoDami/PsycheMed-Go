package visualization

import (
	"fmt"
	"image/color"
	"math"

	"github.com/go-gota/gota/dataframe"
	"github.com/wcharczuk/go-chart/v2"
)

// PlotReceptorInteractions generates a visualization of receptor interaction patterns for a given compound
func PlotReceptorInteractions(compound Compound) error {
	// Create a simple bar chart showing receptor binding affinities
	bindings := compound.GetReceptorBindings()
	
	if len(bindings) == 0 {
		return fmt.Errorf("no receptor bindings found for compound")
	}
	
	// Prepare data for chart
	labels := make([]string, len(bindings))
	values := make([]float64, len(bindings))
	
	for i, binding := range bindings {
		labels[i] = binding.ReceptorType
		values[i] = binding.Affinity
	}
	
	// Create the chart
	graph := chart.Chart{
		Title:      fmt.Sprintf("Receptor Interactions for %s", compound.Name),
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		XAxis: chart.XAxis{
			Name:      "Receptor Type",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Binding Affinity",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name: "Binding Affinity",
				Style: chart.Style{
					StrokeColor: chart.ColorBlue,
					FillColor:   chart.ColorBlue.WithAlpha(100),
				},
				XValues: chart.TimeSeriesToFloat64Slice(labels),
				YValues: values,
			},
		},
	}
	
	// Render the chart to a file
	filename := fmt.Sprintf("%s_interactions.png", compound.Name)
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()
	
	err = graph.Render(chart.PNG, f)
	if err != nil {
		return fmt.Errorf("failed to render chart: %w", err)
	}
	
	return nil
}

// Helper function to convert time series to float64 slice
func TimeSeriesToFloat64Slice(timeSeries []string) []float64 {
	result := make([]float64, len(timeSeries))
	for i, _ := range timeSeries {
		result[i] = float64(i)
	}
	return result
}