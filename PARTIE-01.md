**NOM:** 
**Prénom:** 

## Questions de cours
*15 points*

01. Qu'est-ce que la librairie standard ? *(1 point)*
La librairie standard est la librairie de packages développée et fournie par Go. Chacun des packages de cette librairie est conçu pour être directement utilisable à l'import.


02. Dans un package main, qu'est-ce que la fonction main() ? Quelle est la
    particularité de la signature de cette fonction ? *(1 point)*
La fonction main est la fonction est la principale fonction d'une application, celle qui détermine si le package doit être compilé comme programme éxécutable. Elle ne prend aucune valeur en paramètre et n'en retourne aucune (signature vide), elle ne fait qu'éxécuter des tâches.


03. Est-il possible d'importer un package main ? *(0,5 point)*
Non, il est impossible d'importer un package main. De plus, sa fonction main() n'est pas exportable puisque non-capitalisée.


04. Est-il possible qu'une valeur de type string, boolean, integer ou float ait
    une valeur undefined ? Pourquoi ? *(1 point)* 
Non, le type undefined n'existe pas en Go. La "zero-value" (valeur vide) d'un chaîne de caractères est une chaîne vide, celle d'un booléen est false, celle d'un nombre (entier ou flottant) est 0.


05. Est-il possible de changer le type d'une variable ? Pourquoi ? *(1 point)*
Go est un langage typé. Il est donc impossible de changer le type d'une variable puisque celui-ci est défini implicitement ou déduit (inferred) par le compilateur au moment de l'initialisation de celle-ci.


06. Une fois qu'un map est initialisé, connait-on le type des clés et le type des
    valeurs de celui-ci ? Est-ce que ces types peuvent changer ? *(0,5 point)*
Au moment de l'initialisation d'un map, il est nécessaire de préciser le type des clés et des valeurs. Si l'on ne connaît pas à l'avance le type de ces derniers, il est possible d'y faire abstraction en utilisant interface{} comme type sous-jacent (underlying type) afin de contenir des clés de types différents/inconnus ou des valeurs de types différents/inconnus. En revanche, si l'on définit un type autre qu'interface{} en clé ou en valeur, celui-ci ne peut pas changer.


07. Quelle est la différence entre un slice et un array ? Pourquoi fait-on une
    distinction entre ces deux types ? *(1 point)*
Un slice est un tableau/une liste de valeurs. Il peut contenir un nombre indéterminé de valeurs. Un array est un tableau/une liste de valeurs dont le nombre de rangs/valeurs est défini à l'avance. Il y a une distinction entre ces deux types pour que l'on puisse limiter le nombre de valeurs que l'on stocke dans un tableau si le programme le requiert, ou que l'on le puisse stocker des valeurs sans en connaître le nombre final.


08. Qu'est-ce que le blank identifier ? A quoi sert-il ? *(1 point)*
Le blank identifier est utilisé afin d'ignorer une des valeurs retournées par une fonction qui en retourne plusieurs. Exemple : 
premiereValeurRetournee, _, troisiemeValeurRetournee := fonctionQuiRetourneTroisValeurs()
On souhaite récupérer la première et la troisième valeur que cette fonction retourne, la seconde doit donc être définie comme ignorée avec le blank identifier.


09. Expliquez simplement le concept et l'utilité des pointeurs. *(3 points)*
Un pointeur permet de cibler une valeur unique afin de la conserver en mémoire pour la modifier, la supprimer ou simplement y faire référence dans une fonction. Il est également possible qu'une fonction retourne un pointeur d'un type pour en créer une nouvelle valeur unique.


10. Qu'est-ce qu'une interface ? Quels sont les avantages d'utiliser un tel
    concept ? *(2 points)*
Un interface peut, comme indiqué à la question 6, servir de type sous-jacent lorsqu'on souhaite faire abstraction d'un type ou lorsque le type d'une valeur est amené à changer.
On peut également implémenter des méthodes à une interface afin d'y implémenter implicitement un type (souvent une struct) qui pourra utiliser ces méthodes.


11. Expliquez comment nous avons mis en place l'éxécution de tâches asynchrones
    en détaillant les concepts qui nous ont permis de faire cela. *(3 points)*
Une tâche asynchrone peut se mettre en place grâce au type WaitGroup du package sync de la librairie standard. Il suffit de créer une fonction dans laquelle on définit les tâches asynchrones suivies de defer Done() (defer éxécute une tâche une fois que tout le code du bloc actuel a été éxécuté et Done désincrémente le nombre de 1 de tâches restantes dans le WaitGroup, c'est une fonction du package sync qui est une factory d'un WaitGroup). Puis, on utilise le mot clé go suivi de l'appel de cette fonction (goroutine) afin de l'éxécuter en de manière asynchrone (en arrière-plan pour pouvoir continuer la suite de l'éxécution du code). Il est également possible d'utiliser une goroutine pour une fonction anonyme. Si l'on souhaite imposer l'attente de la fin de ce process pour continuer l'éxécution du code, on utilise la fonction Wait() du package sync, qui est une factory d'un WaitGroup.
