# Ruby Class


ref:


[http://ruby-doc.com/docs/ProgrammingRuby/](http://ruby-doc.com/docs/ProgrammingRuby/)


## classes


和JAVA中的语法类似，class names start with an uppercase letter，while method names start with a lowercase letter
```ruby
class Song
  @@plays = 0
  def initialize(name, artist, duration)
    @name     = name
    @artist   = artist
    @duration = duration
    @plays    = 0
  end
  def play
    @plays += 1
    @@plays += 1
    "This  song: #@plays plays. Total #@@plays plays."
  end
end
```
## variables
和JAVA的变量一样，只是一个Object的reference，定义格式和Python、GO类似会自动识别变量类型
### class varibales
类变量，以`@@varibble`的格式定义
```ruby
class Song
  @@plays = 0
  def initialize(name, artist, duration)
    @name     = name
    @artist   = artist
    @duration = duration
    @plays    = 0
  end
  def play
    @plays += 1
    @@plays += 1
    "This  song: #@plays plays. Total #@@plays plays."
  end
end

s1 = Song.new("Song1", "Artist1", 234)  # test songs..
s2 = Song.new("Song2", "Artist2", 345)
s1.play	»	"This  song: 1 plays. Total 1 plays."
s2.play	»	"This  song: 1 plays. Total 2 plays."
s1.play	»	"This  song: 2 plays. Total 3 plays."
s1.play	»	"This  song: 3 plays. Total 4 plays."
```
和JAVA的类变量有区别，不能通过`ClassName.variableName`的方式调用
### access control
这部分的逻辑和JAVA一样但是Ruby缺省默认使用public修饰变量和方法，可以通过如下两种方式来定义
```ruby
class MyClass

      def method1    # default is 'public'
        #...
      end

  protected          # subsequent methods will be 'protected'

      def method2    # will be 'protected'
        #...
      end

  private            # subsequent methods will be 'private'

      def method3    # will be 'private'
        #...
      end

  public             # subsequent methods will be 'public'

      def method4    # and this will be 'public'
        #...
      end
end
```
```ruby
class MyClass

  def method1
  end

  # ... and so on

  public    :method1, :method4
  protected :method2
  private   :method3
end
```
### freeze
ruby中对变量有一个特殊的用法，`freeze`可以方式变量被修改
```ruby
person1 = "Tim"
person2 = person1
person1.freeze       # prevent modifications to the object
person2[0] = "J"

prog.rb:4:in `=': can't modify frozen string (TypeError)
	from prog.rb:4
```
## method
### initalize


```ruby
class Song
  def initialize(name, artist, duration)
    @name     = name
    @artist   = artist
    @duration = duration
  end
end
```


initalize 即JAVA中的构造函数区别Ruby使用private修饰构造函数，而name、artist、duration即成员变量(以@开头，你也可以称为实例变量，ruby 中默认等同与JAVA private)，


### getting


```ruby
class Song
  def name
    @name
  end
  def artist
    @artist
  end
  def duration
    @duration
  end
end
aSong = Song.new("Bicylops", "Fleck", 260)
aSong.artist	»	"Fleck"
aSong.name	»	"Bicylops"
aSong.duration	»	260
```


可以上述定义getting函数，也可以通过`attr_reader`来实现同样的效果


```
class Song
  attr_reader :name, :artist, :duration
end
aSong = Song.new("Bicylops", "Fleck", 260)
aSong.artist	»	"Fleck"
aSong.name	»	"Bicylops"
aSong.duration	»	260
```


对应的函数名不变，属于是自带loombook了


### setting


有了getting一样的也必须有setting


```
class Song
  def duration=(newDuration)
    @duration = newDuration
  end
end
aSong = Song.new("Bicylops", "Fleck", 260)
aSong.duration	»	260
aSong.duration = 257   # set attribute with updated value
aSong.duration	»	257
```


也可以通过如下方式来设定


```
class Song
  attr_writer :duration
end
aSong = Song.new("Bicylops", "Fleck", 260)
aSong.duration = 257
```


### to_s


```
aSong = Song.new("Bicylops", "Fleck", 260)
aSong.to_s	»	"#<Song:0x401b499c>"
```


同样的ruby中也有`toString`函数即`to_s`
### class methods
类函数和JAVA中的一样
```ruby
class SongList
  MaxTime = 5*60           #  5 minutes
  def SongList.isTooLong(aSong)
    return aSong.duration > MaxTime
  end
end
song1 = Song.new("Bicylops", "Fleck", 260)
SongList.isTooLong(song1)	»	false
song2 = Song.new("The Calling", "Santana", 468)
SongList.isTooLong(song2)	»	true
```


## inheritance
继承逻辑和所有编程语言一样


```ruby
class KaraokeSong < Song
  def initialize(name, artist, duration, lyrics)
    super(name, artist, duration)
    @lyrics = lyrics
  end
end
```


这里的`<`即表示继承自Song，super含义和JAVA，javascript中一样。


但是需要注意的是虽然ruby和JAVA一样都是单继承，但是ruby支持mixins(继承类的一部分)
​

## arrays
和Python、Javascript中的数组类似
```ruby

a = [ 3.14159, "pie", 99 ]
a.type	»	Array
a.length	»	3
a[0]	»	3.14159
a[1]	»	"pie"
a[2]	»	99
a[3]	»	nil
b = Array.new
b.type	»	Array
b.length	»	0
b[0] = "second"
b[1] = "array"
b	»	["second", "array"]
```
ruby可以使用 negative index 来获取，这一点和python、GO类似，但是取范围的语法有一点出入
```ruby
a = [ 1, 3, 5, 7, 9 ]
a[-1]	»	9
a[-2]	»	7
a[-99]	»	nil
```

- `[start,count]`

这个语法在其他编程语言中表示`[startidx,endidx]`，这里的count从startidx开始算起
```ruby
a = [ 1, 3, 5, 7, 9 ]
a[1, 3]	»	[3, 5, 7]
a[3, 1]	»	[7]
a[-3, 2]	»	[5, 7]

```

- `[startidx..endidx]|[startidx...endidx)`

two periods 表示左闭区间、右闭区间
three periods 表示左闭区间、右开区间
```ruby
a = [ 1, 3, 5, 7, 9 ]
a[1..3]	»	[3, 5, 7]
a[1...3]	»	[3, 5]
a[3..3]	»	[7]
a[-3..-1]	»	[5, 7, 9]
```
## hashes
即map，区别于dict 有序字典，和JAVA中的map一样是无序的
```ruby
h = { 'dog' => 'canine', 'cat' => 'feline', 'donkey' => 'asinine' }
h.length	»	3
h['dog']	»	"canine"
h['cow'] = 'bovine'
h[12]    = 'dodecine'
h['cat'] = 99
h	»	{"cow"=>"bovine", "cat"=>99, 12=>"dodecine", "donkey"=>"asinine", "dog"=>"canine"}
```
