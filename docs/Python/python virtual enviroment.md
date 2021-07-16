# python virtual enviroment

假如A需要lib version1.0，但是B需要lib version2.0。如果将两个包都安装就会导致应用无法运行。这时可以在应用自己的目录下创建venv来解决这个问题。

```
cpl in /etc λ pip freeze
apparmor==3.0.1
appdirs==1.4.4
application-utility==1.3.2
attrs==21.2.0
Brlapi==0.8.2
btrfsutil==5.12.1


cpl in /tmp λ python -m venv tutorial-env
cpl in /etc λ cd /tmp 
cpl in /tmp λ cd tutorial-env 
cpl in /tmp/tutorial-env λ source bin/activate
(tutorial-env) cpl in /tmp/tutorial-env λ pip freeze
(tutorial-env) cpl in /tmp/tutorial-env λ 
```

这里可以看到pip freeze显示的包为空了