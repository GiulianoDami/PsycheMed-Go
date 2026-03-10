package analysis

import (
	"math"
)

// Compound represents a psychedelic compound structure
type Compound struct {
	Name              string
	MolecularWeight   float64
	HydrogenBondCount int
	Charge            float64
	ShapeFactor       float64
}

// AnalysisResult contains the binding affinity and therapeutic score
type AnalysisResult struct {
	CompoundName         string
	BindingAffinity      float64
	TherapeuticScore     float64
	ReceptorSpecificity  map[string]float64
	SideEffectPrediction float64
}

// AnalyzeCompound calculates binding affinities and therapeutic scores for a compound
func AnalyzeCompound(c Compound) AnalysisResult {
	// Calculate binding affinity based on molecular properties
	bindingAffinity := calculateBindingAffinity(c)
	
	// Calculate therapeutic score
	therapeuticScore := CalculateTherapeuticScore(c)
	
	// Calculate receptor specificity
	receptorSpecificity := map[string]float64{
		"5-HT2A": 0.8 * bindingAffinity,
		"5-HT2B": 0.3 * bindingAffinity,
		"5-HT2C": 0.6 * bindingAffinity,
		"5-HT1A": 0.9 * bindingAffinity,
	}
	
	// Calculate side effect prediction
	sideEffectPrediction := calculateSideEffectPrediction(c)
	
	return AnalysisResult{
		CompoundName:         c.Name,
		BindingAffinity:      bindingAffinity,
		TherapeuticScore:     therapeuticScore,
		ReceptorSpecificity:  receptorSpecificity,
		SideEffectPrediction: sideEffectPrediction,
	}
}

// CalculateTherapeuticScore computes a therapeutic score based on binding affinity and safety factors
func CalculateTherapeuticScore(c Compound) float64 {
	// Base binding affinity calculation
	baseScore := calculateBindingAffinity(c)
	
	// Safety factor based on hydrogen bond count and charge
	safetyFactor := math.Exp(-math.Abs(c.HydrogenBondCount-3) * 0.5) * math.Exp(-math.Abs(c.Charge) * 0.1)
	
	// Shape factor influence
	shapeInfluence := 1.0 / (1.0 + math.Exp(-c.ShapeFactor+2.0))
	
	// Combined therapeutic score
	therapeuticScore := baseScore * safetyFactor * shapeInfluence
	
	// Normalize to 0-1 range
	maxPossible := 100.0
	if therapeuticScore > maxPossible {
		therapeuticScore = maxPossible
	}
	
	return therapeuticScore / maxPossible
}

// calculateBindingAffinity computes the binding affinity based on molecular characteristics
func calculateBindingAffinity(c Compound) float64 {
	// Simplified model based on molecular weight, hydrogen bonds, and charge
	affinity := 0.0
	
	// Weight contribution (optimal around 300-500 Da)
	weightContribution := math.Exp(-math.Pow((c.MolecularWeight-400)/100, 2))
	
	// Hydrogen bond contribution (optimal around 3-5)
	hbContribution := math.Exp(-math.Pow(float64(c.HydrogenBondCount)-4, 2)/2)
	
	// Charge contribution (neutral is optimal)
	chargeContribution := math.Exp(-math.Pow(c.Charge, 2) * 0.5)
	
	// Combine contributions
	affinity = weightContribution * hbContribution * chargeContribution * 100
	
	return affinity
}

// calculateSideEffectPrediction estimates potential side effects
func calculateSideEffectPrediction(c Compound) float64 {
	// Predict side effects based on molecular features
	sideEffect := 0.0
	
	// Higher hydrogen bond count generally reduces hallucinogenic effects
	hbReduction := math.Exp(-float64(c.HydrogenBondCount) * 0.3)
	
	// Charge affects receptor binding specificity
	chargeImpact := math.Exp(-math.Abs(c.Charge) * 0.2)
	
	// Shape factor influences selectivity
	shapeSelectivity := 1.0 / (1.0 + math.Exp(-c.ShapeFactor+1.0))
	
	// Combine factors
	sideEffect = (hbReduction * 0.4 + chargeImpact * 0.3 + shapeSelectivity * 0.3)
	
	return sideEffect
}