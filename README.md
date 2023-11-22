Эта практика направлена на отработку навыка создания кластера k8s, соответствующего следующим параметрам:

- мультизональный кластер (три зоны), в котором пять нод
- приложение требует около 5-10 секунд для инициализации
- по результатам нагрузочного теста известно, что 4 пода справляются с пиковой нагрузкой
- на первые запросы приложению требуется значительно больше ресурсов CPU, в дальнейшем потребление ровное в районе 0.1 CPU. По памяти всегда “ровно” в районе 128M memory
- приложение имеет дневной цикл по нагрузке – ночью запросов на порядки меньше, пик – днём
- максимально отказоустойчивый deployment
- минимальное потребление ресурсов от этого deployment’а

Была использована программа "Hello world!, написанная на языке GO.

This practice is aimed at practicing the skill of creating a k8s cluster that meets the following parameters:

- multi-zone cluster (three zones) in which five nodes
- the application requires about 5-10 seconds to initialize
- according to the results of the load test, it is known that 4 pods cope with the peak load
- for the first requests, the application requires significantly more CPU resources, in the future the consumption is equal in the area of 0.1 CPU. Memory is always "level" in the area of 128M memory
- the application has a daily load cycle - at night there are fewer orders of magnitude requests, peak - during the day
- the most fault-tolerant deployment
- minimum resource consumption from this deployment

The program "Hello world!, written in GO.
