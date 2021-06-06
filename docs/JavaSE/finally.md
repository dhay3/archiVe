# finally

只有try代码块被执行到了, finally才会执行

finally 一定是在return之前执行的

如果finally中带有return那么最后的返回值将被finally中的return修改

异常被catch后也会执行