# Hacer CPU Profiling: pprof

Para averiguar cuales funciones llevan más tiempo (y detectar posibles cuellos de botella) se debe hacer lo que llaman un "profile": lo que en realidad sirve para ver cuanto tiempo de CPU (procesador) se asigna a cada función durante un determinado tiempo de ejecución de nuestro programa.

Esto se logra basicamente, agregando un conjunto especifico de endpoint (URLS) cuya funcionalidad viene provista en la biblioteca del lenguaje, en particular en el paquete "net/http/pprof".

Para ver como lidiar con esto de modificar el servidor web para que tenga estos endpoints especiales, se pueden consultar [las fuentes que están abajo](#fuentes) o bien el código que esta en [/presentation/web/server.go](https://github.com/vituchon/labora-golang-course/blob/master/meeting-crud-api/presentation/web/server.go#L21)

## Pasos

1. Arrancar la app web que queda funcionado en <app_hostname>, tipicamente un "localhost:8080"
    1. Si levantamos un servidor para hacer profile en otro puerto entonces tener en mente ese otro puerto para realizar las peticiones a los endpoints que hacen profiling.
2. Ejecutar `curl --output profile.out "http://<app_hostname>/pprof/profile?seconds=10"` y se guardará en el archivo `profile.out` el informe.
    1. Notar el parametro `seconds=10` con valor igual 10, cuanto más tiempo se le dé Y cuanto más se  lo use durante ese tiempo, más presiso será el informe que generemos con la herramienta.
    2. 👁️ Realmente hay que hacerlo trabajar al servidor para que salga el reporte, o sea inundenlo de requests!!!
        1. Consideren usar el truco de hacer muchos hits (peticiones a un endpoint) usando los comandos que vienen con el interprete de comandos (bash), yo hice algo como esto: `for((i=1;i<=100;i+=1)); do curl "http://<app_hostname>/endpoint"; done`, y sí! se ejecuta 100 veces un mismo curl!! que mejor forma de hacer trabajar al servidor que usando la propia computadora! y que fácil es bombardear la red de peticiones!!!! no lo hagan en casa! (pueden probar el comando `for((i=1;i<=100;i+=1)); do echo "hola bash!!"; done` que saluda al bash 100 veces!)
3. Acá tenemos dos alternativas:
    1. Ejecutar`go tool pprof profile.out` para arrancar un programa que espera nuestra entrada por teclado, podemos escribir `web`(*) y enter para visualizar gráficamente o bien podemos escribir `top` (para ver las 10 funciones que más tiempo usaron el procesador)
    2. Ejecutar `go tool profile -http <app_hostname>/ profile.out` para que se vea lindo a travéz de una pestaña del navegador, tal como se logra con el comando "web"(*) mencionado anteriormente.

👁️ En el reporte de profile verán que aparecen MUCHISIMAS funciones que son de los paquetes que importamos y no tenemos idea de que existian hasta ahora... ufff, bueno vale ignorar! se puede buscar por nombre de función y van a ver que se remarcan los cuadros de sus funciones con las métricas!!!!

## Fuentes

* <https://www.jajaldoang.com/post/profiling-go-app-with-pprof/>
* <https://groups.google.com/g/golang-nuts/c/TjDMXyBDYG4>
* <https://stackoverflow.com/a/34000544/903998>

