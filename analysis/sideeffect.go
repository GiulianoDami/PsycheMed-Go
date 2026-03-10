package analysis

import (
	"math"
)

// Compound represents a psychedelic compound with receptor binding data
type Compound struct {
	Name              string
	PS5HT2ABinding    float64 // Binding affinity to 5-HT2A receptor (nM)
	PS5HT2BBinding    float64 // Binding affinity to 5-HT2B receptor (nM)
	PS5HT1ABinding    float64 // Binding affinity to 5-HT1A receptor (nM)
	PS5HT2CBinding    float64 // Binding affinity to 5-HT2C receptor (nM)
	PS5HT6Binding     float64 // Binding affinity to 5-HT6 receptor (nM)
	PS5HT7Binding     float64 // Binding affinity to 5-HT7 receptor (nM)
	PS5HT1BBinding    float64 // Binding affinity to 5-HT1B receptor (nM)
	PS5HT1DBinding    float64 // Binding affinity to 5-HT1D receptor (nM)
	PS5HT2AICBinding  float64 // Binding affinity to 5-HT2AI receptor (nM)
	PS5HT2AIBinding   float64 // Binding affinity to 5-HT2AI receptor (nM)
}

// AnalysisResult holds the results of the hallucination risk analysis
type AnalysisResult struct {
	CompoundName       string
	HallucinationRisk  float64
	ReceptorBindings   map[string]float64
	Recommendations    []string
}

// PredictHallucinationRisk calculates the hallucinogenic potential of a compound
// based on its binding affinities to serotonin receptors
func PredictHallucinationRisk(compound Compound) float64 {
	// Normalize binding affinities (convert nM to pKi values)
	pki2A := -math.Log10(compound.PS5HT2ABinding * 1e-9)
	pki2B := -math.Log10(compound.PS5HT2BBinding * 1e-9)
	pki1A := -math.Log10(compound.PS5HT1ABinding * 1e-9)
	pki2C := -math.Log10(compound.PS5HT2CBinding * 1e-9)
	pki6 := -math.Log10(compound.PS5HT6Binding * 1e-9)
	pki7 := -math.Log10(compound.PS5HT7Binding * 1e-9)
	pki1B := -math.Log10(compound.PS5HT1BBinding * 1e-9)
	pki1D := -math.Log10(compound.PS5HT1DBinding * 1e-9)
	pki2AI := -math.Log10(compound.PS5HT2AICBinding * 1e-9)
	pki2AIb := -math.Log10(compound.PS5HT2AIBinding * 1e-9)

	// Weighted scoring system based on known hallucinogenic activity
	// 5-HT2A is the primary target for hallucinogenic effects
	// 5-HT2B and 5-HT2C also contribute significantly
	// 5-HT1A and 5-HT7 have protective effects against hallucinations
	// 5-HT6 and 5-HT1B/1D have minimal direct hallucinogenic impact

	hallucinationScore := 0.0

	// Primary hallucinogenic targets (higher weight)
	hallucinationScore += pki2A * 0.45 // 5-HT2A receptor
	hallucinationScore += pki2B * 0.25 // 5-HT2B receptor
	hallucinationScore += pki2C * 0.15 // 5-HT2C receptor

	// Protective targets (reduce hallucination risk)
	hallucinationScore -= pki1A * 0.10 // 5-HT1A receptor
	hallucinationScore -= pki7 * 0.05  // 5-HT7 receptor

	// Normalize score to 0-1 range
	maxPotential := 10.0 // Maximum possible pKi value
	minPotential := 0.0  // Minimum possible pKi value

	if hallucinationScore < minPotential {
		hallucinationScore = minPotential
	}
	if hallucinationScore > maxPotential {
		hallucinationScore = maxPotential
	}

	// Convert to risk score (0 = no risk, 1 = high risk)
	risk := hallucinationScore / maxPotential

	// Ensure risk is within bounds
	if risk < 0 {
		risk = 0
	}
	if risk > 1 {
		risk = 1
	}

	return risk
}

// GenerateRecommendations provides safety recommendations based on analysis results
func GenerateRecommendations(result AnalysisResult) []string {
	var recommendations []string

	if result.HallucinationRisk > 0.8 {
		recommendations = append(recommendations,
			"High hallucination risk detected. Consider structural modifications to reduce 5-HT2A binding.",
			"Explore substituents that decrease affinity at 5-HT2A receptor sites.",
			"Consider adding 5-HT1A agonist properties to counteract hallucinogenic effects.",
		)
	} else if result.HallucinationRisk > 0.5 {
		recommendations = append(recommendations,
			"Moderate hallucination risk detected. Evaluate current receptor binding profile.",
			"Consider optimizing 5-HT2A binding affinity to reduce psychoactive effects.",
			"Review 5-HT2B and 5-HT2C contributions to overall risk.",
		)
	} else if result.HallucinationRisk > 0.3 {
		recommendations = append(recommendations,
			"Low hallucination risk detected. Continue monitoring receptor binding patterns.",
			"Maintain focus on 5-HT2A selectivity to preserve therapeutic benefits.",
			"Consider enhancing 5-HT1A activity to improve safety profile.",
		)
	} else {
		recommendations = append(recommendations,
			"Very low hallucination risk detected. Compound shows promising therapeutic profile.",
			"Continue with current design approach for depression treatment applications.",
			"Monitor for any unexpected receptor interactions during testing.",
		)
	}

	// Add general recommendations
	if result.ReceptorBindings["5-HT2A"] > 8.0 {
		recommendations = append(recommendations,
			"High 5-HT2A binding detected. This is the primary driver of hallucinogenic effects.",
		)
	}

	if result.ReceptorBindings["5-HT1A"] > 7.0 {
		recommendations = append(recommendations,
			"High 5-HT1A binding detected. This may provide neuroprotective benefits.",
		)
	}

	if result.ReceptorBindings["5-HT2B"] > 7.0 {
		recommendations = append(recommendations,
			"Moderate 5-HT2B binding detected. Monitor for potential cardiac effects.",
		)
	}

	return recommendations
}