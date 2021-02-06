# getResource() 和 getClassLoader().getResource()

## getResource()

```java
/**
 * @Author: CHZ
 * @DateTime: 2020/8/13 23:35
 * @Description: TODO
 * src\io\path\getResourcePathTest.java
 */
public class getResourcePathTest {
    @Test
    public void testGetResource(){
        //获取相对getResourcePathTest包下的a.txt
        //file:/D:/workspace_for_idea/Demo/Test/out/production/Test/io/path/a.txt
        System.out.println(getResourcePathTest.class.getResource("a.txt"));
        
        //获取类加载根目录Source Root
        //file:/D:/workspace_for_idea/Demo/Test/out/production/Test/
        System.out.println(getResourcePathTest.class.getResource("/"));
    }
}
```

## getClassLoader().getResource()

```java
    @Test
    public void testGetClassLoaderResource(){
        //相对于Source Root
        //file:/D:/workspace_for_idea/Demo/Test/out/production/Test/
        System.out.println(getResourcePathTest.class.getClassLoader().getResource(""));

        //file:/D:/workspace_for_idea/Demo/Test/out/production/Test/a.txt
        System.out.println(getResourcePathTest.class.getClassLoader().getResource("a.txt"));
        
        //以 / 开头 返回null
        System.out.println(getResourcePathTest.class.getClassLoader().getResource("/a.txt"));
    }
```

