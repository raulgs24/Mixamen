package domain

type Ejercicio struct { //Representa el problema matemático y actúa como la entidad principal del problema
	id int
	conceptos map[conceptoMatematico]struct{} //Se pone struct{} puesto que se necesita clave y valor, pero en este caso lo que es realmente importante es la clave, creando así es hashmap, lo mismo ocurre con criterios
	criterios map[criterioEvaluacion]struct{}
}