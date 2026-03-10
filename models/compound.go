package models

// Compound represents a molecular compound with its properties and receptor bindings
type Compound struct {
    ID              string            `json:"id"`
    Name            string            `json:"name"`
    MolecularWeight float64           `json:"molecular_weight"`
    SMILES          string            `json:"smiles"`
    InChI           string            `json:"inchi"`
    ReceptorBindings []ReceptorBinding `json:"receptor_bindings"`
}

// ReceptorBinding represents the binding affinity of a compound to a specific receptor
type ReceptorBinding struct {
    ReceptorType string  `json:"receptor_type"`
    Affinity     float64 `json:"affinity"`
    Units        string  `json:"units"`
}