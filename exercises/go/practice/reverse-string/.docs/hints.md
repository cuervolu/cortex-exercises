## 💡 Pistas para Gophers Inversores

1. 🧵 En Go, los strings son inmutables, pero puedes convertirlos fácilmente a un slice de runes (caracteres Unicode) con `[]rune(s)`.

2. 🔄 Piensa en cómo podrías intercambiar los caracteres desde los extremos hacia el centro.

3. 📏 La función `len()` te da la longitud de un string en bytes, pero `len([]rune(s))` te da la cantidad de caracteres.

4. 🏗️ Puedes construir un nuevo string a partir de un slice de runes usando `string(runeSlice)`.

5. 📚 Investiga el paquete `strings` de Go. Podría tener algunas funciones útiles, aunque no es necesario para resolver este problema.

6. 🧠 Recuerda: en Go, la eficiencia es importante. ¿Puedes pensar en una solución que solo recorra la cadena una vez?

Recuerda: En Go, la claridad y la simplicidad son virtudes. Intenta encontrar una solución elegante y eficiente.

¡Adelante, mago de las palabras inversas en Go! 🧙‍♂️✨