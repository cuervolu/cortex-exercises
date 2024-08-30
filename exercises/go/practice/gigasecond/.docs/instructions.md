# 🎂 ¡Feliz Gigacumpleaños en Go! 🎉

¿Alguna vez te has preguntado cuándo cumplirás mil millones de segundos? ¡Eso es un gigasegundo!
Suena como algo sacado de una película de ciencia ficción, ¿verdad? Pues es real, y tú vas a
calcularlo usando Go.

## 🤔 ¿Qué es un gigasegundo?

Un gigasegundo es 1,000,000,000 segundos. Sí, ¡eso es un 1 con nueve ceros detrás! Es muchísimo
tiempo:

- Es más de 31 años
- Es como ver "The Simpsons" sin parar durante casi 12 años
- Es el tiempo que tardarías en contar hasta mil millones si contaras un número por segundo (¡no lo
  intentes en casa!)

## 🚀 Tu misión, si decides aceptarla...

Tu tarea es crear una máquina del tiempo en código Go. Bueno, no exactamente, pero casi:

1. Te daremos una fecha y hora de nacimiento.
2. Tu programa debe calcular exactamente cuándo esa persona cumplirá un gigasegundo de edad.

Por ejemplo:

- Si "naciste" el 24 de enero de 2015 a las 22:00 (10:00 PM)
- Cumplirías un gigasegundo el 2 de octubre de 2046 a las 23:46:40 (11:46:40 PM)

## 🛠 Herramientas para tu máquina del tiempo en Go

- Usarás el paquete `time` de Go, que proporciona tipos y funciones para trabajar con fechas y
  horas.
- Tu función se llamará `AddGigasecond` y recibirá un `time.Time` como fecha de inicio.
- Debe devolver otro `time.Time` que represente la fecha y hora exactas un gigasegundo después.

