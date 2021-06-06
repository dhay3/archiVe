# su

- -c command

  不分配tty，直接将执行的命令结果返回

  ```
  root in /home/ubuntu λ su -c "top" ubuntu
  ```

- -s shell

  使用指定的shell，全路径

  ```
  root in /home/ubuntu λ su ubuntu -s /bin/zsh
  ```

- `- | -l`

  登入后到用户的家目录

  ```
  sudo - ubuntu
  ```

  