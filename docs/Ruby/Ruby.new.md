# Ruby Class



ref:

http://ruby-doc.com/docs/ProgrammingRuby/

## classes

和JAVA中的语法类似，class names start with an uppercase letter，while method names start with a lowercase letter

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

initalize 即JAVA中的构造函数，而name、artist、duration即成员变量(以@开头，你也可以称为实例变量，==ruby 中默认等同与JAVA private==)

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

### virtual attr



### to_s

```
aSong = Song.new("Bicylops", "Fleck", 260)
aSong.to_s	»	"#<Song:0x401b499c>"
```

同样的ruby中也有`toString`函数即`to_s`

### inheritance

继承逻辑和所有编程语言一样

```
class KaraokeSong < Song
  def initialize(name, artist, duration, lyrics)
    super(name, artist, duration)
    @lyrics = lyrics
  end
end
```

这里的`<`即表示继承自Song，super含义和JAVA，javascript中一样。

但是需要注意的是虽然ruby和JAVA一样都是单继承，但是ruby支持mixins(继承类的一部分)