### Домашнее задание
 
**Задача:** обработать N-заказов M-воркерами с применением паттернов Pipeline и WorkerPool.
Параметры N >= 5 и M >= 2 задаются как константы в глобальной области видимости.

Структура заказа состоит из следующих полей:
- ID товара
- ID склада
- ID пункта выдачи заказа
- ID воркера, который обработал заказ
- Массив состояний, в которых находился заказ в рамках пайплайна

Структура состояния заказа состоит из следующих полей:
- Наименование: [Создан, Обработан, Выполнен]
- Время перехода в данное состояние

Каждое состояние является шагом пайплайна.
Входные данные пайплайна:
- Канал, содержащий структуры заказов с инициализированным полем ID товара

Шаг "Создан" пайплайна:
- В массив состояний добавляется новое состояние "Создан"

Шаг "Обработан" пайплайна:
- Инициализируется склад для заказа - результат взятия ID товара по модулю 2
- В массив состояний добавляется новое состояние "Обработан"

Шаг "Завершен" пайплайна:
- Инициализируется пункт выдачи заказа - результат суммы ID товара и ID склада
- В массив состояний добавляется новое состояние "Завершен"

Обработка канала с заказами осуществляется M-воркерами.
Все шаги пайплайна конкретного заказа обрабатываются в рамках одного воркера.
Перед началом выполнения пайплайна каждому заказу присваивается ID воркера, который его обрабатывает.
После того, как заказ прошел все стадии обработки, необходимо передать заказ в канал с обработанными заказами.

Основной поток приложения передает в стандартный поток вывода обработанные заказы в JSON формате структуры заказа.

💎 Использовать паттерны FanIn/FanOut для шага с обработкой заказа 
