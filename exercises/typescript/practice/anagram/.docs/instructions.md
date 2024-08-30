# 🔤 ¡Revoltijo de Letras! 🔠

¿Alguna vez has jugado a revolver las letras de una palabra para formar otra? ¡Eso es un anagrama! Hoy vamos a enseñar a nuestra computadora a ser una experta en este juego.

## 🎯 Tu Misión

Tu tarea es crear un detector de anagramas supersónico. Aquí está lo que debe hacer:

1. Recibir una palabra (llamémosla la "palabra original").
2. Recibir una lista de palabras posibles.
3. ¡Encontrar todas las palabras de la lista que son anagramas de la palabra original!

Por ejemplo:

- Palabra original: "listen"
- Lista de palabras: ["enlists", "google", "inlets", "banana"]
- Tu programa debería decir: ¡"inlets" es un anagrama de "listen"!

## 🤓 Reglas del Juego

1. Los anagramas deben usar todas las letras de la palabra original, ni más ni menos.
2. La misma letra puede usarse solo tantas veces como aparezca en la palabra original.
3. Las mayúsculas y minúsculas no importan (así que "TypeScript" y "typescript" se consideran iguales).
4. La palabra original no es su propio anagrama (¡sería muy fácil!).

## 🚀 ¡A Programar

Tu misión, si decides aceptarla, es implementar la función `anagrams`. Esta función debe:

- Recibir la palabra original como un `string`.
- Recibir la lista de posibles anagramas como un `string[]`.
- Devolver un `string[]` con todos los anagramas encontrados.

¡Buena suerte, maestro de las palabras revueltas! 🧙‍♂️✨
