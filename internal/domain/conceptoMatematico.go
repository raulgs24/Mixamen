package domain

type conceptoMatematico string //Concepto Matemático es un objeto valor modelado como un tipo en Go

const ( //Definimos los conceptos permitidos como un enum
	conceptoAlgebra      conceptoMatematico = "Álgebra"
	conceptoGeometria    conceptoMatematico = "Geometría"
	conceptoAritmetica   conceptoMatematico = "Aritmética"
	conceptoProbabilidad conceptoMatematico = "Probabilidad"
)
