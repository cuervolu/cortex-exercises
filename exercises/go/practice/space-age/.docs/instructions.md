# 🚀 ¡Viaje Intergaláctico de Cumpleaños en Go! 🎂

¿Alguna vez te has preguntado cuántos años tendrías si vivieras en Marte? ¿O cuántas velas tendría
tu pastel de cumpleaños en Júpiter? ¡Hoy vamos a crear una máquina del tiempo espacial en Go para
averiguarlo!

## 🌟 Tu Misión Cósmica

Tu tarea es crear un calculador de edad intergaláctico en Go. Aquí está lo que debe hacer:

1. Recibir una edad en segundos (¡sí, segundos! Los científicos espaciales son muy precisos).
2. Recibir el nombre de un planeta.
3. Calcular cuántos años en ese planeta son esos segundos.

## 🌍 Datos del Centro de Control Espacial

- Un año terrestre = 365.25 días terrestres = 31,557,600 segundos
- Si alguien tiene 1,000,000,000 segundos de edad, tendría 31.69 años terrestres.

Pero, ¡atención astronauta! Cada planeta tiene su propio "año", basado en cuánto tarda en dar una
vuelta alrededor del Sol:

| Planeta  | Período orbital en años terrestres |
|----------|------------------------------------|
| Mercurio | 0.2408467                          |
| Venus    | 0.61519726                         |
| Tierra   | 1.0                                |
| Marte    | 1.8808158                          |
| Júpiter  | 11.862615                          |
| Saturno  | 29.447498                          |
| Urano    | 84.016846                          |
| Neptuno  | 164.79132                          |

## 🛸 Instrucciones de Vuelo

1. Implementa la función `Age` que recibe:
    - `seconds` como `float64`: la edad en segundos
    - `planet` como `Planet` (que es un alias para `string`): el nombre del planeta
2. La función `Age` debe devolver la edad en años del planeta correspondiente como un `float64`.

## 🌠 ¡Despegamos!

¿Listo para tu viaje intergaláctico en Go? ¡Abróchate el cinturón y empieza a programar! Recuerda,
en el espacio nadie puede oír tus errores de compilación, ¡así que prueba tu código con cuidado!
