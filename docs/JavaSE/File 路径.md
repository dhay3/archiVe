# File 路径

By default the classes in the `java.io` package always resolve relative pathnames against the
 current user directory.  This directory is named by the system property `user.dir`

```java
    /**
     * 获取当前项目的路径
     * D:\workspace_for_idea\Demo\Test
     */
    @Test
    public void testSystemPro(){
        System.out.println(System.getProperty("user.dir"));
    }
```

> As Follow

```java
 	/**
     * File的相对路径是项目路径 Content Root
     */
    @Test
     public void testFile(){
        //获取当前的项目目录 D:\workspace_for_idea\Demo\Test\
        System.out.println(new File("").getAbsolutePath());//等价于new File(".")

        //获取当前项目目录下的test.txt D:\workspace_for_idea\Demo\Test\test.txt
        System.out.println(new File("test.txt").getAbsolutePath());

        //获取根分路径 D:\
        System.out.println(new File("/").getAbsolutePath());
    }
```

