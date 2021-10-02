# Seminario-Golang

Para la resolución del trabajo propuesto se creo un modulo con go mod TPEGolang.

Se decidio que en el main la cadena pueda ser creada por teclado, una vez guardada la cadena ingresada en una variable, se llama a la funcion GetResult pasandole la cadena, de no dar error se imprime el resultado, sino se imprime lo ingresado por teclado mas un mensaje de error ("cadena INVALIDA").

Dentro del paquete model se creo una estructura Result la cual almacena el resultado del parseo conteniendo el tipo, valor y longitud.
Tambien se creo una funcion GetResult la cual recibe una cadena de caracteres y devuelve dos parametros: un puntero a una estructura de tipo resultado y un error. La misma se encarga del parseo y de encontrar errores en el formato de la cadena recibida. Para esto, primero se fija en si el largo de la cadena es mayor que 5, de ser asi, lo siguiente que hace es dividir la cadena en 3 variables(las primeras 2 para el tipo, las siguientes 2 para el largo y todas las demas para el valor), se chequea que las variables de el largo y la del valor coincidan (que tengan las mismas cantidad de lugares-letras/numeros), si coinciden se pasa a chequear el tipo de la cadena, para esto se llama a la funcion checkType (recibe un string y devuelve un error) que se encarga de comprobar que lo que recibe sea un tipo valido ("NN/nn" o "TX/tx"). Si todos los chequeos no dan errores se llama a la funcion para crear el resultado pasandole las 3 variables.  

Adicionalmente se realizaron los test correspondientes (dando un 100% de coverage, sin contar el main) como se puede visualizar de forma gráfica en el html generado (out.html).