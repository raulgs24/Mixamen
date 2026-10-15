package domain

type conceptoMatematico string //Concepto Matemático es un objeto valor modelado como un tipo en Go

const ( //Definimos los conceptos permitidos como un enum
	ConceptoAlgebra      ConceptoMatematico = "Álgebra"
	ConceptoGeometria    ConceptoMatematico = "Geometría"
	ConceptoAritmetica   ConceptoMatematico = "Aritmética"
	ConceptoProbabilidad ConceptoMatematico = "Probabilidad"
)
