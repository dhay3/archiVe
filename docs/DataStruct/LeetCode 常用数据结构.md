### LeetCode 常用数据结构

[TOC]

#### #Stack 栈

#### #List 链表

- 静态链表 ArrayList

  便于查找

- 动态链表 LinkedList

  便于删除和插入

#### #Set 

去重

#### #Map

键值对, 去重

#### # Pair 配对

```java
	Pair<Integer, String> pair = new Pair<>(1, "One");
    Integer key = pair.getKey();
    String value = pair.getValue();
```

相当于`map`的一对键值对

#### #Queue 队列

- LinkedList

#### #PriorityQueue 优先队列

给定Comparator或是Comparable, 按照定义的比较规则插入队列头

```java
Queue<User> q = new PriorityQueue<>(Comparator<T>);
```



#### #Deque 双端队列

支持在队列两端插入和删除, 可以取代队列和栈

- DequeArray

  

| Queue方法 | 等效Deque方法 | Stack方法(头插) |
| --------- | ------------- | --------------- |
| add(e)    | addLast(e)    |                 |
|           | addFirst(e)   | push()          |
| offer(e)  | offerLast(e)  |                 |
| remove()  | removeFirst() | pop()           |
|           | removeLast()  |                 |
| poll()    | pollFirst()   |                 |
| peek()    | peekFirst()   | peek()          |
|           |               |                 |
|           |               |                 |

