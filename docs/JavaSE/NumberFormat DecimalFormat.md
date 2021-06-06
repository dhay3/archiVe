# NumberFormat/ DecimalFormat

`NuberFormat`是`DecimalFormat`的父类

```java
        //获取实例
        NumberFormat instance = NumberFormat.getInstance();
        //获取带有金额的实例
        NumberFormat currencyInstance = NumberFormat.getCurrencyInstance();
        //获取带有百分比的实例
        NumberFormat percentInstance = NumberFormat.getPercentInstance();
        instance.setGroupingUsed(true);//三个数字为一组
        instance.setMaximumIntegerDigits(3);//设置整数部分的位数, 从小数点处开始计算
        instance.setMaximumFractionDigits(1);//设置小数部分的位数, 从小数点处开始计算
        //同理setMinimum
        String num = instance.format(00100000.1200);//会裁去开头的0和末尾的0
        System.out.println(num + instance.getCurrency());//获取当前系统所在位置的货币单位

```

`DecimalFormat`

```java
DecimalFormat decimalFormat = new DecimalFormat("pattern"); 
DecimalFormat.getInstance();
DecimalFormat.getCurrencyInstance();
System.out.println(decimalFormat.format(0031.230));
```

小数部分都四舍五入

- `#`, 开头的0和末尾的0不显示, `,`分组

  如下表示整数部分3个位一组, 不限位数, 小数部分最多2位

```java
DecimalFormat decimalFormat = new DecimalFormat(",###.##");
System.out.println(decimalFormat.format(02222.20));//2,222.2
```

- `0`, 整数或小数部分少位数补0

```java
DecimalFormat decimalFormat = new DecimalFormat("0,000.00");
System.out.println(decimalFormat.format(22.26));//0,022.26
```

- 如果输入小于1的小数

```java
DecimalFormat decimalFormat = new DecimalFormat("#.00");
System.out.println(decimalFormat.format(0.26));//.26
```

- `%`, 乘100然后加%

```java
DecimalFormat decimalFormat = new DecimalFormat("%");
System.out.println(decimalFormat.format(0.26));//26%
```

