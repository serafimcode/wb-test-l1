# Questions

1. Самый эффективный способ конкатенации строк string.Builder
2. Интерфейс - это коллекция сигнатур методов. Чтобы реализовать интерфейс,
тип должен реализовать методы, т.е. происходит неявная реализация интерфейса.
Интерфейсы используются для уменьшения сцепленности кода, уменьшения зависимости кода,
т.к. работая с интерфейсом мы абстрагируемся от конкретных реализаций.
3. RWMutex в отличие от Mutex не запрещает параллельного чтения. Т.е. несколько горутин
могут читать данные объекта, 
4. Буферизированный канал, в отличие от буферизированного не блочит поток выполнения.
Обычному каналу для работы необходимо наличие получателя и отправителя в виде запущенных 
горутин. Буферизированный канал, позволяет писать и читать данные в одном потоке.
5. Размер структуры равен сумме размеров ее полей, следовательно, размер структуры 
struct{}{} равен 0.
6. В go нет перегрузки ни операторов ни методов/функций
7. Порядок вывода данных из map не гарантируется, поэтому при каждом запуске
мы можем получить разный вывод
8. make применяется только к каналам, слайсам и мапе, возвращает нам инициализированное зн-е соответствующего типа
new алоцирует память, но никак ее не инициализирует, и возвращает нам указатель на эту область
9. Для слайса и мапы существует 3 способа задать зн-е 
   1. Строковый литерал `a := []int{1,2,3,4}`- получаем слайс, указывающий на созданный под него массив
   2. make `a := make([]int, 3, 5)` - создается массив и слайс, указывающий на этот массив(массив заполнен нулевыми зн-ми)
   3. new `a := new([]int)` - получаем слайс, указывающий на nil
Аналогично те же методы работают с мапой
10. Два раза выведется 1, потому что внутри `update(p *int)` p - это локальная переменная, в которой содержится
значение типа указатель. Поэтому, когда мы делаем `p = &b` мы просто меняем внутри одно зн-е на другое. Чтобы, добиться
предполагаемого эффекта сделаем `*p = b`, тут мы идем по указателю лежащему в p, попадаем в область памяти, где хранится
зн-е, и в эту область памяти пишем новое зн-е
11. Мы получим deadlock, основная горутина заблокируется, потому что ни одна горутина не закончит свое выполнения или точнее,
не сообщит о своем завершении главной горутине. Так происходит, потому что мы передаем зн-е WaitGroup, соответственно,
в каждой горутине будет свой wg. Чтобы все заработало, передадим wg по указателю
12. Выведется 0, потому что в go область видимости ограничена {}. Т.е. n объявленная в теле if, это не та же самая n, которую
объявили выше в main
13. Выведется `[100 2 3 4 5]`. Когда мы объявили слайс `var a = []int8{1, 2, 3, 4, 5}` мы так же его проинициализировали.
В результате такой инициализации в памяти создается массив. `v[0] = 100` сработает,
потому что внутри `someFunction` v - это копия изначального слайса, но с тем же указателем на массив в памяти.
Когда мы делаем `v = append(v, b)` мы работаем с локальной переменной поэтому за пределами ф-ии мы не видим изменений
14. Выведется `[b b a] [a a]`. Причина та же, что и в предыдущем задании. Слайс внутри ф-ия это копия нашего слайса, помещенная 
в локальную переменную. Эта локальная переменная перезаписывается и далее модифицируется. Поэтому снаружи мы не видим изменений.
