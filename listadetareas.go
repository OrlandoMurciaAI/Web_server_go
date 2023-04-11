package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl) // toma como parametro un slice y agrega uno elemento
}

func (t *taskList) eliminarDeLista(index int) {
	//Los tres puntos al final de t.task[index + 1:]… (operador ellipsis)
	// es porque el segundo parámetro del append no es un slice y la
	// función append recibe un item, con este operador lo que hacemos es
	// decirle a go que tome ese slice y lo “desempaquete” para que sean
	// muchos parámetros de 1 solo item y no un slice.
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) mostrarTareasCompletadas() {
	for index, tasks := range t.tasks {
		if tasks.completado {
			fmt.Printf("Tarea %d nombre: %s \n \t descripcion: %s \n",
				index, tasks.nombre, tasks.descripcion)
		}
	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t := &task{
		nombre:      "Completar mi curso de GO",
		descripcion: "Tratar de completar el curso de go",
	}

	t1 := &task{
		nombre:      "Completar mi curso de javascript",
		descripcion: "Tratar de completar el curso de javascript",
	}

	t2 := &task{
		nombre:      "Completar mi curso de python",
		descripcion: "Tratar de completar el curso de python",
	}

	t3 := &task{
		nombre:      "Completar mi curso de docker",
		descripcion: "Tratar de completar el curso de docker",
	}

	// %+v nos dice que nos va a imprimir la propiedad con su valor
	lista := &taskList{
		tasks: []*task{
			t, t1, t2,
		},
	}

	//fmt.Println(lista.tasks[1])
	lista.agregarALista(t3)
	//fmt.Println(len(lista.tasks))
	//lista.eliminarDeLista(1)
	fmt.Println(len(lista.tasks))
	lista.tasks[0].marcarCompleta()
	lista.mostrarTareasCompletadas()

	mapaTareas := make(map[string]*taskList)

	t4 := &task{
		nombre:      "Completar mi curso de JAVA",
		descripcion: "Tratar de completar el curso de JAVA",
	}

	t5 := &task{
		nombre:      "Completar mi curso de C#",
		descripcion: "Tratar de completar el curso de C#",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}
	fmt.Println("===============================")
	mapaTareas["Orlando"] = lista
	mapaTareas["Pedro"] = lista2
	fmt.Println("Tareas de Orlando")
	mapaTareas["Orlando"].mostrarTareasCompletadas()
	lista2.tasks[0].marcarCompleta()
	fmt.Println("Tareas de Pedro")
	mapaTareas["Pedro"].mostrarTareasCompletadas()

}
