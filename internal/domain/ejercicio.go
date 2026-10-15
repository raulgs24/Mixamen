package domain

type Ejercicio struct { //Representa el problema matemático y actúa como la entidad principal del problema
	id int
	concepto *conceptoMatematico
	criterio *criterioEvaluacion
}